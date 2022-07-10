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

// AuditAuditLog is an auto generated low-level Go binding around an user-defined struct.
type AuditAuditLog struct {
	Result bool
	Chal   uint32
	K1     []byte
	K2     []byte
	Myu    []byte
	Gamma  []byte
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spa\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tpa\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"AuditLogs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"Result\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"Chal\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"K1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"K2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Myu\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Gamma\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentId\",\"type\":\"string\"}],\"name\":\"GetAuditLog\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"Result\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"Chal\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"K1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"K2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Myu\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Gamma\",\"type\":\"bytes\"}],\"internalType\":\"structAudit.AuditLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentID\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_chal\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_k1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_k2\",\"type\":\"bytes\"}],\"name\":\"InitAuditLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_contentIDs\",\"type\":\"string[]\"}],\"name\":\"ListAuditLog\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"Result\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"Chal\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"K1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"K2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Myu\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Gamma\",\"type\":\"bytes\"}],\"internalType\":\"structAudit.AuditLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_myu\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_gamma\",\"type\":\"bytes\"}],\"name\":\"SetProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contentID\",\"type\":\"string\"}],\"name\":\"SetResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161133738038061133783398101604081905261002f9161008b565b600180546001600160a01b03199081166001600160a01b0394851617909155600380548216948416949094179093556002805490931691161790556100ce565b80516001600160a01b038116811461008657600080fd5b919050565b6000806000606084860312156100a057600080fd5b6100a98461006f565b92506100b76020850161006f565b91506100c56040850161006f565b90509250925092565b61125a806100dd6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806342eccaf71461006757806346e8ecbb1461009057806368e9b6dc146100b0578063d805317e146100d5578063f4cb7a8d146100ea578063ff164f94146100fd575b600080fd5b61007a610075366004610d4b565b610110565b604051610087919061104d565b60405180910390f35b6100a361009e366004610e13565b610404565b604051610087919061111d565b6100c36100be366004610e13565b6106dc565b604051610087969594939291906110af565b6100e86100e3366004610e50565b610947565b005b6100e86100f8366004610ed8565b610a08565b6100e861010b366004610e13565b610bc1565b8051606090819060005b818110156103fb576000858281518110610136576101366111f8565b602002602001015160405161014b9190611031565b908152604080516020928190038301812060c082018352805460ff811615158352610100900463ffffffff169382019390935260018301805491939284019161019390611194565b80601f01602080910402602001604051908101604052809291908181526020018280546101bf90611194565b801561020c5780601f106101e15761010080835404028352916020019161020c565b820191906000526020600020905b8154815290600101906020018083116101ef57829003601f168201915b5050505050815260200160028201805461022590611194565b80601f016020809104026020016040519081016040528092919081815260200182805461025190611194565b801561029e5780601f106102735761010080835404028352916020019161029e565b820191906000526020600020905b81548152906001019060200180831161028157829003601f168201915b505050505081526020016003820180546102b790611194565b80601f01602080910402602001604051908101604052809291908181526020018280546102e390611194565b80156103305780601f1061030557610100808354040283529160200191610330565b820191906000526020600020905b81548152906001019060200180831161031357829003601f168201915b5050505050815260200160048201805461034990611194565b80601f016020809104026020016040519081016040528092919081815260200182805461037590611194565b80156103c25780601f10610397576101008083540402835291602001916103c2565b820191906000526020600020905b8154815290600101906020018083116103a557829003601f168201915b5050505050815250508382815181106103dd576103dd6111f8565b602002602001018190525080806103f3906111cf565b91505061011a565b50909392505050565b6104456040518060c00160405280600015158152602001600063ffffffff168152602001606081526020016060815260200160608152602001606081525090565b6000826040516104559190611031565b908152604080516020928190038301812060c082018352805460ff811615158352610100900463ffffffff169382019390935260018301805491939284019161049d90611194565b80601f01602080910402602001604051908101604052809291908181526020018280546104c990611194565b80156105165780601f106104eb57610100808354040283529160200191610516565b820191906000526020600020905b8154815290600101906020018083116104f957829003601f168201915b5050505050815260200160028201805461052f90611194565b80601f016020809104026020016040519081016040528092919081815260200182805461055b90611194565b80156105a85780601f1061057d576101008083540402835291602001916105a8565b820191906000526020600020905b81548152906001019060200180831161058b57829003601f168201915b505050505081526020016003820180546105c190611194565b80601f01602080910402602001604051908101604052809291908181526020018280546105ed90611194565b801561063a5780601f1061060f5761010080835404028352916020019161063a565b820191906000526020600020905b81548152906001019060200180831161061d57829003601f168201915b5050505050815260200160048201805461065390611194565b80601f016020809104026020016040519081016040528092919081815260200182805461067f90611194565b80156106cc5780601f106106a1576101008083540402835291602001916106cc565b820191906000526020600020905b8154815290600101906020018083116106af57829003601f168201915b5050505050815250509050919050565b80516020818301810180516000825292820191909301209152805460018201805460ff83169361010090930463ffffffff1692919061071a90611194565b80601f016020809104026020016040519081016040528092919081815260200182805461074690611194565b80156107935780601f1061076857610100808354040283529160200191610793565b820191906000526020600020905b81548152906001019060200180831161077657829003601f168201915b5050505050908060020180546107a890611194565b80601f01602080910402602001604051908101604052809291908181526020018280546107d490611194565b80156108215780601f106107f657610100808354040283529160200191610821565b820191906000526020600020905b81548152906001019060200180831161080457829003601f168201915b50505050509080600301805461083690611194565b80601f016020809104026020016040519081016040528092919081815260200182805461086290611194565b80156108af5780601f10610884576101008083540402835291602001916108af565b820191906000526020600020905b81548152906001019060200180831161089257829003601f168201915b5050505050908060040180546108c490611194565b80601f01602080910402602001604051908101604052809291908181526020018280546108f090611194565b801561093d5780601f106109125761010080835404028352916020019161093d565b820191906000526020600020905b81548152906001019060200180831161092057829003601f168201915b5050505050905086565b6001546001600160a01b031633146109975760405162461bcd60e51b815260206004820152600e60248201526d0596f7520617265206e6f742053560941b60448201526064015b60405180910390fd5b806000846040516109a89190611031565b908152602001604051809103902060040190805190602001906109cc929190610c42565b50816000846040516109de9190611031565b90815260200160405180910390206003019080519060200190610a02929190610c42565b50505050565b6002546001600160a01b03163314610a545760405162461bcd60e51b815260206004820152600f60248201526e596f7520617265206e6f742054504160881b604482015260640161098e565b82600085604051610a659190611031565b908152602001604051809103902060000160016101000a81548163ffffffff021916908363ffffffff16021790555081600085604051610aa59190611031565b90815260200160405180910390206001019080519060200190610ac9929190610c42565b5080600085604051610adb9190611031565b90815260200160405180910390206002019080519060200190610aff929190610c42565b5060405180602001604052806000815250600085604051610b209190611031565b90815260200160405180910390206004019080519060200190610b44929190610c42565b5060405180602001604052806000815250600085604051610b659190611031565b90815260200160405180910390206003019080519060200190610b89929190610c42565b5060008085604051610b9b9190611031565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b6002546001600160a01b03163314610c0d5760405162461bcd60e51b815260206004820152600f60248201526e596f7520617265206e6f742054504160881b604482015260640161098e565b6001600082604051610c1f9190611031565b908152604051908190036020019020805491151560ff1990921691909117905550565b828054610c4e90611194565b90600052602060002090601f016020900481019282610c705760008555610cb6565b82601f10610c8957805160ff1916838001178555610cb6565b82800160010185558215610cb6579182015b82811115610cb6578251825591602001919060010190610c9b565b50610cc2929150610cc6565b5090565b5b80821115610cc25760008155600101610cc7565b600082601f830112610cec57600080fd5b813567ffffffffffffffff811115610d0657610d0661120e565b610d19601f8201601f1916602001611137565b818152846020838601011115610d2e57600080fd5b816020850160208301376000918101602001919091529392505050565b60006020808385031215610d5e57600080fd5b823567ffffffffffffffff80821115610d7657600080fd5b818501915085601f830112610d8a57600080fd5b813581811115610d9c57610d9c61120e565b8060051b610dab858201611137565b8281528581019085870183870188018b1015610dc657600080fd5b60009350835b85811015610e0357813587811115610de2578586fd5b610df08d8b838c0101610cdb565b8552509288019290880190600101610dcc565b50909a9950505050505050505050565b600060208284031215610e2557600080fd5b813567ffffffffffffffff811115610e3c57600080fd5b610e4884828501610cdb565b949350505050565b600080600060608486031215610e6557600080fd5b833567ffffffffffffffff80821115610e7d57600080fd5b610e8987838801610cdb565b94506020860135915080821115610e9f57600080fd5b610eab87838801610cdb565b93506040860135915080821115610ec157600080fd5b50610ece86828701610cdb565b9150509250925092565b60008060008060808587031215610eee57600080fd5b843567ffffffffffffffff80821115610f0657600080fd5b610f1288838901610cdb565b95506020870135915063ffffffff82168214610f2d57600080fd5b90935060408601359080821115610f4357600080fd5b610f4f88838901610cdb565b93506060870135915080821115610f6557600080fd5b50610f7287828801610cdb565b91505092959194509250565b60008151808452610f96816020860160208601611168565b601f01601f19169290920160200192915050565b80511515825263ffffffff60208201511660208301526000604082015160c06040850152610fdb60c0850182610f7e565b905060608301518482036060860152610ff48282610f7e565b9150506080830151848203608086015261100e8282610f7e565b91505060a083015184820360a08601526110288282610f7e565b95945050505050565b60008251611043818460208701611168565b9190910192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156110a257603f19888603018452611090858351610faa565b94509285019290850190600101611074565b5092979650505050505050565b861515815263ffffffff8616602082015260c0604082015260006110d660c0830187610f7e565b82810360608401526110e88187610f7e565b905082810360808401526110fc8186610f7e565b905082810360a08401526111108185610f7e565b9998505050505050505050565b6020815260006111306020830184610faa565b9392505050565b604051601f8201601f1916810167ffffffffffffffff811182821017156111605761116061120e565b604052919050565b60005b8381101561118357818101518382015260200161116b565b83811115610a025750506000910152565b600181811c908216806111a857607f821691505b602082108114156111c957634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156111f157634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea264697066735822122067fd1c69a680282838fa32aa9d7646586daa47750dac82f4c3ba7cb692ef136a64736f6c63430008070033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _user common.Address, _spa common.Address, _tpa common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _user, _spa, _tpa)
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

// AuditLogs is a free data retrieval call binding the contract method 0x68e9b6dc.
//
// Solidity: function AuditLogs(string ) view returns(bool Result, uint32 Chal, bytes K1, bytes K2, bytes Myu, bytes Gamma)
func (_Contracts *ContractsCaller) AuditLogs(opts *bind.CallOpts, arg0 string) (struct {
	Result bool
	Chal   uint32
	K1     []byte
	K2     []byte
	Myu    []byte
	Gamma  []byte
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "AuditLogs", arg0)

	outstruct := new(struct {
		Result bool
		Chal   uint32
		K1     []byte
		K2     []byte
		Myu    []byte
		Gamma  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Chal = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.K1 = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.K2 = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.Myu = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Gamma = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// AuditLogs is a free data retrieval call binding the contract method 0x68e9b6dc.
//
// Solidity: function AuditLogs(string ) view returns(bool Result, uint32 Chal, bytes K1, bytes K2, bytes Myu, bytes Gamma)
func (_Contracts *ContractsSession) AuditLogs(arg0 string) (struct {
	Result bool
	Chal   uint32
	K1     []byte
	K2     []byte
	Myu    []byte
	Gamma  []byte
}, error) {
	return _Contracts.Contract.AuditLogs(&_Contracts.CallOpts, arg0)
}

// AuditLogs is a free data retrieval call binding the contract method 0x68e9b6dc.
//
// Solidity: function AuditLogs(string ) view returns(bool Result, uint32 Chal, bytes K1, bytes K2, bytes Myu, bytes Gamma)
func (_Contracts *ContractsCallerSession) AuditLogs(arg0 string) (struct {
	Result bool
	Chal   uint32
	K1     []byte
	K2     []byte
	Myu    []byte
	Gamma  []byte
}, error) {
	return _Contracts.Contract.AuditLogs(&_Contracts.CallOpts, arg0)
}

// GetAuditLog is a free data retrieval call binding the contract method 0x46e8ecbb.
//
// Solidity: function GetAuditLog(string _contentId) view returns((bool,uint32,bytes,bytes,bytes,bytes))
func (_Contracts *ContractsCaller) GetAuditLog(opts *bind.CallOpts, _contentId string) (AuditAuditLog, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "GetAuditLog", _contentId)

	if err != nil {
		return *new(AuditAuditLog), err
	}

	out0 := *abi.ConvertType(out[0], new(AuditAuditLog)).(*AuditAuditLog)

	return out0, err

}

// GetAuditLog is a free data retrieval call binding the contract method 0x46e8ecbb.
//
// Solidity: function GetAuditLog(string _contentId) view returns((bool,uint32,bytes,bytes,bytes,bytes))
func (_Contracts *ContractsSession) GetAuditLog(_contentId string) (AuditAuditLog, error) {
	return _Contracts.Contract.GetAuditLog(&_Contracts.CallOpts, _contentId)
}

// GetAuditLog is a free data retrieval call binding the contract method 0x46e8ecbb.
//
// Solidity: function GetAuditLog(string _contentId) view returns((bool,uint32,bytes,bytes,bytes,bytes))
func (_Contracts *ContractsCallerSession) GetAuditLog(_contentId string) (AuditAuditLog, error) {
	return _Contracts.Contract.GetAuditLog(&_Contracts.CallOpts, _contentId)
}

// ListAuditLog is a free data retrieval call binding the contract method 0x42eccaf7.
//
// Solidity: function ListAuditLog(string[] _contentIDs) view returns((bool,uint32,bytes,bytes,bytes,bytes)[])
func (_Contracts *ContractsCaller) ListAuditLog(opts *bind.CallOpts, _contentIDs []string) ([]AuditAuditLog, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ListAuditLog", _contentIDs)

	if err != nil {
		return *new([]AuditAuditLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]AuditAuditLog)).(*[]AuditAuditLog)

	return out0, err

}

// ListAuditLog is a free data retrieval call binding the contract method 0x42eccaf7.
//
// Solidity: function ListAuditLog(string[] _contentIDs) view returns((bool,uint32,bytes,bytes,bytes,bytes)[])
func (_Contracts *ContractsSession) ListAuditLog(_contentIDs []string) ([]AuditAuditLog, error) {
	return _Contracts.Contract.ListAuditLog(&_Contracts.CallOpts, _contentIDs)
}

// ListAuditLog is a free data retrieval call binding the contract method 0x42eccaf7.
//
// Solidity: function ListAuditLog(string[] _contentIDs) view returns((bool,uint32,bytes,bytes,bytes,bytes)[])
func (_Contracts *ContractsCallerSession) ListAuditLog(_contentIDs []string) ([]AuditAuditLog, error) {
	return _Contracts.Contract.ListAuditLog(&_Contracts.CallOpts, _contentIDs)
}

// InitAuditLog is a paid mutator transaction binding the contract method 0xf4cb7a8d.
//
// Solidity: function InitAuditLog(string _contentID, uint32 _chal, bytes _k1, bytes _k2) returns()
func (_Contracts *ContractsTransactor) InitAuditLog(opts *bind.TransactOpts, _contentID string, _chal uint32, _k1 []byte, _k2 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "InitAuditLog", _contentID, _chal, _k1, _k2)
}

// InitAuditLog is a paid mutator transaction binding the contract method 0xf4cb7a8d.
//
// Solidity: function InitAuditLog(string _contentID, uint32 _chal, bytes _k1, bytes _k2) returns()
func (_Contracts *ContractsSession) InitAuditLog(_contentID string, _chal uint32, _k1 []byte, _k2 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.InitAuditLog(&_Contracts.TransactOpts, _contentID, _chal, _k1, _k2)
}

// InitAuditLog is a paid mutator transaction binding the contract method 0xf4cb7a8d.
//
// Solidity: function InitAuditLog(string _contentID, uint32 _chal, bytes _k1, bytes _k2) returns()
func (_Contracts *ContractsTransactorSession) InitAuditLog(_contentID string, _chal uint32, _k1 []byte, _k2 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.InitAuditLog(&_Contracts.TransactOpts, _contentID, _chal, _k1, _k2)
}

// SetProof is a paid mutator transaction binding the contract method 0xd805317e.
//
// Solidity: function SetProof(string _contentID, bytes _myu, bytes _gamma) returns()
func (_Contracts *ContractsTransactor) SetProof(opts *bind.TransactOpts, _contentID string, _myu []byte, _gamma []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetProof", _contentID, _myu, _gamma)
}

// SetProof is a paid mutator transaction binding the contract method 0xd805317e.
//
// Solidity: function SetProof(string _contentID, bytes _myu, bytes _gamma) returns()
func (_Contracts *ContractsSession) SetProof(_contentID string, _myu []byte, _gamma []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetProof(&_Contracts.TransactOpts, _contentID, _myu, _gamma)
}

// SetProof is a paid mutator transaction binding the contract method 0xd805317e.
//
// Solidity: function SetProof(string _contentID, bytes _myu, bytes _gamma) returns()
func (_Contracts *ContractsTransactorSession) SetProof(_contentID string, _myu []byte, _gamma []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetProof(&_Contracts.TransactOpts, _contentID, _myu, _gamma)
}

// SetResult is a paid mutator transaction binding the contract method 0xff164f94.
//
// Solidity: function SetResult(string _contentID) returns()
func (_Contracts *ContractsTransactor) SetResult(opts *bind.TransactOpts, _contentID string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "SetResult", _contentID)
}

// SetResult is a paid mutator transaction binding the contract method 0xff164f94.
//
// Solidity: function SetResult(string _contentID) returns()
func (_Contracts *ContractsSession) SetResult(_contentID string) (*types.Transaction, error) {
	return _Contracts.Contract.SetResult(&_Contracts.TransactOpts, _contentID)
}

// SetResult is a paid mutator transaction binding the contract method 0xff164f94.
//
// Solidity: function SetResult(string _contentID) returns()
func (_Contracts *ContractsTransactorSession) SetResult(_contentID string) (*types.Transaction, error) {
	return _Contracts.Contract.SetResult(&_Contracts.TransactOpts, _contentID)
}
