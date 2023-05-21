HISTORY_DIR := ~/.cache/online-judge-tools
HISTORY_FILE := download-history.jsonl

OJ_HISTORY := $(HISTORY_DIR)/$(HISTORY_FILE)

.PHONY: init test submit save restore

init:
	oj download "${url}"
	@cp .vscode/template.go main.go
	@code ./main.go

test:
	@go build -o a.out main.go
	@oj test
	@rm a.out

submit:
	@gottani | oj submit /dev/stdin

save:
	@$(eval target = `./bin/dirname.py`)
	@mkdir -p submissions/$(target)
	@mv -f test submissions/$(target)
	@mv -f main.go submissions/$(target)/main.go.out
	@cp -f $(OJ_HISTORY) submissions/$(target)/

restore:
	@[ -z "$(dir)" ] && echo 'no dir' && exit 1 || true
	@[ -f main.go ] && echo 'main.go exists. do make save' && exit 1 || true
	@rm -rf ./test
	@if [ -d $(dir)/test ]; then \
		cp -r $(dir)/test . ;\
		cp -f $(dir)/$(HISTORY_FILE) $(OJ_HISTORY) ;\
	else \
		oj download `./bin/url.py $(dir)`; \
	fi
	@cp $(dir)/main.go.out ./main.go
	@code ./main.go
