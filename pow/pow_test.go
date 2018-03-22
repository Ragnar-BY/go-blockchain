package pow

import (
	"testing"
)

type Test struct {
	header [32]byte
	nonce  [8]byte
	hash   [32]byte
}

var tests = []Test{
	{[32]byte{194, 121, 220, 122, 10, 137, 141, 24, 160, 173, 149, 254, 22, 143, 209, 173, 247, 153, 111, 249, 0, 246, 95, 173, 35, 112, 125, 171, 191, 56, 234, 200}, [8]byte{0, 0, 0, 0, 0, 0, 111, 2},
		[32]byte{0, 0, 151, 254, 10, 145, 143, 70, 193, 107, 2, 226, 191, 203, 241, 76, 228, 127, 123, 219, 135, 78, 90, 17, 177, 172, 25, 24, 230, 40, 47, 9},
	},
	{[32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, [8]byte{0, 0, 0, 0, 0, 0, 251, 233},
		[32]byte{0, 0, 179, 85, 240, 194, 189, 150, 133, 179, 17, 49, 34, 121, 243, 234, 153, 254, 174, 18, 22, 118, 248, 120, 215, 141, 217, 71, 105, 140, 110, 179},
	},
}

func TestRun(t *testing.T) {

	pow := NewProofOfWork()

	for i := 0; i < len(tests); i++ {
		n, h := pow.Run(tests[i].header)
		if n != tests[i].nonce {
			t.Errorf("Nonce:Expected %v, received %v", tests[i].nonce, n)
		}
		if h != tests[i].hash {
			t.Errorf("Hash:Expected %v, received %v", tests[i].hash, h)
		}
	}
}

func TestIsValid(t *testing.T) {
	pow := NewProofOfWork()

	for i := 0; i < len(tests); i++ {
		if pow.IsValid(tests[i].header, tests[i].nonce) == false {

			t.Errorf("received false,expected true")
		}
	}

	if pow.IsValid(tests[0].header, [8]byte{0, 0, 0, 0, 0, 0, 0, 0}) == true {
		t.Errorf("received true, expected false")
	}
}
