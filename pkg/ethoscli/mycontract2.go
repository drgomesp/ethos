// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethoscli

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

// MyContract2MetaData contains all meta data concerning the MyContract2 contract.
var MyContract2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"purpose\",\"type\":\"string\"}],\"name\":\"SetPurpose\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"purpose\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newPurpose\",\"type\":\"string\"}],\"name\":\"setPurpose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526040518060400160405280601c81526020017f4275696c64696e6720556e73746f707061626c652041707073212121000000008152506000908051906020019061004f929190610055565b50610158565b82805461006190610127565b90600052602060002090601f01602090048101928261008357600085556100ca565b82601f1061009c57805160ff19168380011785556100ca565b828001600101855582156100ca579182015b828111156100c95782518255916020019190600101906100ae565b5b5090506100d791906100db565b5090565b5b808211156100f45760008160009055506001016100dc565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061013f57607f821691505b602082108103610152576101516100f8565b5b50919050565b6105f8806101676000396000f3fe60806040526004361061002d5760003560e01c806370740aab14610036578063eb68757f1461006157610034565b3661003457005b005b34801561004257600080fd5b5061004b61008a565b60405161005891906102a8565b60405180910390f35b34801561006d57600080fd5b5061008860048036038101906100839190610413565b610118565b005b600080546100979061048b565b80601f01602080910402602001604051908101604052809291908181526020018280546100c39061048b565b80156101105780601f106100e557610100808354040283529160200191610110565b820191906000526020600020905b8154815290600101906020018083116100f357829003601f168201915b505050505081565b806000908051906020019061012e92919061016c565b507f6ea5d6383a120235c7728a9a6751672a8ac068e4ed34dcca2ee444182c1812de336000604051610161929190610592565b60405180910390a150565b8280546101789061048b565b90600052602060002090601f01602090048101928261019a57600085556101e1565b82601f106101b357805160ff19168380011785556101e1565b828001600101855582156101e1579182015b828111156101e05782518255916020019190600101906101c5565b5b5090506101ee91906101f2565b5090565b5b8082111561020b5760008160009055506001016101f3565b5090565b600081519050919050565b600082825260208201905092915050565b60005b8381101561024957808201518184015260208101905061022e565b83811115610258576000848401525b50505050565b6000601f19601f8301169050919050565b600061027a8261020f565b610284818561021a565b935061029481856020860161022b565b61029d8161025e565b840191505092915050565b600060208201905081810360008301526102c2818461026f565b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103208261025e565b810181811067ffffffffffffffff8211171561033f5761033e6102e8565b5b80604052505050565b60006103526102ca565b905061035e8282610317565b919050565b600067ffffffffffffffff82111561037e5761037d6102e8565b5b6103878261025e565b9050602081019050919050565b82818337600083830152505050565b60006103b66103b184610363565b610348565b9050828152602081018484840111156103d2576103d16102e3565b5b6103dd848285610394565b509392505050565b600082601f8301126103fa576103f96102de565b5b813561040a8482602086016103a3565b91505092915050565b600060208284031215610429576104286102d4565b5b600082013567ffffffffffffffff811115610447576104466102d9565b5b610453848285016103e5565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806104a357607f821691505b6020821081036104b6576104b561045c565b5b50919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104e7826104bc565b9050919050565b6104f7816104dc565b82525050565b60008190508160005260206000209050919050565b6000815461051f8161048b565b610529818661021a565b94506001821660008114610544576001811461055657610589565b60ff1983168652602086019350610589565b61055f856104fd565b60005b8381101561058157815481890152600182019150602081019050610562565b808801955050505b50505092915050565b60006040820190506105a760008301856104ee565b81810360208301526105b98184610512565b9050939250505056fea2646970667358221220c4721efcac35de3f60424ad5183634edec0daf4ba271644c0b2fdbd0a42bf60f64736f6c634300080d0033",
}

// MyContract2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MyContract2MetaData.ABI instead.
var MyContract2ABI = MyContract2MetaData.ABI

// MyContract2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyContract2MetaData.Bin instead.
var MyContract2Bin = MyContract2MetaData.Bin

// DeployMyContract2 deploys a new Ethereum contract, binding an instance of MyContract2 to it.
func DeployMyContract2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyContract2, error) {
	parsed, err := MyContract2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyContract2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyContract2{MyContract2Caller: MyContract2Caller{contract: contract}, MyContract2Transactor: MyContract2Transactor{contract: contract}, MyContract2Filterer: MyContract2Filterer{contract: contract}}, nil
}

// MyContract2 is an auto generated Go binding around an Ethereum contract.
type MyContract2 struct {
	MyContract2Caller     // Read-only binding to the contract
	MyContract2Transactor // Write-only binding to the contract
	MyContract2Filterer   // Log filterer for contract events
}

// MyContract2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyContract2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContract2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyContract2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContract2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyContract2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContract2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyContract2Session struct {
	Contract     *MyContract2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyContract2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyContract2CallerSession struct {
	Contract *MyContract2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MyContract2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyContract2TransactorSession struct {
	Contract     *MyContract2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MyContract2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyContract2Raw struct {
	Contract *MyContract2 // Generic contract binding to access the raw methods on
}

// MyContract2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyContract2CallerRaw struct {
	Contract *MyContract2Caller // Generic read-only contract binding to access the raw methods on
}

// MyContract2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyContract2TransactorRaw struct {
	Contract *MyContract2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyContract2 creates a new instance of MyContract2, bound to a specific deployed contract.
func NewMyContract2(address common.Address, backend bind.ContractBackend) (*MyContract2, error) {
	contract, err := bindMyContract2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyContract2{MyContract2Caller: MyContract2Caller{contract: contract}, MyContract2Transactor: MyContract2Transactor{contract: contract}, MyContract2Filterer: MyContract2Filterer{contract: contract}}, nil
}

// NewMyContract2Caller creates a new read-only instance of MyContract2, bound to a specific deployed contract.
func NewMyContract2Caller(address common.Address, caller bind.ContractCaller) (*MyContract2Caller, error) {
	contract, err := bindMyContract2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyContract2Caller{contract: contract}, nil
}

// NewMyContract2Transactor creates a new write-only instance of MyContract2, bound to a specific deployed contract.
func NewMyContract2Transactor(address common.Address, transactor bind.ContractTransactor) (*MyContract2Transactor, error) {
	contract, err := bindMyContract2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyContract2Transactor{contract: contract}, nil
}

// NewMyContract2Filterer creates a new log filterer instance of MyContract2, bound to a specific deployed contract.
func NewMyContract2Filterer(address common.Address, filterer bind.ContractFilterer) (*MyContract2Filterer, error) {
	contract, err := bindMyContract2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyContract2Filterer{contract: contract}, nil
}

// bindMyContract2 binds a generic wrapper to an already deployed contract.
func bindMyContract2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyContract2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract2 *MyContract2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract2.Contract.MyContract2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract2 *MyContract2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract2.Contract.MyContract2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract2 *MyContract2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract2.Contract.MyContract2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract2 *MyContract2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract2 *MyContract2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract2 *MyContract2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract2.Contract.contract.Transact(opts, method, params...)
}

// Purpose is a free data retrieval call binding the contract method 0x70740aab.
//
// Solidity: function purpose() view returns(string)
func (_MyContract2 *MyContract2Caller) Purpose(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyContract2.contract.Call(opts, &out, "purpose")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Purpose is a free data retrieval call binding the contract method 0x70740aab.
//
// Solidity: function purpose() view returns(string)
func (_MyContract2 *MyContract2Session) Purpose() (string, error) {
	return _MyContract2.Contract.Purpose(&_MyContract2.CallOpts)
}

// Purpose is a free data retrieval call binding the contract method 0x70740aab.
//
// Solidity: function purpose() view returns(string)
func (_MyContract2 *MyContract2CallerSession) Purpose() (string, error) {
	return _MyContract2.Contract.Purpose(&_MyContract2.CallOpts)
}

// SetPurpose is a paid mutator transaction binding the contract method 0xeb68757f.
//
// Solidity: function setPurpose(string newPurpose) returns()
func (_MyContract2 *MyContract2Transactor) SetPurpose(opts *bind.TransactOpts, newPurpose string) (*types.Transaction, error) {
	return _MyContract2.contract.Transact(opts, "setPurpose", newPurpose)
}

// SetPurpose is a paid mutator transaction binding the contract method 0xeb68757f.
//
// Solidity: function setPurpose(string newPurpose) returns()
func (_MyContract2 *MyContract2Session) SetPurpose(newPurpose string) (*types.Transaction, error) {
	return _MyContract2.Contract.SetPurpose(&_MyContract2.TransactOpts, newPurpose)
}

// SetPurpose is a paid mutator transaction binding the contract method 0xeb68757f.
//
// Solidity: function setPurpose(string newPurpose) returns()
func (_MyContract2 *MyContract2TransactorSession) SetPurpose(newPurpose string) (*types.Transaction, error) {
	return _MyContract2.Contract.SetPurpose(&_MyContract2.TransactOpts, newPurpose)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MyContract2 *MyContract2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MyContract2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MyContract2 *MyContract2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MyContract2.Contract.Fallback(&_MyContract2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MyContract2 *MyContract2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MyContract2.Contract.Fallback(&_MyContract2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyContract2 *MyContract2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyContract2 *MyContract2Session) Receive() (*types.Transaction, error) {
	return _MyContract2.Contract.Receive(&_MyContract2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyContract2 *MyContract2TransactorSession) Receive() (*types.Transaction, error) {
	return _MyContract2.Contract.Receive(&_MyContract2.TransactOpts)
}

// MyContract2SetPurposeIterator is returned from FilterSetPurpose and is used to iterate over the raw logs and unpacked data for SetPurpose events raised by the MyContract2 contract.
type MyContract2SetPurposeIterator struct {
	Event *MyContract2SetPurpose // Event containing the contract specifics and raw log

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
func (it *MyContract2SetPurposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContract2SetPurpose)
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
		it.Event = new(MyContract2SetPurpose)
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
func (it *MyContract2SetPurposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContract2SetPurposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContract2SetPurpose represents a SetPurpose event raised by the MyContract2 contract.
type MyContract2SetPurpose struct {
	Sender  common.Address
	Purpose string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetPurpose is a free log retrieval operation binding the contract event 0x6ea5d6383a120235c7728a9a6751672a8ac068e4ed34dcca2ee444182c1812de.
//
// Solidity: event SetPurpose(address sender, string purpose)
func (_MyContract2 *MyContract2Filterer) FilterSetPurpose(opts *bind.FilterOpts) (*MyContract2SetPurposeIterator, error) {

	logs, sub, err := _MyContract2.contract.FilterLogs(opts, "SetPurpose")
	if err != nil {
		return nil, err
	}
	return &MyContract2SetPurposeIterator{contract: _MyContract2.contract, event: "SetPurpose", logs: logs, sub: sub}, nil
}

// WatchSetPurpose is a free log subscription operation binding the contract event 0x6ea5d6383a120235c7728a9a6751672a8ac068e4ed34dcca2ee444182c1812de.
//
// Solidity: event SetPurpose(address sender, string purpose)
func (_MyContract2 *MyContract2Filterer) WatchSetPurpose(opts *bind.WatchOpts, sink chan<- *MyContract2SetPurpose) (event.Subscription, error) {

	logs, sub, err := _MyContract2.contract.WatchLogs(opts, "SetPurpose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContract2SetPurpose)
				if err := _MyContract2.contract.UnpackLog(event, "SetPurpose", log); err != nil {
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

// ParseSetPurpose is a log parse operation binding the contract event 0x6ea5d6383a120235c7728a9a6751672a8ac068e4ed34dcca2ee444182c1812de.
//
// Solidity: event SetPurpose(address sender, string purpose)
func (_MyContract2 *MyContract2Filterer) ParseSetPurpose(log types.Log) (*MyContract2SetPurpose, error) {
	event := new(MyContract2SetPurpose)
	if err := _MyContract2.contract.UnpackLog(event, "SetPurpose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
