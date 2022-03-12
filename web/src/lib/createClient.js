import log from './log';
import Client from './Client'

class FakeWebsocket {
  addEventListener(id, callback) {
    if (id === 'open') {
      callback();
    } else if (id === 'message') {
      this._message_callback = callback;
    }
  }

  send(data) {
    let message;
    try {
      message = JSON.parse(data);
    } catch (_) {
      const error = {
        type: 'error',
        details: {
          reason: "invalid request json",
        },
      }
      setTimeout(() => {
        this._message_callback(JSON.stringify(error));
      }, 100);
    }

    let obj;
    switch (message.type) {
      case 'battle-start':
        obj = {
          type: 'battle-stop',
          details: {
            output: 'some-battle-output',
          },
        };
        break;
      case 'spirit-req':
        obj = {
          type: 'spirit-rsp',
          details: {
            spirits: [{name: 'a', health: 3, power: 1}, {name: 'b', health: 3, power: 2}],
          },
        };
        break;
      default:
        obj = {
          type: 'error',
          details: {
            reason: "invalid request type: " + message.type,
          },
        };
    }
    setTimeout(() => {
      this._message_callback({data: JSON.stringify(obj)});
    }, 100);
  }
}

const tryWs = (wsDetails) => {
  return new Promise((resolve) => {
    const ws = wsDetails.create();
    ws.addEventListener('open', () => {
      wsDetails.ws = ws;
      resolve(wsDetails);
    });
  });
};

const createClient = (callback) => {
  const promises = [
    tryWs({
      name: 'production',
      create: () => {
        const scheme = (window.location.protocol === 'https:' ? 'wss://' : 'ws://');
        return new WebSocket(scheme + window.location.host + '/api/battle');
      },
    }),
    tryWs({
      name: 'development',
      create: () => {
        return new WebSocket('ws://localhost:12345/api/battle');
      },
    }),
  ];
  Promise.any(promises).then(wsDetails => {
    log(`choosing websocket: ${wsDetails.name}`);
    callback(new Client(wsDetails.ws));
  }).catch(error => {
    log('falling back to fake websocket');
    callback(new Client(new FakeWebsocket()));
  });
};

export default createClient;