const path = require('path');
const t = require('tap');

const index = require.resolve('../index.js');
const packageIndex = require.resolve('../');

function runProcess(fixture, callback) {
  const {spawn} = require('child_process');
  const argv = [index, path.join(__dirname, 'fixture', fixture)];

  const p = spawn(process.execPath, argv);

  const out = [];
  const err = [];
  p.stdout.on('data', (c) => out.push(c));
  p.stderr.on('data', (c) => err.push(c));

  p.on('close', (code, signal) => {
    const stdout = Buffer.concat(out).toString();
    const stderr = Buffer.concat(err).toString();
    callback(code, signal, stdout, stderr);
  });
}

t.equal(index, packageIndex, 'index is main package require() export');
t.throws(() => require(index), {
  message: 'error: must run this module as main',
});

t.test('runs a battle between two spirits', (t) => {
  runProcess('good-spirits.json', (code, signal, stdout, stderr) => {
    t.equal(code, 0);
    t.equal(signal, null);
    const expectedStdout = `> summary
  a: 5
  b: 5
> summary
  a: 5
  b: 4
> summary
  a: 3
  b: 4
> summary
  a: 3
  b: 3
> summary
  a: 1
  b: 3
> summary
  a: 1
  b: 2
> summary
  a: 0
  b: 2
`;
    t.equal(stdout, expectedStdout);
    t.equal(stderr, '');
    t.end();
  });
});

t.test('throws error when there are more than 2 spirits', (t) => {
  runProcess('3-spirits.json', (code, signal, stdout, stderr) => {
    t.equal(code, 1, stderr);
    t.equal(signal, null);
    t.equal(stdout, '');
    t.equal(stderr, 'error: must pass 2 spirits\n');
    t.end();
  });
});

t.test('throws error when there is less than 2 spirits', (t) => {
  runProcess('1-spirit.json', (code, signal, stdout, stderr) => {
    t.equal(code, 1, stderr);
    t.equal(signal, null);
    t.equal(stdout, '');
    t.equal(stderr, 'error: must pass 2 spirits\n');
    t.end();
  });
});
