run() {
  echo "running '$@'"
  $@
  echo
}

run curl localhost:12345

run curl localhost:12345/rooms

run curl localhost:12345/rooms/whatever

run curl -X POST localhost:12345/rooms --data '{"name":"whatever"}'

run curl localhost:12345/rooms/whatever

run curl localhost:12345/rooms/whatever/manifests

run curl localhost:12345/rooms/whatever/manifests/my-manifest

run curl -X POST localhost:12345/rooms/whatever/manifests --data @internal/test/testdata/good-manifest.json

run curl localhost:12345/rooms/whatever/manifests

run curl localhost:12345/rooms/whatever/manifests/my-manifest
