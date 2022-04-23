class Strategy {
  constructor(spirits) {
    this._spirits = spirits;
    this._next = 0;
  }

  hasNext() {
    return this._spirits.filter((spirit) => spirit.stats.health > 0).length > 1;
  }

  next() {
    return [this._spirits[this._next++ & 1], this._spirits[this._next & 1]];
  }
};

class Battle {
  constructor(spirits, callback) {
    this._spirits = spirits;
    this._callback = callback;
  }

  start() {
    this._callback(this._spirits);

    const strategy = new Strategy(this._spirits);
    while (strategy.hasNext()) {
      const [from, to] = strategy.next();
      from.action(from, to);

      this._callback(this._spirits);
    }
  }
};

module.exports = {
  Battle: Battle,
}