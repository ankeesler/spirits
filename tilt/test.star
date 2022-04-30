load('./globals.star', 'go_srcs')

def test_all():
  local_resource(
    'go-test',
    'go test ./...',
    deps=go_srcs,
    resource_deps=['external-api'],
    auto_init=False,
    allow_parallel=True,
    labels=['test'],
  )

  local_resource(
    'go-vet',
    'go vet ./...',
    deps=go_srcs,
    resource_deps=['external-api'],
    auto_init=False,
    allow_parallel=True,
    labels=['test'],
  )
