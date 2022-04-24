const log = require('../../log');
const conditions = require('../conditions');

const reconcileBattle = (battlesClient, spiritsClient, battle) => {
  log(`battle-controller: reconciling ${JSON.stringify(battle)}`);
};

const reconcileSpirit = (spiritsClient, spirit) => {
  log(`battle-controller: reconciling ${JSON.stringify(battle)}`);
};

module.exports = {
  make: (battleClient, spiritClient, battlesInformer, spiritsInformer) => {
    const battlesReconcileFn = (obj) => reconcileBattle(battleClient, spiritClient, obj);
    battlesInformer.on('add', battlesReconcileFn);
    battlesInformer.on('update', battlesReconcileFn);

    const spiritsReconcileFn = (obj) => reconcileSpirit(spiritClient, obj);
    spiritsInformer.on('update', spiritsReconcileFn);
    spiritsInformer.on('delete', spiritsReconcileFn);
  },
};
