package btc
import (
  "testing"
  "github.com/stretchr/testify/require"
  "errors"
)
type testBtc struct {
  priv []byte//the value of private key with the type of []byte
  pub []byte//the value of public key with the type of []byte
  address string//the value of address with the type of string
}

//the correct test data
var BtcTest = []testBtc {
  testBtc {[]byte{194, 125, 101, 129, 185, 39, 133, 131, 75, 56, 31, 166, 151, 196, 176, 255, 196, 87, 75, 73, 87, 67, 114, 46, 10, 203, 118, 1, 177, 182, 139, 153}, 
           []byte{2, 216, 248, 132, 104, 197, 162, 232, 225, 129, 95, 175, 85, 95, 89, 203, 209, 151, 158, 61, 189, 248, 35, 248, 12, 39, 27, 111, 183, 13, 45, 81, 155},
           "14wtcepMNiEazuN7YosWY8bwD9tcCtxXRB"},
  testBtc {[]byte{210, 6, 195, 255, 97, 203, 104, 69, 146, 219, 13, 40, 27, 9, 17, 177, 78, 4, 219, 226, 180, 102, 112, 44, 157, 12, 8, 121, 191, 109, 112, 110}, 
           []byte{3, 186, 1, 208, 74, 83, 165, 98, 78, 113, 101, 201, 0, 73, 94, 245, 210, 133, 125, 241, 229, 43, 89, 197, 90, 138, 241, 102, 241, 190, 81, 130, 154},
           "1PebMaea4HGr2b2PDKepLi8Xdvj8fqRLJ3"},
  testBtc {[]byte{24,243,177,145,1,158,131,135,138,129,85,122,190,187,42,253,161,153,227,29,34,225,80,216,191,77,244,86,22,113,190,108},
           []byte{2, 185, 193, 117, 185, 8, 98, 79, 138, 142, 170, 194, 39, 208, 232, 199, 124, 14, 236, 50, 123, 140, 81, 42, 209, 184, 183, 164, 181, 182, 118, 151, 31},
           "1N3D8jy2aQuUsKBsDgZ6ZPTVR9VhHgJYpE"},
         }

//the test of the return value of error
//the first data has a wrong length of private key, it is 31, lacking 194
var errorPriv = []byte {125, 101, 129, 185, 39, 133, 131, 75, 56, 31, 166, 151, 196, 176, 255, 196, 87, 75, 73, 87, 67, 114, 46, 10, 203, 118, 1, 177, 182, 139, 153}
//the second data has a wrong length of public key, it is 32, lacking 216
var errorPub1 = []byte{2, 248, 132, 104, 197, 162, 232, 225, 129, 95, 175, 85, 95, 89, 203, 209, 151, 158, 61, 189, 248, 35, 248, 12, 39, 27, 111, 183, 13, 45, 81, 155}
//the third data has a wrong format of public key, the first is 4, it is irregular
var errorPub2 = []byte{4, 216, 248, 132, 104, 197, 162, 232, 225, 129, 95, 175, 85, 95, 89, 203, 209, 151, 158, 61, 189, 248, 35, 248, 12, 39, 27, 111, 183, 13, 45, 81, 155}

func TestPrivKeyToPub(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    for _, v := range(BtcTest) {
    pubTest, err := btc.PrivKeyToPub(v.priv)
    require.Nil(err)
    require.Equal(pubTest, v.pub)
    }      
    })
  t.Run("test the err", func(t *testing.T){
    _, err := btc.PrivKeyToPub(errorPriv)
    require.Equal("the length of the imported private key is wrong, please check the import", err.Error())
    })  
  //fmt.Println("The function of PrivKeyToPub works well")
}
func TestPubkeyToAddress(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    for _, v := range(BtcTest) {
    addrTest, err:= btc.PubKeyToAddress(v.pub)
    require.Nil(err)
    require.Equal(addrTest, v.address)
    }
  })
  t.Run("test the err", func(t *testing.T){
    _, err1 := btc.PubKeyToAddress(errorPub1)
    _, err2 := btc.PubKeyToAddress(errorPub2)
    require.Equal(errors.New("the length of the imported public key is wrong, please check the import"), err1)
    require.Equal(errors.New("the format of the imported public key is wrong, please check the import"), err2)
    })
 
  //fmt.Println("The function of PrivKeyToPub works well")
}
func BenchmarkPrivKeyToPub(b *testing.B) {
  for i := 0; i < b.N; i++ {
    btc.PrivKeyToPub(BtcTest[1].priv)
  }
}
func BenchmarkPubkeyToAddress(b *testing.B) {
  for i := 0; i < b.N; i++ {
    btc.PubKeyToAddress(BtcTest[1].pub)
  }
}

