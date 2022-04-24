const date = require('./date');

const upsert = (objStatus, type, status, reason) => {
  // If conditions are unset on status, initialize them.
  if (!objStatus.conditions) {
    objStatus.conditions = [];
  }
  const conditions = objStatus.conditions;

  // Try to find condition type in existing conditions to see if we need to create or update.
  const condition = conditions.find((condition) => condition.type === type);

  if (condition) {
    // Only update the condition if status or reason is not correct.
    if (condition.status !== status || condition.reason !== reason) {
      condition.status = status;
      condition.reason = reason;
      condition.lastTransitionTime = date();
      return true;
    }
  } else {
    // On create, add a new conditions to the list.
    conditions.push({
      type: type,
      status: status,
      reason: reason,
      lastTransitionTime: date(),
    });
    return true;
  }

  return false;
};

module.exports = {
  upsert: upsert,
};