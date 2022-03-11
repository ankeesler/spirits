import RealClient from './RealClient';

describe('RealClient', () => {
  describe('generateSpirit()', () => {
    [
      {
        name: 'when the ws is not open',
        calls: 1,
        open: false,
        messages: [],
        wantSends: [],
        wantCallbackCalls: [['client not ready']],
      },
      {
        name: 'when the ws is open',
        calls: 1,
        open: true,
        messages: [{type: 'spirit-rsp', details: {spirits: ['some-spirits']}}],
        wantSends: [{type: "spirit-req", details: {}}],
        wantCallbackCalls: [['', '["some-spirits"]']],
      },
      {
        name: 'when the ws is open and there is an error',
        calls: 1,
        open: true,
        messages: [{type: 'error', details: {reason: 'some error'}}],
        wantSends: [{type: "spirit-req", details: {}}],
        wantCallbackCalls: [['some error']],
      },
      {
        name: 'when the ws is open and there is an unexpected battle-stop',
        calls: 1,
        open: true,
        messages: [
          {type: 'battle-stop', details: {}},
          {type: 'spirit-rsp', details: {spirits: ['some-spirits']}},
        ],
        wantSends: [{type: "spirit-req", details: {}}],
        wantCallbackCalls: [['', '["some-spirits"]']],
      },
      {
        name: 'when the ws is open and we call generateSpirits twice',
        calls: 2,
        open: true,
        messages: [{type: 'spirit-rsp', details: {spirits: ['some-spirits']}}],
        wantSends: [{type: "spirit-req", details: {}}],
        wantCallbackCalls: [['spirit request already in flight'], ['', '["some-spirits"]']],
      },
    ].forEach(test => {
      const wsListeners = new Map();
      const ws = {
        addEventListener: (id, cb) => {
          wsListeners.set(id, cb);
        },
      };
      const client = new RealClient(ws);
      const f = () => {
        const wsSends = [];
        ws.send = (d) => {
          wsSends.push(d);
        };

        if (test.open) {
          wsListeners.get('open')();
        }

        if (test.open) {
          wsListeners.get('open')();
        }

        const callback = jest.fn();
        for (let i = 0; i < test.calls; i++) {
          client.generateSpirits(callback);
        }

        test.messages
          .map(JSON.stringify)
          .map((m) => {return {data: m}})
          .forEach(wsListeners.get('message'));

        expect(wsSends).toEqual(test.wantSends.map(JSON.stringify));
        expect(callback.mock.calls).toEqual(test.wantCallbackCalls);
      };
      it(test.name + ' (run 0/1)', f)
      it(test.name + ' (run 1/1)', f)
   });
  });

  describe('startBattle()', () => {
    [
      {
        name: 'when the ws is not open',
        calls: 1,
        open: false,
        messages: [],
        wantSends: [],
        wantCallbackCalls: [['client not ready']],
      },
      {
        name: 'when the ws is open',
        calls: 1,
        open: true,
        spirits: JSON.stringify(['some-spirits']),
        messages: [{type: 'battle-stop', details: {output: 'some-output'}}],
        wantSends: [{type: "battle-start", details: {spirits: ['some-spirits']}}],
        wantCallbackCalls: [['', 'some-output']],
      },
      {
        name: 'when the ws is open and we send bad spirits',
        calls: 1,
        open: true,
        spirits: 'invalid [ json',
        messages: [],
        wantSends: [],
        wantCallbackCalls: [['invalid spirits JSON']],
      },
      {
        name: 'when the ws is open and there is an error',
        calls: 1,
        open: true,
        spirits: JSON.stringify(['some-spirits']),
        messages: [{type: 'error', details: {reason: 'some error'}}],
        wantSends: [{type: "battle-start", details: {spirits: ['some-spirits']}}],
        wantCallbackCalls: [['some error']],
      },
      {
        name: 'when the ws is open and there is an unexpected spirit-rsp',
        calls: 1,
        open: true,
        spirits: JSON.stringify(['some-spirits']),
        messages: [
          {type: 'spirit-rsp', details: {spirits: ['some-spirits']}},
          {type: 'battle-stop', details: {output: 'some-output'}},
        ],
        wantSends: [{type: "battle-start", details: {spirits: ['some-spirits']}}],
        wantCallbackCalls: [['', 'some-output']],
      },
      {
        name: 'when the ws is open and we call startBattle twice',
        calls: 2,
        open: true,
        spirits: JSON.stringify(['some-spirits']),
        messages: [{type: 'battle-stop', details: {output: 'some-output'}}],
        wantSends: [{type: "battle-start", details: {spirits: ['some-spirits']}}],
        wantCallbackCalls: [['battle already running'], ['', 'some-output']],
      },
    ].forEach(test => {
      const wsListeners = new Map();
      const ws = {
        addEventListener: (id, cb) => {
          wsListeners.set(id, cb);
        },
      };
      const client = new RealClient(ws);
      const f = () => {
        const wsSends = [];
        ws.send = (d) => {
          wsSends.push(d);
        };

        if (test.open) {
          wsListeners.get('open')();
        }

        const callback = jest.fn();
        for (let i = 0; i < test.calls; i++) {
          client.startBattle(test.spirits, callback);
        }

        test.messages
          .map(JSON.stringify)
          .map((m) => {return {data: m}})
          .forEach(wsListeners.get('message'));

        expect(wsSends).toEqual(test.wantSends.map(JSON.stringify));
        expect(callback.mock.calls).toEqual(test.wantCallbackCalls);
      };
      it(test.name + ' (run 0/1)', f)
      it(test.name + ' (run 1/1)', f)
    });
  });
});
