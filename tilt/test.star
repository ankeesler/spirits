load(
  './globals.star',
  'go_srcs',
)

_go_tests = ['test/api']

def test_all():
  local_resource(
    'go-test',
    'go test ./...',
    deps=go_srcs + _go_tests,
    allow_parallel=True,
    labels=['test'],
  )

  local_resource(
    'go-vet',
    'go vet ./...',
    deps=go_srcs + _go_tests,
    allow_parallel=True,
    labels=['test'],
  )
