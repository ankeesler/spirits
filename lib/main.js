const server = require('./server');

function reallyMain(process) {
  const port = process.env.PORT || 12345;
  server.createServer().listen(port, () => {
    console.log('listening on port %d', port);
  });
}

module.exports = (process) => {
  try {
    reallyMain(process);
  } catch (e) {
    console.error('error: %s', e.message);
    process.exitCode = 1;
  }
};
