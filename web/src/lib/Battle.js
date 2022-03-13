class Battle {
  constructor(client) {
    this._client = client;
    this._promises = [];

    client.on('battle-stop', (message) => {
      if (this._promises.length === 0) {
        return;
      }

      const [resolve] = this._promises.shift();
      resolve(message.details);
    });

    client.on('action-req', (message) => {
      if (this._promises.length === 0) {
        return;
      }

      const [resolve] = this._promises.shift();
      resolve(message.details);
    });

    client.on('error', (message) => {
      if (this._promises.length === 0) {
        return;
      }

      const [,reject] = this._promises.shift();
      reject(message.details.reason);
    });
  }

  start(spirits) {
    return new Promise((resolve, reject) => {
      this._client.send('battle-start', {spirits: spirits});
      this._promises.push([resolve, reject]);
    });
  }

  stop() {
    return new Promise((resolve, reject) => {
      this._client.send('battle-stop', {});
      this._promises.push([resolve, reject]);
    });
  }

  action(spirit, id) {
    return new Promise((resolve, reject) => {
      this._client.send('action-rsp', {spirit: spirit, id: id});
      this._promises.push([resolve, reject]);
    });
  }
}

export default Battle;