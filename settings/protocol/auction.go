package protocol

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	protocoldao "github.com/multisig-labs/gogopool-go/dao/protocol"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/avax"
)

// Config
const AuctionSettingsContractName = "gogoDAOProtocolSettingsAuction"

// Lot creation currently enabled
func GetCreateLotEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(ggp)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := auctionSettingsContract.Call(opts, value, "getCreateLotEnabled"); err != nil {
		return false, fmt.Errorf("Could not get lot creation enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapCreateLotEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(ggp, AuctionSettingsContractName, "auction.lot.create.enabled", value, opts)
}

// Lot bidding currently enabled
func GetBidOnLotEnabled(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (bool, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(ggp)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := auctionSettingsContract.Call(opts, value, "getBidOnLotEnabled"); err != nil {
		return false, fmt.Errorf("Could not get lot bidding enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapBidOnLotEnabled(ggp *gogopool.GoGoPool, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(ggp, AuctionSettingsContractName, "auction.lot.bidding.enabled", value, opts)
}

// The minimum lot size in ETH value
func GetLotMinimumEthValue(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(ggp)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getLotMinimumEthValue"); err != nil {
		return nil, fmt.Errorf("Could not get lot minimum ETH value: %w", err)
	}
	return *value, nil
}
func BootstrapLotMinimumEthValue(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, AuctionSettingsContractName, "auction.lot.value.minimum", value, opts)
}

// The maximum lot size in ETH value
func GetLotMaximumEthValue(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(ggp)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getLotMaximumEthValue"); err != nil {
		return nil, fmt.Errorf("Could not get lot maximum ETH value: %w", err)
	}
	return *value, nil
}
func BootstrapLotMaximumEthValue(ggp *gogopool.GoGoPool, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, AuctionSettingsContractName, "auction.lot.value.maximum", value, opts)
}

// The lot duration in blocks
func GetLotDuration(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getLotDuration"); err != nil {
		return 0, fmt.Errorf("Could not get lot duration: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapLotDuration(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, AuctionSettingsContractName, "auction.lot.duration", big.NewInt(int64(value)), opts)
}

// The starting price relative to current ETH price, as a fraction
func GetLotStartingPriceRatio(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getStartingPriceRatio"); err != nil {
		return 0, fmt.Errorf("Could not get lot starting price ratio: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapLotStartingPriceRatio(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, AuctionSettingsContractName, "auction.price.start", avax.EthToWei(value), opts)
}

// The reserve price relative to current ETH price, as a fraction
func GetLotReservePriceRatio(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getReservePriceRatio"); err != nil {
		return 0, fmt.Errorf("Could not get lot reserve price ratio: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapLotReservePriceRatio(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, AuctionSettingsContractName, "auction.price.reserve", avax.EthToWei(value), opts)
}

// Get contracts
var auctionSettingsContractLock sync.Mutex

func getAuctionSettingsContract(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	auctionSettingsContractLock.Lock()
	defer auctionSettingsContractLock.Unlock()
	return ggp.GetContract(AuctionSettingsContractName)
}
