#!/bin/sh

find . -type f \( -name "*.go" -o -name "*.yaml" -o -name "*.jet.html" \) | grep -v vendor | entr -r ./run.sh
