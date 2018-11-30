package utils

import (
	"bytes"
	"encoding/hex"

	"math"
	"testing"
)

func TestUintToHex(t *testing.T) {

	type Test struct {
		in  uint64
		out []byte
	}

	tests := []Test{
		{0, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{12345, []byte{0, 0, 0, 0, 0, 0, 48, 57}},
		{1234567890, []byte{0, 0, 0, 0, 73, 150, 2, 210}},
		{72340172838076673, []byte{1, 1, 1, 1, 1, 1, 1, 1}},
		{math.MaxUint64, []byte{255, 255, 255, 255, 255, 255, 255, 255}},
	}
	for i := 0; i < len(tests); i++ {
		if !bytes.Equal(UintToHex(tests[i].in), tests[i].out) {
			t.Errorf("Expected %v, received %v", tests[i].out, UintToHex(tests[i].in))
		}
	}
}

func TestHash(t *testing.T) {

	type Test struct {
		in  []byte
		out string
	}
	tests := []Test{
		{[]byte{}, "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a"},
		// 'abc'
		{[]byte{97, 98, 99}, "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532"},
		//32 symbols 'a'
		{[]byte{97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97}, "2eee42b299cb44bb3d5dd0a02210ce29debc2110b7c4644ff38c53084940be21"},
		// '1234'
		{[]byte{49, 50, 51, 52}, "1d6442ddcfd9db1ff81df77cbefcd5afcc8c7ca952ab3101ede17a84b866d3f3"},
	}

	for i := 0; i < len(tests); i++ {

		dst := make([]byte, hex.DecodedLen(len(tests[i].out)))
		n, _ := hex.Decode(dst, []byte(tests[i].out))

		h, err := Hash(tests[i].in)
		if err != nil {
			t.Errorf("Can not count hash %v", err)
		}
		if !bytes.Equal(h[:], dst[:n]) {
			t.Errorf("Expected %v, received %s", tests[i].out, h[:])
		}
	}

}
