#!/bin/sh

pkill Khmer2Romanize && echo "Sent kill"
rm -f ./Khmer2Romanize && echo "Remove old binary"

echo "Build & Run..."
go build && ./Khmer2Romanize
