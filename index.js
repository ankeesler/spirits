if (require.main === module) {
  require('./src/main').main(process)
} else {
  throw new Error('Must run this module as main program')
}