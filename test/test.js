const child_process = require('child_process');
const path = require('path');
const process = require('process');

const main = require('../src/main');

const spiritsYamlPath = path.join(__dirname, 'fixture', 'spirits.yaml');

const kubectl = (...args) => {
  console.log('running:', 'kubectl', args);
  child_process.execFileSync('kubectl', args);
};

let teardownMainFn;

beforeEach(async () => {
  // Run main to set everything up.
  await main.main((fn) => {
    teardownMainFn = fn;
  });
});

afterEach(() => {
  // Teardown main.
  teardownMainFn();
});

describe('spirits', () => {
  beforeEach(() => {
    // Create resources.
    kubectl('apply', '-f', spiritsYamlPath);
  });

  afterEach(() => {
    // Delete resources.
    kubectl('delete', '-f', spiritsYamlPath);
  });

  it('upserts spirits API', () => {
    // Wait for spirits to be ready.
    kubectl('wait', '-f', spiritsYamlPath, '--for', 'condition=Ready', '--timeout', '3s');
  });
});
