load('./globals.star', 'go_srcs')

def run_all():
  local_resource(
    'crds',
    [
      'go',
      'run',
      'sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0',
      'paths=./pkg/apis',
      'crd',
      'output:crd:artifacts:config=./config/zz_generated_crd'
    ],
    deps=[os.path.join('pkg', 'controller')],
    auto_init=False,
    allow_parallel=True,
    labels=['build'],
  )

  local_resource(
    'rbac',
    [
      'go',
      'run',
      'sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0',
      'paths=./pkg/controller',
      '+rbac:roleName=zz-generated-spirits-manager',
      'output:rbac:dir=./config',
    ],
    deps=[os.path.join('pkg', 'controller')],
    auto_init=False,
    allow_parallel=True,
    labels=['build'],
  )

  k8s_yaml(
    listdir('config', recursive=True),
  )

  k8s_resource(
    'spirits-manager',
    labels=['run'],
  )
