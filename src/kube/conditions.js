const date = require('./date');

const get = (obj, type) => {
  if (!obj.status) {
    return false;
  }

  if (!obj.status.conditions) {
    return false;
  }

  const condition = obj.status.conditions.find((condition) => condition.type === type);
  if (condition) {
    return condition.status === 'True';
  }

  return false;
};

const upsert = (obj, type, status, reason, message) => {
  // Ensure conditions array is initialized.
  if (!obj.status) {
    obj.status = {};
  }
  if (!obj.status.conditions) {
    obj.status.conditions = [];
  }
  const conditions = obj.status.conditions;

  // Try to find condition type in existing conditions to see if we need to create or update.
  const condition = conditions.find((condition) => condition.type === type);

  if (condition) {
    // Only update the condition if status or reason is not correct.
    if (condition.status !== status
      || condition.reason !== reason
      || condition.message !== message
      || condition.observedGeneration !== obj.metadata.generation) {
      condition.status = status;
      condition.reason = reason;
      condition.message = message;
      condition.observedGeneration = obj.metadata.generation;
      condition.lastTransitionTime = date();
      return true;
    }
  } else {
    // On create, add a new conditions to the list.
    conditions.push({
      type: type,
      status: status,
      reason: reason,
      message: message,
      observedGeneration: obj.metadata.generation,
      lastTransitionTime: date(),
    });
    return true;
  }

  return false;
};

module.exports = {
  get: get,
  upsert: upsert,
};