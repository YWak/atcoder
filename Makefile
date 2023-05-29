HISTORY_DIR := ~/.cache/online-judge-tools
HISTORY_FILE := download-history.jsonl

OJ_HISTORY := $(HISTORY_DIR)/$(HISTORY_FILE)

.PHONY: clean init test submit save restore

clean:
	rm -rf \
	_main.go \
	a.out \
	main.go \
	test \
	sandbox

init:
	oj download "${url}"
	cp .vscode/template.go main.go
	code ./main.go

test:
	mkdir -p sandbox
	cp -f go.mod go.sum sandbox
	gottani > sandbox/main.go
	go build -o ./a.out sandbox/main.go
	oj test
	rm -rf a.out

submit:
	gottani > _main.go 
	oj submit _main.go
	rm -rf _main.go

save:
	! [ -f main.go ] && echo 'main.go does not exist' && exit 1 || true
	$(eval target = submissions/`./bin/dirname.py`)
	mkdir -p $(target)
	rm -rf $(target)/test
	mv -f test $(target)
	mv -f main.go $(target)/main.go.out
	cp -f $(OJ_HISTORY) $(target)/
	git add .
	git commit

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
