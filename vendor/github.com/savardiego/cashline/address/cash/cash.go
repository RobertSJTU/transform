package cash

import (
	"fmt"
	"github.com/savardiego/cashline/address/keys"
	"math"
)

// AddressTypeP2KH value for the Pay-To-Key-Hash type address
const AddressTypeP2KH int8 = 0

// FromPrivKey derivates the cashaddress from a private key, in compressed or uncompressed format
func FromPrivKey(privKey []byte, compressed bool) (string, error) {
	publicKeyBytes := keys.Public(privKey, compressed)
	withprefix, err := FromPubKey(publicKeyBytes)
	return withprefix, err
}

// FromWIF derivates a legacy address (version 1, the oldest) from a base58 encoded WIF private key, compressed/uncompressed depending on the WIF format.
func FromWIF(privKeyWIF string) (string, error) {
	decodedPrivKey, compressed, err := keys.PrivateFromWIF(privKeyWIF)
	if err != nil {
		return "", fmt.Errorf("cannot decode private key from base58 string: %v due to %v", privKeyWIF, err)
	}
	publicKey := keys.Public(decodedPrivKey, compressed)
	withprefix, err := FromPubKey(publicKey)
	return withprefix, nil
}

//FromPubKey returns a P2KH (ripemd160) mainnet (prefix:bitcoincash) bchaddress (withprefix, without prefix)
func FromPubKey(pubKey []byte) (string, error) {
	hashed := keys.Hashed(pubKey)
	prefix := "bitcoincash"
	withPrefix, _, err := addressFromHash(prefix, AddressTypeP2KH, hashed)
	return withPrefix, err
}

// calculate the address given the prefix, the hash and the address type
func addressFromHash(prefix string, addrType int8, hash []byte) (string, string, error) {
	hashSize, err := getHashSize(hash) //ripemd160 is 160 bit -> 0
	if err != nil {
		return "", "", fmt.Errorf("cannot get hash size because %v", err)
	}
	versionByte := byte(addrType + hashSize)
	prefixBytes := fullPrefixTo5Bit(prefix)
	data := append([]byte{versionByte}, hash...)
	data5bit, err := convert(data, 8, 5, false)
	if err != nil {
		return "", "", fmt.Errorf("cannot convert data to 5 bit due to %v", err)
	}
	checksumData := append(append(prefixBytes, data5bit...), []byte{0, 0, 0, 0, 0, 0, 0, 0}...)
	checksum5bit := getChecksum(checksumData)
	if err != nil {
		return "", "", fmt.Errorf("cannot convert checksum to 5 bit due to %v", err)
	}
	addressPayload := append(data5bit, checksum5bit...)
	if err != nil {
		return "", "", fmt.Errorf("cannot encode to Base32 due to %v", err)
	}
	encodedAddress, err := Base32Encode(addressPayload)
	return prefix + ":" + encodedAddress, encodedAddress, nil
}

// PolyMod calculates 40 bit checksum
// Reference: https://github.com/bitcoincashorg/bitcoincash.org/blob/master/spec/cashaddr.md
// Credits to https://github.com/bitcoincashjs/cashaddrjs/ and https://github.com/cpacia/bchutil/blob/master/cashaddr.go
func polyMod(v []byte) uint64 {
	var c uint64 = 1
	for _, d := range v {
		c0 := byte(c >> 35)
		c = ((c & 0x07ffffffff) << 5) ^ uint64(d)
		if c0&0x01 > 0 {
			c ^= 0x98f2bc8e61
		}
		if c0&0x02 > 0 {
			c ^= 0x79b76d99e2
		}
		if c0&0x04 > 0 {
			c ^= 0xf33e5fb3c4
		}
		if c0&0x08 > 0 {
			c ^= 0xae2eabe2a8
		}
		if c0&0x10 > 0 {
			c ^= 0x1e4f43e470
		}
	}
	return c ^ 1
}

// getChcksum calculates the byte array of the checksum
func getChecksum(checksumData []byte) []byte {
	mod := polyMod(checksumData)
	check := make([]byte, 8)
	for i := 0; i < 8; i++ {
		// Convert the 5-bit groups in mod to checksum values.
		check[i] = byte((mod >> uint(5*(7-i))) & 0x1f)
	}
	return check
}

// fullPrefixTo5Bit returns an array of byte with with the lower 5 bit of every prefix char plus a 0 for the colon separator
func fullPrefixTo5Bit(prefix string) []byte {
	ret := make([]byte, len(prefix)+1) // one more for the separator
	for i := 0; i < len(prefix); i++ {
		ret[i] = byte(prefix[i]) & 0x1f
	}
	ret[len(prefix)] = 0 // separator
	return ret
}

// converts an array of byte with inSize bits to an array of byte with toSize bits
func convert(data []byte, inSize uint, toSize uint, strict bool) ([]byte, error) {
	var outLen int
	if strict {
		outLen = int(math.Floor((float64(len(data)) * float64(inSize)) / float64(toSize)))
	} else {
		outLen = int(math.Ceil((float64(len(data)) * float64(inSize)) / float64(toSize)))
	}
	mask := uint((1 << toSize) - 1)
	result := make([]byte, outLen, outLen)
	index := 0
	var accumulator uint
	var bits uint
	for i := 0; i < len(data); i++ {
		value := uint(data[i])
		if (value < 0) || ((value >> inSize) != 0) {
			return nil, fmt.Errorf("invalid value %x", value)
		}
		accumulator = (accumulator << inSize) | value
		bits += inSize
		for bits >= toSize {
			bits -= toSize
			result[index] = byte((accumulator >> bits) & mask)
			index++
		}
	}
	end := byte((accumulator << (toSize - bits)) & mask)
	if !strict {
		if bits > 0 {
			result[index] = end
			index++
		}
	} else {
		if (bits > inSize) || (end != 0) {
			return nil, fmt.Errorf("strict mode required but input connot be converted to %d bits without padding", toSize)
		}
	}
	return result, nil
}

func getHashSize(hash []byte) (int8, error) {
	switch len(hash) * 8 {
	case 160:
		return 0, nil
	case 192:
		return 1, nil
	case 224:
		return 2, nil
	case 256:
		return 3, nil
	case 320:
		return 4, nil
	case 384:
		return 5, nil
	case 448:
		return 6, nil
	case 512:
		return 7, nil
	default:
		return -1, fmt.Errorf("invalid hash size: %d ", len(hash))
	}
}
