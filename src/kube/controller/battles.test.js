const battlesController = require('./battles');

jest.mock('../date', () => {
  return () => 'some-date';
});

describe('spirits controller', () => {
  let battlesClient, spiritsClient, battlesInformer, spiritsInformer;

  beforeEach(() => {
    battlesClient = {
      updateStatus: jest.fn(),
    }
    spiritsClient = {
      create: jest.fn(),
      update: jest.fn(),
    }

    battlesInformer = {
      on: jest.fn(),
    }
    spiritsInformer = {
      on: jest.fn(),
    }

    battlesController.make(battlesClient, spiritsClient, battlesInformer, spiritsInformer);
  });

  it('registers itself with the informer', () => {
    expect(battlesInformer.on.mock.calls.length).toEqual(2);

    expect(battlesInformer.on.mock.calls[0][0]).toEqual('add');
    expect(typeof battlesInformer.on.mock.calls[0][1]).toEqual('function');

    expect(battlesInformer.on.mock.calls[1][0]).toEqual('update');
    expect(typeof battlesInformer.on.mock.calls[1][1]).toEqual('function');

    expect(spiritsInformer.on.mock.calls.length).toEqual(2);

    expect(spiritsInformer.on.mock.calls[0][0]).toEqual('update');
    expect(typeof spiritsInformer.on.mock.calls[0][1]).toEqual('function');

    expect(spiritsInformer.on.mock.calls[1][0]).toEqual('delete');
    expect(typeof spiritsInformer.on.mock.calls[1][1]).toEqual('function');
  });
});