package sc

import (
   "transform"
   "errors"
   "github.com/NebulousLabs/Sia/types"

)

type SC struct {
} 

var sc transform.Transform = new(SC)
func Init() {
  transform.Register("SC",sc)
}

func (sc SC) PrivKeyToPub(priv []byte) (pub []byte, err error) {
  if len(priv) != 64 {//to judge the length of the input, the default is 32
    //fmt.Println("the length of the imported private key is wrong, please check the import")
    return nil, errors.New("the length of the imported private key is wrong, please check the import")
   }
  return priv[32:], nil
}


func (sc SC) PubKeyToAddress(pub []byte) (add string, err error) {
  if len(pub) != 32 {
    return "", errors.New("the length of the imported public key is wrong, please check the import")
    }
    var SiaPublicKey1 types.SiaPublicKey
    SiaPublicKey1.Algorithm =  types.SignatureEd25519
    SiaPublicKey1.Key = pub

    var UnlockConditions types.UnlockConditions
    UnlockConditions.Timelock = 0
    UnlockConditions.SignaturesRequired = 1
    var var1 = []types.SiaPublicKey{SiaPublicKey1}
    UnlockConditions.PublicKeys = var1

    unlockhash := UnlockConditions.UnlockHash()

    return unlockhash.String(), nil
}

