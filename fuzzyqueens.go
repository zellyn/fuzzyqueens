// "Solve" the 8-queens problem, the way we always do: isolate them.
//
// To Aphyr, who brings us joy in unexpected ways.
// https://news.ycombinator.com/item?id=14084239
//
// rm -rf workdir/ && \
// go-fuzz-build github.com/zellyn/fuzzyqueens \
// && go-fuzz -bin=./fuzzyqueens-fuzz.zip -workdir=workdir
//
//
// $ rm -rf workdir/ && go-fuzz-build github.com/zellyn/fuzzyqueens && go-fuzz -bin=./fuzzyqueens-fuzz.zip -workdir=workdir
// 2017/04/14 09:52:46 slaves: 8, corpus: 30 (1s ago), crashers: 2, restarts: 1/0, execs: 0 (0/sec), cover: 0, uptime: 3s
// 2017/04/14 09:52:49 slaves: 8, corpus: 30 (4s ago), crashers: 2, restarts: 1/6014, execs: 30072 (5005/sec), cover: 143, uptime: 6s
// ^C2017/04/14 09:52:49 shutting down...
// $ for i in workdir/crashers/*.output; do head -1 $i; done
// panic: (0,5), (1, 2), (2, 4), (3, 7), (4, 0), (5, 3), (6, 1), (7, 6)
// panic: (0,5), (1, 2), (2, 4), (3, 6), (4, 0), (5, 3), (6, 1), (7, 7)

package fuzzyqueens

import "fmt"

type pos struct {
	y    int
	sum  int
	diff int
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
		queens[x] = pos{y: y, sum: x + y, diff: x - y}
	}

	if connected(queens) {
		return 0
	}

	// Eight women -- perhaps the only eight -- occupy the same space,
	// yet are cut off from each other, unable to communicate or lend
	// support. Panic.
	panic(fmt.Sprintf("(%d,%d), (%d, %d), (%d, %d), (%d, %d), (%d, %d), (%d, %d), (%d, %d), (%d, %d)",
		0, queens[0].y, 1, queens[1].y, 2, queens[2].y, 3, queens[3].y,
		4, queens[4].y, 5, queens[5].y, 6, queens[6].y, 7, queens[7].y,
	))
}

func connected(queens [8]pos) bool {
	// Alas, no kind of un-unrolled loops work as well.

	if queens[1].y == queens[0].y || queens[1].sum == queens[0].sum || queens[1].diff == queens[0].diff {
		return true
	}

	if queens[2].y == queens[0].y || queens[2].sum == queens[0].sum || queens[2].diff == queens[0].diff {
		return true
	}
	if queens[2].y == queens[1].y || queens[2].sum == queens[1].sum || queens[2].diff == queens[1].diff {
		return true
	}

	if queens[3].y == queens[0].y || queens[3].sum == queens[0].sum || queens[3].diff == queens[0].diff {
		return true
	}
	if queens[3].y == queens[1].y || queens[3].sum == queens[1].sum || queens[3].diff == queens[1].diff {
		return true
	}
	if queens[3].y == queens[2].y || queens[3].sum == queens[2].sum || queens[3].diff == queens[2].diff {
		return true
	}

	if queens[4].y == queens[0].y || queens[4].sum == queens[0].sum || queens[4].diff == queens[0].diff {
		return true
	}
	if queens[4].y == queens[1].y || queens[4].sum == queens[1].sum || queens[4].diff == queens[1].diff {
		return true
	}
	if queens[4].y == queens[2].y || queens[4].sum == queens[2].sum || queens[4].diff == queens[2].diff {
		return true
	}
	if queens[4].y == queens[3].y || queens[4].sum == queens[3].sum || queens[4].diff == queens[3].diff {
		return true
	}

	if queens[5].y == queens[0].y || queens[5].sum == queens[0].sum || queens[5].diff == queens[0].diff {
		return true
	}
	if queens[5].y == queens[1].y || queens[5].sum == queens[1].sum || queens[5].diff == queens[1].diff {
		return true
	}
	if queens[5].y == queens[2].y || queens[5].sum == queens[2].sum || queens[5].diff == queens[2].diff {
		return true
	}
	if queens[5].y == queens[3].y || queens[5].sum == queens[3].sum || queens[5].diff == queens[3].diff {
		return true
	}
	if queens[5].y == queens[4].y || queens[5].sum == queens[4].sum || queens[5].diff == queens[4].diff {
		return true
	}

	if queens[6].y == queens[0].y || queens[6].sum == queens[0].sum || queens[6].diff == queens[0].diff {
		return true
	}
	if queens[6].y == queens[1].y || queens[6].sum == queens[1].sum || queens[6].diff == queens[1].diff {
		return true
	}
	if queens[6].y == queens[2].y || queens[6].sum == queens[2].sum || queens[6].diff == queens[2].diff {
		return true
	}
	if queens[6].y == queens[3].y || queens[6].sum == queens[3].sum || queens[6].diff == queens[3].diff {
		return true
	}
	if queens[6].y == queens[4].y || queens[6].sum == queens[4].sum || queens[6].diff == queens[4].diff {
		return true
	}
	if queens[6].y == queens[5].y || queens[6].sum == queens[5].sum || queens[6].diff == queens[5].diff {
		return true
	}

	if queens[7].y == queens[0].y || queens[7].sum == queens[0].sum || queens[7].diff == queens[0].diff {
		return true
	}
	if queens[7].y == queens[1].y || queens[7].sum == queens[1].sum || queens[7].diff == queens[1].diff {
		return true
	}
	if queens[7].y == queens[2].y || queens[7].sum == queens[2].sum || queens[7].diff == queens[2].diff {
		return true
	}
	if queens[7].y == queens[3].y || queens[7].sum == queens[3].sum || queens[7].diff == queens[3].diff {
		return true
	}
	if queens[7].y == queens[4].y || queens[7].sum == queens[4].sum || queens[7].diff == queens[4].diff {
		return true
	}
	if queens[7].y == queens[5].y || queens[7].sum == queens[5].sum || queens[7].diff == queens[5].diff {
		return true
	}
	if queens[7].y == queens[6].y || queens[7].sum == queens[6].sum || queens[7].diff == queens[6].diff {
		return true
	}

	return false
}
