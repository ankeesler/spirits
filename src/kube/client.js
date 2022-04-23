class Client {
  constructor(options) {
    this._baseUrl = options.url;

    if (options.group) {
      this._baseUrl += `/apis/${options.group}/`
    } else {
      this._baseUrl += '/api/';
    }

    this._baseUrl += `${options.version}/`;

    if (options.namespace) {
      this._baseUrl += `namespaces/${options.namespace}/`;
    }

    this._baseUrl += `${options.resource}/`;
  }

  create(obj) {
    return fetch(this._baseUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      body: JSON.stringify(obj),
    }).then((rsp) => rsp.json());
  }

  list() {
    return fetch(this._baseUrl, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
      },
    }).then((rsp) => rsp.json());
  }

  get(name) {
    return this._client.getNamespacedCustomObject(this._group, this._version, this._namespace, this._resource, name);
  }

  update(obj) {
    return this._client.replaceNamespacedCustomObject(this._group, this._version, this._namespace, this._resource, obj.metadata.name, obj);
  }

  updateStatus(obj) {
    return this._client.replaceNamespacedCustomObjectStatus(this._group, this._version, this._namespace, this._resource, obj.metadata.name, obj);
  }

  delete(name) {
    return fetch(this._baseUrl, {
      method: 'DELETE',
      headers: {
        'Accept': 'application/json',
      },
    }).then((rsp) => rsp.json());
  }
};

module.exports = {
  Client: Client,
};
