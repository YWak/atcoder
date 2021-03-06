#!/bin/sh

if [ "$1" = "" ]; then
    echo "$0 dirname" >&1
fi

mkdir -p "$1"
! [ -f "$1/main.go" ] && cat ./.vscode/template.go > "$1/main.go"

touch "$1/ex1.txt"
touch "$1/ex2.txt"
touch "$1/ex3.txt"
touch "$1/ex4.txt"
touch "$1/ex5.txt"
touch "$1/ans1.txt"
touch "$1/ans2.txt"
touch "$1/ans3.txt"
touch "$1/ans4.txt"
touch "$1/ans5.txt"

code -r $1/main.go
