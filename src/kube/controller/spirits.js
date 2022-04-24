const log = require('../../log');
const conditions = require('../conditions');

const validate = (spirit, callback) => {
  callback();
};

const reconcile = (client, spirit) => {
  log(`spirit-controller: reconciling ${JSON.stringify(spirit)}`);

  if (!spirit.status) {
    spirit.status = {};
  }

  validate(spirit, (error) => {
    if (conditions.upsert(
      spirit.status,
      'Ready',
      error ? 'False' : 'True',
      error ? 'Error' : 'Valid',
    )) {
      client.updateStatus(spirit);
    }
  });
};

module.exports = {
  make: (client, informer) => {
    const reconcileFn = (obj) => reconcile(client, obj);
    informer.on('add', reconcileFn);
    informer.on('update', reconcileFn);
  },
};
