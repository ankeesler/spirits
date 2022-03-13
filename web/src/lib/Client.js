import log from './log';

class Client {
  constructor(ws) {
    this._ws = ws;

    this._listeners = new Map();

    this._ws.addEventListener('open', (e) => {
      log('ws opened...again?');
    });
    this._ws.addEventListener('close', (e) => {
      log('ws closed');
    });
    this._ws.addEventListener('error', (e) => {
      log(`ws error: ${e}`);
    });
    this._ws.addEventListener('message', (e) => {
      log(`ws message`);

      let message;
      try {
        message = JSON.parse(e.data);
      } catch (error) {
        log(`could not parse message data (${e.data}): ${error}`);
        return;
      }

      if (this._listeners.has(message.type)) {
        this._listeners.get(message.type).forEach((l) => l(message));
      }
    });
  };

  on(type, callback) {
    if (!this._listeners.has(type)) {
      this._listeners.set(type, []);
    }
    this._listeners.get(type).push(callback);
  }

  send(type, details) {
    const message = {
      type: type,
      details: details,
    };
    this._ws.send(JSON.stringify(message));
  };
};

export default Client;