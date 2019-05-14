package bty

import (
   "transform"
)

type BTY struct {
} 

var bty transform.Transform = new(BTY)

func Init() {
  transform.Register("BTY",bty)
}

func (bty BTY) PrivKeyToPub(priv []byte) (pub []byte, err error) {
  return transform.PrivKeyToPub1(priv)
}

func (bty BTY) PubKeyToAddress(pub []byte) (add string, err error) {
  return transform.PubKeyToAddress1(pub)
}

