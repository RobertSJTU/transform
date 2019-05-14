package zec

import (
   "transform"
   "errors"
   "github.com/blackkeyboard/zgenerate/base58"
   "github.com/btcsuite/btcutil"
)

type ZEC struct {
} 

var zec transform.Transform = new(ZEC)
var netId = [2]byte {0x1C, 0xB8}
func Init() {
  transform.Register("ZEC",zec)
}

func (zec ZEC) PrivKeyToPub(priv []byte) (pub []byte, err error) {
  return transform.PrivKeyToPub1(priv)
}


func (zec ZEC) PubKeyToAddress(pub []byte) (add string, err error) {
  if len(pub) != 33 {
    return "", errors.New("the length of the imported public key is wrong, please check the import")
    }
  if pub[0]!=0x02 && pub[0]!=0x03 {
    return "", errors.New("the format of the imported public key is wrong, please check the import")
   }
   pubhash := btcutil.Hash160(pub)
   address := base58.CheckEncode(pubhash[:20], netId)
   return address, nil
}

