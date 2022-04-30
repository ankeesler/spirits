docker_build('spirits-controller-manager', '.', dockerfile='Dockerfile')
k8s_yaml('config/deployment.yaml')
k8s_resource('spirits-controller-manager', port_forwards=8000)

load('ext://tests/golang', 'test_go')
test_go('unit-tests', '.', '.', recursive=True)
