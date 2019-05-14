package bch
import (
  "testing"
  "github.com/stretchr/testify/require"
  "errors"
)
type testBch struct {
  priv []byte//the value of private key with the type of []byte
  pub []byte//the value of public key with the type of []byte
  address string//the value of address with the type of string
}

//the correct test data
var BchTest = testBch {[]byte {78, 25, 183, 99, 11, 198, 222, 40, 11, 126, 6, 190, 214, 93, 228, 151, 24, 166, 90, 196, 142, 60, 150, 40, 65, 219, 131, 71, 158, 57, 76, 100}, 
           []byte{2, 168, 244, 8, 9, 195, 187, 36, 245, 88, 13, 68, 125, 124, 191, 73, 83, 232, 19, 101, 108, 243, 122, 150, 89, 251, 221, 201, 183, 184, 38, 96, 78},
           "bitcoincash:qz5u8hq0nms68n06fuyu7ez4cv7sp9nj6gw2gn4c9x"}


//the test of the return value of error
//the first data has a wrong length of private key, it is 31, lacking 25
var errorPriv = []byte {25, 183, 99, 11, 198, 222, 40, 11, 126, 6, 190, 214, 93, 228, 151, 24, 166, 90, 196, 142, 60, 150, 40, 65, 219, 131, 71, 158, 57, 76, 100}
//the second data has a wrong length of public key, it is 32, lacking 168
var errorPub1 = []byte {2,  244, 8, 9, 195, 187, 36, 245, 88, 13, 68, 125, 124, 191, 73, 83, 232, 19, 101, 108, 243, 122, 150, 89, 251, 221, 201, 183, 184, 38, 96, 78}
//the third data has a wrong format of public key, the first is 4, it is irregular
var errorPub2 = []byte {4, 168, 244, 8, 9, 195, 187, 36, 245, 88, 13, 68, 125, 124, 191, 73, 83, 232, 19, 101, 108, 243, 122, 150, 89, 251, 221, 201, 183, 184, 38, 96, 78}
//{4, 216, 248, 132, 104, 197, 162, 232, 225, 129, 95, 175, 85, 95, 89, 203, 209, 151, 158, 61, 189, 248, 35, 248, 12, 39, 27, 111, 183, 13, 45, 81, 155}

func TestPrivKeyToPub(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    pubTest, err := bch.PrivKeyToPub(BchTest.priv)
    require.Nil(err)
    require.Equal( BchTest.pub, pubTest)      
    })
  t.Run("test the err", func(t *testing.T){
    _, err := bch.PrivKeyToPub(errorPriv)
    require.Equal("the length of the imported private key is wrong, please check the import", err.Error())
    })  
  //fmt.Println("The function of PrivKeyToPub works well")
}
func TestPubkeyToAddress(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    addrTest, err:= bch.PubKeyToAddress(BchTest.pub)
    require.Nil(err)
    require.Equal(BchTest.address, addrTest)
  })
  t.Run("test the err", func(t *testing.T){
    _, err1 := bch.PubKeyToAddress(errorPub1)
    _, err2 := bch.PubKeyToAddress(errorPub2)
    require.Equal(errors.New("the length of the imported public key is wrong, please check the import"), err1)
    require.Equal(errors.New("the format of the imported public key is wrong, please check the import"), err2)
    })
 
  //fmt.Println("The function of PrivKeyToPub works well")
}
func BenchmarkPrivKeyToPub(b *testing.B) {
  for i := 0; i < b.N; i++ {
    bch.PrivKeyToPub(BchTest.priv)
  }
}
func BenchmarkPubkeyToAddress(b *testing.B) {
  for i := 0; i < b.N; i++ {
    bch.PubKeyToAddress(BchTest.pub)
  }
}
