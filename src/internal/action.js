const attack = (from, to) => {
  to.stats.health -= from.stats.power;
  if (to.stats.health <= 0) {
    to.stats.health = 0;
  }
};

module.exports = {
  attack: attack,
};