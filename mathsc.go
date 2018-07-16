// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_val\",\"type\":\"int256\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"int256\"},{\"name\":\"_b\",\"type\":\"int256\"}],\"name\":\"sum\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x608060405234801561001057600080fd5b506000805560ed806100236000396000f30060806040526004361060525763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632ddbd13a8114605757806387db03b714607b5780639967062d146090575b600080fd5b348015606257600080fd5b50606960a8565b60408051918252519081900360200190f35b348015608657600080fd5b50606960043560ae565b348015609b57600080fd5b50606960043560243560bd565b60005481565b60008054820190819055919050565b01905600a165627a7a723058200c8cc82516ae032f2f523f0f80c20c3118e19f639def23dca35fa8ca8de939960029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// Sum is a free data retrieval call binding the contract method 0x9967062d.
//
// Solidity: function sum(_a int256, _b int256) constant returns(int256)
func (_Math *MathCaller) Sum(opts *bind.CallOpts, _a *big.Int, _b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Math.contract.Call(opts, out, "sum", _a, _b)
	return *ret0, err
}

// Sum is a free data retrieval call binding the contract method 0x9967062d.
//
// Solidity: function sum(_a int256, _b int256) constant returns(int256)
func (_Math *MathSession) Sum(_a *big.Int, _b *big.Int) (*big.Int, error) {
	return _Math.Contract.Sum(&_Math.CallOpts, _a, _b)
}

// Sum is a free data retrieval call binding the contract method 0x9967062d.
//
// Solidity: function sum(_a int256, _b int256) constant returns(int256)
func (_Math *MathCallerSession) Sum(_a *big.Int, _b *big.Int) (*big.Int, error) {
	return _Math.Contract.Sum(&_Math.CallOpts, _a, _b)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() constant returns(int256)
func (_Math *MathCaller) Total(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Math.contract.Call(opts, out, "total")
	return *ret0, err
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() constant returns(int256)
func (_Math *MathSession) Total() (*big.Int, error) {
	return _Math.Contract.Total(&_Math.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() constant returns(int256)
func (_Math *MathCallerSession) Total() (*big.Int, error) {
	return _Math.Contract.Total(&_Math.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x87db03b7.
//
// Solidity: function add(_val int256) returns(int256)
func (_Math *MathTransactor) Add(opts *bind.TransactOpts, _val *big.Int) (*types.Transaction, error) {
	return _Math.contract.Transact(opts, "add", _val)
}

// Add is a paid mutator transaction binding the contract method 0x87db03b7.
//
// Solidity: function add(_val int256) returns(int256)
func (_Math *MathSession) Add(_val *big.Int) (*types.Transaction, error) {
	return _Math.Contract.Add(&_Math.TransactOpts, _val)
}

// Add is a paid mutator transaction binding the contract method 0x87db03b7.
//
// Solidity: function add(_val int256) returns(int256)
func (_Math *MathTransactorSession) Add(_val *big.Int) (*types.Transaction, error) {
	return _Math.Contract.Add(&_Math.TransactOpts, _val)
}
