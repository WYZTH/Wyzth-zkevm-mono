// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wyzth_zkevml2

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

// WYzth_zkevmL2EIP1559Params is an auto generated low-level Go binding around an user-defined struct.
type WYzth_zkevmL2EIP1559Params struct {
	Basefee            uint64
	GasIssuedPerSecond uint64
	GasExcessMax       uint64
	GasTarget          uint64
	Ratio2x1x          uint64
}

// WYzth_zkevmL2MetaData contains all meta data concerning the WYzth_zkevmL2 contract.
var WYzth_zkevmL2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expected\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"actual\",\"type\":\"uint64\"}],\"name\":\"L2_BASEFEE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_1559_PARAMS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_CHAIN_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_GOLDEN_TOUCH_K\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_SENDER\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L2_PUBLIC_INPUT_HASH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_TOO_LATE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"M1559_OUT_OF_STOCK\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"M1559_OUT_OF_STOCK\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expected\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"actual\",\"type\":\"uint64\"}],\"name\":\"M1559_UNEXPECTED_CHANGE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expected\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"actual\",\"type\":\"uint64\"}],\"name\":\"M1559_UNEXPECTED_CHANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"number\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"basefee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gaslimit\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevrandao\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainid\",\"type\":\"uint32\"}],\"name\":\"Anchored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GOLDEN_TOUCH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOLDEN_TOUCH_PRIVATEKEY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__reserved1\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l1SignalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"parentGasUsed\",\"type\":\"uint64\"}],\"name\":\"anchor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasExcess\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasIssuedPerSecond\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"timeSinceParent\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"parentGasUsed\",\"type\":\"uint64\"}],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getCrossChainBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getCrossChainSignalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"basefee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasIssuedPerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasExcessMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ratio2x1x\",\"type\":\"uint64\"}],\"internalType\":\"structWYzth_zkevmL2.EIP1559Params\",\"name\":\"_param1559\",\"type\":\"tuple\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSyncedL1Height\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"k\",\"type\":\"uint8\"}],\"name\":\"signAnchor\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xscale\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yscale\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WYzth_zkevmL2ABI is the input ABI used to generate the binding from.
// Deprecated: Use WYzth_zkevmL2MetaData.ABI instead.
var WYzth_zkevmL2ABI = WYzth_zkevmL2MetaData.ABI

// WYzth_zkevmL2 is an auto generated Go binding around an Ethereum contract.
type WYzth_zkevmL2 struct {
	WYzth_zkevmL2Caller     // Read-only binding to the contract
	WYzth_zkevmL2Transactor // Write-only binding to the contract
	WYzth_zkevmL2Filterer   // Log filterer for contract events
}

// WYzth_zkevmL2Caller is an auto generated read-only Go binding around an Ethereum contract.
type WYzth_zkevmL2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WYzth_zkevmL2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WYzth_zkevmL2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WYzth_zkevmL2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WYzth_zkevmL2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WYzth_zkevmL2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WYzth_zkevmL2Session struct {
	Contract     *WYzth_zkevmL2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WYzth_zkevmL2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WYzth_zkevmL2CallerSession struct {
	Contract *WYzth_zkevmL2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WYzth_zkevmL2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WYzth_zkevmL2TransactorSession struct {
	Contract     *WYzth_zkevmL2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WYzth_zkevmL2Raw is an auto generated low-level Go binding around an Ethereum contract.
type WYzth_zkevmL2Raw struct {
	Contract *WYzth_zkevmL2 // Generic contract binding to access the raw methods on
}

// WYzth_zkevmL2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WYzth_zkevmL2CallerRaw struct {
	Contract *WYzth_zkevmL2Caller // Generic read-only contract binding to access the raw methods on
}

// WYzth_zkevmL2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WYzth_zkevmL2TransactorRaw struct {
	Contract *WYzth_zkevmL2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWYzth_zkevmL2 creates a new instance of WYzth_zkevmL2, bound to a specific deployed contract.
func NewWYzth_zkevmL2(address common.Address, backend bind.ContractBackend) (*WYzth_zkevmL2, error) {
	contract, err := bindWYzth_zkevmL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL2{WYzth_zkevmL2Caller: WYzth_zkevmL2Caller{contract: contract}, WYzth_zkevmL2Transactor: WYzth_zkevmL2Transactor{contract: contract}, WYzth_zkevmL2Filterer: WYzth_zkevmL2Filterer{contract: contract}}, nil
}

// NewWYzth_zkevmL2Caller creates a new read-only instance of WYzth_zkevmL2, bound to a specific deployed contract.
func NewWYzth_zkevmL2Caller(address common.Address, caller bind.ContractCaller) (*WYzth_zkevmL2Caller, error) {
	contract, err := bindWYzth_zkevmL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL2Caller{contract: contract}, nil
}

// NewWYzth_zkevmL2Transactor creates a new write-only instance of WYzth_zkevmL2, bound to a specific deployed contract.
func NewWYzth_zkevmL2Transactor(address common.Address, transactor bind.ContractTransactor) (*WYzth_zkevmL2Transactor, error) {
	contract, err := bindWYzth_zkevmL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL2Transactor{contract: contract}, nil
}

// NewWYzth_zkevmL2Filterer creates a new log filterer instance of WYzth_zkevmL2, bound to a specific deployed contract.
func NewWYzth_zkevmL2Filterer(address common.Address, filterer bind.ContractFilterer) (*WYzth_zkevmL2Filterer, error) {
	contract, err := bindWYzth_zkevmL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL2Filterer{contract: contract}, nil
}

// bindWYzth_zkevmL2 binds a generic wrapper to an already deployed contract.
func bindWYzth_zkevmL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WYzth_zkevmL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WYzth_zkevmL2 *WYzth_zkevmL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WYzth_zkevmL2.Contract.WYzth_zkevmL2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WYzth_zkevmL2 *WYzth_zkevmL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.WYzth_zkevmL2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WYzth_zkevmL2 *WYzth_zkevmL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.WYzth_zkevmL2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WYzth_zkevmL2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WYzth_zkevmL2 *WYzth_zkevmL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WYzth_zkevmL2 *WYzth_zkevmL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.contract.Transact(opts, method, params...)
}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) GOLDENTOUCHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "GOLDEN_TOUCH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) GOLDENTOUCHADDRESS() (common.Address, error) {
	return _WYzth_zkevmL2.Contract.GOLDENTOUCHADDRESS(&_WYzth_zkevmL2.CallOpts)
}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) GOLDENTOUCHADDRESS() (common.Address, error) {
	return _WYzth_zkevmL2.Contract.GOLDENTOUCHADDRESS(&_WYzth_zkevmL2.CallOpts)
}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) GOLDENTOUCHPRIVATEKEY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "GOLDEN_TOUCH_PRIVATEKEY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) GOLDENTOUCHPRIVATEKEY() (*big.Int, error) {
	return _WYzth_zkevmL2.Contract.GOLDENTOUCHPRIVATEKEY(&_WYzth_zkevmL2.CallOpts)
}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) GOLDENTOUCHPRIVATEKEY() (*big.Int, error) {
	return _WYzth_zkevmL2.Contract.GOLDENTOUCHPRIVATEKEY(&_WYzth_zkevmL2.CallOpts)
}

// Reserved1 is a free data retrieval call binding the contract method 0xf5d11edc.
//
// Solidity: function __reserved1() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) Reserved1(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "__reserved1")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Reserved1 is a free data retrieval call binding the contract method 0xf5d11edc.
//
// Solidity: function __reserved1() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) Reserved1() (uint64, error) {
	return _WYzth_zkevmL2.Contract.Reserved1(&_WYzth_zkevmL2.CallOpts)
}

// Reserved1 is a free data retrieval call binding the contract method 0xf5d11edc.
//
// Solidity: function __reserved1() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) Reserved1() (uint64, error) {
	return _WYzth_zkevmL2.Contract.Reserved1(&_WYzth_zkevmL2.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) AddressManager() (common.Address, error) {
	return _WYzth_zkevmL2.Contract.AddressManager(&_WYzth_zkevmL2.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) AddressManager() (common.Address, error) {
	return _WYzth_zkevmL2.Contract.AddressManager(&_WYzth_zkevmL2.CallOpts)
}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) GasExcess(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "gasExcess")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) GasExcess() (uint64, error) {
	return _WYzth_zkevmL2.Contract.GasExcess(&_WYzth_zkevmL2.CallOpts)
}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) GasExcess() (uint64, error) {
	return _WYzth_zkevmL2.Contract.GasExcess(&_WYzth_zkevmL2.CallOpts)
}

// GasIssuedPerSecond is a free data retrieval call binding the contract method 0x210c9fe8.
//
// Solidity: function gasIssuedPerSecond() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) GasIssuedPerSecond(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "gasIssuedPerSecond")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GasIssuedPerSecond is a free data retrieval call binding the contract method 0x210c9fe8.
//
// Solidity: function gasIssuedPerSecond() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) GasIssuedPerSecond() (uint64, error) {
	return _WYzth_zkevmL2.Contract.GasIssuedPerSecond(&_WYzth_zkevmL2.CallOpts)
}

// GasIssuedPerSecond is a free data retrieval call binding the contract method 0x210c9fe8.
//
// Solidity: function gasIssuedPerSecond() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) GasIssuedPerSecond() (uint64, error) {
	return _WYzth_zkevmL2.Contract.GasIssuedPerSecond(&_WYzth_zkevmL2.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0xe1848cb0.
//
// Solidity: function getBasefee(uint32 timeSinceParent, uint64 gasLimit, uint64 parentGasUsed) view returns(uint256 _basefee)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) GetBasefee(opts *bind.CallOpts, timeSinceParent uint32, gasLimit uint64, parentGasUsed uint64) (*big.Int, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "getBasefee", timeSinceParent, gasLimit, parentGasUsed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasefee is a free data retrieval call binding the contract method 0xe1848cb0.
//
// Solidity: function getBasefee(uint32 timeSinceParent, uint64 gasLimit, uint64 parentGasUsed) view returns(uint256 _basefee)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) GetBasefee(timeSinceParent uint32, gasLimit uint64, parentGasUsed uint64) (*big.Int, error) {
	return _WYzth_zkevmL2.Contract.GetBasefee(&_WYzth_zkevmL2.CallOpts, timeSinceParent, gasLimit, parentGasUsed)
}

// GetBasefee is a free data retrieval call binding the contract method 0xe1848cb0.
//
// Solidity: function getBasefee(uint32 timeSinceParent, uint64 gasLimit, uint64 parentGasUsed) view returns(uint256 _basefee)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) GetBasefee(timeSinceParent uint32, gasLimit uint64, parentGasUsed uint64) (*big.Int, error) {
	return _WYzth_zkevmL2.Contract.GetBasefee(&_WYzth_zkevmL2.CallOpts, timeSinceParent, gasLimit, parentGasUsed)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) GetBlockHash(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "getBlockHash", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) GetBlockHash(number *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL2.Contract.GetBlockHash(&_WYzth_zkevmL2.CallOpts, number)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) GetBlockHash(number *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL2.Contract.GetBlockHash(&_WYzth_zkevmL2.CallOpts, number)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 number) view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) GetCrossChainBlockHash(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "getCrossChainBlockHash", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 number) view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) GetCrossChainBlockHash(number *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL2.Contract.GetCrossChainBlockHash(&_WYzth_zkevmL2.CallOpts, number)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 number) view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) GetCrossChainBlockHash(number *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL2.Contract.GetCrossChainBlockHash(&_WYzth_zkevmL2.CallOpts, number)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 number) view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) GetCrossChainSignalRoot(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "getCrossChainSignalRoot", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 number) view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) GetCrossChainSignalRoot(number *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL2.Contract.GetCrossChainSignalRoot(&_WYzth_zkevmL2.CallOpts, number)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 number) view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) GetCrossChainSignalRoot(number *big.Int) ([32]byte, error) {
	return _WYzth_zkevmL2.Contract.GetCrossChainSignalRoot(&_WYzth_zkevmL2.CallOpts, number)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) LatestSyncedL1Height(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "latestSyncedL1Height")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) LatestSyncedL1Height() (uint64, error) {
	return _WYzth_zkevmL2.Contract.LatestSyncedL1Height(&_WYzth_zkevmL2.CallOpts)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) LatestSyncedL1Height() (uint64, error) {
	return _WYzth_zkevmL2.Contract.LatestSyncedL1Height(&_WYzth_zkevmL2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) Owner() (common.Address, error) {
	return _WYzth_zkevmL2.Contract.Owner(&_WYzth_zkevmL2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) Owner() (common.Address, error) {
	return _WYzth_zkevmL2.Contract.Owner(&_WYzth_zkevmL2.CallOpts)
}

// ParentTimestamp is a free data retrieval call binding the contract method 0x539b8ade.
//
// Solidity: function parentTimestamp() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) ParentTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "parentTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ParentTimestamp is a free data retrieval call binding the contract method 0x539b8ade.
//
// Solidity: function parentTimestamp() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) ParentTimestamp() (uint64, error) {
	return _WYzth_zkevmL2.Contract.ParentTimestamp(&_WYzth_zkevmL2.CallOpts)
}

// ParentTimestamp is a free data retrieval call binding the contract method 0x539b8ade.
//
// Solidity: function parentTimestamp() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) ParentTimestamp() (uint64, error) {
	return _WYzth_zkevmL2.Contract.ParentTimestamp(&_WYzth_zkevmL2.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) PublicInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "publicInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) PublicInputHash() ([32]byte, error) {
	return _WYzth_zkevmL2.Contract.PublicInputHash(&_WYzth_zkevmL2.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) PublicInputHash() ([32]byte, error) {
	return _WYzth_zkevmL2.Contract.PublicInputHash(&_WYzth_zkevmL2.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) Resolve(opts *bind.CallOpts, chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _WYzth_zkevmL2.Contract.Resolve(&_WYzth_zkevmL2.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _WYzth_zkevmL2.Contract.Resolve(&_WYzth_zkevmL2.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _WYzth_zkevmL2.Contract.Resolve0(&_WYzth_zkevmL2.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _WYzth_zkevmL2.Contract.Resolve0(&_WYzth_zkevmL2.CallOpts, name, allowZeroAddress)
}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) SignAnchor(opts *bind.CallOpts, digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "signAnchor", digest, k)

	outstruct := new(struct {
		V uint8
		R *big.Int
		S *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.S = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) SignAnchor(digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _WYzth_zkevmL2.Contract.SignAnchor(&_WYzth_zkevmL2.CallOpts, digest, k)
}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) SignAnchor(digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _WYzth_zkevmL2.Contract.SignAnchor(&_WYzth_zkevmL2.CallOpts, digest, k)
}

// Xscale is a free data retrieval call binding the contract method 0xf5c97740.
//
// Solidity: function xscale() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) Xscale(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "xscale")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Xscale is a free data retrieval call binding the contract method 0xf5c97740.
//
// Solidity: function xscale() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) Xscale() (uint64, error) {
	return _WYzth_zkevmL2.Contract.Xscale(&_WYzth_zkevmL2.CallOpts)
}

// Xscale is a free data retrieval call binding the contract method 0xf5c97740.
//
// Solidity: function xscale() view returns(uint64)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) Xscale() (uint64, error) {
	return _WYzth_zkevmL2.Contract.Xscale(&_WYzth_zkevmL2.CallOpts)
}

// Yscale is a free data retrieval call binding the contract method 0x3fa85350.
//
// Solidity: function yscale() view returns(uint128)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Caller) Yscale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WYzth_zkevmL2.contract.Call(opts, &out, "yscale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Yscale is a free data retrieval call binding the contract method 0x3fa85350.
//
// Solidity: function yscale() view returns(uint128)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) Yscale() (*big.Int, error) {
	return _WYzth_zkevmL2.Contract.Yscale(&_WYzth_zkevmL2.CallOpts)
}

// Yscale is a free data retrieval call binding the contract method 0x3fa85350.
//
// Solidity: function yscale() view returns(uint128)
func (_WYzth_zkevmL2 *WYzth_zkevmL2CallerSession) Yscale() (*big.Int, error) {
	return _WYzth_zkevmL2.Contract.Yscale(&_WYzth_zkevmL2.CallOpts)
}

// Anchor is a paid mutator transaction binding the contract method 0x3d384a4b.
//
// Solidity: function anchor(bytes32 l1Hash, bytes32 l1SignalRoot, uint64 l1Height, uint64 parentGasUsed) returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2Transactor) Anchor(opts *bind.TransactOpts, l1Hash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint64) (*types.Transaction, error) {
	return _WYzth_zkevmL2.contract.Transact(opts, "anchor", l1Hash, l1SignalRoot, l1Height, parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0x3d384a4b.
//
// Solidity: function anchor(bytes32 l1Hash, bytes32 l1SignalRoot, uint64 l1Height, uint64 parentGasUsed) returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) Anchor(l1Hash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint64) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.Anchor(&_WYzth_zkevmL2.TransactOpts, l1Hash, l1SignalRoot, l1Height, parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0x3d384a4b.
//
// Solidity: function anchor(bytes32 l1Hash, bytes32 l1SignalRoot, uint64 l1Height, uint64 parentGasUsed) returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2TransactorSession) Anchor(l1Hash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint64) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.Anchor(&_WYzth_zkevmL2.TransactOpts, l1Hash, l1SignalRoot, l1Height, parentGasUsed)
}

// Init is a paid mutator transaction binding the contract method 0x8f3ca30d.
//
// Solidity: function init(address _addressManager, (uint64,uint64,uint64,uint64,uint64) _param1559) returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2Transactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _param1559 WYzth_zkevmL2EIP1559Params) (*types.Transaction, error) {
	return _WYzth_zkevmL2.contract.Transact(opts, "init", _addressManager, _param1559)
}

// Init is a paid mutator transaction binding the contract method 0x8f3ca30d.
//
// Solidity: function init(address _addressManager, (uint64,uint64,uint64,uint64,uint64) _param1559) returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) Init(_addressManager common.Address, _param1559 WYzth_zkevmL2EIP1559Params) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.Init(&_WYzth_zkevmL2.TransactOpts, _addressManager, _param1559)
}

// Init is a paid mutator transaction binding the contract method 0x8f3ca30d.
//
// Solidity: function init(address _addressManager, (uint64,uint64,uint64,uint64,uint64) _param1559) returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2TransactorSession) Init(_addressManager common.Address, _param1559 WYzth_zkevmL2EIP1559Params) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.Init(&_WYzth_zkevmL2.TransactOpts, _addressManager, _param1559)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WYzth_zkevmL2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) RenounceOwnership() (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.RenounceOwnership(&_WYzth_zkevmL2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.RenounceOwnership(&_WYzth_zkevmL2.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WYzth_zkevmL2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.TransferOwnership(&_WYzth_zkevmL2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WYzth_zkevmL2 *WYzth_zkevmL2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WYzth_zkevmL2.Contract.TransferOwnership(&_WYzth_zkevmL2.TransactOpts, newOwner)
}

// WYzth_zkevmL2AnchoredIterator is returned from FilterAnchored and is used to iterate over the raw logs and unpacked data for Anchored events raised by the WYzth_zkevmL2 contract.
type WYzth_zkevmL2AnchoredIterator struct {
	Event *WYzth_zkevmL2Anchored // Event containing the contract specifics and raw log

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
func (it *WYzth_zkevmL2AnchoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL2Anchored)
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
		it.Event = new(WYzth_zkevmL2Anchored)
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
func (it *WYzth_zkevmL2AnchoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL2AnchoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL2Anchored represents a Anchored event raised by the WYzth_zkevmL2 contract.
type WYzth_zkevmL2Anchored struct {
	Number     uint64
	Basefee    uint64
	Gaslimit   uint64
	Timestamp  uint64
	ParentHash [32]byte
	Prevrandao *big.Int
	Coinbase   common.Address
	Chainid    uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAnchored is a free log retrieval operation binding the contract event 0x4dcb01f99c4a2c27a16ab38d00ec92434f8231be81fa62e058f260d3c7156029.
//
// Solidity: event Anchored(uint64 number, uint64 basefee, uint64 gaslimit, uint64 timestamp, bytes32 parentHash, uint256 prevrandao, address coinbase, uint32 chainid)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) FilterAnchored(opts *bind.FilterOpts) (*WYzth_zkevmL2AnchoredIterator, error) {

	logs, sub, err := _WYzth_zkevmL2.contract.FilterLogs(opts, "Anchored")
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL2AnchoredIterator{contract: _WYzth_zkevmL2.contract, event: "Anchored", logs: logs, sub: sub}, nil
}

// WatchAnchored is a free log subscription operation binding the contract event 0x4dcb01f99c4a2c27a16ab38d00ec92434f8231be81fa62e058f260d3c7156029.
//
// Solidity: event Anchored(uint64 number, uint64 basefee, uint64 gaslimit, uint64 timestamp, bytes32 parentHash, uint256 prevrandao, address coinbase, uint32 chainid)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) WatchAnchored(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL2Anchored) (event.Subscription, error) {

	logs, sub, err := _WYzth_zkevmL2.contract.WatchLogs(opts, "Anchored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL2Anchored)
				if err := _WYzth_zkevmL2.contract.UnpackLog(event, "Anchored", log); err != nil {
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

// ParseAnchored is a log parse operation binding the contract event 0x4dcb01f99c4a2c27a16ab38d00ec92434f8231be81fa62e058f260d3c7156029.
//
// Solidity: event Anchored(uint64 number, uint64 basefee, uint64 gaslimit, uint64 timestamp, bytes32 parentHash, uint256 prevrandao, address coinbase, uint32 chainid)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) ParseAnchored(log types.Log) (*WYzth_zkevmL2Anchored, error) {
	event := new(WYzth_zkevmL2Anchored)
	if err := _WYzth_zkevmL2.contract.UnpackLog(event, "Anchored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WYzth_zkevmL2CrossChainSyncedIterator is returned from FilterCrossChainSynced and is used to iterate over the raw logs and unpacked data for CrossChainSynced events raised by the WYzth_zkevmL2 contract.
type WYzth_zkevmL2CrossChainSyncedIterator struct {
	Event *WYzth_zkevmL2CrossChainSynced // Event containing the contract specifics and raw log

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
func (it *WYzth_zkevmL2CrossChainSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL2CrossChainSynced)
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
		it.Event = new(WYzth_zkevmL2CrossChainSynced)
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
func (it *WYzth_zkevmL2CrossChainSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL2CrossChainSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL2CrossChainSynced represents a CrossChainSynced event raised by the WYzth_zkevmL2 contract.
type WYzth_zkevmL2CrossChainSynced struct {
	SrcHeight  *big.Int
	BlockHash  [32]byte
	SignalRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) FilterCrossChainSynced(opts *bind.FilterOpts, srcHeight []*big.Int) (*WYzth_zkevmL2CrossChainSyncedIterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _WYzth_zkevmL2.contract.FilterLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL2CrossChainSyncedIterator{contract: _WYzth_zkevmL2.contract, event: "CrossChainSynced", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL2CrossChainSynced, srcHeight []*big.Int) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _WYzth_zkevmL2.contract.WatchLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL2CrossChainSynced)
				if err := _WYzth_zkevmL2.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
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

// ParseCrossChainSynced is a log parse operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) ParseCrossChainSynced(log types.Log) (*WYzth_zkevmL2CrossChainSynced, error) {
	event := new(WYzth_zkevmL2CrossChainSynced)
	if err := _WYzth_zkevmL2.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WYzth_zkevmL2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WYzth_zkevmL2 contract.
type WYzth_zkevmL2InitializedIterator struct {
	Event *WYzth_zkevmL2Initialized // Event containing the contract specifics and raw log

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
func (it *WYzth_zkevmL2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL2Initialized)
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
		it.Event = new(WYzth_zkevmL2Initialized)
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
func (it *WYzth_zkevmL2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL2Initialized represents a Initialized event raised by the WYzth_zkevmL2 contract.
type WYzth_zkevmL2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) FilterInitialized(opts *bind.FilterOpts) (*WYzth_zkevmL2InitializedIterator, error) {

	logs, sub, err := _WYzth_zkevmL2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL2InitializedIterator{contract: _WYzth_zkevmL2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL2Initialized) (event.Subscription, error) {

	logs, sub, err := _WYzth_zkevmL2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL2Initialized)
				if err := _WYzth_zkevmL2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) ParseInitialized(log types.Log) (*WYzth_zkevmL2Initialized, error) {
	event := new(WYzth_zkevmL2Initialized)
	if err := _WYzth_zkevmL2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WYzth_zkevmL2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WYzth_zkevmL2 contract.
type WYzth_zkevmL2OwnershipTransferredIterator struct {
	Event *WYzth_zkevmL2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WYzth_zkevmL2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WYzth_zkevmL2OwnershipTransferred)
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
		it.Event = new(WYzth_zkevmL2OwnershipTransferred)
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
func (it *WYzth_zkevmL2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WYzth_zkevmL2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WYzth_zkevmL2OwnershipTransferred represents a OwnershipTransferred event raised by the WYzth_zkevmL2 contract.
type WYzth_zkevmL2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WYzth_zkevmL2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WYzth_zkevmL2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WYzth_zkevmL2OwnershipTransferredIterator{contract: _WYzth_zkevmL2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WYzth_zkevmL2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WYzth_zkevmL2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WYzth_zkevmL2OwnershipTransferred)
				if err := _WYzth_zkevmL2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WYzth_zkevmL2 *WYzth_zkevmL2Filterer) ParseOwnershipTransferred(log types.Log) (*WYzth_zkevmL2OwnershipTransferred, error) {
	event := new(WYzth_zkevmL2OwnershipTransferred)
	if err := _WYzth_zkevmL2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
