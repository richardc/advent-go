#!/bin/bash
set -e

# Small helper to bootstrap a new days solution - errors early
# to avoid stomping anything you already did
DAY="${1:-$(date +%d)}"
PADDAY="$(printf %02d "$DAY")"
YEAR="${2:-$(date +%Y)}"

echo "Bootstrapping ${YEAR} ${DAY}"

mkdir "${YEAR}/day${PADDAY}"

# https://github.com/GreenLightning/advent-of-code-downloader
# go install github.com/GreenLightning/advent-of-code-downloader/aocdl@latest
aocdl -year "${YEAR}" -day "${DAY}" -output "${YEAR}/day${PADDAY}/input.txt"

cat > "${YEAR}/day${PADDAY}/day${PADDAY}.go" << HERE
package day${PADDAY}

import (
    _ "embed"

    "github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string
func init() {
    runner.Register(
        runner.Solution{
            Year: ${YEAR},
            Day: ${DAY},
            Part1: func(any) any { return solve(puzzle) },
        },
    )
}

func solve(puzzle string) int {
    return 0
}
HERE

go fmt "./${YEAR}/day${PADDAY}/..."
