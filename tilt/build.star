load(
  './globals.star',
  'go_srcs',
  'hack_generate',
)

_go_external_api_src=os.path.join('pkg', 'apis')
_go_internal_api_src=os.path.join('internal', 'apis')

def _join(l):
  if len(l) == 0:
    return ''

  s = l[0]
  for e in l[1:]:
    s += ' ' + e

  return s

def build_all():
  local_resource(
    'go-api',
    [hack_generate, 'generate_groups'],
    deps=[hack_generate, _go_external_api_src],
    ignore=['**zz_generated**'],
    auto_init=False,
    allow_parallel=True,
    labels=['build'],
  )

  local_resource(
    'go-internal-api',
    [hack_generate, 'generate_internal_groups'],
    deps=[hack_generate, _go_external_api_src, _go_internal_api_src],
    ignore=['**zz_generated**'],
    auto_init=False,
    allow_parallel=True,
    labels=['build'],
  )

  local_resource(
    'spirits-manager-compile',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o manager .',
    deps=go_srcs,
    labels=['build'],
  )

  docker_build(
    'spirits-manager',
    '.',
    dockerfile='tilt/Dockerfile',
    live_update=[
      sync('manager', '/manager'),
    ],
    only=[
      'manager',
    ],
  )
