const {spawn, spawnSync} = require('child_process');
const http = require('http');
const path = require('path');

let p = null;

const build = () => {
  const image = 'spirits:spirits-test';
  const rootDir = path.join(__dirname, '..');
  const p = spawnSync('docker', ['build', '-t', image, rootDir]);
  if (p.error) {
    throw new Error('server build failed: ' + p.error.message);
  }
  return image;
};

const reallyStart = (image, onClose) => {
  const out = [];
  const err = [];
  p = spawn('docker', ['run', '-p', '12345:12345', '--name', 'spirits-under-test', image]);
  p.stdout.on('data', c => out.push(c));
  p.stderr.on('data', c => err.push(c));
  p.on('close', (code, signal) => {
    const details = {
      code: code,
      signal: signal,
      stdout: Buffer.concat(out).toString(),
      stderr: Buffer.concat(err).toString(),
    };
    onClose(details);
  });
  return 'http://localhost:12345';
};

const wait = (baseURL) => {
  return new Promise((resolve, reject) => {
    let tries = 3;
    const get = () => {
      try {
        http.get(baseURL, (rsp) => {
          if (rsp.statusCode !== 200) {
            if (!--tries) {
              reject('server never came up');
            } else {
              console.log(`server not up yet (error: response status code was ${rsp.statusCode})`);
              setTimeout(get, 1000);
            }
          } else {
            resolve();
          }
        });
      } catch (error) {
        console.log(`server not up yet (error: ${error})`);
        --tries;
        setTimeout(get, 1000);
      }
    };
    setTimeout(get, 1000);
  });
};

exports.start = async (onClose) => {
  if (process.env.SPIRITS_TEST_URL) {
    return process.env.SPIRITS_TEST_URL;
  }

  if (p) {
    throw new Error('server already started');
  }

  const image = build();
  const baseURL = reallyStart(image, onClose);
  await wait(baseURL);
  return baseURL;
};

exports.stop = () => {
  if (process.env.SPIRITS_TEST_URL) {
    return;
  }

  if (!p) {
    throw new Error('server already stopped');
  }

  const stopP = spawnSync('docker', ['stop', 'spirits-under-test']);
  if (stopP.error) {
    throw new Error('server build failed: ' + stopP.error.message);
  }

  p = null;
};