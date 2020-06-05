#!/bin/sh

if [ $# = 0 ] || [ ! -d "$1" ] || [ ! -f "$1/main.go" ]; then
    echo "usage: $0 target_dir" >&2
    exit 1
fi

cd "$1" || exit 1
EXE="./main"
TEMPFILE1=$(mktemp)
TEMPFILE2=$(mktemp)
go1.6 build -o $EXE main.go

trap 'rm -rf main $TEMPFILE1 $TEMPFILE2 $EXE' 0

for i in $(seq 5); do
    INPUT="./ex${i}.txt"
    OUTPUT="./ans${i}.txt"

    if [ ! -f "$INPUT" ] || [ ! -f "$OUTPUT" ] || [ ! -s "$INPUT" ]; then
        continue
    fi

    time -f '(%Us %MKB)' $EXE < "$INPUT" 2> "$TEMPFILE2" | diff --side-by-side - "$OUTPUT" > "$TEMPFILE1"

    if [ $? = 0 ]; then
        echo "$i => OK $(cat $TEMPFILE2)" >&2
    else
        echo "$i => NG $(cat $TEMPFILE2)" >&2
        cat "$TEMPFILE1" >&2
    fi
done
