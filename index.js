if (require.main === module) {
  require('./lib/main.js')(process);
} else {
  throw new Error('error: must run this module as main');
}
