#!/bin/sh

pkill Khmer2Romance && echo "Sent kill"
rm -f ./Khmer2Romance && echo "Remove old binary"

echo "Build & Run..."
go build && ./Khmer2Romance
