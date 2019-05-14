package etc
import (
  "testing"
  "github.com/stretchr/testify/require"
  "errors"
)
type testEtc struct {
  priv []byte//the value of private key with the type of []byte
  pub []byte//the value of public key with the type of []byte
  address string//the value of address with the type of string
}

//the correct test data
var EtcTest = testEtc {[]byte {40, 156, 40, 87, 212, 89, 142, 55, 251, 150, 71, 80, 126, 71, 163, 9, 214, 19, 53, 57, 191, 33, 168, 185, 203, 109, 248, 143, 213, 35, 32, 50}, 
           []byte{3, 125, 178, 39, 215, 9, 76, 226, 21, 195, 160, 245, 126, 27, 204, 115, 37, 81, 254, 53, 31, 148, 36, 148, 113, 147, 69, 103, 224, 245, 220, 27, 247},
           "0x970e8128ab834e8eac17ab8e3812f010678cf791"}


//the test of the return value of error
//the first data has a wrong length of private key, it is 31, lacking 40
var errorPriv = []byte {156, 40, 87, 212, 89, 142, 55, 251, 150, 71, 80, 126, 71, 163, 9, 214, 19, 53, 57, 191, 33, 168, 185, 203, 109, 248, 143, 213, 35, 32, 50}
//the second data has a wrong length of public key, it is 32, lacking 125
var errorPub1 = []byte {3, 178, 39, 215, 9, 76, 226, 21, 195, 160, 245, 126, 27, 204, 115, 37, 81, 254, 53, 31, 148, 36, 148, 113, 147, 69, 103, 224, 245, 220, 27, 247}
//the third data has a wrong format of public key, the first is 4, it is irregular
var errorPub2 = []byte {4, 125, 178, 39, 215, 9, 76, 226, 21, 195, 160, 245, 126, 27, 204, 115, 37, 81, 254, 53, 31, 148, 36, 148, 113, 147, 69, 103, 224, 245, 220, 27, 247}
//{4, 216, 248, 132, 104, 197, 162, 232, 225, 129, 95, 175, 85, 95, 89, 203, 209, 151, 158, 61, 189, 248, 35, 248, 12, 39, 27, 111, 183, 13, 45, 81, 155}

func TestPrivKeyToPub(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    pubTest, err := etc.PrivKeyToPub(EtcTest.priv)
    require.Nil(err)
    require.Equal( EtcTest.pub, pubTest)      
    })
  t.Run("test the err", func(t *testing.T){
    _, err := etc.PrivKeyToPub(errorPriv)
    require.Equal("the length of the imported private key is wrong, please check the import", err.Error())
    })  
  //fmt.Println("The function of PrivKeyToPub works well")
}
func TestPubkeyToAddress(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    addrTest, err:= etc.PubKeyToAddress(EtcTest.pub)
    require.Nil(err)
    require.Equal(EtcTest.address, addrTest)
  })
  t.Run("test the err", func(t *testing.T){
    _, err1 := etc.PubKeyToAddress(errorPub1)
    _, err2 := etc.PubKeyToAddress(errorPub2)
    require.Equal(errors.New("the length of the imported public key is wrong, please check the import"), err1)
    require.Equal(errors.New("the format of the imported public key is wrong, please check the import"), err2)
    })
 
  //fmt.Println("The function of PrivKeyToPub works well")
}
func BenchmarkPrivKeyToPub(b *testing.B) {
  for i := 0; i < b.N; i++ {
    etc.PrivKeyToPub(EtcTest.priv)
  }
}
func BenchmarkPubkeyToAddress(b *testing.B) {
  for i := 0; i < b.N; i++ {
    etc.PubKeyToAddress(EtcTest.pub)
  }
}