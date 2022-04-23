const spiritsController = require('./spirits');

jest.mock('../date', () => {
  return () => 'some date';
});

describe('spirits controller', () => {
  let client, informer;

  beforeEach(() => {
    client = {
      updateStatus: jest.fn(),
    }
    informer = {
      on: jest.fn(),
    }
    spiritsController.make(client, informer);
  });

  it('registers itself with the informer', () => {
    expect(informer.on.mock.calls.length).toEqual(2);

    expect(informer.on.mock.calls[0][0]).toEqual('add');
    expect(typeof informer.on.mock.calls[0][1]).toEqual('function');

    expect(informer.on.mock.calls[1][0]).toEqual('update');
    expect(typeof informer.on.mock.calls[1][1]).toEqual('function');
  });

  describe('spirit created', () => {
    let spirit;
    beforeEach(() => {
      spirit = {
        metadata: {
          name: 'some-name',
        },
        spec: {
          stats: {
            health: 5,
          },
        },
        status: {},
      };
      informer.on.mock.calls.find(call => call[0] === 'add')[1](spirit);
    });

    it('updates the conditions', () => {
      const updatedSpirit = JSON.parse(JSON.stringify(spirit));;
      updatedSpirit.status.conditions = [
        {
          type: 'Ready',
          status: 'True',
          lastTransitionTime: 'some date',
        },
      ];
      expect(client.updateStatus.mock.calls).toEqual([[updatedSpirit]]);
    });
  });
});