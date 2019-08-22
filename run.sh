#!/bin/sh

pkill Khmer2Romanized && echo "Sent kill"
rm -f ./Khmer2Romanized && echo "Remove old binary"

echo "Build & Run..."
go build && ./Khmer2Romanized
