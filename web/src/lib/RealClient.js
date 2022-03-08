class RealClient {
  constructor() {
    this._ready = false;
    this._battle_callback = null;
    this._spirit_callback = null;

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
        case 'spirit-rsp':
          this._onSpiritRsp(message);
          break;
        case 'error':
          this._onError(message);
          break;
        default:
          console.log(`unexpected message type: ${message.type}`);
      };
    });
  };

  _onBattleStop(message) {
    if (!this._battle_callback) {
      console.log(`unexpected battle-stop: ${message.details}`);
      return;
    } 
    
    this._battle_callback('', message.details.output);
    this._battle_callback = null;
  }

  _onSpiritRsp(message) {
    if (!this._spirit_callback) {
      console.log(`unexpected spirit-rsp: ${message.details}`);
      return;
    }

    this._spirit_callback('', JSON.stringify(message.details.spirits));
    this._spirit_callback = null;
  }

  _onError(message) {
    console.log(`error message: ${message.details.reason}`);
    this._battle_callback = null;
    this._spirit_callback = null;
  }

  startBattle(spirits, callback) {
    if (!this._ready) {
      callback('client not ready');
      return;
    }
    if (this._battle_callback) {
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
    this._battle_callback = callback;
  };

  generateSpirits(callback) {
    if (!this._ready) {
      callback('client not ready');
      return;
    }
    if (this._spirit_callback) {
      callback('spirit request already in flight');
      return;
    }

    const spiritReq = {
      type: 'spirit-req',
      details: {},
    };
    this._ws.send(JSON.stringify(spiritReq));
    this._spirit_callback = callback;
  };
};

export default RealClient;