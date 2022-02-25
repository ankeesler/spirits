const express = require('express');
const fastify = require('fastify')({logger: true});

const battle = require('./battle');
const ui = require('./ui');


function createServer() {
  fastify.post('/battle', (req, res) {
    const spirits = JSON.parse(data);
    if (spirits.length != 2) {
      throw new Error('must pass 2 spirits');
    }
    const b = new battle.Battle();
    b.on('spirits', ui.onSpirits);
    b.run(spirits);
  });
  return fastify;
};

exports.createServer = createServer;
