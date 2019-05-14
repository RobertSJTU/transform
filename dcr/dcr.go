package dcr

import (
   "transform"
   "errors"
   "github.com/decred/base58"
   "github.com/decred/dcrd/dcrutil"
)

type DCR struct {
} 

var dcr transform.Transform = new(DCR)
var netId = [2]byte {0x07, 0x3F}
func Init() {
  transform.Register("DCR",dcr)
}

func (dcr DCR) PrivKeyToPub(priv []byte) (pub []byte, err error) {
  return transform.PrivKeyToPub1(priv)
}


func (dcr DCR) PubKeyToAddress(pub []byte) (add string, err error) {
  if len(pub) != 33 {
    return "", errors.New("the length of the imported public key is wrong, please check the import")
    }
  if pub[0]!=0x02 && pub[0]!=0x03 {
    return "", errors.New("the format of the imported public key is wrong, please check the import")
   }
  pubhash := dcrutil.Hash160(pub)
  address := base58.CheckEncode(pubhash[:20], netId)
   return address, nil
}

