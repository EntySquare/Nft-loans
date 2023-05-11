package tron

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/common/decimals"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
)

// 地址生成，注意校验和的计算
func publicKeyToTronAddresstest(publicKey *ecdsa.PublicKey) string {
	// 获取公钥的字节表示形式
	pubBytes := crypto.FromECDSAPub(publicKey)

	// 使用 Keccak-256（SHA-3）哈希算法计算公钥哈希
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubBytes[1:])
	address := hash.Sum(nil)[12:]

	// 将波场地址字节前添加 0x41（十进制 65）
	tronAddressBytes := append([]byte{0x41}, address...)

	// 计算波场地址的校验和
	hash256 := sha256.New()
	hash256.Write(tronAddressBytes)
	firstHash := hash256.Sum(nil)
	hash256.Reset()
	hash256.Write(firstHash)
	secondHash := hash256.Sum(nil)
	checksum := secondHash[:4]

	// 将波场地址字节与校验和合并，生成原始地址
	rawAddress := append(tronAddressBytes, checksum...)

	// 使用 Base58 编码生成波场地址字符串表示形式
	tronAddress := base58.Encode(rawAddress)
	return tronAddress
}

// 测试tronlink兼容的助记词和地址生成
func TestTron(t *testing.T) {
	// 生成 256 位熵
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		panic(err)
	}

	// 使用熵生成助记词
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mnemonic:", mnemonic)

	// 使用助记词和空密码生成种子
	seed := bip39.NewSeed("mimic link field aisle nut tail endorse witness business garlic mean carry churn yard narrow owner oyster fix dash canyon position nurse pond police", "")

	// 使用种子生成 BIP32 主密钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}

	// 生成 BIP32 子密钥
	childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild)
	if err != nil {
		panic(err)
	}

	// 将子密钥转换为 ECDSA 私钥
	privateKey, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		panic(err)
	}

	// 获取 ECDSA 公钥
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// 使用公钥生成波场地址
	tronAddress := publicKeyToTronAddress(publicKey)

	// 输出私钥和波场地址
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
	fmt.Println("Private Key:", privateKeyHex)
	fmt.Println("TRON Address:", tronAddress)

	//Mnemonic: mimic link field aisle nut tail endorse witness business garlic mean carry churn yard narrow owner oyster fix dash canyon position nurse pond police
	//Private Key: affd4b82b39133318eb708b62b400bb0cbbdabe7f32b4b90a734de37dc29f6b4
	//TRON Address: TS3gv9gATfVfXLiNEfmojmD8jGhvYPGnGH
	//TRON Address: TS3gv9gATfVfXLiNEfmojmD8jGhvYPGnGH
}

// 测试波场本币转账
func TestSendAndSign(t *testing.T) {
	// 从助记词恢复私钥
	mnemonicStr := "mimic link field aisle nut tail endorse witness business garlic mean carry churn yard narrow owner oyster fix dash canyon position nurse pond police"

	seed := bip39.NewSeed(mnemonicStr, "")

	// 使用种子生成 BIP32 主密钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}

	// 生成 BIP32 子密钥
	childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild)
	if err != nil {
		panic(err)
	}

	// 将子密钥转换为 ECDSA 私钥
	privateKey, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		panic(err)
	}

	// 获取 ECDSA 公钥
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// 使用公钥生成波场地址
	tronAddress := publicKeyToTronAddress(publicKey)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
	fmt.Println("Private Key:", privateKeyHex)
	fmt.Println(tronAddress)

	// 创建客户端连接到波场节点
	grpcClient := client.NewGrpcClient("grpc.trongrid.io:50051")
	if grpcClient == nil {
		log.Fatal("Failed to create grpcClient:", err)
	}

	defer grpcClient.Stop()
	err = grpcClient.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	// 获取当前账户余额
	account, err := grpcClient.GetAccount(tronAddress)
	if err != nil {
		log.Fatal("Failed to get account:", err)
	}
	fmt.Printf("Account balance: %d\n", account.Balance)

	// 构建交易 - 1 TRX = 1000000 sun 转1个TRX
	toAddress, err := common.DecodeCheck("TQ7Zft6PKXJQTH3K6oFAQ5D1m5eGyDiJz8")
	if err != nil {
		log.Fatal("Failed to decode destination address:", err, toAddress)
	}
	amount := int64(1000000) // 设置转账金额（单位：sun）
	//tx, err := grpcClient.Transfer(tronAddress, "TQ7Zft6PKXJQTH3K6oFAQ5D1m5eGyDiJz8", account.Balance)
	tx, err := grpcClient.Transfer(tronAddress, "TQ7Zft6PKXJQTH3K6oFAQ5D1m5eGyDiJz8", amount)
	if err != nil {
		log.Fatal("Failed to create transaction:", err)
	}

	// 签名交易
	signature, err := crypto.Sign(tx.Txid, privateKey)
	if err != nil {
		log.Fatal("Failed to sign transaction:", err)
	}
	tx.Transaction.Signature = append(tx.Transaction.Signature, signature)

	// 广播交易
	result, err := grpcClient.Broadcast(tx.Transaction)
	if err != nil {
		log.Fatal("Failed to broadcast transaction:", err)
	}
	fmt.Printf("Transaction %x result: %v\n", tx.Txid, result)
}

//// 测试波场本币转账
//func TestSendAndSign(t *testing.T) {
//	// 从助记词恢复私钥
//	mnemonicStr := "mimic link field aisle nut tail endorse witness business garlic mean carry churn yard narrow owner oyster fix dash canyon position nurse pond police"
//
//	seed := bip39.NewSeed(mnemonicStr, "")
//
//	// 使用种子生成 BIP32 主密钥
//	masterKey, err := bip32.NewMasterKey(seed)
//	if err != nil {
//		panic(err)
//	}
//
//	// 生成 BIP32 子密钥
//	childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild)
//	if err != nil {
//		panic(err)
//	}
//
//	// 将子密钥转换为 ECDSA 私钥
//	privateKey, err := crypto.ToECDSA(childKey.Key)
//	if err != nil {
//		panic(err)
//	}
//
//	// 获取 ECDSA 公钥
//	publicKey := privateKey.Public().(*ecdsa.PublicKey)
//
//	// 使用公钥生成波场地址
//	tronAddress := publicKeyToTronAddress(publicKey)
//
//	privateKeyBytes := crypto.FromECDSA(privateKey)
//	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
//	fmt.Println("Private Key:", privateKeyHex)
//	fmt.Println(tronAddress)
//
//	// 创建客户端连接到波场节点
//	grpcClient := client.NewGrpcClient("grpc.trongrid.io:50051")
//	if grpcClient == nil {
//		log.Fatal("Failed to create grpcClient:", err)
//	}
//
//	defer grpcClient.Stop()
//	err = grpcClient.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		return
//	}
//	// 获取当前账户余额
//	account, err := grpcClient.GetAccount(tronAddress)
//	if err != nil {
//		log.Fatal("Failed to get account:", err)
//	}
//	fmt.Printf("Account balance: %d\n", account.Balance)
//
//	// 构建交易 - 1 TRX = 1000000 sun 转1个TRX
//	toAddress, err := common.DecodeCheck("TQ7Zft6PKXJQTH3K6oFAQ5D1m5eGyDiJz8")
//	if err != nil {
//		log.Fatal("Failed to decode destination address:", err, toAddress)
//	}
//	amount := int64(1000000) // 设置转账金额（单位：sun）
//	//tx, err := grpcClient.Transfer(tronAddress, "TQ7Zft6PKXJQTH3K6oFAQ5D1m5eGyDiJz8", account.Balance)
//	tx, err := grpcClient.Transfer(tronAddress, "TQ7Zft6PKXJQTH3K6oFAQ5D1m5eGyDiJz8", amount)
//	if err != nil {
//		log.Fatal("Failed to create transaction:", err)
//	}
//
//	// 签名交易
//	signature, err := crypto.Sign(tx.Txid, privateKey)
//	if err != nil {
//		log.Fatal("Failed to sign transaction:", err)
//	}
//	tx.Transaction.Signature = append(tx.Transaction.Signature, signature)
//
//	// 广播交易
//	result, err := grpcClient.Broadcast(tx.Transaction)
//	if err != nil {
//		log.Fatal("Failed to broadcast transaction:", err)
//	}
//	fmt.Printf("Transaction %x result: %v\n", tx.Txid, result)
//}

// 测试合约调用USDT TRC20合约地址：TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t 的转账
func TestUSDTSend(t *testing.T) {
	// 从助记词恢复私钥
	mnemonicStr := "mimic link field aisle nut tail endorse witness business garlic mean carry churn yard narrow owner oyster fix dash canyon position nurse pond police"

	seed := bip39.NewSeed(mnemonicStr, "")

	// 使用种子生成 BIP32 主密钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}

	// 生成 BIP32 子密钥
	childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild)
	if err != nil {
		panic(err)
	}

	// 将子密钥转换为 ECDSA 私钥
	privateKey, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		panic(err)
	}

	// 获取 ECDSA 公钥
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// 使用公钥生成波场地址
	tronAddress := publicKeyToTronAddress(publicKey)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
	fmt.Println("Private Key:", privateKeyHex)
	fmt.Println(tronAddress)

	// 创建客户端连接到波场节点
	grpcClient := client.NewGrpcClient("grpc.trongrid.io:50051")
	if grpcClient == nil {
		log.Fatal("Failed to create grpcClient:", err)
	}

	defer grpcClient.Stop()
	err = grpcClient.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	// 获取当前账户余额
	account, err := grpcClient.GetAccount(tronAddress)
	if err != nil {
		log.Fatal("Failed to get account:", err)
	}
	fmt.Printf("Account balance: %d\n", account.Balance)

	// USDT 的合约地址
	usdtAddress := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"

	// 获取 USDT 代币余额信息
	usdtBalance, err := grpcClient.TRC20ContractBalance(tronAddress, usdtAddress)
	if err != nil {
		log.Fatal("Failed to create USDT token:", err)
	}
	fmt.Printf("USDT balance: %s\n", usdtBalance.String())

	// 构建 USDT 转账交易
	value, ok := decimals.FromString("3")
	if !ok {
		log.Fatal("Failed to parse amount")
	}
	tokenDecimals, err := grpcClient.TRC20GetDecimals(usdtAddress)
	if err != nil {
		log.Fatal("Failed to get USDT decimals:", err)
	}
	amount, _ := decimals.ApplyDecimals(value, tokenDecimals.Int64())
	txe, err := grpcClient.TRC20Send(tronAddress, "TQ7Zft6PKXJQTH3K6oFAQ5D1m5eGyDiJz8", usdtAddress, amount, 100000000) //feeLimit 100TRX
	if err != nil {
		log.Fatal("Failed to send USDT :", err)
	}

	// 签名交易
	signature, err := crypto.Sign(txe.Txid, privateKey)
	if err != nil {
		log.Fatal("Failed to sign transaction:", err)
	}
	txe.Transaction.Signature = append(txe.Transaction.Signature, signature)

	// 广播交易
	result, err := grpcClient.Broadcast(txe.Transaction)
	if err != nil {
		log.Fatal("Failed to broadcast transaction:", err)
	}
	fmt.Printf("Transaction %x result: %v\n", txe.Txid, result)
}

const (
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func doubleSHA256(input []byte) []byte {
	hasher := sha256.New()
	hasher.Write(input)
	hash := hasher.Sum(nil)

	hasher.Reset()
	hasher.Write(hash)
	hash = hasher.Sum(nil)

	return hash
}

func encodeBase58(input []byte) string {
	alphabetIndex := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(alphabet)))
	zero := big.NewInt(0)
	remainder := big.NewInt(0)
	result := ""

	for alphabetIndex.Cmp(zero) > 0 {
		alphabetIndex, remainder = alphabetIndex.DivMod(alphabetIndex, base, remainder)
		result = string(alphabet[remainder.Int64()]) + result
	}

	for _, b := range input {
		if b == 0x00 {
			result = string(alphabet[0]) + result
		} else {
			break
		}
	}

	return result
}

func TestBl(t *testing.T) {
	hexAddress := "410cb969ca690be18ff2de100355f3a11cd0b578a4"
	addressBytes, err := hex.DecodeString(hexAddress)
	if err != nil {
		panic(err)
	}

	// 计算地址的校验和
	checksum := doubleSHA256(addressBytes)[:4]

	// 添加校验和
	addressBytes = append(addressBytes, checksum...)

	// 转换为Base58
	base58Address := encodeBase58(addressBytes)

	fmt.Printf("Base58地址：%s\n", base58Address)
}

type Block struct {
	BlockID      string        `json:"blockID"`
	BlockHeader  BlockHeader   `json:"block_header"`
	Transactions []Transaction `json:"transactions"`
}

type BlockHeader struct {
	RawData          RawData `json:"raw_data"`
	WitnessSignature string  `json:"witness_signature"`
}

type RawData struct {
	Number         uint64 `json:"number"`
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witness_address"`
	ParentHash     string `json:"parentHash"`
	Version        int    `json:"version"`
	Timestamp      int64  `json:"timestamp"`
}

type Transaction struct {
	Ret        []Ret     `json:"ret"`
	Signature  []string  `json:"signature"`
	TxID       string    `json:"txID"`
	RawData    TxRawData `json:"raw_data"`
	RawDataHex string    `json:"raw_data_hex"`
}

type Ret struct {
	ContractRet string `json:"contractRet"`
}

type TxRawData struct {
	Contract      []Contract `json:"contract"`
	RefBlockBytes string     `json:"ref_block_bytes"`
	RefBlockHash  string     `json:"ref_block_hash"`
	Expiration    int64      `json:"expiration"`
	Timestamp     int64      `json:"timestamp"`
	FeeLimit      int64      `json:"fee_limit,omitempty"`
}

type Contract struct {
	Parameter Parameter `json:"parameter"`
	Type      string    `json:"type"`
}

type Parameter struct {
	Value   Value  `json:"value"`
	TypeURL string `json:"type_url"`
}

type Value struct {
	Amount          int64  `json:"amount,omitempty"`
	OwnerAddress    string `json:"owner_address,omitempty"`
	ToAddress       string `json:"to_address,omitempty"`
	AssetName       string `json:"asset_name,omitempty"`
	Data            string `json:"data,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
}

func TestBlockJson(t *testing.T) {
	jsonData := []byte(`{
    "blockID": "00000000030113d6fc39a5a7662b0e935207e94fad803b6b125ff7511a6378d9",
    "block_header": {
        "raw_data": {
            "number": 50402262,
            "txTrieRoot": "8fa3fc2a8c4a21cf91ddd3c3a455dad4ee1a0ef4c6f596bdbc30774791baa74f",
            "witness_address": "4178c842ee63b253f8f0d2955bbc582c661a078c9d",
            "parentHash": "00000000030113d5c310d69d1bbb2e093c5e865f5099911f6654910730c04c75",
            "version": 27,
            "timestamp": 1681820022000
        },
        "witness_signature": "08749962ac09326ebca887b38e868ee02b0d5159b76b65f5ac2cf27bb6e1a6ed0cbbbcefcf48cd2f250b15f36521d18380e4426f0a1e83973d781884e85cc4b301"
    },
    "transactions": [
        {
            "ret": [
                {
                    "contractRet": "SUCCESS"
                }
            ],
            "signature": [
                "da0f75dc98c23c46573f41f11b183cb21ed624433352d07efcaa3ca01219c72810b770a3286d8f4cc6522166325a7835cdf61230c7c7a3b30260cadc4305cd3c00"
            ],
            "txID": "191179a2a42dc93b09e7c50c9519dca64eaab17d441433e25a3d57ef8c8ad731",
            "raw_data": {
                "contract": [
                    {
                        "parameter": {
                            "value": {
                                "amount": 5,
                                "owner_address": "4187acffe45c0ef0bd2658f97cd689f8ea8b939163",
                                "to_address": "41370887f5a83ec09cacc7e4a510c43df87723d992"
                            },
                            "type_url": "type.googleapis.com/protocol.TransferContract"
                        },
                        "type": "TransferContract"
                    }
                ],
                "ref_block_bytes": "13c2",
                "ref_block_hash": "05b5c63cf65e5bf0",
                "expiration": 1681820076000,
                "timestamp": 1681820018907
            },
            "raw_data_hex": "0a0213c2220805b5c63cf65e5bf040e08fa7a2f9305a65080112610a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412300a154187acffe45c0ef0bd2658f97cd689f8ea8b939163121541370887f5a83ec09cacc7e4a510c43df87723d992180570dbd1a3a2f930"
        },
        {
            "ret": [
                {
                    "contractRet": "SUCCESS"
                }
            ],
            "signature": [
                "981b91e16938b393352c75b87be8b509a1ea67727da72bf6f32938170dd1cf4472292a4d94831d4a5a2204ec503dc014050943f65724ade549830c0172044d1800"
            ],
            "txID": "849f76b137c895e2d71ebbe231acc54aae9b86150aedf7048fb43ccc4f0be03b",
            "raw_data": {
                "contract": [
                    {
                        "parameter": {
                            "value": {
                                "amount": 10000,
                                "asset_name": "31303034393335",
                                "owner_address": "411c8099ed397bab0ffdb1ccbbd7a9b598e40dd497",
                                "to_address": "41ae743a1a0006dfce57bbd3ff18edf06d3a3d0d65"
                            },
                            "type_url": "type.googleapis.com/protocol.TransferAssetContract"
                        },
                        "type": "TransferAssetContract"
                    }
                ],
                "ref_block_bytes": "13c2",
                "ref_block_hash": "05b5c63cf65e5bf0",
                "expiration": 1681820076000,
                "timestamp": 1681820017945
            },
            "raw_data_hex": "0a0213c2220805b5c63cf65e5bf040e08fa7a2f9305a74080212700a32747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e736665724173736574436f6e7472616374123a0a07313030343933351215411c8099ed397bab0ffdb1ccbbd7a9b598e40dd4971a1541ae743a1a0006dfce57bbd3ff18edf06d3a3d0d6520904e7099caa3a2f930"
        },
        {
            "ret": [
                {
                    "contractRet": "SUCCESS"
                }
            ],
            "signature": [
                "03b5d1bd3b62022b67cf15bd4dee94ed8b7f0c4e40ac2301a2b66c51fab27815cc5681910a3188e936b450e539cca43c228b89a6aacc267bf7ae46f5c0b191bb00"
            ],
            "txID": "2f301dd4209a48dc6a4a282f244d1a734f535dc580de1ba94cade4da6756e2ac",
            "raw_data": {
                "contract": [
                    {
                        "parameter": {
                            "value": {
                                "amount": 10000,
                                "asset_name": "31303034393335",
                                "owner_address": "419d9e4982048cd3761c61dd2f3a3915c0ec785e9f",
                                "to_address": "416438378469c9e5bf0cf07320c3c8c20eb9424bd7"
                            },
                            "type_url": "type.googleapis.com/protocol.TransferAssetContract"
                        },
                        "type": "TransferAssetContract"
                    }
                ],
                "ref_block_bytes": "13c2",
                "ref_block_hash": "05b5c63cf65e5bf0",
                "expiration": 1681820076000,
                "timestamp": 1681820017461
            },
            "raw_data_hex": "0a0213c2220805b5c63cf65e5bf040e08fa7a2f9305a74080212700a32747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e736665724173736574436f6e7472616374123a0a07313030343933351215419d9e4982048cd3761c61dd2f3a3915c0ec785e9f1a15416438378469c9e5bf0cf07320c3c8c20eb9424bd720904e70b5c6a3a2f930"
        },
        {
            "ret": [
                {
                    "contractRet": "SUCCESS"
                }
            ],
            "signature": [
                "f652d059c26f95a53068f2985e62b991579b055eeed0c5450f8f8ab6e343aba2d04eaa66ffef5c24b34964584b4f135afb823ec11e096f4159bb4d909109c4d701"
            ],
            "txID": "c387876cb53384649302c32ebf0cf9984d7b9d68a0032c1220b417e8aa6bda8a",
            "raw_data": {
                "contract": [
                    {
                        "parameter": {
                            "value": {
                                "data": "23b872dd000000000000000000000041d530633434a0d3ffe26395d5e10738fb72175ec5000000000000000000000041817691d0553e54d21bb925e67764b5e166194ca0000000000000000000000000000000000000000000000173729a208c50900000",
                                "owner_address": "41cf7991c2268235dc2839a17aa9bec57000b2827c",
                                "contract_address": "4155459d5cc2974e618b57ace6ce92094c6ca77780"
                            },
                            "type_url": "type.googleapis.com/protocol.TriggerSmartContract"
                        },
                        "type": "TriggerSmartContract"
                    }
                ],
                "ref_block_bytes": "13c2",
                "ref_block_hash": "05b5c63cf65e5bf0",
                "expiration": 1681820076000,
                "fee_limit": 50000000,
                "timestamp": 1681820018227
            },
            "raw_data_hex": "0a0213c2220805b5c63cf65e5bf040e08fa7a2f9305acf01081f12ca010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e74726163741294010a1541cf7991c2268235dc2839a17aa9bec57000b2827c12154155459d5cc2974e618b57ace6ce92094c6ca77780226423b872dd000000000000000000000041d530633434a0d3ffe26395d5e10738fb72175ec5000000000000000000000041817691d0553e54d21bb925e67764b5e166194ca0000000000000000000000000000000000000000000000173729a208c5090000070b3cca3a2f930900180e1eb17"
        },
        {
            "ret": [
                {
                    "contractRet": "SUCCESS"
                }
            ],
            "signature": [
                "b8e481ba1cb15f8d09cfa0c5a5fda2e783255c4db0944d07d6aed3f4102565b00dba2ae0f738ecafaec11bb5408b7c4c39e600848a347142d02b2cfe4ddae7a700"
            ],
            "txID": "d1837bb550a046d23119a83385efb1e8c6bd0b384314e4493fcfd82a1fbcd1f3",
            "raw_data": {
                "contract": [
                    {
                        "parameter": {
                            "value": {
                                "amount": 1,
                                "owner_address": "410c46387f011a232264d987bbefe4d7ad1f78f454",
                                "to_address": "41f196f172b171f07e0ad96c81bd1db9d39e105a69"
                            },
                            "type_url": "type.googleapis.com/protocol.TransferContract"
                        },
                        "type": "TransferContract"
                    }
                ],
                "ref_block_bytes": "13c2",
                "ref_block_hash": "05b5c63cf65e5bf0",
                "expiration": 1681820076000,
                "timestamp": 1681820017810
            },
            "raw_data_hex": "0a0213c2220805b5c63cf65e5bf040e08fa7a2f9305a65080112610a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412300a15410c46387f011a232264d987bbefe4d7ad1f78f454121541f196f172b171f07e0ad96c81bd1db9d39e105a6918017092c9a3a2f930"
        },
        {
            "ret": [
                {
                    "contractRet": "SUCCESS"
                }
            ],
            "signature": [
                "b8e443b4574d0a9de5576e21d0467d2ffbc58da0d56e5de3d4a59713987f844b10d25e301902a997f534009a87caf096914346df06c98224949f4451b3a04adf00"
            ],
            "txID": "dd371b7578620eadeb29e9513d1ebb0bcda475bb0a754c4b6a45aadcde219052",
            "raw_data": {
                "contract": [
                    {
                        "parameter": {
                            "value": {
                                "amount": 5,
                                "owner_address": "412617cc9e204b212ff887cbe2e2f8dd9102fd7e7c",
                                "to_address": "419e6b7348467d5583ac62d5e90e077d79516f3bed"
                            },
                            "type_url": "type.googleapis.com/protocol.TransferContract"
                        },
                        "type": "TransferContract"
                    }
                ],
                "ref_block_bytes": "13c2",
                "ref_block_hash": "05b5c63cf65e5bf0",
                "expiration": 1681820076000,
                "timestamp": 1681820019005
            },
            "raw_data_hex": "0a0213c2220805b5c63cf65e5bf040e08fa7a2f9305a65080112610a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412300a15412617cc9e204b212ff887cbe2e2f8dd9102fd7e7c1215419e6b7348467d5583ac62d5e90e077d79516f3bed180570bdd2a3a2f930"
        },
        {
            "ret": [
                {
                    "contractRet": "SUCCESS"
                }
            ],
            "signature": [
                "57855c45bd1cd215eed0cba39ae8ddab17a16839224046afeb8f05ccc2beee0159149c4547a46f9b98a046f21a96f1f3ec0e766435d1831315530e031d713cfc01"
            ],
            "txID": "1c9d61cb96495475e17c92eeb30ed2bf70c3ed2bb789425ad90e4fe1f301a80a",
            "raw_data": {
                "contract": [
                    {
                        "parameter": {
                            "value": {
                                "amount": 5,
                                "owner_address": "415a9341aad5db1a65165c85064b273c08f47eaf29",
                                "to_address": "416ebe496c56374a60b1b7212c858ce008a7e106a0"
                            },
                            "type_url": "type.googleapis.com/protocol.TransferContract"
                        },
                        "type": "TransferContract"
                    }
                ],
                "ref_block_bytes": "13c2",
                "ref_block_hash": "05b5c63cf65e5bf0",
                "expiration": 1681820076000,
                "timestamp": 1681820019005
            },
            "raw_data_hex": "0a0213c2220805b5c63cf65e5bf040e08fa7a2f9305a65080112610a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412300a15415a9341aad5db1a65165c85064b273c08f47eaf291215416ebe496c56374a60b1b7212c858ce008a7e106a0180570bdd2a3a2f930"
        }]}`)
	var block Block
	err := json.Unmarshal(jsonData, &block)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("%+v\n", block)
}
