// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NGTTokenData is an auto generated low-level Go binding around an user-defined struct.
type NGTTokenData struct {
	TokenId      *big.Int
	ReceivedTime *big.Int
	Flag         *big.Int
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fund\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ChangeMarketStatusEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedTime\",\"type\":\"uint256\"}],\"name\":\"TokenReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"loansCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"loansList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flag\",\"type\":\"uint256\"}],\"internalType\":\"structNGT.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftContract\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receivedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flag\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewtotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Contracts *ContractsCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Contracts *ContractsSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Contracts *ContractsCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Contracts *ContractsCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Contracts *ContractsSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Contracts *ContractsCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCallerSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// LoansCount is a free data retrieval call binding the contract method 0xffc49f38.
//
// Solidity: function loansCount(address user) view returns(uint256)
func (_Contracts *ContractsCaller) LoansCount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "loansCount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoansCount is a free data retrieval call binding the contract method 0xffc49f38.
//
// Solidity: function loansCount(address user) view returns(uint256)
func (_Contracts *ContractsSession) LoansCount(user common.Address) (*big.Int, error) {
	return _Contracts.Contract.LoansCount(&_Contracts.CallOpts, user)
}

// LoansCount is a free data retrieval call binding the contract method 0xffc49f38.
//
// Solidity: function loansCount(address user) view returns(uint256)
func (_Contracts *ContractsCallerSession) LoansCount(user common.Address) (*big.Int, error) {
	return _Contracts.Contract.LoansCount(&_Contracts.CallOpts, user)
}

// LoansList is a free data retrieval call binding the contract method 0x27d88a78.
//
// Solidity: function loansList(address user) view returns((uint256,uint256,uint256)[])
func (_Contracts *ContractsCaller) LoansList(opts *bind.CallOpts, user common.Address) ([]NGTTokenData, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "loansList", user)

	if err != nil {
		return *new([]NGTTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]NGTTokenData)).(*[]NGTTokenData)

	return out0, err

}

// LoansList is a free data retrieval call binding the contract method 0x27d88a78.
//
// Solidity: function loansList(address user) view returns((uint256,uint256,uint256)[])
func (_Contracts *ContractsSession) LoansList(user common.Address) ([]NGTTokenData, error) {
	return _Contracts.Contract.LoansList(&_Contracts.CallOpts, user)
}

// LoansList is a free data retrieval call binding the contract method 0x27d88a78.
//
// Solidity: function loansList(address user) view returns((uint256,uint256,uint256)[])
func (_Contracts *ContractsCallerSession) LoansList(user common.Address) ([]NGTTokenData, error) {
	return _Contracts.Contract.LoansList(&_Contracts.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_Contracts *ContractsCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_Contracts *ContractsSession) NftContract() (common.Address, error) {
	return _Contracts.Contract.NftContract(&_Contracts.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_Contracts *ContractsCallerSession) NftContract() (common.Address, error) {
	return _Contracts.Contract.NftContract(&_Contracts.CallOpts)
}

// ReceivedTokens is a free data retrieval call binding the contract method 0xf8fe54fa.
//
// Solidity: function receivedTokens(address , uint256 ) view returns(uint256 tokenId, uint256 receivedTime, uint256 flag)
func (_Contracts *ContractsCaller) ReceivedTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	TokenId      *big.Int
	ReceivedTime *big.Int
	Flag         *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "receivedTokens", arg0, arg1)

	outstruct := new(struct {
		TokenId      *big.Int
		ReceivedTime *big.Int
		Flag         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReceivedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Flag = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReceivedTokens is a free data retrieval call binding the contract method 0xf8fe54fa.
//
// Solidity: function receivedTokens(address , uint256 ) view returns(uint256 tokenId, uint256 receivedTime, uint256 flag)
func (_Contracts *ContractsSession) ReceivedTokens(arg0 common.Address, arg1 *big.Int) (struct {
	TokenId      *big.Int
	ReceivedTime *big.Int
	Flag         *big.Int
}, error) {
	return _Contracts.Contract.ReceivedTokens(&_Contracts.CallOpts, arg0, arg1)
}

// ReceivedTokens is a free data retrieval call binding the contract method 0xf8fe54fa.
//
// Solidity: function receivedTokens(address , uint256 ) view returns(uint256 tokenId, uint256 receivedTime, uint256 flag)
func (_Contracts *ContractsCallerSession) ReceivedTokens(arg0 common.Address, arg1 *big.Int) (struct {
	TokenId      *big.Int
	ReceivedTime *big.Int
	Flag         *big.Int
}, error) {
	return _Contracts.Contract.ReceivedTokens(&_Contracts.CallOpts, arg0, arg1)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// ViewtotalSupply is a free data retrieval call binding the contract method 0x8951cf59.
//
// Solidity: function viewtotalSupply() view returns(uint256)
func (_Contracts *ContractsCaller) ViewtotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "viewtotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewtotalSupply is a free data retrieval call binding the contract method 0x8951cf59.
//
// Solidity: function viewtotalSupply() view returns(uint256)
func (_Contracts *ContractsSession) ViewtotalSupply() (*big.Int, error) {
	return _Contracts.Contract.ViewtotalSupply(&_Contracts.CallOpts)
}

// ViewtotalSupply is a free data retrieval call binding the contract method 0x8951cf59.
//
// Solidity: function viewtotalSupply() view returns(uint256)
func (_Contracts *ContractsCallerSession) ViewtotalSupply() (*big.Int, error) {
	return _Contracts.Contract.ViewtotalSupply(&_Contracts.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Contracts *ContractsTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Contracts *ContractsSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Contracts *ContractsTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, _spender, _value)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Contracts *ContractsTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Contracts *ContractsSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnERC721Received(&_Contracts.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Contracts *ContractsTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnERC721Received(&_Contracts.TransactOpts, operator, from, tokenId, data)
}

// OwnerOf is a paid mutator transaction binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) returns(address)
func (_Contracts *ContractsTransactor) OwnerOf(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "ownerOf", _tokenId)
}

// OwnerOf is a paid mutator transaction binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) returns(address)
func (_Contracts *ContractsSession) OwnerOf(_tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.OwnerOf(&_Contracts.TransactOpts, _tokenId)
}

// OwnerOf is a paid mutator transaction binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) returns(address)
func (_Contracts *ContractsTransactorSession) OwnerOf(_tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.OwnerOf(&_Contracts.TransactOpts, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Contracts *ContractsTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Contracts *ContractsSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Contracts *ContractsTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Contracts *ContractsTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Contracts *ContractsSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Contracts *ContractsTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, _from, _to, _value)
}

// WithdrawNFT is a paid mutator transaction binding the contract method 0x6088e93a.
//
// Solidity: function withdrawNFT(address to, uint256 _tokenId) payable returns()
func (_Contracts *ContractsTransactor) WithdrawNFT(opts *bind.TransactOpts, to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawNFT", to, _tokenId)
}

// WithdrawNFT is a paid mutator transaction binding the contract method 0x6088e93a.
//
// Solidity: function withdrawNFT(address to, uint256 _tokenId) payable returns()
func (_Contracts *ContractsSession) WithdrawNFT(to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawNFT(&_Contracts.TransactOpts, to, _tokenId)
}

// WithdrawNFT is a paid mutator transaction binding the contract method 0x6088e93a.
//
// Solidity: function withdrawNFT(address to, uint256 _tokenId) payable returns()
func (_Contracts *ContractsTransactorSession) WithdrawNFT(to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawNFT(&_Contracts.TransactOpts, to, _tokenId)
}

// ContractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contracts contract.
type ContractsApprovalIterator struct {
	Event *ContractsApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApproval represents a Approval event raised by the Contracts contract.
type ContractsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_Contracts *ContractsFilterer) FilterApproval(opts *bind.FilterOpts) (*ContractsApprovalIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalIterator{contract: _Contracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_Contracts *ContractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractsApproval) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApproval)
				if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_Contracts *ContractsFilterer) ParseApproval(log types.Log) (*ContractsApproval, error) {
	event := new(ContractsApproval)
	if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsChangeMarketStatusEventIterator is returned from FilterChangeMarketStatusEvent and is used to iterate over the raw logs and unpacked data for ChangeMarketStatusEvent events raised by the Contracts contract.
type ContractsChangeMarketStatusEventIterator struct {
	Event *ContractsChangeMarketStatusEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsChangeMarketStatusEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsChangeMarketStatusEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsChangeMarketStatusEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsChangeMarketStatusEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsChangeMarketStatusEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsChangeMarketStatusEvent represents a ChangeMarketStatusEvent event raised by the Contracts contract.
type ContractsChangeMarketStatusEvent struct {
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChangeMarketStatusEvent is a free log retrieval operation binding the contract event 0x8853484ae4a06c3395f689a3fa189b683f6542d0f6af1c3aa07b25e219cd08da.
//
// Solidity: event ChangeMarketStatusEvent(uint8 status)
func (_Contracts *ContractsFilterer) FilterChangeMarketStatusEvent(opts *bind.FilterOpts) (*ContractsChangeMarketStatusEventIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ChangeMarketStatusEvent")
	if err != nil {
		return nil, err
	}
	return &ContractsChangeMarketStatusEventIterator{contract: _Contracts.contract, event: "ChangeMarketStatusEvent", logs: logs, sub: sub}, nil
}

// WatchChangeMarketStatusEvent is a free log subscription operation binding the contract event 0x8853484ae4a06c3395f689a3fa189b683f6542d0f6af1c3aa07b25e219cd08da.
//
// Solidity: event ChangeMarketStatusEvent(uint8 status)
func (_Contracts *ContractsFilterer) WatchChangeMarketStatusEvent(opts *bind.WatchOpts, sink chan<- *ContractsChangeMarketStatusEvent) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ChangeMarketStatusEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsChangeMarketStatusEvent)
				if err := _Contracts.contract.UnpackLog(event, "ChangeMarketStatusEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangeMarketStatusEvent is a log parse operation binding the contract event 0x8853484ae4a06c3395f689a3fa189b683f6542d0f6af1c3aa07b25e219cd08da.
//
// Solidity: event ChangeMarketStatusEvent(uint8 status)
func (_Contracts *ContractsFilterer) ParseChangeMarketStatusEvent(log types.Log) (*ContractsChangeMarketStatusEvent, error) {
	event := new(ContractsChangeMarketStatusEvent)
	if err := _Contracts.contract.UnpackLog(event, "ChangeMarketStatusEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTokenReceivedIterator is returned from FilterTokenReceived and is used to iterate over the raw logs and unpacked data for TokenReceived events raised by the Contracts contract.
type ContractsTokenReceivedIterator struct {
	Event *ContractsTokenReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsTokenReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTokenReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsTokenReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsTokenReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTokenReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTokenReceived represents a TokenReceived event raised by the Contracts contract.
type ContractsTokenReceived struct {
	From         common.Address
	TokenId      *big.Int
	ReceivedTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenReceived is a free log retrieval operation binding the contract event 0xcbd44b327242a19152f06fd696fc88ca79d777e3de08a110708838867427bd51.
//
// Solidity: event TokenReceived(address indexed from, uint256 indexed tokenId, uint256 receivedTime)
func (_Contracts *ContractsFilterer) FilterTokenReceived(opts *bind.FilterOpts, from []common.Address, tokenId []*big.Int) (*ContractsTokenReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TokenReceived", fromRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTokenReceivedIterator{contract: _Contracts.contract, event: "TokenReceived", logs: logs, sub: sub}, nil
}

// WatchTokenReceived is a free log subscription operation binding the contract event 0xcbd44b327242a19152f06fd696fc88ca79d777e3de08a110708838867427bd51.
//
// Solidity: event TokenReceived(address indexed from, uint256 indexed tokenId, uint256 receivedTime)
func (_Contracts *ContractsFilterer) WatchTokenReceived(opts *bind.WatchOpts, sink chan<- *ContractsTokenReceived, from []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "TokenReceived", fromRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTokenReceived)
				if err := _Contracts.contract.UnpackLog(event, "TokenReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenReceived is a log parse operation binding the contract event 0xcbd44b327242a19152f06fd696fc88ca79d777e3de08a110708838867427bd51.
//
// Solidity: event TokenReceived(address indexed from, uint256 indexed tokenId, uint256 receivedTime)
func (_Contracts *ContractsFilterer) ParseTokenReceived(log types.Log) (*ContractsTokenReceived, error) {
	event := new(ContractsTokenReceived)
	if err := _Contracts.contract.UnpackLog(event, "TokenReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contracts contract.
type ContractsTransferIterator struct {
	Event *ContractsTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer represents a Transfer event raised by the Contracts contract.
type ContractsTransfer struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_Contracts *ContractsFilterer) FilterTransfer(opts *bind.FilterOpts) (*ContractsTransferIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &ContractsTransferIterator{contract: _Contracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_Contracts *ContractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractsTransfer) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer)
				if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_Contracts *ContractsFilterer) ParseTransfer(log types.Log) (*ContractsTransfer, error) {
	event := new(ContractsTransfer)
	if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
