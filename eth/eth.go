package eth

import (
   "transform"
   "github.com/ethereum/go-ethereum/crypto"
   "encoding/hex"
   "errors"
)

type ETH struct {
} 

var eth transform.Transform = new(ETH)

func Init() {
  transform.Register("ETH",eth)
}

func (eth ETH) PrivKeyToPub(priv []byte) (pub []byte, err error) {
  return transform.PrivKeyToPub1(priv)
}

func (eth ETH) PubKeyToAddress(pub []byte) (add string, err error) {
   if len(pub) != 33 {//to judge the length of the public key, the default is 33
    return "", errors.New("the length of the imported public key is wrong, please check the import")
    }
   if pub[0]!=0x02 && pub[0]!=0x03 {//to judge whether the public key's format is right
    return "", errors.New("the format of the imported public key is wrong, please check the import")
   }
	pubkey, _ := crypto.DecompressPubkey(pub)//get the uncompressed pubkey
	address := crypto.PubkeyToAddress(*pubkey)//get the address
	return "0x" + hex.EncodeToString(address[:]),nil//get the address in base58
}