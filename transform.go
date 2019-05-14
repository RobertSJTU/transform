package transform

import (
   "fmt"
   "errors"
   "sync"
   "github.com/btcsuite/btcd/btcec"
   "github.com/btcsuite/btcutil"
   "github.com/btcsuite/btcutil/base58"
   "golang.org/x/crypto/ripemd160"
)

type Transform interface {
PrivKeyToPub(priv []byte) (pub []byte, err error)
PubKeyToAddress(pub []byte) (add string, err error)
}

//lock the map
var lock = struct{
  sync.RWMutex
  drivers map[string] Transform
} {drivers : make(map[string] Transform)}


func Register(name string, driver Transform) {
    if driver == nil {
        panic("transform: Register driver is nil")
    }
    if _, dup := lock.drivers[name]; dup {
        panic("transform: Register called twice for driver " + name)
    }
    lock.Lock()
    lock.drivers[name] = driver
    defer lock.Unlock()
}

func New(name string) (t Transform, err error) {
    lock.RLock()
    t, ok := lock.drivers[name]
    defer lock.RUnlock()
    if !ok {
        err = fmt.Errorf("unknown driver %q", name)
        return
    }

    return t, nil
}
//PrivKeyToPub1 transforms the private key to public key for all the type of coins
func PrivKeyToPub1(priv []byte) (pub []byte, err error) {
   if len(priv) != 32 {
    //to judge the length of the input, the default is 32
    //fmt.Println("the length of the imported private key is wrong, please check the import")
    return nil, errors.New("the length of the imported private key is wrong, please check the import")
   }
   privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), priv)//to get the data of private key by the function from btcec
   pubKeySerial := privKey.PubKey().SerializeCompressed()//to get the compressed public key with the type of []byte, and the length is 33, the first byte is the symbol byte 
   return pubKeySerial, nil
}
//PubKeyToAddress1 transform public key to address for BTC USDT BTY
func PubKeyToAddress1(pub []byte) (add string, err error) {
   if len(pub) != 33 {//to judge the length of the public key, the default is 33
    //fmt.Println("the length of the imported public key is wrong, please check the import")
    return "", errors.New("the length of the imported public key is wrong, please check the import")
    }
   if pub[0]!=0x02 && pub[0]!=0x03 {//to judge whether the public key's format is right
    //fmt.Println("the format of the imported public key is wrong, please check the import")
    return "", errors.New("the format of the imported public key is wrong, please check the import")
   }
    pubhash := btcutil.Hash160(pub)//to get the second hash(ripemd160) value of the public key
    return base58.CheckEncode(pubhash[:ripemd160.Size], 000),nil//transform the address' style to string
}
