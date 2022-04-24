const child_process = require('child_process');
const fs = require('fs');
const path = require('path');

const main = require('../src/main');

const spiritsYamlPath = path.join(__dirname, 'fixture', 'spirits.yaml');

describe('spirits', () => {
  beforeEach(async () => {
    // Run main to set everything up.
    await main.main();

    // Create resources.
    child_process.execFileSync(
      'kubectl',
      ['apply', '-f', spiritsYamlPath],
    );
  });

  afterEach(async () => {
    // Delete resources.
    child_process.execFileSync(
      'kubectl',
      ['delete', '-f', spiritsYamlPath],
    );
  });

  it('upserts spirits API', async () => {
    // Wait for spirits to be ready.
    child_process.execFileSync(
      'kubectl',
      ['wait', '-f', spiritsYamlPath, '--for', 'condition=Ready'],
    );
  });
});