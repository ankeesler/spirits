load(
  './globals.star',
  'go_srcs',
  'spirits_server_resource',
)

_go_tests = [os.path.join('test', 'api')]
_hack_test_unit = [os.path.join('hack', 'test-unit.sh')]
_hack_test_integration = [os.path.join('hack', 'test-integration.sh')]

def test_all():
  local_resource(
    'test-unit',
    _hack_test_unit,
    deps=go_srcs + _hack_test_unit,
    auto_init=False,
    labels=['test'],
  )

  local_resource(
    'test-integration',
    _hack_test_integration,
    deps=go_srcs + _go_tests + _hack_test_integration,
    auto_init=False,
    trigger_mode=TRIGGER_MODE_MANUAL,
    resource_deps=[spirits_server_resource],
    labels=['test'],
  )