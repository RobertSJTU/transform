package sc
import (
  "testing"
  "github.com/stretchr/testify/require"
  "errors"
)
type testSc struct {
  priv []byte//the value of private key with the type of []byte
  pub []byte//the value of public key with the type of []byte
  address string//the value of address with the type of string
}

//the correct test data
var ScTest = testSc {[]byte {7, 143, 39, 64, 144, 76, 88, 69, 132, 99, 70, 84, 192, 98 ,83 ,137, 148, 171, 138, 99, 209, 165, 60 ,129 ,26 ,156, 100, 24 ,230 ,165, 235, 249, 136, 173, 245, 147 ,180 ,212, 87, 204, 185, 243, 24 ,178 ,236 ,227 ,130, 109, 100, 69, 240, 222, 46, 80, 159, 222, 99, 8 ,183 ,141, 157, 89, 154, 59}, 
           []byte {136, 173, 245, 147, 180, 212, 87, 204, 185, 243, 24, 178, 236, 227, 130, 109, 100 ,69 ,240, 222, 46, 80, 159, 222, 99, 8 ,183, 141, 157, 89 ,154, 59},
           "fa986e9499482306d81e14d64c2ded0a00dfc5f6e2af5b814f1592d59c9740f71d462e783d9a"}


//the test of the return value of error
//the first data has a wrong length of private key, it is 31, lacking 7
var errorPriv = []byte { 143, 39, 64, 144, 76, 88, 69, 132, 99, 70, 84, 192, 98 ,83 ,137, 148, 171, 138, 99, 209, 165, 60 ,129 ,26 ,156, 100, 24 ,230 ,165, 235, 249, 136, 173, 245, 147 ,180 ,212, 87, 204, 185, 243, 24 ,178 ,236 ,227 ,130, 109, 100, 69, 240, 222, 46, 80, 159, 222, 99, 8 ,183 ,141, 157, 89, 154, 59}
//the second data has a wrong length of public key, it is 32, lacking 136
var errorPub1 = []byte { 173, 245, 147, 180, 212, 87, 204, 185, 243, 24, 178, 236, 227, 130, 109, 100 ,69 ,240, 222, 46, 80, 159, 222, 99, 8 ,183, 141, 157, 89 ,154, 59}

func TestPrivKeyToPub(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    pubTest, err := sc.PrivKeyToPub(ScTest.priv)
    require.Nil(err)
    require.Equal( ScTest.pub, pubTest)      
    })
  t.Run("test the err", func(t *testing.T){
    _, err := sc.PrivKeyToPub(errorPriv)
    require.Equal("the length of the imported private key is wrong, please check the import", err.Error())
    })  
  //fmt.Println("The function of PrivKeyToPub works well")
}
func TestPubkeyToAddress(t *testing.T) {
  require := require.New(t)
  t.Run("test the correct result", func(t *testing.T){
    addrTest, err:= sc.PubKeyToAddress(ScTest.pub)
    require.Nil(err)
    require.Equal(ScTest.address, addrTest)
  })
  t.Run("test the err", func(t *testing.T){
    _, err1 := sc.PubKeyToAddress(errorPub1)
    require.Equal(errors.New("the length of the imported public key is wrong, please check the import"), err1)
    })
 
  //fmt.Println("The function of PrivKeyToPub works well")
}
func BenchmarkPrivKeyToPub(b *testing.B) {
  for i := 0; i < b.N; i++ {
    sc.PrivKeyToPub(ScTest.priv)
  }
}
func BenchmarkPubkeyToAddress(b *testing.B) {
  for i := 0; i < b.N; i++ {
    sc.PubKeyToAddress(ScTest.pub)
  }
}
