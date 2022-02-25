const EventsEmitter = require('events');

class Strategy {
  constructor(spirits) {
    this.spirits = spirits;
    this.first = false;
  }

  hasNext() {
    return this.spirits.filter((s) => s.health > 0).length > 1;
  }

  next() {
    this.first = !this.first;
    if (this.first) {
      return [this.spirits[0], this.spirits[1]];
    } else {
      return [this.spirits[1], this.spirits[0]];
    }
  }
};

class Battle extends EventsEmitter {
  constructor() {
    super();
  }

  run(spirits) {
    this.emit('spirits', spirits);

    const strategy = new Strategy(spirits);
    while (strategy.hasNext()) {
      const [from, to] = strategy.next();
      to.health -= from.power;
      if (to.health < 0) {
        to.health = 0;
      }
      this.emit('spirits', spirits);
    }

    this.emit('end');
  }
};

exports.Battle = Battle;
