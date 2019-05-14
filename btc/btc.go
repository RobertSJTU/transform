package btc

import (
   "transform"
)

type BTC struct {
} 

var btc transform.Transform = new(BTC)

func Init() {
  transform.Register("BTC",btc)
}

func (btc BTC) PrivKeyToPub(priv []byte) (pub []byte, err error) {
  return transform.PrivKeyToPub1(priv)
}

func (btc BTC) PubKeyToAddress(pub []byte) (add string, err error) {
  return transform.PubKeyToAddress1(pub)
}

