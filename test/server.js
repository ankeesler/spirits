const fs = require('fs');
const path = require('path');

const request = require('supertest');
const t = require('tap');

const server = require('../lib/server');

function readSpirits(fixture) {
  const data = fs.readFileSync(path.join(__dirname, 'fixture', fixture));
  const spirits = JSON.parse(data);
  return spirits;
}

t.test('it can create battles', (t) => {
  const s = server.createServer();
  request(s)
    .post('/battles')
    .send(readSpirits('good-spirits.json'))
    .set('Content-Type', 'application/json')
    .set('Accept', 'application/json')
    .expect(201)
    .expect('Content-Type', 'application/json')
    .expect('Location')
    .end((err, res) => {
      t.equal(err, null);
    });
});