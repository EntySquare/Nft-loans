package contracts

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type config struct {
	Dial            string
	ContractAddress string
	Pk              string
}

func NewInstance(network string) (*Contracts, *bind.TransactOpts, error) {
	var config config
	if network == "polygon" {
		config.Dial = "https://bsc-dataseed1.binance.org"
		config.ContractAddress = "0x60A3Cff47fCA4eA4cBf28ff47E001F3a4468527b"
		config.Pk = "privateKey"
	} else if network == "eth" {
		config.Dial = "https://mainnet.infura.io/v3/a936bfa4553a4a95862326edddc46306"
		config.ContractAddress = "0x286699858aBAbA49Cace3681C1Ca3defDDC91868"
		config.Pk = "privateKey"
	}

	client, err := ethclient.Dial(config.Dial)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	privateKey, err := crypto.HexToECDSA(config.Pk)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
		return nil, nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units  auth.GasLimit = uint64(4200000)
	auth.GasPrice = gasPrice
	address := common.HexToAddress(config.ContractAddress)
	if network == "eth" {
		var gasPriceNew *big.Int
		gasPriceNew = big.NewInt(gasPrice.Int64() + 3*1000000000) // 额外加3gwei
		auth.GasPrice = gasPriceNew
	}

	instance, err := NewContracts(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance, auth, nil

}

func TransferFrom(_from common.Address, _to common.Address, _value *big.Int, chain string) string {
	instance, _, err := NewInstance(chain)
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))
	tx, err := instance.TransferFrom(&bind.TransactOpts{}, _from, _to, _value)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("tx : %s", tx.Hash().String())
	return tx.Hash().String()

}

func WithdrawNft(_to common.Address, _tokenId *big.Int, chain string) string {

	instance, auth, err := NewInstance(chain)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))
	//tx, err := instance.SetPool(auth, "german", "japan", "1", false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("tx : %s", tx.Hash().String())
	tx, err := instance.WithdrawNFT(auth, _to, _tokenId)
	if err != nil {
		log.Fatal(err)
	}

	//poolKey := "0x" + pass
	return tx.Hash().String()

	//fmt.Printf("tx : %s", poolKey) //1d9f0539884bd60a5899c82d431c811b5f20f58be9b8c54bf97eeaaee71cacb0
	//0x1d9f0539884bd60a5899c82d431c811b5f20f58be9b8c54bf97eeaaee71cacb0
}
