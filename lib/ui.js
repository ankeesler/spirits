class UI {
  constructor(out) {
    this._out = out;
  }

  onSpirits(spirits) {
    this._out.write('> summary\n');
    spirits.forEach((spirit) => {
      this._out.write(`  ${spirit.name}: ${spirit.health}\n`);
    });
  }
};

exports.UI = UI;
