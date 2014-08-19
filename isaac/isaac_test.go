package isaac

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	Randinit(true)
	for i := 0; i < 2; i++ {
		Isaac()
		var randrsl = Randrsl()
		for j := 0; j < 256; j++ {
			fmt.Printf("%.8x", randrsl[j])
			if (j & 7) == 7 {
				fmt.Printf("\n")
			}
		}
	}
}
