package network

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/avax"
)

// Get the current network node demand in ETH
func GetNodeDemand(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoNetworkFees, err := getGoGoNetworkFees(ggp)
	if err != nil {
		return nil, err
	}
	nodeDemand := new(*big.Int)
	if err := gogoNetworkFees.Call(opts, nodeDemand, "getNodeDemand"); err != nil {
		return nil, fmt.Errorf("Could not get network node demand: %w", err)
	}
	return *nodeDemand, nil
}

// Get the current network node commission rate
func GetNodeFee(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	gogoNetworkFees, err := getGoGoNetworkFees(ggp)
	if err != nil {
		return 0, err
	}
	nodeFee := new(*big.Int)
	if err := gogoNetworkFees.Call(opts, nodeFee, "getNodeFee"); err != nil {
		return 0, fmt.Errorf("Could not get network node fee: %w", err)
	}
	return avax.WeiToEth(*nodeFee), nil
}

// Get the network node fee for a node demand value
func GetNodeFeeByDemand(ggp *gogopool.GoGoPool, nodeDemand *big.Int, opts *bind.CallOpts) (float64, error) {
	gogoNetworkFees, err := getGoGoNetworkFees(ggp)
	if err != nil {
		return 0, err
	}
	nodeFee := new(*big.Int)
	if err := gogoNetworkFees.Call(opts, nodeFee, "getNodeFeeByDemand", nodeDemand); err != nil {
		return 0, fmt.Errorf("Could not get node fee by node demand: %w", err)
	}
	return avax.WeiToEth(*nodeFee), nil
}

// Get contracts
var gogoNetworkFeesLock sync.Mutex

func getGoGoNetworkFees(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoNetworkFeesLock.Lock()
	defer gogoNetworkFeesLock.Unlock()
	return ggp.GetContract("rocketNetworkFees")
}
