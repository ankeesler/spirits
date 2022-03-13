class Battle {
  constructor(client) {
    this._client = client;
    this._promises = [];

    client.on('battle-stop', (message) => {
      if (this._promises.length === 0) {
        return;
      }

      const [resolve] = this._promises.shift();
      resolve(message.details.output);
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
}

export default Battle;