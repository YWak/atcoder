#!/bin/sh -e

if [ "$1" = "" ]; then
    echo "$0 dirname" >&1
fi

# cd $(dirname $(readlin -f $0)) || exit

target=$(dirname "$(pwd)/$1/a")
base=$(dirname $(readlink -f $0))
mkdir -p "$target"
! [ -f "$target/main.go" ] && cat "$base/.vscode/template.go" > "$target/main.go"

touch "$target/ex1.txt"
touch "$target/ex2.txt"
touch "$target/ex3.txt"
touch "$target/ex4.txt"
touch "$target/ex5.txt"
touch "$target/ans1.txt"
touch "$target/ans2.txt"
touch "$target/ans3.txt"
touch "$target/ans4.txt"
touch "$target/ans5.txt"

code -r $target/main.go
