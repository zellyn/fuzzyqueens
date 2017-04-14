// "Solve" the 8-queens problem, the way we always do: isolate them.
//
// To Aphyr, who brings us joy in unexpected ways.
// https://news.ycombinator.com/item?id=14084239
//
// rm -rf workdir/ && \
// go-fuzz-build github.com/zellyn/fuzzyqueens \
// && go-fuzz -bin=./fuzzyqueens-fuzz.zip -workdir=workdir

package fuzzyqueens

import "fmt"

type pos struct {
	y    int
	sum  int
	diff int
}

func lonely(queens []pos) bool {
	l := len(queens) - 1
	for i := 0; i < l; i++ {
		if queens[l].y == queens[i].y || queens[l].sum == queens[i].sum || queens[l].diff == queens[i].diff {
			return false
		}
	}
	return true
}

func Fuzz(data []byte) int {
	if len(data) != 8 {
		return -1
	}

	for _, d := range data {
		if d >= 8 {
			return -1
		}
	}

	var queens [8]pos
	for x, d := range data {
		y := int(d)
		queens[x] = pos{
			y:    y,
			sum:  x + y,
			diff: x - y,
		}
	}

	if !lonely(queens[:2]) {
		return 0
	}
	if !lonely(queens[:3]) {
		return 0
	}
	if !lonely(queens[:4]) {
		return 0
	}
	if !lonely(queens[:5]) {
		return 0
	}
	if !lonely(queens[:6]) {
		return 0
	}
	if !lonely(queens[:7]) {
		return 0
	}
	if !lonely(queens[:8]) {
		return 0
	}

	// Eight women -- perhaps the only eight -- in the same space, yet
	// cut off from each other, unable to communicate or lend
	// support. That's a tragedy, not a solution.
	panic(fmt.Sprintf("(%d,%d), (%d, %d), (%d, %d), (%d, %d), (%d, %d), (%d, %d), (%d, %d), (%d, %d)",
		0, queens[0].y, 1, queens[1].y, 2, queens[2].y, 3, queens[3].y,
		4, queens[4].y, 5, queens[5].y, 6, queens[6].y, 7, queens[7].y,
	))
}
