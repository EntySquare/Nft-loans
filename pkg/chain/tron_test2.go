package tron

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gorilla/mux"
//	"github.com/parnurzeal/gorequest"
//	"github.com/valyala/fasthttp"
//	"log"
//	"net/http"
//	"testing"
//)
//
//// 测试扫链
//type ResponseData struct {
//	Total             int64
//	RangeTotal        int64
//	WholeChainTxCount int64
//	ContractMap       map[string]bool
//	ContractInfo      map[string]ContractInfo
//	Data              []TransactionInfo
//}
//
//type ContractInfo struct {
//	Tag1    string `json:"tag1"`
//	Tag1Url string `json:"tag1Url"`
//	Name    string `json:"name"`
//	Vip     bool   `json:"vip"`
//}
//
//type TransactionInfo struct {
//	Block         int64
//	Hash          string
//	Timestamp     int64
//	OwnerAddress  string
//	ToAddressList []string
//	ToAddress     string
//	ContractType  int
//	Confirmed     bool
//	Revert        bool
//	ContractData  ContractData
//	SmartCalls    string
//	Events        string
//	ID            string
//	Data          string
//	Fee           string
//	ContractRet   string
//	Result        string
//	Amount        string
//	Cost          Cost
//	TokenInfo     TokenInfo
//	TokenType     string
//}
//
//type ContractData struct {
//	Amount       int64
//	AssetName    string
//	OwnerAddress string
//	ToAddress    string
//	TokenInfo    TokenInfo
//}
//
//type TokenInfo struct {
//	TokenID      string
//	TokenAbbr    string
//	TokenName    string
//	TokenDecimal int
//	TokenCanShow int
//	TokenType    string
//	TokenLogo    string
//	TokenLevel   string
//	Vip          bool
//}
//
//type Cost struct {
//	NetFee             int
//	EnergyPenaltyTotal int
//	EnergyUsage        int
//	Fee                int
//	EnergyFee          int
//	EnergyUsageTotal   int
//	OriginEnergyUsage  int
//	NetUsage           int
//}
//
//func getBlockTransactions(blockNumber int64) []TransactionInfo {
//	url := fmt.Sprintf("https://apilist.tronscan.org/api/transaction?sort=-timestamp&count=true&start=0&limit=300&block=%d", blockNumber)
//
//	req := fasthttp.AcquireRequest()
//	resp := fasthttp.AcquireResponse()
//	defer fasthttp.ReleaseRequest(req)
//	defer fasthttp.ReleaseResponse(resp)
//
//	req.SetRequestURI(url)
//	req.Header.SetMethod("GET")
//
//	if err := fasthttp.Do(req, resp); err != nil {
//		panic("Failed to get block transactions")
//	}
//
//	bodyBytes := resp.Body()
//	var transactionData ResponseData
//	if err := json.Unmarshal(bodyBytes, &transactionData); err != nil {
//		panic("Failed to unmarshal block transactions")
//	}
//
//	return transactionData.Data
//}
//
////func findTransactionsWithAddress(blockNumber int64, targetAddress string) {
////	transactions := getBlockTransactions(blockNumber)
////
////	for _, tx := range transactions {
////		if tx.From == targetAddress || tx.To == targetAddress {
////			fmt.Printf("Transaction ID: %s\nFrom: %s\nTo: %s\nValue: %d\n\n", tx.TxID, tx.From, tx.To, tx.Value)
////		}
////	}
////}
//
//func TestScan(t *testing.T) {
//	blockNumber := int64(50157874)
//	//targetAddress := "TL3QBVVzyXkowKcPd5W685MMBndrMkGm59"
//
//	getBlockTransactions(blockNumber)
//}
//
//const apiURL = "https://api.trongrid.io/v1/accounts/{address}/transactions"
//
//type Transaction struct {
//	ID       string `json:"txID"`
//	Contract []struct {
//		Type    string `json:"type"`
//		From    string `json:"from"`
//		To      string `json:"to"`
//		Value   int64  `json:"value"`
//		Address string `json:"contract_address"`
//	} `json:"contract"`
//}
//
//type ResponseData2 struct {
//	Data []Transaction `json:"data"`
//}
//
//func TestBlockScan(t *testing.T) {
//	router := mux.NewRouter()
//	router.HandleFunc("/transactions/{address}", handleRequest)
//	log.Fatal(http.ListenAndServe(":8080", router))
//}
//
//func handleRequest(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	address := vars["address"]
//
//	response, _, errs := gorequest.New().Get(apiURL).Query("address=" + address).End()
//	if errs != nil {
//		http.Error(w, "Error getting transactions", http.StatusInternalServerError)
//		return
//	}
//
//	var responseData ResponseData2
//	err := json.NewDecoder(response.Body).Decode(&responseData)
//	if err != nil {
//		http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
//		return
//	}
//
//	usdtContractAddress := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
//
//	for _, tx := range responseData.Data {
//		for _, contract := range tx.Contract {
//			if contract.Type == "TriggerSmartContract" && contract.Address == usdtContractAddress {
//				fmt.Printf("From: %s\nTo: %s\nUSDT Transfer Value: %d\n", contract.From, contract.To, contract.Value)
//			} else {
//				fmt.Printf("From: %s\nTo: %s\nValue: %d\n", contract.From, contract.To, contract.Value)
//			}
//		}
//	}
//}
