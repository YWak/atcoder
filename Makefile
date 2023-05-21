.PHONY: init test submit save restore

init:
	oj download "${url}"
	cp .vscode/template.go main.go
	code ./main.go

test:
	go build -o a.out main.go
	oj test

submit:
	oj submit main.go

save:
