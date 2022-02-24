const fs = require('fs');

const battle = require('./battle');
const ui = require('./ui');

function reallyMain(process) {
  const data = fs.readFileSync(process.argv[2], 'utf8');
  const spirits = JSON.parse(data);
  if (spirits.length != 2) {
    throw new Error('must pass 2 spirits');
  }
  const b = new battle.Battle();
  b.on('spirits', ui.onSpirits);
  b.run(spirits);
}

module.exports = (process) => {
  try {
    reallyMain(process);
  } catch (e) {
    console.error('error: %s', e.message);
    process.exitCode = 1;
  }
};
