package contracts

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	loansconfig "nft-loans/config"
)

type ChainConfig struct {
	Dial            string
	ContractAddress string
	Pk              string
}

// GetAddressLoansList 查询这个钱包的 抵押记录
func NewContractsApi() (*Contracts, *bind.TransactOpts, error) {
	var config2 = ChainConfig{
		Dial:            loansconfig.Config("CHAIN_RPC_URL"),
		ContractAddress: loansconfig.Config("CONTRACT_ADDRESS"),
		Pk:              loansconfig.Config("CONTRACT_PRIVATE_KEY"),
	}

	client, err := ethclient.Dial(config2.Dial)
	if err != nil {
		return nil, nil, err
	}
	//defer client.Close()

	privateKey, err := crypto.HexToECDSA(config2.Pk)
	if err != nil {
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		//log.Fatal("error casting public key to ECDSA")
		return nil, nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		//log.Fatal(err)
		return nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		//log.Fatal(err)
		return nil, nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units  auth.GasLimit = uint64(4200000)
	auth.GasPrice = gasPrice
	address := common.HexToAddress(config2.ContractAddress)
	//if network == "eth" {
	//	var gasPriceNew *big.Int
	//	gasPriceNew = big.NewInt(gasPrice.Int64() + 3*1000000000) // 额外加3gwei
	//	auth.GasPrice = gasPriceNew
	//}

	contracts, err := NewContracts(address, client)
	if err != nil {
		return nil, nil, err
	}
	//return instance, auth, nil
	return contracts, auth, nil
}

// GetAddressLoansList 查询这个钱包的 抵押记录
func GetAddressLoansList(address string) ([]NGTTokenData, error) {
	contracts, _, err := NewContractsApi()
	if err != nil {
		return nil, err
	}
	user := common.HexToAddress(address)
	list, err := contracts.LoansList(&bind.CallOpts{}, user)
	if err != nil {
		return nil, err
	}
	return list, nil
}
