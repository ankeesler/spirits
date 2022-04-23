const log = require('../log');
const date = require('../date');

const validate = (spirit, callback) => {
  callback();
};

const reconcile = (client, informer, spirit) => {
  log(`reconciling ${JSON.stringify(spirit)}`);

  validate(spirit, (error) => {
    let condition = {
      type: 'Ready',
      status: 'True',
      // lastTransitionTime: date(),
    };
    if (error) {
      condition.status = 'False';
      condition.Reason = 'Error';
      condition.Message = error;
    }

    spirit.status = {
      conditions: [condition],
    };
    client.updateStatus(spirit);
  });
};

module.exports = {
  make: (client, informer) => {
    const reconcileFn = (obj) => reconcile(client, informer, obj);
    informer.on('add', reconcileFn);
    informer.on('update', reconcileFn);
  },
};
