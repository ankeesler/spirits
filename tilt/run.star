load(
  './globals.star',
  'go_srcs',
  'hack_generate',
  'spirits_manager_resource',
)

_controller_gen_pkg='sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0'

def run_all():
  local_resource(
    'crds',
    [hack_generate, 'generate_crds'],
    deps=[hack_generate, os.path.join('pkg', 'apis')],
    auto_init=False,
    labels=['run'],
  )

  local_resource(
    'rbac',
    [hack_generate, 'generate_rbac'],
    deps=[hack_generate, os.path.join('pkg', 'controller')],
    auto_init=False,
    labels=['run'],
  )

  k8s_yaml(
    listdir('config', recursive=True),
  )

  k8s_resource(
    spirits_manager_resource,
    labels=['run'],
  )
