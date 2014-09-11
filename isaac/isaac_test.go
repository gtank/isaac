package isaac

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
)

/*
 * This tests against the official randvect.txt, which outputs two rounds with no key (zeros)
 *
 * I am not sure about the secondary call to isaac() after randInit() - the demo code does it,
 * but not for any clear reason. My own code just calls isaac() directly from the rand()
 * function when it's run out of available numbers. Nevertheless, the extra isaac() call is
 * included here for compatibility with the test vector.
 */
func TestRand(t *testing.T) {
	vect, err := os.Open("randvect.txt")
	if err != nil {
		t.Error("could not find text vector file")
	}
	defer vect.Close()
	scanner := bufio.NewScanner(vect)
	scanner.Scan()

	var rng isaac
	rng.randInit(true)
	rng.isaac()

	var buf bytes.Buffer
	for i := 0; i < 2; i++ {
		for j := 0; j < 256; j++ {
			buf.WriteString(fmt.Sprintf("%.8x", rng.Rand()))
			if (j & 7) == 7 {
				var output = buf.String()
				if scanner.Text() == output {
					scanner.Scan()
					buf.Reset()
				} else {
					t.Errorf("output did not match test vector at line %d")
				}
			}
		}
	}
}
