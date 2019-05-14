package cash

import (
	"fmt"
)

var (
	// DECODEMAP is the deconding map for base32 BCH encoding
	DECODEMAP map[byte]int
	ENCODEMAP map[int]byte
)

func init() {
	DECODEMAP = make(map[byte]int)
	DECODEMAP[byte('q')] = 0
	DECODEMAP[byte('p')] = 1
	DECODEMAP[byte('z')] = 2
	DECODEMAP[byte('r')] = 3
	DECODEMAP[byte('y')] = 4
	DECODEMAP[byte('9')] = 5
	DECODEMAP[byte('x')] = 6
	DECODEMAP[byte('8')] = 7
	DECODEMAP[byte('g')] = 8
	DECODEMAP[byte('f')] = 9
	DECODEMAP[byte('2')] = 10
	DECODEMAP[byte('t')] = 11
	DECODEMAP[byte('v')] = 12
	DECODEMAP[byte('d')] = 13
	DECODEMAP[byte('w')] = 14
	DECODEMAP[byte('0')] = 15
	DECODEMAP[byte('s')] = 16
	DECODEMAP[byte('3')] = 17
	DECODEMAP[byte('j')] = 18
	DECODEMAP[byte('n')] = 19
	DECODEMAP[byte('5')] = 20
	DECODEMAP[byte('4')] = 21
	DECODEMAP[byte('k')] = 22
	DECODEMAP[byte('h')] = 23
	DECODEMAP[byte('c')] = 24
	DECODEMAP[byte('e')] = 25
	DECODEMAP[byte('6')] = 26
	DECODEMAP[byte('m')] = 27
	DECODEMAP[byte('u')] = 28
	DECODEMAP[byte('a')] = 29
	DECODEMAP[byte('7')] = 30
	DECODEMAP[byte('l')] = 31

	ENCODEMAP = make(map[int]byte)
	for k, v := range DECODEMAP {
		ENCODEMAP[v] = k
	}
}

// Base32Decode return a cashaddress base32 the decoded array of byte
func Base32Decode(bch32 string) ([]byte, error) {
	l := len(bch32)
	decoded := make([]byte, l, l)
	for i := 0; i < l; i++ {
		c := bch32[i]
		v, ok := DECODEMAP[c]
		if !ok {
			return nil, fmt.Errorf("char not allowed in the address payload %c", c)
		}
		decoded[i] = byte(v)
	}
	return decoded, nil
}

// Base32Encode return the given byte array encoded to a cashaddress base32 string
func Base32Encode(bytes []byte) (string, error) {
	l := len(bytes)
	var encoded string
	for i := 0; i < l; i++ {
		n := int(bytes[i])
		c, ok := ENCODEMAP[n]
		if !ok {
			return "", fmt.Errorf("value not allowed for base32 encoding %d", n)
		}
		encoded += string(c)
	}
	return encoded, nil
}
