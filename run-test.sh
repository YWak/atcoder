#!/bin/sh

if [ $# = 0 ] || [ ! -d "$1" ] || [ ! -f "$1/main.go" ]; then
    echo "usage: $0 target_dir" >&2
    exit 1
fi

cd "$1" || exit 1
EXE="./main"
TEMPFILE=$(mktemp)
go1.6 build -o $EXE main.go

trap 'rm -rf main $TEMPFILE' 0

for i in $(seq 5); do
    INPUT="./ex${i}.txt"
    OUTPUT="./ans${i}.txt"

    if [ ! -f "$INPUT" ] || [ ! -f "$OUTPUT" ] || [ ! -s "$INPUT" ]; then
        continue
    fi

    $EXE < "$INPUT" | diff --side-by-side - "$OUTPUT" > "$TEMPFILE"

    if [ $? = 0 ]; then
        echo "$i => OK" >&2
    else
        echo "$i => NG" >&2
        cat "$TEMPFILE" >&2
    fi
done
