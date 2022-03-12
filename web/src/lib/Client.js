import log from './log';

class Client {
  constructor(ws) {
    this._ws = ws;

    this._battle_callback = null;
    this._spirit_callback = null;

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
          log(`unexpected message type: ${message.type}`);
      };
    });
  };

  _onBattleStop(message) {
    if (!this._battle_callback) {
      log(`unexpected battle-stop: ${message.details}`);
      return;
    } 
    
    this._battle_callback('', message.details.output);
    this._battle_callback = null;
  }

  _onSpiritRsp(message) {
    if (!this._spirit_callback) {
      log(`unexpected spirit-rsp: ${message.details}`);
      return;
    }

    this._spirit_callback('', message.details.spirits);
    this._spirit_callback = null;
  }

  _onError(message) {
    log(`error message: ${message.details.reason}`);
    if (this._battle_callback) {
      this._battle_callback(message.details.reason);
      this._battle_callback = null;
    }
    if (this._spirit_callback) {
      this._spirit_callback(message.details.reason);
      this._spirit_callback = null;
    }
  }

  startBattle(spirits, callback) {
    if (this._battle_callback) {
      callback('battle already running');
      return;
    }

    const battleStart = {
      type: 'battle-start',
      details: {
        spirits: spirits,
      },
    };
    this._ws.send(JSON.stringify(battleStart));
    this._battle_callback = callback;
  };

  generateSpirits(callback) {
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

export default Client;