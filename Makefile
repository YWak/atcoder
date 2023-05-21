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
	@gottani > _main.go 
	@oj submit _main.go
	@rm _main.go

save:
	@$(eval target = submissions/`./bin/dirname.py`)
	@mkdir -p $(target)
	@rm -rf $(target)/test
	@mv -f test $(target)
	@mv -f main.go $(target)/main.go.out
	@cp -f $(OJ_HISTORY) $(target)/

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
