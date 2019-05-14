package ltc

import (
   "transform"
   "errors"
   "github.com/btcsuite/btcutil/base58"
   "github.com/btcsuite/btcutil"
)

type LTC struct {
} 

var ltc transform.Transform = new(LTC)

func Init() {
  transform.Register("LTC",ltc)
}

func (ltc LTC) PrivKeyToPub(priv []byte) (pub []byte, err error) {
  return transform.PrivKeyToPub1(priv)
}


func (ltc LTC) PubKeyToAddress(pub []byte) (add string, err error) {
  if len(pub) != 33 {
    return "", errors.New("the length of the imported public key is wrong, please check the import")
    }
  if pub[0]!=0x02 && pub[0]!=0x03 {
    return "", errors.New("the format of the imported public key is wrong, please check the import")
   }
   pubhash := btcutil.Hash160(pub)
   address := base58.CheckEncode(pubhash[:20], 0x30)
   return address, nil
}

