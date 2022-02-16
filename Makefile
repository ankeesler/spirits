SPIRITS ?= .bin/spirits
BTOOL ?= .bin/btool

GENERATED_FILES =         \
  source/data/locations.h \
  source/data/locations.c

all: $(SPIRITS)

.PHONY: registry
registry: $(BTOOL)
	test -n "$(shell docker ps -q -f name=btoolregistry)" || docker run --name btoolregistry -d --rm -it -p 80:80 "ankeesler/btoolregistry"

clean: .bin/btool registry
	$< -clean -root source -target main -registry http://localhost:80
	docker stop btoolregistry
	rm -rf $(<D)

$(BTOOL):
	mkdir -p $(@D)
	curl -L -o $(BTOOL).gz https://github.com/ankeesler/btool/releases/download/0.8/btool-0.8-macos-x86_64.gz
	gunzip $(BTOOL).gz
	chmod +x $@

source/main: $(BTOOL) registry # $(GENERATED_FILES)
	$< -root source -target main -registry http://localhost:80

$(SPIRITS): source/main
	mkdir -p $(@D)
	cp $< $@

source/data/%.h: script/%.awk data/%.txt
	awk -F "::" -v source=$< -v output_type=h -f $^ >$@

source/data/%.c: script/%.awk data/%.txt
	awk -F "::" -v source=$< -v output_type=c -f $^ >$@
