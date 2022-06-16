package fp

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFMap(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}

	ys := FMap(xs, func(x int) string {
		return strconv.Itoa(x * x)
	})

	fmt.Println(ys)
}
