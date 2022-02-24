exports.onSpirits = (spirits) => {
  console.log('> summary');
  spirits.forEach((spirit) => {
    console.log('  %s: %s', spirit.name, spirit.health);
  });
};
