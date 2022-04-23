const k8s = require('@kubernetes/client-node');

const crds = require('./crds');
const log = require('./log');
const client = require('./client')
const spiritsController = require('./controller/spirits');

const loadKubeconfig = () => {
  const kc = new k8s.KubeConfig();
  kc.loadFromDefault();
  return kc;
};

const installCrds = (kc) => {
  const crdsClient = kc.makeApiClient(k8s.ApiextensionsV1Api);
  return crds.install(crdsClient);
}

const getNamespace = (kc) => {
  const kcNamespace = kc.getContextObject(kc.getCurrentContext()).namespace;
  return kcNamespace ? kcNamespace : 'default';
};

const makeClients = (kc, installedCrdsRsps, namespace) => {
  return installedCrdsRsps.reduce((map, rsp) => {
    const crd = rsp.body;
    map.set(crd.spec.names.plural, new client.Client(
      kc.makeApiClient(k8s.CustomObjectsApi),
      crd.spec.group,
      crd.status.storedVersions[0],
      crd.spec.names.plural,
      namespace,
    ));
    return map;
  }, new Map());
};

const makeInformers = (kc, installedCrdsRsps, namespace, clients) => {
  return installedCrdsRsps.reduce((map, rsp) => {
    const crd = rsp.body;
    const path = `/apis/${crd.spec.group}/${crd.status.storedVersions[0]}/namespaces/${namespace}/spirits`;
    const listFn = () => clients.get(crd.spec.names.plural).list();
    map.set(crd.spec.names.plural, k8s.makeInformer(kc, path, listFn));
    return map;
  }, new Map());
};

const makeControllers = (clients, informers) => {
  spiritsController.make(clients.get('spirits'), informers.get('spirits'));
};

const startInformers = (informers) => {
  informers.forEach((informer, type) => {
    informer.on('error', (error) => {
      log(`${type} informer error: ${error}`);
      setTimeout(() => {
        log(`restarting ${type} informer`);
        informer.start();
      }, 3000);
    });

    log(`starting ${type} informer`);
    informer.start();
  });
};

module.exports = async () => {
  const kc = loadKubeconfig();
  log('loaded kubeconfig');

  const installedCrdsRsps = await installCrds(kc);
  log(`installed ${installedCrdsRsps.length} crds`);

  const namespace = getNamespace(kc);
  log(`using namespace ${namespace}`);

  const clients = makeClients(kc, installedCrdsRsps, namespace);
  log(`created ${clients.size} clients`);

  const informers = makeInformers(kc, installedCrdsRsps, namespace, clients);
  log(`created ${informers.size} informers`);

  makeControllers(clients, informers);
  log(`created controllers`);

  startInformers(informers);
  log(`started informers`);
}
