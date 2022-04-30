load('./globals.star', 'go_srcs')

go_external_api_src=os.path.join('pkg', 'apis')
go_internal_api_src=os.path.join('internal', 'apis')

spirits_go_pkg_bash='$(awk \'/^module / {print $2}\' <go.mod)'
codegen_go_pkg_bash='$(awk \'/k8s.io.code-generator/ {gsub(" ", "@"); gsub("\t", ""); print}\' <go.mod)'

def _join(l):
  s = ''
  for e in l:
    s += ' ' + e
  return s

def build_all():
  local_resource(
    'external-api',
    [
      'bash',
      '-c',
      _join([
        'eval',
        '"$(go env)"',
        '&&',
        os.path.join('${GOMODCACHE}', codegen_go_pkg_bash, 'generate-internal-groups.sh'),
        'deepcopy,defaulter,conversion',
        os.path.join(spirits_go_pkg_bash, go_internal_api_src),
        os.path.join(spirits_go_pkg_bash, go_internal_api_src),
        os.path.join(spirits_go_pkg_bash, go_external_api_src),
        'spirits:v1alpha1',
        '--go-header-file', 'hack/boilerplate.go.txt',
        '-v', '1',
      ]),
    ],
    deps=[go_external_api_src],
    auto_init=False,
    allow_parallel=True,
    labels=['build'],
  )

  local_resource(
    'internal-api',
    [
      'bash',
      '-c',
      _join([
        'eval',
        '"$(go env)"',
        '&&',
        os.path.join('${GOMODCACHE}', codegen_go_pkg_bash, 'generate-internal-groups.sh'),
        'deepcopy,defaulter,conversion',
        os.path.join(spirits_go_pkg_bash, go_internal_api_src),
        os.path.join(spirits_go_pkg_bash, go_internal_api_src),
        os.path.join(spirits_go_pkg_bash, go_external_api_src),
        'spirits:v1alpha1',
        '--go-header-file', 'hack/boilerplate.go.txt',
        '-v', '1',
      ]),
    ],
    deps=[go_internal_api_src],
    auto_init=False,
    allow_parallel=True,
    labels=['build'],
  )

  docker_build(
    'spirits-manager',
    '.',
    only=go_srcs,
  )
