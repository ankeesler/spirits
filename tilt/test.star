load(
  './globals.star',
  'go_srcs',
)

def test_all():
  local_resource(
    'go-test',
    'go test ./...',
    deps=go_srcs,
    allow_parallel=True,
    labels=['test'],
  )

  local_resource(
    'go-vet',
    'go vet ./...',
    deps=go_srcs,
    allow_parallel=True,
    labels=['test'],
  )
