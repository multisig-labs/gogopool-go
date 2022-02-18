package minipool

import (
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/minipool"
	"github.com/multisig-labs/gogopool-go/network"
	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
	"github.com/multisig-labs/gogopool-go/tests/testutils/validator"
)

// Minipool created event
type minipoolCreated struct {
	Minipool common.Address
	Node     common.Address
	Time     *big.Int
}

// Create a minipool
func CreateMinipool(t *testing.T, ggp *gogopool.GoGoPool, ownerAccount, nodeAccount *accounts.Account, depositAmount *big.Int, pubkey int) (*minipool.Minipool, error) {

	// Mint & stake GGP required for mininpool
	ggpRequired, err := GetMinipoolGGPRequired(ggp)
	if err != nil {
		return nil, err
	}
	if err := nodeutils.StakeGGP(ggp, ownerAccount, nodeAccount, ggpRequired); err != nil {
		return nil, err
	}

	// Do the node deposit to generate the minipool
	expectedMinipoolAddress, txReceipt, err := nodeutils.Deposit(t, ggp, nodeAccount, depositAmount, pubkey)
	if err != nil {
		return nil, fmt.Errorf("Could not do node deposit: %w", err)
	}

	// Get minipool manager contract
	gogoMinipoolManager, err := ggp.GetContract("gogoMinipoolManager")
	if err != nil {
		return nil, err
	}

	// Get created minipool address
	minipoolCreatedEvents, err := gogoMinipoolManager.GetTransactionEvents(txReceipt, "MinipoolCreated", minipoolCreated{})
	if err != nil || len(minipoolCreatedEvents) == 0 {
		return nil, errors.New("Could not get minipool created event")
	}
	minipoolAddress := minipoolCreatedEvents[0].(minipoolCreated).Minipool

	// Sanity check to verify the created minipool is at the expected address
	if expectedMinipoolAddress != minipoolAddress {
		return nil, errors.New(fmt.Sprintf("Expected minipool address %s but got %s", expectedMinipoolAddress.Hex(), minipoolAddress.Hex()))
	}

	// Return minipool instance
	return minipool.NewMinipool(ggp, minipoolAddress)

}

// Stake a minipool
func StakeMinipool(ggp *gogopool.GoGoPool, mp *minipool.Minipool, nodeAccount *accounts.Account) error {

	// Get validator & deposit data
	validatorPubkey, err := validator.GetValidatorPubkey(1)
	if err != nil {
		return err
	}
	withdrawalCredentials, err := minipool.GetMinipoolWithdrawalCredentials(ggp, mp.Address, nil)
	if err != nil {
		return err
	}
	validatorSignature, err := validator.GetValidatorSignature(1)
	if err != nil {
		return err
	}
	depositDataRoot, err := validator.GetDepositDataRoot(validatorPubkey, withdrawalCredentials, validatorSignature)
	if err != nil {
		return err
	}

	// Stake minipool & return
	_, err = mp.Stake(validatorSignature, depositDataRoot, nodeAccount.GetTransactor())
	return err

}

// Get the GGP required per minipool
func GetMinipoolGGPRequired(ggp *gogopool.GoGoPool) (*big.Int, error) {

	// Get data
	depositUserAmount, err := protocol.GetMinipoolHalfDepositUserAmount(ggp, nil)
	if err != nil {
		return nil, err
	}
	minimumPerMinipoolStake, err := protocol.GetMinimumPerMinipoolStake(ggp, nil)
	if err != nil {
		return nil, err
	}
	ggpPrice, err := network.GetGGPPrice(ggp, nil)
	if err != nil {
		return nil, err
	}

	// Calculate and return GGP required
	var tmp big.Int
	var ggpRequired big.Int
	tmp.Mul(depositUserAmount, avax.EthToWei(minimumPerMinipoolStake))
	ggpRequired.Quo(&tmp, ggpPrice)
	return &ggpRequired, nil

}
