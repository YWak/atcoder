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
	target := `date +%Y/%m/%d`${dir}
	mkdir -p submissions/$(target)
	mv main.go test submissions/$(target)
