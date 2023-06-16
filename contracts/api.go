package contracts

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
	"math"
	"math/big"
	loansconfig "nft-loans/config"
	"strings"
	"time"
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

func CheckHash(hashStr string) (_from string, _to string, _v string, f float64, err error) {
	if len(hashStr) < 30 {
		return "", "", "", 0, nil
	}
	const transferFnSignature = "transfer(address,uint256)"
	client, err := ethclient.Dial(loansconfig.Config("CHAIN_RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 请替换为您要查询的交易哈希
	txHash := common.HexToHash(hashStr)
	// USDT合约地址
	usdtContractAddress := common.HexToAddress(loansconfig.Config("CONTRACT_ADDRESS"))

	ticker := time.NewTicker(10 * time.Second)

	// 计算 transfer 函数签名的 Keccak-256 哈希的前8个字符
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(transferFnSignature))
	transferFnID := hex.EncodeToString(hash.Sum(nil)[:4])

	i := 0
	for {
		select {
		case <-ticker.C:
			tx, pending, err := client.TransactionByHash(context.Background(), txHash)
			if pending {
				continue
			}
			if err != nil {
				if err == ethereum.NotFound {
					i++
					if i == 30 {
						return "", "", "", 0, nil
					}
					log.Println("Transaction not found", i)
					continue
				}
				log.Fatal(err)
			}

			if strings.ToLower(tx.To().Hex()) == strings.ToLower(usdtContractAddress.Hex()) {
				receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					log.Fatal(err)
					continue
				}

				if receipt.Status == types.ReceiptStatusSuccessful {
					txDataHex := hex.EncodeToString(tx.Data())
					if strings.HasPrefix(txDataHex, transferFnID) {
						from, err := client.TransactionSender(context.Background(), tx, receipt.BlockHash, 0)
						if err != nil {
							log.Fatal(err)
						}

						parsedData, err := abi.JSON(strings.NewReader(
							`[
	{
		"constant": false,
		"inputs": [
			{
				"name": "_to",
				"type": "address"
			},
			{
				"name": "_value",
				"type": "uint256"
			}
		],
		"name": "transfer",
		"outputs": [
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`))
						if err != nil {
							log.Fatal(err)
						}

						inputData, err := hex.DecodeString(txDataHex[8:])
						if err != nil {
							log.Fatal(err)
						}
						unpackedValues, err := parsedData.Methods["transfer"].Inputs.Unpack(inputData)
						if err != nil {
							log.Fatal(err)
						}

						to := unpackedValues[0].(common.Address)
						value := unpackedValues[1].(*big.Int)

						fmt.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
						fmt.Printf("From: %s\n", from.Hex())
						fmt.Printf("To: %s\n", to.Hex())
						f, acu := new(big.Float).Quo(new(big.Float).SetInt(value), new(big.Float).SetFloat64(math.Pow10(4))).Float64()
						fmt.Println("NFT Value: ", f, acu)
						_from = from.Hex()
						_to = to.Hex()
						_v = value.Text(10)
						//if _v != "3000000000" {
						//	return _from, _to, _v, errors.New("_v != '3000000000'")
						//}
						//fmt.Println("----------------------")
						return _from, _to, _v, f, nil
					} else {
						log.Println("Transaction is not a USDT transfer")
					}
				} else {
					log.Println("Transaction failed")
				}
				// 交易已完成，退出程序
				return "", "", "", 0, errors.New("CheckHash err 001")
			} else {
				log.Println("Transaction still pending or not a USDT transfer")
				return "", "", "", 0, errors.New("CheckHash err 002")
			}
		}
	}
}
