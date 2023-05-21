.PHONY: init test submit save restore

init:
	oj download "${url}"
	cp .vscode/template.go main.go
	code ./main.go

test:
	go build -o a.out main.go
	oj test
	rm a.out

submit:
	oj submit main.go

save:
	$(eval target = `date +%Y/%m/%d`/${dir})
	mkdir -p submissions/$(target)
	mv test submissions/$(target)
	mv main.go submissions/$(target)/main.go.out
	cp ~/.local/share/online-judge-tools/download-history.jsonl
