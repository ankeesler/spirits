import Generator from './Generator';

describe('Generator', () => {
  let g;
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
    g = new Generator(client);
  });

  describe('generate()', () => {
    let p;

    beforeEach(() => {
      p = g.generate();
    });

    it('calls client.send()', () => {
      expect(client.send.mock.calls).toEqual([['spirit-req', {}]])
    });

    describe('a spirit-rsp arrives', () => {
      let spirits;

      beforeEach(() => {
        spirits = ['some spirits'];
        clientListeners.get('spirit-rsp')({
          type: 'spirit-rsp',
          details: {
            spirits: spirits,
          },
        });
      });

      it('resolves the promise with the spirits', async () => {
        expect.assertions(1);
        await expect(p).resolves.toEqual(spirits);
      });

      describe('another spirit-rsp arrives', () => {
        let spirits;
  
        beforeEach(() => {
          spirits = ['some spirits'];
          clientListeners.get('spirit-rsp')({
            type: 'spirit-rsp',
            details: {
              spirits: spirits,
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
