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
	Hash     [][]byte
	Provider common.Address
	LogId    string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ContentLogs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"LogId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"}],\"name\":\"GetContentLog\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"Hash\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"Provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"LogId\",\"type\":\"string\"}],\"internalType\":\"structContent.ContentLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"}],\"name\":\"InitContentLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ListContentID\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ListContentLog\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"Hash\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"Provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"LogId\",\"type\":\"string\"}],\"internalType\":\"structContent.ContentLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_hash\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"SetContentLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610fea380380610fea83398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610f57806100936000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630728df4114610067578063242d3b4b1461007c57806366754e45146100a55780639a4790f5146100c7578063a33a9072146100dc578063d2056ab9146100ef575b600080fd5b61007a610075366004610a63565b610104565b005b61008f61008a366004610b72565b610258565b60405161009c9190610e21565b60405180910390f35b6100b86100b3366004610b72565b610438565b60405161009c93929190610d3e565b6100cf6104fc565b60405161009c9190610dcc565b61007a6100ea366004610b72565b610719565b6100f76107a0565b60405161009c9190610d6a565b6001546001600160a01b031633146101535760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b604482015260640160405180910390fd5b336000836040516101649190610c86565b908152602001604051809103902060020160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550826000836040516101aa9190610c86565b908152602001604051809103902060010190805190602001906101ce92919061087c565b50806000836040516101e09190610c86565b9081526040516020918190038201902080546001600160a01b0319166001600160a01b039390931692909217909155600480546001810182556000919091528351610252927f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201918501906108d9565b50505050565b604080516080810182526000808252606060208301819052928201528181019190915260008260405161028b9190610c86565b90815260408051918290036020908101832060808401835280546001600160a01b0316845260018101805484518185028101850190955280855291938584019390929060009084015b828210156103805783829060005260206000200180546102f390610e91565b80601f016020809104026020016040519081016040528092919081815260200182805461031f90610e91565b801561036c5780601f106103415761010080835404028352916020019161036c565b820191906000526020600020905b81548152906001019060200180831161034f57829003601f168201915b5050505050815260200190600101906102d4565b5050509082525060028201546001600160a01b031660208201526003820180546040909201916103af90610e91565b80601f01602080910402602001604051908101604052809291908181526020018280546103db90610e91565b80156104285780601f106103fd57610100808354040283529160200191610428565b820191906000526020600020905b81548152906001019060200180831161040b57829003601f168201915b5050505050815250509050919050565b8051808201602090810180516000825292820191909301209152805460028201546003830180546001600160a01b0393841694929093169261047990610e91565b80601f01602080910402602001604051908101604052809291908181526020018280546104a590610e91565b80156104f25780601f106104c7576101008083540402835291602001916104f2565b820191906000526020600020905b8154815290600101906020018083116104d557829003601f168201915b5050505050905083565b600454606090819060005b818110156107115760006004828154811061052457610524610ef5565b9060005260206000200160405161053b9190610ca2565b90815260408051918290036020908101832060808401835280546001600160a01b0316845260018101805484518185028101850190955280855291938584019390929060009084015b828210156106305783829060005260206000200180546105a390610e91565b80601f01602080910402602001604051908101604052809291908181526020018280546105cf90610e91565b801561061c5780601f106105f15761010080835404028352916020019161061c565b820191906000526020600020905b8154815290600101906020018083116105ff57829003601f168201915b505050505081526020019060010190610584565b5050509082525060028201546001600160a01b0316602082015260038201805460409092019161065f90610e91565b80601f016020809104026020016040519081016040528092919081815260200182805461068b90610e91565b80156106d85780601f106106ad576101008083540402835291602001916106d8565b820191906000526020600020905b8154815290600101906020018083116106bb57829003601f168201915b5050505050815250508382815181106106f3576106f3610ef5565b6020026020010181905250808061070990610ecc565b915050610507565b509092915050565b3360008260405161072a9190610c86565b9081526040516020918190038201902080546001600160a01b0319166001600160a01b03939093169290921790915560048054600181018255600091909152825161079c927f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201918401906108d9565b5050565b6060806004805480602002602001604051908101604052809291908181526020016000905b828210156108715783829060005260206000200180546107e490610e91565b80601f016020809104026020016040519081016040528092919081815260200182805461081090610e91565b801561085d5780601f106108325761010080835404028352916020019161085d565b820191906000526020600020905b81548152906001019060200180831161084057829003601f168201915b5050505050815260200190600101906107c5565b509295945050505050565b8280548282559060005260206000209081019282156108c9579160200282015b828111156108c957825180516108b99184916020909101906108d9565b509160200191906001019061089c565b506108d5929150610959565b5090565b8280546108e590610e91565b90600052602060002090601f016020900481019282610907576000855561094d565b82601f1061092057805160ff191683800117855561094d565b8280016001018555821561094d579182015b8281111561094d578251825591602001919060010190610932565b506108d5929150610976565b808211156108d557600061096d828261098b565b50600101610959565b5b808211156108d55760008155600101610977565b50805461099790610e91565b6000825580601f106109a7575050565b601f0160209004906000526020600020908101906109c59190610976565b50565b600067ffffffffffffffff8311156109e2576109e2610f0b565b6109f5601f8401601f1916602001610e34565b9050828152838383011115610a0957600080fd5b828260208301376000602084830101529392505050565b80356001600160a01b0381168114610a3757600080fd5b919050565b600082601f830112610a4d57600080fd5b610a5c838335602085016109c8565b9392505050565b600080600060608486031215610a7857600080fd5b833567ffffffffffffffff80821115610a9057600080fd5b818601915086601f830112610aa457600080fd5b8135602082821115610ab857610ab8610f0b565b8160051b610ac7828201610e34565b8381528281019086840183880185018d1015610ae257600080fd5b600093505b85841015610b3557803587811115610afe57600080fd5b8801603f81018e13610b0f57600080fd5b610b208e87830135604084016109c8565b84525060019390930192918401918401610ae7565b509850505087013592505080821115610b4d57600080fd5b50610b5a86828701610a3c565b925050610b6960408501610a20565b90509250925092565b600060208284031215610b8457600080fd5b813567ffffffffffffffff811115610b9b57600080fd5b610ba784828501610a3c565b949350505050565b60008151808452610bc7816020860160208601610e65565b601f01601f19169290920160200192915050565b60006080830160018060a01b03835116845260208084015160808287015282815180855260a08801915060a08160051b8901019450838301925060005b81811015610c4657609f19898703018352610c34868551610baf565b95509284019291840191600101610c18565b50505050506040830151610c6560408601826001600160a01b03169052565b5060608301518482036060860152610c7d8282610baf565b95945050505050565b60008251610c98818460208701610e65565b9190910192915050565b600080835481600182811c915080831680610cbe57607f831692505b6020808410821415610cde57634e487b7160e01b86526022600452602486fd5b818015610cf25760018114610d0357610d30565b60ff19861689528489019650610d30565b60008a81526020902060005b86811015610d285781548b820152908501908301610d0f565b505084890196505b509498975050505050505050565b6001600160a01b03848116825283166020820152606060408201819052600090610c7d90830184610baf565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610dbf57603f19888603018452610dad858351610baf565b94509285019290850190600101610d91565b5092979650505050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610dbf57603f19888603018452610e0f858351610bdb565b94509285019290850190600101610df3565b602081526000610a5c6020830184610bdb565b604051601f8201601f1916810167ffffffffffffffff81118282101715610e5d57610e5d610f0b565b604052919050565b60005b83811015610e80578181015183820152602001610e68565b838111156102525750506000910152565b600181811c90821680610ea557607f821691505b60208210811415610ec657634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415610eee57634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212202ba16bcb8b4054993ef1f62037468cfa4b448c582e1086da0a978bc92c27fc0064736f6c63430008070033",
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
// Solidity: function GetContentLog(string _contentId) view returns((address,bytes[],address,string))
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
// Solidity: function GetContentLog(string _contentId) view returns((address,bytes[],address,string))
func (_Contracts *ContractsSession) GetContentLog(_contentId string) (ContentContentLog, error) {
	return _Contracts.Contract.GetContentLog(&_Contracts.CallOpts, _contentId)
}

// GetContentLog is a free data retrieval call binding the contract method 0x242d3b4b.
//
// Solidity: function GetContentLog(string _contentId) view returns((address,bytes[],address,string))
func (_Contracts *ContractsCallerSession) GetContentLog(_contentId string) (ContentContentLog, error) {
	return _Contracts.Contract.GetContentLog(&_Contracts.CallOpts, _contentId)
}

// ListContentID is a free data retrieval call binding the contract method 0xd2056ab9.
//
// Solidity: function ListContentID() view returns(string[])
func (_Contracts *ContractsCaller) ListContentID(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ListContentID")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ListContentID is a free data retrieval call binding the contract method 0xd2056ab9.
//
// Solidity: function ListContentID() view returns(string[])
func (_Contracts *ContractsSession) ListContentID() ([]string, error) {
	return _Contracts.Contract.ListContentID(&_Contracts.CallOpts)
}

// ListContentID is a free data retrieval call binding the contract method 0xd2056ab9.
//
// Solidity: function ListContentID() view returns(string[])
func (_Contracts *ContractsCallerSession) ListContentID() ([]string, error) {
	return _Contracts.Contract.ListContentID(&_Contracts.CallOpts)
}

// ListContentLog is a free data retrieval call binding the contract method 0x9a4790f5.
//
// Solidity: function ListContentLog() view returns((address,bytes[],address,string)[])
func (_Contracts *ContractsCaller) ListContentLog(opts *bind.CallOpts) ([]ContentContentLog, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ListContentLog")

	if err != nil {
		return *new([]ContentContentLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]ContentContentLog)).(*[]ContentContentLog)

	return out0, err

}

// ListContentLog is a free data retrieval call binding the contract method 0x9a4790f5.
//
// Solidity: function ListContentLog() view returns((address,bytes[],address,string)[])
func (_Contracts *ContractsSession) ListContentLog() ([]ContentContentLog, error) {
	return _Contracts.Contract.ListContentLog(&_Contracts.CallOpts)
}

// ListContentLog is a free data retrieval call binding the contract method 0x9a4790f5.
//
// Solidity: function ListContentLog() view returns((address,bytes[],address,string)[])
func (_Contracts *ContractsCallerSession) ListContentLog() ([]ContentContentLog, error) {
	return _Contracts.Contract.ListContentLog(&_Contracts.CallOpts)
}

// InitContentLog is a paid mutator transaction binding the contract method 0xa33a9072.
//
// Solidity: function InitContentLog(string _contentId) returns()
func (_Contracts *ContractsTransactor) InitContentLog(opts *bind.TransactOpts, _contentId string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "InitContentLog", _contentId)
}

// InitContentLog is a paid mutator transaction binding the contract method 0xa33a9072.
//
// Solidity: function InitContentLog(string _contentId) returns()
func (_Contracts *ContractsSession) InitContentLog(_contentId string) (*types.Transaction, error) {
	return _Contracts.Contract.InitContentLog(&_Contracts.TransactOpts, _contentId)
}

// InitContentLog is a paid mutator transaction binding the contract method 0xa33a9072.
//
// Solidity: function InitContentLog(string _contentId) returns()
func (_Contracts *ContractsTransactorSession) InitContentLog(_contentId string) (*types.Transaction, error) {
	return _Contracts.Contract.InitContentLog(&_Contracts.TransactOpts, _contentId)
}

// SetContentLog is a paid mutator transaction binding the contract method 0x0728df41.
//
// Solidity: function SetContentLog(bytes[] _hash, string _contentId, address _userAddress) returns()
func (_Contracts *ContractsTransactor) SetContentLog(opts *bind.TransactOpts, _hash [][]byte, _contentId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetContentLog", _hash, _contentId, _userAddress)
}

// SetContentLog is a paid mutator transaction binding the contract method 0x0728df41.
//
// Solidity: function SetContentLog(bytes[] _hash, string _contentId, address _userAddress) returns()
func (_Contracts *ContractsSession) SetContentLog(_hash [][]byte, _contentId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetContentLog(&_Contracts.TransactOpts, _hash, _contentId, _userAddress)
}

// SetContentLog is a paid mutator transaction binding the contract method 0x0728df41.
//
// Solidity: function SetContentLog(bytes[] _hash, string _contentId, address _userAddress) returns()
func (_Contracts *ContractsTransactorSession) SetContentLog(_hash [][]byte, _contentId string, _userAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetContentLog(&_Contracts.TransactOpts, _hash, _contentId, _userAddress)
}
