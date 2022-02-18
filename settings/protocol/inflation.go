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
const InflationSettingsContractName = "gogoDAOProtocolSettingsInflation"

// GGP inflation rate per interval
func GetInflationIntervalRate(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (float64, error) {
	inflationSettingsContract, err := getInflationSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := inflationSettingsContract.Call(opts, value, "getInflationIntervalRate"); err != nil {
		return 0, fmt.Errorf("Could not get inflation rate: %w", err)
	}
	return avax.WeiToEth(*value), nil
}
func BootstrapInflationIntervalRate(ggp *gogopool.GoGoPool, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, InflationSettingsContractName, "ggp.inflation.interval.rate", avax.EthToWei(value), opts)
}

// GGP inflation start time
func GetInflationStartTime(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	inflationSettingsContract, err := getInflationSettingsContract(ggp)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := inflationSettingsContract.Call(opts, value, "getInflationIntervalStartTime"); err != nil {
		return 0, fmt.Errorf("Could not get inflation start time: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapInflationStartTime(ggp *gogopool.GoGoPool, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(ggp, InflationSettingsContractName, "ggp.inflation.interval.start", big.NewInt(int64(value)), opts)
}

// Get contracts
var inflationSettingsContractLock sync.Mutex

func getInflationSettingsContract(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	inflationSettingsContractLock.Lock()
	defer inflationSettingsContractLock.Unlock()
	return ggp.GetContract(InflationSettingsContractName)
}
