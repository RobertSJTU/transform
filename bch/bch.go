package bch

import (
   "transform"
   "github.com/savardiego/cashline/address/cash"
   "errors"
)

type BCH struct {
} 

var bch transform.Transform = new(BCH)

func Init() {
  transform.Register("BCH",bch)
}

func (bch BCH) PrivKeyToPub(priv []byte) (pub []byte, err error) {
  return transform.PrivKeyToPub1(priv)
}


func (bch BCH) PubKeyToAddress(pub []byte) (add string, err error) {
  if len(pub) != 33 {
    return "", errors.New("the length of the imported public key is wrong, please check the import")
    }
  if pub[0]!=0x02 && pub[0]!=0x03 {
    return "", errors.New("the format of the imported public key is wrong, please check the import")
   }
   address, _ := cash.FromPubKey(pub)
   return address, nil
}

