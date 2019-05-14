package ltc
import (
  "testing"
  "github.com/stretchr/testify/require"
  "errors"
)
type testLtc struct {
  priv []byte//the value of private key with the type of []byte
  pub []byte//the value of public key with the type of []byte
  address string//the value of address with the type of string
}

//the correct test data
var LtcTest = testLtc {[]byte {57, 189, 218, 225, 96 ,250, 181, 231, 93, 178, 85, 213, 49, 1 ,181, 53 ,111 ,70 ,131, 254 ,31, 27, 127, 77, 164, 193, 113, 196, 6 ,43, 229, 97}, 
           []byte{2, 5, 71, 219, 42, 251, 35, 40, 149, 213, 2, 138, 80, 61, 180, 245, 5 ,132, 231, 61 ,102, 32, 81 ,53 ,104, 244, 247, 211, 134, 16 ,236, 113, 127},
           "LUHGh2bYtgPGJAJcUPas8W83woUnBJrXsk"}


//the test of the return value of error
//the first data has a wrong length of private key, it is 31, lacking 57
var errorPriv = []byte {189, 218, 225, 96 ,250, 181, 231, 93, 178, 85, 213, 49, 1 ,181, 53 ,111 ,70 ,131, 254 ,31, 27, 127, 77, 164, 193, 113, 196, 6 ,43, 229, 97}
//the second data has a wrong length of public key, it is 32, lacking 5
var errorPub1 = []byte {2,  71, 219, 42, 251, 35, 40, 149, 213, 2, 138, 80, 61, 180, 245, 5 ,132, 231, 61 ,102, 32, 81 ,53 ,104, 244, 247, 211, 134, 16 ,236, 113, 127}
//the third data has a wrong format of public key, the first is 4, it is irregular
var errorPub2 = []byte {4, 5, 71, 219, 42, 251, 35, 40, 149, 213, 2, 138, 80, 61, 180, 245, 5 ,132, 231, 61 ,102, 32, 81 ,53 ,104, 244, 247, 211, 134, 16 ,236, 113, 127}
//{4, 216, 248, 132, 104, 197, 162, 232, 225, 129, 95, 175, 85, 95, 89, 203, 209, 151, 158, 61, 189, 248, 35, 248, 12, 39, 27, 111, 183, 13, 45, 81, 155}

func TestPrivKeyToPub(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    pubTest, err := ltc.PrivKeyToPub(LtcTest.priv)
    require.Nil(err)
    require.Equal( LtcTest.pub, pubTest)      
    })
  t.Run("test the err", func(t *testing.T){
    _, err := ltc.PrivKeyToPub(errorPriv)
    require.Equal("the length of the imported private key is wrong, please check the import", err.Error())
    })  
  //fmt.Println("The function of PrivKeyToPub works well")
}
func TestPubkeyToAddress(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    addrTest, err:= ltc.PubKeyToAddress(LtcTest.pub)
    require.Nil(err)
    require.Equal(LtcTest.address, addrTest)
  })
  t.Run("test the err", func(t *testing.T){
    _, err1 := ltc.PubKeyToAddress(errorPub1)
    _, err2 := ltc.PubKeyToAddress(errorPub2)
    require.Equal(errors.New("the length of the imported public key is wrong, please check the import"), err1)
    require.Equal(errors.New("the format of the imported public key is wrong, please check the import"), err2)
    })
 
  //fmt.Println("The function of PrivKeyToPub works well")
}
func BenchmarkPrivKeyToPub(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ltc.PrivKeyToPub(LtcTest.priv)
  }
}
func BenchmarkPubkeyToAddress(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ltc.PubKeyToAddress(LtcTest.pub)
  }
}
