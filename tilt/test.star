load(
  './globals.star',
  'go_srcs',
  'spirits_manager_resource',
)

_go_tests = ['test/api']

def test_all():
  local_resource(
    'go-vet',
    'go vet ./...',
    deps=go_srcs + _go_tests,
    labels=['test'],
  )

  local_resource(
    'go-test-unit',
    'go test -v ./...',
    deps=go_srcs,
    labels=['test'],
  )

  local_resource(
    'go-test-integration',
    'go test -count 1 -v ./test/...',
    deps=go_srcs + _go_tests,
    trigger_mode=TRIGGER_MODE_MANUAL,
    resource_deps=[spirits_manager_resource],
    env={
      'SPIRITS_TEST_INTEGRATION': '',
    },
    labels=['test'],
  )