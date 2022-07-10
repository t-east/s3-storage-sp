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
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"GetPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PubKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"SetPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610455806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806309140fb61461004657806376070e7a1461006f578063a154789114610084575b600080fd5b610059610054366004610298565b610097565b6040516100669190610379565b60405180910390f35b61008261007d3660046102c8565b610131565b005b610059610092366004610298565b610153565b600060208190529081526040902080546100b0906103ce565b80601f01602080910402602001604051908101604052809291908181526020018280546100dc906103ce565b80156101295780601f106100fe57610100808354040283529160200191610129565b820191906000526020600020905b81548152906001019060200180831161010c57829003601f168201915b505050505081565b33600090815260208181526040909120825161014f928401906101ff565b5050565b6001600160a01b038116600090815260208190526040902080546060919061017a906103ce565b80601f01602080910402602001604051908101604052809291908181526020018280546101a6906103ce565b80156101f35780601f106101c8576101008083540402835291602001916101f3565b820191906000526020600020905b8154815290600101906020018083116101d657829003601f168201915b50505050509050919050565b82805461020b906103ce565b90600052602060002090601f01602090048101928261022d5760008555610273565b82601f1061024657805160ff1916838001178555610273565b82800160010185558215610273579182015b82811115610273578251825591602001919060010190610258565b5061027f929150610283565b5090565b5b8082111561027f5760008155600101610284565b6000602082840312156102aa57600080fd5b81356001600160a01b03811681146102c157600080fd5b9392505050565b6000602082840312156102da57600080fd5b813567ffffffffffffffff808211156102f257600080fd5b818401915084601f83011261030657600080fd5b81358181111561031857610318610409565b604051601f8201601f19908116603f0116810190838211818310171561034057610340610409565b8160405282815287602084870101111561035957600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156103a65785810183015185820160400152820161038a565b818111156103b8576000604083870101525b50601f01601f1916929092016040019392505050565b600181811c908216806103e257607f821691505b6020821081141561040357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220d7966575e9170928f2b4ee600497b83a02bcd9f947bcbd085581e01ffd9c588b64736f6c63430008070033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

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
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// GetPubkey is a free data retrieval call binding the contract method 0xa1547891.
//
// Solidity: function GetPubkey(address _address) view returns(bytes)
func (_Contracts *ContractsCaller) GetPubkey(opts *bind.CallOpts, _address common.Address) ([]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetPubkey", _address)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPubkey is a free data retrieval call binding the contract method 0xa1547891.
//
// Solidity: function GetPubkey(address _address) view returns(bytes)
func (_Contracts *ContractsSession) GetPubkey(_address common.Address) ([]byte, error) {
	return _Contracts.Contract.GetPubkey(&_Contracts.CallOpts, _address)
}

// GetPubkey is a free data retrieval call binding the contract method 0xa1547891.
//
// Solidity: function GetPubkey(address _address) view returns(bytes)
func (_Contracts *ContractsCallerSession) GetPubkey(_address common.Address) ([]byte, error) {
	return _Contracts.Contract.GetPubkey(&_Contracts.CallOpts, _address)
}

// PubKey is a free data retrieval call binding the contract method 0x09140fb6.
//
// Solidity: function PubKey(address ) view returns(bytes)
func (_Contracts *ContractsCaller) PubKey(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "PubKey", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PubKey is a free data retrieval call binding the contract method 0x09140fb6.
//
// Solidity: function PubKey(address ) view returns(bytes)
func (_Contracts *ContractsSession) PubKey(arg0 common.Address) ([]byte, error) {
	return _Contracts.Contract.PubKey(&_Contracts.CallOpts, arg0)
}

// PubKey is a free data retrieval call binding the contract method 0x09140fb6.
//
// Solidity: function PubKey(address ) view returns(bytes)
func (_Contracts *ContractsCallerSession) PubKey(arg0 common.Address) ([]byte, error) {
	return _Contracts.Contract.PubKey(&_Contracts.CallOpts, arg0)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x76070e7a.
//
// Solidity: function SetPubkey(bytes _pubkey) returns()
func (_Contracts *ContractsTransactor) SetPubkey(opts *bind.TransactOpts, _pubkey []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetPubkey", _pubkey)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x76070e7a.
//
// Solidity: function SetPubkey(bytes _pubkey) returns()
func (_Contracts *ContractsSession) SetPubkey(_pubkey []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetPubkey(&_Contracts.TransactOpts, _pubkey)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x76070e7a.
//
// Solidity: function SetPubkey(bytes _pubkey) returns()
func (_Contracts *ContractsTransactorSession) SetPubkey(_pubkey []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetPubkey(&_Contracts.TransactOpts, _pubkey)
}
