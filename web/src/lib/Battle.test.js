import Battle from './Battle';

describe('Battle', () => {
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

    it('calls client.send() with a battle-start message', () => {
      expect(client.send.mock.calls).toEqual([['battle-start', {spirits: spirits}]]);
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

      it('resolves the promise with the details', async () => {
        expect.assertions(1);
        await expect(p).resolves.toEqual({output: output});
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

    describe('an action-req arrives', () => {
      let spirit, output;

      beforeEach(() => {
        spirit = {name: 'a', actions: ['foo', 'bar']};
        output = 'some output';
        clientListeners.get('action-req')({
          type: 'action-req',
          details: {
            spirit: spirit,
            output: output,
          },
        });
      });

      it('resolves the promise with the spirit and output', async () => {
        expect.assertions(1);
        await expect(p).resolves.toEqual({
          spirit: spirit,
          output: output,
        });
      });

      describe('the client sends an action-rsp', () => {
        let id;
        beforeEach(() => {
          id = 'bar';
          p = b.action(spirit, id);
        });

        it('calls client.send() with a action-rsp message', async () => {
          expect(client.send.mock.calls).toEqual([
            ['battle-start', {spirits: spirits}],
            ['action-rsp', {spirit: spirit, id: id}],
          ]);
        });

        describe('another action-req arrives', () => {
          let anotherSpirit, anotherOutput;

          beforeEach(() => {
            anotherSpirit = {name: 'b', actions: ['bar', 'bat']};
            anotherOutput = 'some more output';
            clientListeners.get('action-req')({
              type: 'action-req',
              details: {
                spirit: anotherSpirit,
                output: anotherOutput,
              },
            });
          });

          it('resolves the promise with the spirit and output again', async () => {
            expect.assertions(1);
            await expect(p).resolves.toEqual({
              spirit: anotherSpirit,
              output: anotherOutput,
            });
          });

          describe('the client sends an action-rsp', () => {
            let anotherId = 'bat';
            beforeEach(() => {
              anotherId = 'bat';
              p = b.action(anotherSpirit, anotherId);
            });

            it('calls client.send() with a action-rsp message again', async () => {
              expect(client.send.mock.calls).toEqual([
                ['battle-start', {spirits: spirits}],
                ['action-rsp', {spirit: spirit, id: id}],
                ['action-rsp', {spirit: anotherSpirit, id: anotherId}],
              ]);
            });

            describe('a battle-stop arrives', () => {
              let finalOutput;

              beforeEach(() => {
                output = 'some output';
                clientListeners.get('battle-stop')({
                  type: 'battle-stop',
                  details: {
                    output: finalOutput,
                  },
                });
              });

              it('resolves the promise with the details', async () => {
                expect.assertions(1);
                await expect(p).resolves.toEqual({output: finalOutput});
              });
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

      describe('the client sends battle-stop', () => {
        beforeEach(() => {
          p = b.stop();
        });

        it('calls client.send() with a action-rsp message', async () => {
          expect(client.send.mock.calls).toEqual([
            ['battle-start', {spirits: spirits}],
            ['battle-stop', {}],
          ]);
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
