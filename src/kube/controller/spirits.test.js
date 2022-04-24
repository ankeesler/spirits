const spiritsController = require('./spirits');

jest.mock('../date', () => {
  return () => 'some-date';
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

  describe('on good spirit create', () => {
    let spirit;
    beforeEach(() => {
      spirit = {
        metadata: {
          name: 'some-name',
          generation: 555,
        },
        spec: {
          stats: {
            health: 5,
          },
        },
      };
      informer.on.mock.calls.find(call => call[0] === 'add')[1](spirit);
    });

    it('updates the conditions', () => {
      const updatedSpirit = JSON.parse(JSON.stringify(spirit));
      updatedSpirit.status.conditions = [
        {
          type: 'Ready',
          status: 'True',
          reason: 'Valid',
          message: 'spirit is valid',
          observedGeneration: 555,
          lastTransitionTime: 'some-date',
        },
      ];
      expect(client.updateStatus.mock.calls).toEqual([[updatedSpirit]]);
    });
  });

  describe('on good spirit update', () => {
    beforeEach(() => {
      spirit = {
        metadata: {
          name: 'some-name',
          generation: 555,
        },
        spec: {
          stats: {
            health: 5,
          },
        },
        status: {
          conditions: [
            {
              type: 'Ready',
              status: 'True',
              reason: 'Valid',
              message: 'spirit is valid',
              observedGeneration: 555,
              lastTransitionTime: 'some-date',
            },
          ]
        },
      };
      informer.on.mock.calls.find(call => call[0] === 'update')[1](spirit);
    });

    it('does not update the conditions', () => {
      expect(client.updateStatus.mock.calls).toEqual([]);
    });
  });
});