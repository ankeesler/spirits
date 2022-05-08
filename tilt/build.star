load('ext://restart_process', 'docker_build_with_restart')

load(
  './globals.star',
  'go_srcs',
  'hack_generate',
)

def _get_non_generated_files(files):
  non_generated_files = []
  for file in files:
    if "zz_generated" not in file:
      non_generated_files += file
  return non_generated_files

_go_external_api_src=_get_non_generated_files(listdir(os.path.join('pkg', 'apis')))
_go_internal_api_src=_get_non_generated_files(listdir(os.path.join('internal', 'apis')))

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
    deps=[hack_generate] + _go_external_api_src,
    ignore=['**zz_generated**'],
    auto_init=False,
    labels=['build'],
  )

  local_resource(
    'go-internal-api',
    [hack_generate, 'generate_internal_groups'],
    deps=[hack_generate] + _go_external_api_src + _go_internal_api_src,
    ignore=['**zz_generated**'],
    auto_init=False,
    labels=['build'],
  )

  local_resource(
    'spirits-server-compile',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server-linux-amd64 ./cmd/server',
    deps=go_srcs,
    labels=['build'],
  )

  docker_build_with_restart(
    'ankeesler/spirits-server',
    '.',
    dockerfile='tilt/Dockerfile',
    entrypoint=['/server', '-v=1', '-logtostderr'],
    live_update=[
      sync('server-linux-amd64', '/server'),
    ],
    only=[
      'server-linux-amd64',
    ],
  )
