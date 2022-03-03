package node

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	ggptypes "github.com/multisig-labs/gogopool-go/types"
	"github.com/multisig-labs/gogopool-go/utils/avax"
)

// Estimate the gas of Deposit
func EstimateDepositGas(ggp *gogopool.GoGoPool, minimumNodeFee float64, validatorPubkey ggptypes.ValidatorPubkey, validatorSignature ggptypes.ValidatorSignature, depositDataRoot common.Hash, salt *big.Int, expectedMinipoolAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoNodeDeposit, err := getGoGoNodeDeposit(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoNodeDeposit.GetTransactionGasInfo(opts, "deposit", avax.EthToWei(minimumNodeFee), validatorPubkey[:], validatorSignature[:], depositDataRoot, salt, expectedMinipoolAddress)
}

// Make a node deposit
func Deposit(ggp *gogopool.GoGoPool, minimumNodeFee float64, validatorPubkey ggptypes.ValidatorPubkey, validatorSignature ggptypes.ValidatorSignature, depositDataRoot common.Hash, salt *big.Int, expectedMinipoolAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	gogoNodeDeposit, err := getGoGoNodeDeposit(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoNodeDeposit.Transact(opts, "deposit", avax.EthToWei(minimumNodeFee), validatorPubkey[:], validatorSignature[:], depositDataRoot, salt, expectedMinipoolAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not make node deposit: %w", err)
	}
	return hash, nil
}

// Get the type of a deposit based on the amount
func GetDepositType(ggp *gogopool.GoGoPool, amount *big.Int, opts *bind.CallOpts) (ggptypes.MinipoolDeposit, error) {
	gogoNodeDeposit, err := getGoGoNodeDeposit(ggp)
	if err != nil {
		return ggptypes.Empty, err
	}

	depositType := new(uint8)
	if err := gogoNodeDeposit.Call(opts, depositType, "getDepositType", amount); err != nil {
		return ggptypes.Empty, fmt.Errorf("Could not get deposit type: %w", err)
	}
	return ggptypes.MinipoolDeposit(*depositType), nil
}

// Get contracts
var gogoNodeDepositLock sync.Mutex

func getGoGoNodeDeposit(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoNodeDepositLock.Lock()
	defer gogoNodeDepositLock.Unlock()
	return ggp.GetContract("rocketNodeDeposit")
}
