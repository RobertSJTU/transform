package usdt

import (
   "transform"
)

type USDT struct {
} 

var usdt transform.Transform = new(USDT)

func Init() {
  transform.Register("USDT",usdt)
}

func (usdt USDT) PrivKeyToPub(priv []byte) (pub []byte, err error) {
  return transform.PrivKeyToPub1(priv)
}

func (usdt USDT) PubKeyToAddress(pub []byte) (add string, err error) {
  return transform.PubKeyToAddress1(pub)
}

