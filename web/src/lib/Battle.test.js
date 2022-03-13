import Battle from './Battle';

describe('Generator', () => {
  let b;
  let client;
  let clientListeners;

  beforeEach(() => {
    clientListeners = new Map();
    client = {
      on: (id, callback) => {
        clientListeners.set(id, callback);
      },
      send: jest.fn(),
    }
    b = new Battle(client);
  });

  describe('start()', () => {
    let spirits;
    let p;

    beforeEach(() => {
      spirits = [{name: 'a'}, {name: 'b'}];
      p = b.start(spirits);
    });

    it('calls client.send()', () => {
      expect(client.send.mock.calls).toEqual([['battle-start', {spirits: spirits}]])
    });

    describe('a battle-stop arrives', () => {
      let output;

      beforeEach(() => {
        output = 'some output';
        clientListeners.get('battle-stop')({
          type: 'battle-stop',
          details: {
            output: output,
          },
        });
      });

      it('resolves the promise with the output', async () => {
        expect.assertions(1);
        await expect(p).resolves.toEqual(output);
      });

      describe('another battle-stop arrives', () => {
        beforeEach(() => {
          spirits = ['some spirits'];
          clientListeners.get('battle-stop')({
            type: 'battle-stop',
            details: {
              output: output,
            },
          });
        });
  
        it('does nothing', async () => {
        });
      });  
    });

    describe('an error arrives', () => {
      beforeEach(() => {
        clientListeners.get('error')({
          type: 'error',
          details: {
            reason: 'some error',
          },
        });
      });

      it('rejects the promise with the error reason', async () => {
        expect.assertions(1);
        await expect(p).rejects.toEqual('some error');
      });
    });
  });
});
