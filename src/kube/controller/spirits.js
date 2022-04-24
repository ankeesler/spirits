const action = require('../../internal/action');
const conditions = require('../conditions');
const log = require('../../log');

const toInternalSpirit = (spirit) => {
  return {
    name: spirit.metadata.name,
    stats: spirit.stats,
    action: action.attack,
  };
};

const reconcile = async (client, spiritsCache, spirit) => {
  log(`spirit-controller: reconciling ${JSON.stringify(spirit)}`);

  let error;
  try {
    spiritsCache.set(spirit.metadata.name, toInternalSpirit(spirit));
  } catch (e) {
    error = e.message;
  }

  if (conditions.upsert(
    spirit,
    'Ready',
    error ? 'False' : 'True',
    error ? 'Error' : 'Valid',
    error ? error : 'spirit is valid',
  )) {
    await client.updateStatus(spirit);
  }
};

module.exports = {
  make: (client, informer, spiritsCache) => {
    const reconcileFn = async (obj) => reconcile(client, spiritsCache, obj);
    informer.on('add', reconcileFn);
    informer.on('update', reconcileFn);
  },
};
