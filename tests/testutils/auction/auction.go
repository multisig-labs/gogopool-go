package auction

import (
	"fmt"
	"testing"

	"github.com/multisig-labs/gogopool-go/deposit"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/minipool"
	"github.com/multisig-labs/gogopool-go/settings/trustednode"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	minipoolutils "github.com/multisig-labs/gogopool-go/tests/testutils/minipool"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

// Create an amount of slashed GGP in the auction contract
func CreateSlashedGGP(t *testing.T, ggp *gogopool.GoGoPool, ownerAccount *accounts.Account, trustedNodeAccount, trustedNodeAccount2 *accounts.Account, userAccount *accounts.Account) error {

	// Stake a large amount of GGP against the node
	if err := nodeutils.StakeGGP(ggp, ownerAccount, trustedNodeAccount, avax.EthToWei(1000000)); err != nil {
		return err
	}

	// Make user deposit
	depositOpts := userAccount.GetTransactor()
	depositOpts.Value = avax.EthToWei(16)
	if _, err := deposit.Deposit(ggp, depositOpts); err != nil {
		return err
	}

	// Create unbonded minipool
	mp, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, trustedNodeAccount, avax.EthToWei(16), 1)
	if err != nil {
		return err
	}

	// Deposit user ETH to minipool
	opts := userAccount.GetTransactor()
	opts.Value = avax.EthToWei(16)
	if _, err := deposit.Deposit(ggp, opts); err != nil {
		return err
	}

	// Delay for the time between depositing and staking
	scrubPeriod, err := trustednode.GetScrubPeriod(ggp, nil)
	if err != nil {
		return err
	}
	err = evm.IncreaseTime(int(scrubPeriod + 1))
	if err != nil {
		return fmt.Errorf("Could not increase time: %w", err)
	}

	// Stake minipool
	if err := minipoolutils.StakeMinipool(ggp, mp, trustedNodeAccount); err != nil {
		return err
	}

	// Mark minipool as withdrawable with zero end balance
	if _, err := minipool.SubmitMinipoolWithdrawable(ggp, mp.Address, trustedNodeAccount.GetTransactor()); err != nil {
		return err
	}
	if _, err := minipool.SubmitMinipoolWithdrawable(ggp, mp.Address, trustedNodeAccount2.GetTransactor()); err != nil {
		return err
	}

	// Distribute balance and finalise pool to send slashed GGP to auction contract
	if _, err := mp.DistributeBalanceAndFinalise(trustedNodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Return
	return nil

}
