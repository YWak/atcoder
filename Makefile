HISTORY_DIR := ~/.cache/online-judge-tools
HISTORY_FILE := download-history.jsonl

OJ_HISTORY := $(HISTORY_DIR)/$(HISTORY_FILE)

.PHONY: clean gottani init test submit save restore

init:
	cp .vscode/template.go main.go
	oj download "${url}"
	code ./main.go

clean:
	rm -rf \
	_main.go \
	a.out \
	main.go \
	test \
	sandbox

init-offline:
	cp .vscode/template.go main.go
	mkdir -p test
	for i in `seq 5`; do touch "test/sample-$$i.in"; touch "test/sample-$$i.out"; done
	code ./main.go

gottani:
	mkdir -p sandbox
	cp -f go.mod go.sum sandbox
	gottani > sandbox/main.go
	cp -f sandbox/main.go _main.go

test: gottani
	go build -o ./a.out sandbox/main.go
	oj test
	rm -rf a.out

testadd:
	for i in `seq 100`; do [ ! -e "test/test-$$i.in" ] && touch "test/test-$$i.in" "test/test-$$i.out" && exit; done

submit: gottani
	oj submit _main.go
	rm -f _main.go

save:
	! [ -f main.go ] && echo 'main.go does not exist' && exit 1 || true
	$(eval target = submissions/`./bin/dirname.py`)
	mkdir -p $(target)
	rm -rf $(target)/test
	[ -d test ] && mv -f test $(target) || true
	mv -f main.go $(target)/main.go.out
	cp -f $(OJ_HISTORY) $(target)/
	git add .
	EDITOR="code --wait" git commit
restore:
	[ -z "$(dir)" ] && echo 'no dir' && exit 1 || true
	[ -f main.go ] && echo 'main.go exists. do make save' && exit 1 || true
	rm -rf ./test
	if [ -d $(dir)/test ]; then \
		cp -r $(dir)/test . ;\
		cp -f $(dir)/$(HISTORY_FILE) $(OJ_HISTORY) ;\
	else \
		oj download `./bin/url.py $(dir)`; \
	fi
	cp -f $(dir)/main.go.out ./main.go
	code ./main.go
