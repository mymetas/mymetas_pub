// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testerc20token

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
)

// Testerc20tokenMetaData contains all meta data concerning the Testerc20token contract.
var Testerc20tokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimalUnits\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620009a6380380620009a6833981016040819052620000349162000203565b336000908152600160209081526040822086905590859055835162000060916003919086019062000090565b506004805460ff191660ff841617905580516200008590600590602084019062000090565b5050505050620002ce565b8280546200009e9062000292565b90600052602060002090601f016020900481019282620000c257600085556200010d565b82601f10620000dd57805160ff19168380011785556200010d565b828001600101855582156200010d579182015b828111156200010d578251825591602001919060010190620000f0565b506200011b9291506200011f565b5090565b5b808211156200011b576000815560010162000120565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200015e57600080fd5b81516001600160401b03808211156200017b576200017b62000136565b604051601f8301601f19908116603f01168101908282118183101715620001a657620001a662000136565b81604052838152602092508683858801011115620001c357600080fd5b600091505b83821015620001e75785820183015181830184015290820190620001c8565b83821115620001f95760008385830101525b9695505050505050565b600080600080608085870312156200021a57600080fd5b845160208601519094506001600160401b03808211156200023a57600080fd5b62000248888389016200014c565b94506040870151915060ff821682146200026157600080fd5b6060870151919350808211156200027757600080fd5b5062000286878288016200014c565b91505092959194509250565b600181811c90821680620002a757607f821691505b602082108103620002c857634e487b7160e01b600052602260045260246000fd5b50919050565b6106c880620002de6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063313ce56711610071578063313ce567146101395780635c6581651461015857806370a082311461018357806395d89b41146101ac578063a9059cbb146101b4578063dd62ed3e146101c757600080fd5b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd1461010657806327e235e314610119575b600080fd5b6100b6610200565b6040516100c391906104e7565b60405180910390f35b6100df6100da366004610558565b61028e565b60405190151581526020016100c3565b6100f860005481565b6040519081526020016100c3565b6100df610114366004610582565b6102fa565b6100f86101273660046105be565b60016020526000908152604090205481565b6004546101469060ff1681565b60405160ff90911681526020016100c3565b6100f86101663660046105e0565b600260209081526000928352604080842090915290825290205481565b6100f86101913660046105be565b6001600160a01b031660009081526001602052604090205490565b6100b6610430565b6100df6101c2366004610558565b61043d565b6100f86101d53660046105e0565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6003805461020d90610613565b80601f016020809104026020016040519081016040528092919081815260200182805461023990610613565b80156102865780601f1061025b57610100808354040283529160200191610286565b820191906000526020600020905b81548152906001019060200180831161026957829003601f168201915b505050505081565b3360008181526002602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906102e99086815260200190565b60405180910390a350600192915050565b6001600160a01b03831660008181526002602090815260408083203384528252808320549383526001909152812054909190831180159061033b5750828110155b61034457600080fd5b6001600160a01b0384166000908152600160205260408120805485929061036c908490610663565b90915550506001600160a01b0385166000908152600160205260408120805485929061039990849061067b565b90915550506001600160a01b0385166000908152600260209081526040808320338452909152812080548592906103d190849061067b565b92505081905550836001600160a01b0316856001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405161041d91815260200190565b60405180910390a3506001949350505050565b6005805461020d90610613565b3360009081526001602052604081205482111561045957600080fd5b336000908152600160205260408120805484929061047890849061067b565b90915550506001600160a01b038316600090815260016020526040812080548492906104a5908490610663565b90915550506040518281526001600160a01b0384169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016102e9565b600060208083528351808285015260005b81811015610514578581018301518582016040015282016104f8565b81811115610526576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461055357600080fd5b919050565b6000806040838503121561056b57600080fd5b6105748361053c565b946020939093013593505050565b60008060006060848603121561059757600080fd5b6105a08461053c565b92506105ae6020850161053c565b9150604084013590509250925092565b6000602082840312156105d057600080fd5b6105d98261053c565b9392505050565b600080604083850312156105f357600080fd5b6105fc8361053c565b915061060a6020840161053c565b90509250929050565b600181811c9082168061062757607f821691505b60208210810361064757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b600082198211156106765761067661064d565b500190565b60008282101561068d5761068d61064d565b50039056fea2646970667358221220a16035c31a4bc3ef40f469aacfa5da322d94c661ba74891a643728f686f8620c64736f6c634300080d0033",
}

// Testerc20tokenABI is the input ABI used to generate the binding from.
// Deprecated: Use Testerc20tokenMetaData.ABI instead.
var Testerc20tokenABI = Testerc20tokenMetaData.ABI

// Testerc20tokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Testerc20tokenMetaData.Bin instead.
var Testerc20tokenBin = Testerc20tokenMetaData.Bin

// DeployTesterc20token deploys a new Ethereum contract, binding an instance of Testerc20token to it.
func DeployTesterc20token(auth *bind.TransactOpts, backend bind.ContractBackend, _initialAmount *big.Int, _tokenName string, _decimalUnits uint8, _tokenSymbol string) (common.Address, *types.Transaction, *Testerc20token, error) {
	parsed, err := Testerc20tokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Testerc20tokenBin), backend, _initialAmount, _tokenName, _decimalUnits, _tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Testerc20token{Testerc20tokenCaller: Testerc20tokenCaller{contract: contract}, Testerc20tokenTransactor: Testerc20tokenTransactor{contract: contract}, Testerc20tokenFilterer: Testerc20tokenFilterer{contract: contract}}, nil
}

// Testerc20token is an auto generated Go binding around an Ethereum contract.
type Testerc20token struct {
	Testerc20tokenCaller     // Read-only binding to the contract
	Testerc20tokenTransactor // Write-only binding to the contract
	Testerc20tokenFilterer   // Log filterer for contract events
}

// Testerc20tokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type Testerc20tokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Testerc20tokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Testerc20tokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Testerc20tokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Testerc20tokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Testerc20tokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Testerc20tokenSession struct {
	Contract     *Testerc20token   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Testerc20tokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Testerc20tokenCallerSession struct {
	Contract *Testerc20tokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Testerc20tokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Testerc20tokenTransactorSession struct {
	Contract     *Testerc20tokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Testerc20tokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type Testerc20tokenRaw struct {
	Contract *Testerc20token // Generic contract binding to access the raw methods on
}

// Testerc20tokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Testerc20tokenCallerRaw struct {
	Contract *Testerc20tokenCaller // Generic read-only contract binding to access the raw methods on
}

// Testerc20tokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Testerc20tokenTransactorRaw struct {
	Contract *Testerc20tokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTesterc20token creates a new instance of Testerc20token, bound to a specific deployed contract.
func NewTesterc20token(address common.Address, backend bind.ContractBackend) (*Testerc20token, error) {
	contract, err := bindTesterc20token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testerc20token{Testerc20tokenCaller: Testerc20tokenCaller{contract: contract}, Testerc20tokenTransactor: Testerc20tokenTransactor{contract: contract}, Testerc20tokenFilterer: Testerc20tokenFilterer{contract: contract}}, nil
}

// NewTesterc20tokenCaller creates a new read-only instance of Testerc20token, bound to a specific deployed contract.
func NewTesterc20tokenCaller(address common.Address, caller bind.ContractCaller) (*Testerc20tokenCaller, error) {
	contract, err := bindTesterc20token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Testerc20tokenCaller{contract: contract}, nil
}

// NewTesterc20tokenTransactor creates a new write-only instance of Testerc20token, bound to a specific deployed contract.
func NewTesterc20tokenTransactor(address common.Address, transactor bind.ContractTransactor) (*Testerc20tokenTransactor, error) {
	contract, err := bindTesterc20token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Testerc20tokenTransactor{contract: contract}, nil
}

// NewTesterc20tokenFilterer creates a new log filterer instance of Testerc20token, bound to a specific deployed contract.
func NewTesterc20tokenFilterer(address common.Address, filterer bind.ContractFilterer) (*Testerc20tokenFilterer, error) {
	contract, err := bindTesterc20token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Testerc20tokenFilterer{contract: contract}, nil
}

// bindTesterc20token binds a generic wrapper to an already deployed contract.
func bindTesterc20token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Testerc20tokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testerc20token *Testerc20tokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testerc20token.Contract.Testerc20tokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testerc20token *Testerc20tokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testerc20token.Contract.Testerc20tokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testerc20token *Testerc20tokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testerc20token.Contract.Testerc20tokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testerc20token *Testerc20tokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testerc20token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testerc20token *Testerc20tokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testerc20token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testerc20token *Testerc20tokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testerc20token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Testerc20token *Testerc20tokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Testerc20token.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Testerc20token *Testerc20tokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Testerc20token.Contract.Allowance(&_Testerc20token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Testerc20token *Testerc20tokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Testerc20token.Contract.Allowance(&_Testerc20token.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Testerc20token *Testerc20tokenCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Testerc20token.contract.Call(opts, &out, "allowed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Testerc20token *Testerc20tokenSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Testerc20token.Contract.Allowed(&_Testerc20token.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Testerc20token *Testerc20tokenCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Testerc20token.Contract.Allowed(&_Testerc20token.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Testerc20token *Testerc20tokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Testerc20token.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Testerc20token *Testerc20tokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Testerc20token.Contract.BalanceOf(&_Testerc20token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Testerc20token *Testerc20tokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Testerc20token.Contract.BalanceOf(&_Testerc20token.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Testerc20token *Testerc20tokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Testerc20token.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Testerc20token *Testerc20tokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Testerc20token.Contract.Balances(&_Testerc20token.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Testerc20token *Testerc20tokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Testerc20token.Contract.Balances(&_Testerc20token.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Testerc20token *Testerc20tokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Testerc20token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Testerc20token *Testerc20tokenSession) Decimals() (uint8, error) {
	return _Testerc20token.Contract.Decimals(&_Testerc20token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Testerc20token *Testerc20tokenCallerSession) Decimals() (uint8, error) {
	return _Testerc20token.Contract.Decimals(&_Testerc20token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Testerc20token *Testerc20tokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Testerc20token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Testerc20token *Testerc20tokenSession) Name() (string, error) {
	return _Testerc20token.Contract.Name(&_Testerc20token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Testerc20token *Testerc20tokenCallerSession) Name() (string, error) {
	return _Testerc20token.Contract.Name(&_Testerc20token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Testerc20token *Testerc20tokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Testerc20token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Testerc20token *Testerc20tokenSession) Symbol() (string, error) {
	return _Testerc20token.Contract.Symbol(&_Testerc20token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Testerc20token *Testerc20tokenCallerSession) Symbol() (string, error) {
	return _Testerc20token.Contract.Symbol(&_Testerc20token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Testerc20token *Testerc20tokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Testerc20token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Testerc20token *Testerc20tokenSession) TotalSupply() (*big.Int, error) {
	return _Testerc20token.Contract.TotalSupply(&_Testerc20token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Testerc20token *Testerc20tokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Testerc20token.Contract.TotalSupply(&_Testerc20token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Testerc20token *Testerc20tokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Testerc20token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Testerc20token *Testerc20tokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Testerc20token.Contract.Approve(&_Testerc20token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Testerc20token *Testerc20tokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Testerc20token.Contract.Approve(&_Testerc20token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Testerc20token *Testerc20tokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Testerc20token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Testerc20token *Testerc20tokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Testerc20token.Contract.Transfer(&_Testerc20token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Testerc20token *Testerc20tokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Testerc20token.Contract.Transfer(&_Testerc20token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Testerc20token *Testerc20tokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Testerc20token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Testerc20token *Testerc20tokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Testerc20token.Contract.TransferFrom(&_Testerc20token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Testerc20token *Testerc20tokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Testerc20token.Contract.TransferFrom(&_Testerc20token.TransactOpts, _from, _to, _value)
}

// Testerc20tokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Testerc20token contract.
type Testerc20tokenApprovalIterator struct {
	Event *Testerc20tokenApproval // Event containing the contract specifics and raw log

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
func (it *Testerc20tokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Testerc20tokenApproval)
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
		it.Event = new(Testerc20tokenApproval)
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
func (it *Testerc20tokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Testerc20tokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Testerc20tokenApproval represents a Approval event raised by the Testerc20token contract.
type Testerc20tokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Testerc20token *Testerc20tokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*Testerc20tokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Testerc20token.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &Testerc20tokenApprovalIterator{contract: _Testerc20token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Testerc20token *Testerc20tokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Testerc20tokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Testerc20token.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Testerc20tokenApproval)
				if err := _Testerc20token.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Testerc20token *Testerc20tokenFilterer) ParseApproval(log types.Log) (*Testerc20tokenApproval, error) {
	event := new(Testerc20tokenApproval)
	if err := _Testerc20token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Testerc20tokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Testerc20token contract.
type Testerc20tokenTransferIterator struct {
	Event *Testerc20tokenTransfer // Event containing the contract specifics and raw log

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
func (it *Testerc20tokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Testerc20tokenTransfer)
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
		it.Event = new(Testerc20tokenTransfer)
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
func (it *Testerc20tokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Testerc20tokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Testerc20tokenTransfer represents a Transfer event raised by the Testerc20token contract.
type Testerc20tokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Testerc20token *Testerc20tokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*Testerc20tokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Testerc20token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &Testerc20tokenTransferIterator{contract: _Testerc20token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Testerc20token *Testerc20tokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Testerc20tokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Testerc20token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Testerc20tokenTransfer)
				if err := _Testerc20token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Testerc20token *Testerc20tokenFilterer) ParseTransfer(log types.Log) (*Testerc20tokenTransfer, error) {
	event := new(Testerc20tokenTransfer)
	if err := _Testerc20token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
