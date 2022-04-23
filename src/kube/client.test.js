import { Client } from './client';

describe("Client", () => {
  let ns;
  let nsClient;

  beforeEach(() => {
    ns = {
      metadata: {
        name: 'test-ns',
      },
    };
    nsClient = new Client({
      url: 'https://127.0.0.1:50202',
      group: '',
      version: 'v1',
      resource: 'namespaces',
    });
    nsClient.create(ns);
  });

  afterEach(() => {
    nsClient.delete(ns.metadata.name);
  });

  it('lists cluster-scoped built-in resources', () => {
    expect.assertions(1);
    return expect(nsClient.list()).resolves.toEqual([ns]);
  });
});
