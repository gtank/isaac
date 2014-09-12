package isaac

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
)

/* This tests against the official randvect.txt, which outputs two rounds with no key (zeros) */
func TestZeros(t *testing.T) {
	vect, err := os.Open("randvect.txt")
	if err != nil {
		t.Error("could not find text vector file")
	}
	defer vect.Close()
	scanner := bufio.NewScanner(vect)
	scanner.Scan()

	var rng ISAAC
	rng.randInit(true)

	var buf bytes.Buffer
	for i := 0; i < 2; i++ {
		rng.isaac()
		for j := 0; j < 256; j++ {
			buf.WriteString(fmt.Sprintf("%.8x", rng.randrsl[j]))
			if (j & 7) == 7 {
				var output = buf.String()
				if scanner.Text() == output {
					scanner.Scan()
					buf.Reset()
				} else {
					fmt.Printf("o: " + output + "\n" + "v: " + scanner.Text() + "\n")
					t.Fail()
					return
				}
			}
		}
	}
}

/* This tests against the randrsl state of randtest.c after randInit()
 * to make sure that my seeding procedure works the same as the reference */
func TestSeed(t *testing.T) {
	keyVector, _ := os.Open("keytest.txt")
	defer keyVector.Close()

	keyScan := bufio.NewScanner(keyVector)
	keyScan.Scan()

	var rng ISAAC
	var key = "This is <i>not</i> the right mytext."
	rng.Seed(key)

	var buf bytes.Buffer
	for i := 0; i < 256; i++ {
		buf.WriteString(fmt.Sprintf("%.8x ", rng.randrsl[i]))
		if (i & 7) == 7 {
			var output = buf.String()
			if keyScan.Text() == output {
				keyScan.Scan()
				buf.Reset()
			} else {
				fmt.Printf("index: %d\n", i)
				fmt.Printf("o: " + output + "\n" + "v: " + keyScan.Text() + "\n")
				t.Fail()
				return
			}
		}
	}
}

/* This tests output against the official randseed.txt */
func TestRand(t *testing.T) {
	testVector, _ := os.Open("randseed.txt")
	defer testVector.Close()

	outScan := bufio.NewScanner(testVector)
	outScan.Scan()

	var rng ISAAC
	var key = "This is <i>not</i> the right mytext."
	rng.Seed(key)

	var buf bytes.Buffer
	k := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 256; j++ {
			buf.WriteString(fmt.Sprintf("%.8x ", rng.Rand()))
			k++
			if k == 8 {
				k = 0
				if outScan.Text() == buf.String() {
					outScan.Scan()
					buf.Reset()
				} else {
					fmt.Printf("o: " + outScan.Text() + "\n" + "v: " + buf.String() + "\n")
					t.Fail()
					return
				}
			}
		}
	}
}

func TestStream(t *testing.T) {
	var key = "This is <i>not</i> the right mytext."
	var plaintext = []byte("Hello, world")
	var ciphertext = make([]byte, len(plaintext))
	var decrypted = make([]byte, len(plaintext))

	enc := NewISAACStream(key)
	enc.XORKeyStream(ciphertext, plaintext)

	dec := NewISAACStream(key)
	dec.XORKeyStream(decrypted, ciphertext)

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Plaintext not equal to decrypted ciphertext")
	}

}
