class Generator {
  constructor(client) {
    this._client = client;
    this._promises = [];

    client.on('spirit-rsp', (message) => {
      if (this._promises.length === 0) {
        return;
      }

      const [resolve] = this._promises.shift();
      resolve(message.details.spirits);
    });

    client.on('error', (message) => {
      if (this._promises.length === 0) {
        return;
      }

      const [,reject] = this._promises.shift();
      reject(message.details.reason);
    });
  }

  generate(callback) {
    return new Promise((resolve, reject) => {
      this._client.send('spirit-req', {});
      this._promises.push([resolve, reject]);
    });
  }
}

export default Generator;