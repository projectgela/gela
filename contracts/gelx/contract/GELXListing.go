// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/projectgela/gela/accounts/abi"
	"github.com/projectgela/gela/accounts/abi/bind"
	"github.com/projectgela/gela/common"
	"github.com/projectgela/gela/core/types"
)

// GELXListingABI is the input ABI used to generate the binding from.
const GELXListingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// GELXListingBin is the compiled bytecode used for deploying new contracts.
const GELXListingBin = `0x608060405234801561001057600080fd5b506102be806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639d63848a811461005b578063a3ff31b5146100c0578063c6b32f34146100f5575b600080fd5b34801561006757600080fd5b5061007061010b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ac578181015183820152602001610094565b505050509050019250505060405180910390f35b3480156100cc57600080fd5b506100e1600160a060020a036004351661016d565b604080519115158252519081900360200190f35b610109600160a060020a036004351661018b565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561016357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610145575b5050505050905090565b600160a060020a031660009081526001602052604090205460ff1690565b80600160a060020a03811615156101a157600080fd5b600160a060020a03811660009081526001602081905260409091205460ff16151514156101cd57600080fd5b683635c9adc5dea0000034146101e257600080fd5b6040516068903480156108fc02916000818181858888f1935050505015801561020f573d6000803e3d6000fd5b505060008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556040805160208082018352838252948452919093529190209051815460ff19169015151790555600a165627a7a723058206d2dc0ce827743c25efa82f99e7830ade39d28e17f4d651573f89e0460a6626a0029`

// DeployGELXListing deploys a new Ethereum contract, binding an instance of GELXListing to it.
func DeployGELXListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GELXListing, error) {
	parsed, err := abi.JSON(strings.NewReader(GELXListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GELXListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GELXListing{GELXListingCaller: GELXListingCaller{contract: contract}, GELXListingTransactor: GELXListingTransactor{contract: contract}, GELXListingFilterer: GELXListingFilterer{contract: contract}}, nil
}

// GELXListing is an auto generated Go binding around an Ethereum contract.
type GELXListing struct {
	GELXListingCaller     // Read-only binding to the contract
	GELXListingTransactor // Write-only binding to the contract
	GELXListingFilterer   // Log filterer for contract events
}

// GELXListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type GELXListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GELXListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GELXListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GELXListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GELXListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GELXListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GELXListingSession struct {
	Contract     *GELXListing     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GELXListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GELXListingCallerSession struct {
	Contract *GELXListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GELXListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GELXListingTransactorSession struct {
	Contract     *GELXListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GELXListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type GELXListingRaw struct {
	Contract *GELXListing // Generic contract binding to access the raw methods on
}

// GELXListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GELXListingCallerRaw struct {
	Contract *GELXListingCaller // Generic read-only contract binding to access the raw methods on
}

// GELXListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GELXListingTransactorRaw struct {
	Contract *GELXListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGELXListing creates a new instance of GELXListing, bound to a specific deployed contract.
func NewGELXListing(address common.Address, backend bind.ContractBackend) (*GELXListing, error) {
	contract, err := bindGELXListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GELXListing{GELXListingCaller: GELXListingCaller{contract: contract}, GELXListingTransactor: GELXListingTransactor{contract: contract}, GELXListingFilterer: GELXListingFilterer{contract: contract}}, nil
}

// NewGELXListingCaller creates a new read-only instance of GELXListing, bound to a specific deployed contract.
func NewGELXListingCaller(address common.Address, caller bind.ContractCaller) (*GELXListingCaller, error) {
	contract, err := bindGELXListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GELXListingCaller{contract: contract}, nil
}

// NewGELXListingTransactor creates a new write-only instance of GELXListing, bound to a specific deployed contract.
func NewGELXListingTransactor(address common.Address, transactor bind.ContractTransactor) (*GELXListingTransactor, error) {
	contract, err := bindGELXListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GELXListingTransactor{contract: contract}, nil
}

// NewGELXListingFilterer creates a new log filterer instance of GELXListing, bound to a specific deployed contract.
func NewGELXListingFilterer(address common.Address, filterer bind.ContractFilterer) (*GELXListingFilterer, error) {
	contract, err := bindGELXListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GELXListingFilterer{contract: contract}, nil
}

// bindGELXListing binds a generic wrapper to an already deployed contract.
func bindGELXListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GELXListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GELXListing *GELXListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GELXListing.Contract.GELXListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GELXListing *GELXListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GELXListing.Contract.GELXListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GELXListing *GELXListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GELXListing.Contract.GELXListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GELXListing *GELXListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GELXListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GELXListing *GELXListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GELXListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GELXListing *GELXListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GELXListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_GELXListing *GELXListingCaller) GetTokenStatus(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GELXListing.contract.Call(opts, out, "getTokenStatus", token)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_GELXListing *GELXListingSession) GetTokenStatus(token common.Address) (bool, error) {
	return _GELXListing.Contract.GetTokenStatus(&_GELXListing.CallOpts, token)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_GELXListing *GELXListingCallerSession) GetTokenStatus(token common.Address) (bool, error) {
	return _GELXListing.Contract.GetTokenStatus(&_GELXListing.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_GELXListing *GELXListingCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GELXListing.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_GELXListing *GELXListingSession) Tokens() ([]common.Address, error) {
	return _GELXListing.Contract.Tokens(&_GELXListing.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_GELXListing *GELXListingCallerSession) Tokens() ([]common.Address, error) {
	return _GELXListing.Contract.Tokens(&_GELXListing.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_GELXListing *GELXListingTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _GELXListing.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_GELXListing *GELXListingSession) Apply(token common.Address) (*types.Transaction, error) {
	return _GELXListing.Contract.Apply(&_GELXListing.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_GELXListing *GELXListingTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _GELXListing.Contract.Apply(&_GELXListing.TransactOpts, token)
}
