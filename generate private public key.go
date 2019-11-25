package main
import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)
func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil{
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey) //转换为字节
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	PublicKey := privateKey.Public()
	PublicKeyECDSA , ok := PublicKey.(*ecdsa.PublicKey)
	if !ok{
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(PublicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	address := crypto.PubkeyToAddress(*PublicKeyECDSA).Hex()

	fmt.Println(address)
	hash := sha256.New()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))


}