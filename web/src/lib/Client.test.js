import Client from './Client';

describe('Client', () => {
  let wsListeners;
  let ws;
  let client;

  beforeEach(() => {
    wsListeners = new Map();
    ws = {
      addEventListener: (id, cb) => {
        wsListeners.set(id, cb);
      },
      send: jest.fn(),
    };
    client = new Client(ws);
  });

  describe('send()', () => {
    let type, details;

    beforeEach(() => {
      type = 'spirit-req';
      details = {some: 'details'};
      client.send(type, details);
    });

    it('calls ws.send()', () => {
      const message = {
        type: type,
        details: details,
      };
      expect(ws.send.mock.calls).toEqual([[JSON.stringify(message)]]);
    });
  });

  describe('on()', () => {
    let battleStopListeners, spiritRspListeners;

    beforeEach(() => {
      battleStopListeners = [jest.fn(), jest.fn()];
      battleStopListeners.forEach((l) => client.on('battle-stop', l));

      spiritRspListeners = [jest.fn(), jest.fn()];
      spiritRspListeners.forEach((l) => client.on('spirit-rsp', l));
    });

    describe('a battle-stop arrives', () => {
      let rsp;
      beforeEach(() => {
        rsp = {
          type: 'battle-stop',
          details: {
            output: 'some output',
          },
        };
        const message = {
          data: JSON.stringify(rsp),
        }
        wsListeners.get('message')(message);
      });

      it('sends events to only battle-stop listeners', async () => {
        battleStopListeners.forEach((l) => expect(l.mock.calls).toEqual([[rsp]]));
        spiritRspListeners.forEach((l) => expect(l.mock.calls).toEqual([]));
      });
    });

    describe('a spirit-rsp arrives', () => {
      let rsp;
      beforeEach(() => {
        rsp = {
          type: 'spirit-rsp',
          details: {
            spirits: ['some spirits'],
          },
        };
        const message = {
          data: JSON.stringify(rsp),
        }
        wsListeners.get('message')(message);
      });

      it('sends events to only spirit-rsp listeners', async () => {
        battleStopListeners.forEach((l) => expect(l.mock.calls).toEqual([]));
        spiritRspListeners.forEach((l) => expect(l.mock.calls).toEqual([[rsp]]));
      });
    });

    describe('a random message type arrives', () => {
      let rsp;
      beforeEach(() => {
        rsp = {
          type: 'random',
          details: {},
        };
        const message = {
          data: JSON.stringify(rsp),
        }
        wsListeners.get('message')(message);
      });

      it('does not send events to any listeners', async () => {
        battleStopListeners.forEach((l) => expect(l.mock.calls).toEqual([]));
        spiritRspListeners.forEach((l) => expect(l.mock.calls).toEqual([]));
      });
    });

    describe('an invalid json message arrives', () => {
      beforeEach(() => {
        const message = {
          data: 'bad [ json',
        }
        wsListeners.get('message')(message);
      });

      it('does not send events to any listeners', async () => {
        battleStopListeners.forEach((l) => expect(l.mock.calls).toEqual([]));
        spiritRspListeners.forEach((l) => expect(l.mock.calls).toEqual([]));
      });
    });
  });
});
