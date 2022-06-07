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

// ContentContentLog is an auto generated low-level Go binding around an user-defined struct.
type ContentContentLog struct {
	Owner    common.Address
	Hash     []string
	Provider common.Address
	LogId    string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ContentLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"LogId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"}],\"name\":\"GetContentLog\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"Hash\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"Provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"LogId\",\"type\":\"string\"}],\"internalType\":\"structContent.ContentLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_hash\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"SetContentLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610a09380380610a0983398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610976806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063242d3b4b1461004657806366754e451461006f57806390ca908a14610091575b600080fd5b61005961005436600461072d565b6100a6565b60405161006691906107e7565b60405180910390f35b61008261007d36600461072d565b610286565b604051610066939291906107b2565b6100a461009f366004610632565b61034a565b005b60408051608081018252600080825260606020830181905292820152818101919091526000826040516100d99190610796565b90815260408051918290036020908101832060808401835280546001600160a01b0316845260018101805484518185028101850190955280855291938584019390929060009084015b828210156101ce578382906000526020600020018054610141906108ef565b80601f016020809104026020016040519081016040528092919081815260200182805461016d906108ef565b80156101ba5780601f1061018f576101008083540402835291602001916101ba565b820191906000526020600020905b81548152906001019060200180831161019d57829003601f168201915b505050505081526020019060010190610122565b5050509082525060028201546001600160a01b031660208201526003820180546040909201916101fd906108ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610229906108ef565b80156102765780601f1061024b57610100808354040283529160200191610276565b820191906000526020600020905b81548152906001019060200180831161025957829003601f168201915b5050505050815250509050919050565b8051808201602090810180516000825292820191909301209152805460028201546003830180546001600160a01b039384169492909316926102c7906108ef565b80601f01602080910402602001604051908101604052809291908181526020018280546102f3906108ef565b80156103405780601f1061031557610100808354040283529160200191610340565b820191906000526020600020905b81548152906001019060200180831161032357829003601f168201915b5050505050905083565b6001546001600160a01b031633146103995760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b604482015260640160405180910390fd5b336000836040516103aa9190610796565b908152602001604051809103902060020160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550826000836040516103f09190610796565b9081526020016040518091039020600101908051906020019061041492919061045a565b50806000836040516104269190610796565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b0319909216919091179055505050565b8280548282559060005260206000209081019282156104a7579160200282015b828111156104a757825180516104979184916020909101906104b7565b509160200191906001019061047a565b506104b3929150610537565b5090565b8280546104c3906108ef565b90600052602060002090601f0160209004810192826104e5576000855561052b565b82601f106104fe57805160ff191683800117855561052b565b8280016001018555821561052b579182015b8281111561052b578251825591602001919060010190610510565b506104b3929150610554565b808211156104b357600061054b8282610569565b50600101610537565b5b808211156104b35760008155600101610555565b508054610575906108ef565b6000825580601f10610585575050565b601f0160209004906000526020600020908101906105a39190610554565b50565b80356001600160a01b03811681146105bd57600080fd5b919050565b600082601f8301126105d357600080fd5b813567ffffffffffffffff8111156105ed576105ed61092a565b610600601f8201601f191660200161088e565b81815284602083860101111561061557600080fd5b816020850160208301376000918101602001919091529392505050565b60008060006060848603121561064757600080fd5b833567ffffffffffffffff8082111561065f57600080fd5b818601915086601f83011261067357600080fd5b81356020828211156106875761068761092a565b8160051b61069682820161088e565b8381528281019086840183880185018d10156106b157600080fd5b600093505b858410156106f0578035878111156106cd57600080fd5b6106db8e87838c01016105c2565b845250600193909301929184019184016106b6565b50985050508701359250508082111561070857600080fd5b50610715868287016105c2565b925050610724604085016105a6565b90509250925092565b60006020828403121561073f57600080fd5b813567ffffffffffffffff81111561075657600080fd5b610762848285016105c2565b949350505050565b600081518084526107828160208601602086016108bf565b601f01601f19169290920160200192915050565b600082516107a88184602087016108bf565b9190910192915050565b6001600160a01b038481168252831660208201526060604082018190526000906107de9083018461076a565b95945050505050565b602080825282516001600160a01b0316828201528281015160806040840152805160a084018190526000929160c0600583901b86018101929184019190860190855b818110156108575760bf1988860301835261084585855161076a565b94509285019291850191600101610829565b5050505060408501516001600160a01b038116606086015291506060850151848203601f1901608086015291506107de818361076a565b604051601f8201601f1916810167ffffffffffffffff811182821017156108b7576108b761092a565b604052919050565b60005b838110156108da5781810151838201526020016108c2565b838111156108e9576000848401525b50505050565b600181811c9082168061090357607f821691505b6020821081141561092457634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220ea5fa84d545994e45fc4ac8ad787761f54d03131475bf81f34e698eb2fd06cc864736f6c63430008070033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _spa common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _spa)
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

// ContentLogs is a free data retrieval call binding the contract method 0x66754e45.
//
// Solidity: function ContentLogs(string ) view returns(address Owner, address Provider, string LogId)
func (_Contracts *ContractsCaller) ContentLogs(opts *bind.CallOpts, arg0 string) (struct {
	Owner    common.Address
	Provider common.Address
	LogId    string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ContentLogs", arg0)

	outstruct := new(struct {
		Owner    common.Address
		Provider common.Address
		LogId    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Provider = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.LogId = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// ContentLogs is a free data retrieval call binding the contract method 0x66754e45.
//
// Solidity: function ContentLogs(string ) view returns(address Owner, address Provider, string LogId)
func (_Contracts *ContractsSession) ContentLogs(arg0 string) (struct {
	Owner    common.Address
	Provider common.Address
	LogId    string
}, error) {
	return _Contracts.Contract.ContentLogs(&_Contracts.CallOpts, arg0)
}

// ContentLogs is a free data retrieval call binding the contract method 0x66754e45.
//
// Solidity: function ContentLogs(string ) view returns(address Owner, address Provider, string LogId)
func (_Contracts *ContractsCallerSession) ContentLogs(arg0 string) (struct {
	Owner    common.Address
	Provider common.Address
	LogId    string
}, error) {
	return _Contracts.Contract.ContentLogs(&_Contracts.CallOpts, arg0)
}

// GetContentLog is a free data retrieval call binding the contract method 0x242d3b4b.
//
// Solidity: function GetContentLog(string _contentId) view returns((address,string[],address,string))
func (_Contracts *ContractsCaller) GetContentLog(opts *bind.CallOpts, _contentId string) (ContentContentLog, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetContentLog", _contentId)

	if err != nil {
		return *new(ContentContentLog), err
	}

	out0 := *abi.ConvertType(out[0], new(ContentContentLog)).(*ContentContentLog)

	return out0, err

}

// GetContentLog is a free data retrieval call binding the contract method 0x242d3b4b.
//
// Solidity: function GetContentLog(string _contentId) view returns((address,string[],address,string))
func (_Contracts *ContractsSession) GetContentLog(_contentId string) (ContentContentLog, error) {
	return _Contracts.Contract.GetContentLog(&_Contracts.CallOpts, _contentId)
}

// GetContentLog is a free data retrieval call binding the contract method 0x242d3b4b.
//
// Solidity: function GetContentLog(string _contentId) view returns((address,string[],address,string))
func (_Contracts *ContractsCallerSession) GetContentLog(_contentId string) (ContentContentLog, error) {
	return _Contracts.Contract.GetContentLog(&_Contracts.CallOpts, _contentId)
}

// SetContentLog is a paid mutator transaction binding the contract method 0x90ca908a.
//
// Solidity: function SetContentLog(string[] _hash, string _contentId, address _userAddress) returns()
func (_Contracts *ContractsTransactor) SetContentLog(opts *bind.TransactOpts, _hash []string, _contentId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetContentLog", _hash, _contentId, _userAddress)
}

// SetContentLog is a paid mutator transaction binding the contract method 0x90ca908a.
//
// Solidity: function SetContentLog(string[] _hash, string _contentId, address _userAddress) returns()
func (_Contracts *ContractsSession) SetContentLog(_hash []string, _contentId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetContentLog(&_Contracts.TransactOpts, _hash, _contentId, _userAddress)
}

// SetContentLog is a paid mutator transaction binding the contract method 0x90ca908a.
//
// Solidity: function SetContentLog(string[] _hash, string _contentId, address _userAddress) returns()
func (_Contracts *ContractsTransactorSession) SetContentLog(_hash []string, _contentId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetContentLog(&_Contracts.TransactOpts, _hash, _contentId, _userAddress)
}
