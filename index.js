if (require.main === module) {
  require('./src/main').main((fn) => {
    process.on('SIGINT', fn);
  });
} else {
  throw new Error('Must run this module as main program')
}