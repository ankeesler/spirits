class RealClient {
  constructor() {
    this._ready = false;
    this._callback = null;

    const scheme = (window.location.protocol === 'https:' ? 'wss://' : 'ws://');
    this._ws = new WebSocket(scheme + window.location.host + '/api/battle');

    this._ws.addEventListener('open', (e) => {
      console.log('ws opened');
      this._ready = true;
    });
    this._ws.addEventListener('close', (e) => {
      console.log('ws closed');
    });
    this._ws.addEventListener('error', (e) => {
      console.log(`ws error: ${e}`);
    });
    this._ws.addEventListener('message', (e) => {
      console.log(`ws message`);
      const message = JSON.parse(e.data);
      switch (message.type) {
        case 'battle-stop':
          this._onBattleStop(message);
          break;
        default:
          console.log(`unexpected message type: ${message.type}`);
      };
    });
  };

  _onBattleStop(message) {
    if (!this._callback) {
      console.log(`unexpected battle-stop: ${message.details}`);
      return;
    } 
    
    this._callback('', message.details.output);
    this._callback = null;
  }

  startBattle(spirits, callback) {
    if (!this._ready) {
      callback('client not ready');
      return;
    }
    if (this._callback) {
      callback('battle already running');
      return;
    }

    let spiritsObj = null;
    try {
      spiritsObj = JSON.parse(spirits);
    } catch (error) {
      spiritsObj = null;
    }
    if (!spiritsObj) {
      callback('invalid spirits JSON');
    }

    const battleStart = {
      type: 'battle-start',
      details: {
        spirits: spiritsObj,
      },
    };
    this._ws.send(JSON.stringify(battleStart));
    this._callback = callback;
  };

  generateSpirits(callback) {
    fetch('/api/spirit', {
      method: 'POST',
    }).then((rsp) => {
      if (!rsp.ok) {
        callback(`unexpected response: ${rsp.status}`);
        return;
      }
      return rsp.text();
    }).then((text) => {
      callback('', text);
    }).catch((error) => {
      callback(error);
    });
  };
};

export default RealClient;