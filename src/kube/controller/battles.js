const crypto = require('crypto');

const conditions = require('../conditions');
const log = require('../../log');

const internalBattleLib = require('./../../internal/battle');

const digest = (obj) => {
  const hash = crypto.createHash('sha256');
  const message = JSON.stringify(obj.spec);
  hash.update(message);
  return hash.digest('hex').substring(0, 6);
};

const createOrGetSpirit = async (spiritsClient, name, fromSpirit, forBattle) => {
  return spiritsClient.get(name).catch((error) => {
    // TODO: de-dup this logic with CRDs.
    if (error.statusCode === 404) {
      const spirit = {
        apiVersion: fromSpirit.apiVersion,
        kind: fromSpirit.kind,
        metadata: {
          namespace: fromSpirit.metadata.namespace,
          name: name,
          ownerReferences: [
            {
              apiVersion: forBattle.apiVersion,
              kind: forBattle.kind,
              name: forBattle.metadata.name,
              uid: forBattle.metadata.uid,
              controller: true,
            },
          ],
        },
        spec: fromSpirit.spec,
      };
      spirit.labels = {};
      spirit.labels['battle.spirits.dev/battle-name'] = forBattle.metadata.name;
      return spiritsClient.create(spirit);
    }

    // Unexpected error - propagate.
    return Promise.reject(error);
  });
};

const validateBattle = (battle) => {
  const valid = battle.spec.spirits.length === 2;
  return conditions.upsert(
    battle,
    'Valid',
    valid ? 'True' : 'False',
    valid ? 'BattleIsValid' : 'Error',
  );
};

const prepareBattleSpirits = async (spiritsClient, battle) => {
  // TODO: set a non-controller owner reference on the spirit so we can trace back to ourselves?
  let battleSpiritNames = [];
  for (let i = 0; i < battle.spec.spirits.length; i++) {
    const spiritName = battle.spec.spirits[i];
    const spiritRsp = await spiritsClient.get(spiritName);
    const spirit = spiritRsp.body;
    const battleSpiritName = `${battle.metadata.name}-${spirit.metadata.name}-${digest(spirit)}`;
    await createOrGetSpirit(spiritsClient, battleSpiritName, spirit, battle);
    battleSpiritNames.push(battleSpiritName);
  }
  return battleSpiritNames;
};

const readyBattle = async (spiritsClient, battle) => {
  let error;

  // Ensure proper number of spirits.
  if (battle.spec.spirits.length !== 2) {
    error = `expected 2 spirits, got ${battle.spec.spirits.length}`;
  }

  // Prepare battle spirits.
  let battleSpiritNames;
  if (!error) {
    try {
      battleSpiritNames = await prepareBattleSpirits(spiritsClient, battle);
    } catch (e) {
      error = e.message;
      if (e.body) {
        error += ` (${e.body.message})`;
      }
    }
  }

  return conditions.upsert(
    battle,
    'Ready',
    error ? 'False' : 'True',
    error ? 'Error' : 'Success',
    error ? error : `readied battle with spirits ${battleSpiritNames}`,
  );
};

const progressBattle = async (spiritsClient, battleCache, spiritCache, battle) => {
  const ready = conditions.get(battle.status, 'Ready');
  const alreadyStarted = conditions.get(battle, 'Progressing');
  let progressing = alreadyStarted;

  // If the battle is not ready, but already started, stop it.
  if (!ready && alreadyStarted) {
    // Stop battle.
    battleCache.get(battle.metadata.name).stop();
    progressing = false;
  }

  // If the battle is ready, but not started, start it.
  if (ready && !alreadyStarted) {
    // Start battle.
    const internalSpirits = battle.spec.spirits.map(spiritCache.get);
    const internalBattle = new internalBattleLib.Battle(internalSpirits, async (internalSpirits) => {
      internalSpirits.forEach(async (internalSpirit) => {
        const externalSpiritRsp = await spiritsClient.get(internalSpirit.name);
        externalSpiritRsp.body.spec.stats = internalSpirit.stats;
        await spiritsClient.update(externalSpiritRsp.body);
      });
    })
    internalBattle.start();

    // Store battle in cache.
    battleCache.set(battle.metadata.name, internalBattle);

    progressing = true;
  }

  return conditions.upsert(
    battle,
    'Progressing',
    progressing ? 'True' : 'False',
    progressing ? 'Running' : 'NotReady',
    progressing ? 'battle is running' : 'battle not ready',
  );
};

const calculatePhase = (battle) => {
  return battle.status.conditions.find((condition) => condition.status === 'False')
    ? 'Error'
    : 'Ready';
};

const reconcileBattle = async (battlesClient, spiritsClient, battleCache, spiritCache, battle) => {
  log(`battle-controller: reconciling battle ${JSON.stringify(battle)}`);

  // Process conditions.
  let needsUpdate = false;
  needsUpdate ||= await readyBattle(spiritsClient, battle);
  needsUpdate ||= await progressBattle(spiritsClient, battleCache, spiritCache, battle);

  // Set human-readable phase, if necessary.
  const phase = calculatePhase(battle);
  if (battle.status.phase !== phase) {
    needsUpdate = true;
    battle.status.phase = phase;
  }

  if (needsUpdate) {
    await battlesClient.updateStatus(battle);
  }
};

const reconcileSpirit = async (battlesClient, spiritsClient, battleCache, spiritCache, spirit) => {
  log(`battle-controller: reconciling spirit ${JSON.stringify(spirit)}`);

  // Is this spirit owned by a battle? If so, reconcile the battle.
  if (spirit.metadata.ownerReferences) {
    spirit.metadata.ownerReferences.filter((ownerReference) => {
      return ownerReference.apiVersion === 'spirits.dev/v1alpha1' && ownerReference.kind === 'Battle';
    }).forEach(async (ownerReference) => {
      const battleRsp = await battlesClient.get(ownerReference.name);
      await reconcileBattle(battlesClient, spiritsClient, battleCache, spiritCache, battleRsp.body);
    });
  }

  // Is this spirit used by a battle? If so, reconcile the battle.
  // TODO: this seems wasteful - is there a better way here? Use the informer?
  const battlesRsp = await battlesClient.list();
  battlesRsp.body.items.filter((battle) => {
    return battle.spec.spirits.includes(spirit.name)
  }).forEach(async (battle) => {
    await reconcileBattle(battlesClient, spiritsClient, battleCache, spiritCache, battle);
  });
};

module.exports = {
  make: (battlesClient, spiritsClient, battlesInformer, spiritsInformer, spiritCache) => {
    const battleCache = new Map();
    const battlesReconcileFn = async (obj) => reconcileBattle(battlesClient, spiritsClient, battleCache, spiritCache, obj);
    battlesInformer.on('add', battlesReconcileFn);
    battlesInformer.on('update', battlesReconcileFn);

    // If spirits get changed during the course of a battle, the battle logic should get rerun.
    const spiritsReconcileFn = (obj) => reconcileSpirit(battlesClient, spiritsClient, battleCache, spiritCache, obj);
    spiritsInformer.on('add', spiritsReconcileFn);
    spiritsInformer.on('update', spiritsReconcileFn);
    spiritsInformer.on('delete', spiritsReconcileFn);
  },
};
