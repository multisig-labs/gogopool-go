package rewards

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/multisig-labs/gogopool-go/deposit"
	"github.com/multisig-labs/gogopool-go/node"
	"github.com/multisig-labs/gogopool-go/rewards"
	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/settings/trustednode"
	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	minipoolutils "github.com/multisig-labs/gogopool-go/tests/testutils/minipool"
	"github.com/multisig-labs/gogopool-go/tokens"
	"github.com/multisig-labs/gogopool-go/utils/avax"
)

func TestNodeRewards(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Constants
	oneDay := 24 * 60 * 60
	rewardInterval := oneDay

	// Register node
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Set network parameters
	if _, err := protocol.BootstrapRewardsClaimIntervalTime(ggp, uint64(rewardInterval), ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check node claims enabled status
	if claimsEnabled, err := rewards.GetNodeClaimsEnabled(ggp, nil); err != nil {
		t.Error(err)
	} else if !claimsEnabled {
		t.Error("Incorrect node claims enabled status")
	}

	// Get & check initial node claim possible status
	if nodeClaimPossible, err := rewards.GetNodeClaimPossible(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeClaimPossible {
		t.Error("Incorrect initial node claim possible status")
	}

	// Increase time until node claims are possible
	if err := evm.IncreaseTime(rewardInterval); err != nil {
		t.Fatal(err)
	}

	// Get & check updated node claim possible status
	if nodeClaimPossible, err := rewards.GetNodeClaimPossible(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if !nodeClaimPossible {
		t.Error("Incorrect updated node claim possible status")
	}

	// Get & check initial node claim rewards percent
	if rewardsPerc, err := rewards.GetNodeClaimRewardsPerc(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if rewardsPerc != 0 {
		t.Errorf("Incorrect initial node claim rewards perc %f", rewardsPerc)
	}

	// Stake GGP & create a minipool
	mp, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, nodeAccount, avax.EthToWei(16), 1)
	if err != nil {
		t.Fatal(err)
	}

	// Deposit user ETH to minipool
	opts := nodeAccount.GetTransactor()
	opts.Value = avax.EthToWei(16)
	if _, err := deposit.Deposit(ggp, opts); err != nil {
		t.Error(err)
	}

	// Delay for the time between depositing and staking
	scrubPeriod, err := trustednode.GetScrubPeriod(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = evm.IncreaseTime(int(scrubPeriod + 1))
	if err != nil {
		t.Fatal(fmt.Errorf("Could not increase time: %w", err))
	}

	// Stake minipool
	if err := minipoolutils.StakeMinipool(ggp, mp, nodeAccount); err != nil {
		t.Error(err)
	}

	// Get & check updated node claim rewards percent
	if rewardsPerc, err := rewards.GetNodeClaimRewardsPerc(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if rewardsPerc != 1 {
		t.Errorf("Incorrect updated node claim rewards perc %f", rewardsPerc)
	}

	// Get & check initial node claim rewards amount
	if rewardsAmount, err := rewards.GetNodeClaimRewardsAmount(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if rewardsAmount.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial node claim rewards amount %s", rewardsAmount.String())
	}

	// Start GGP inflation
	if header, err := ggp.Client.HeaderByNumber(context.Background(), nil); err != nil {
		t.Fatal(err)
	} else if _, err := protocol.BootstrapInflationStartTime(ggp, header.Time+uint64(oneDay), ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Increase time until rewards are available
	if err := evm.IncreaseTime(oneDay + oneDay); err != nil {
		t.Fatal(err)
	}

	// Get & check updated node claim rewards amount
	if rewardsAmount, err := rewards.GetNodeClaimRewardsAmount(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if rewardsAmount.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Incorrect updated node claim rewards amount %s", rewardsAmount.String())
	}

	// Get & check initial node GGP balance
	if ggpBalance, err := tokens.GetGGPBalance(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if ggpBalance.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial node GGP balance %s", ggpBalance.String())
	}

	// Claim node rewards
	if _, err := rewards.ClaimNodeRewards(ggp, nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated node GGP balance
	if ggpBalance, err := tokens.GetGGPBalance(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if ggpBalance.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Incorrect updated node GGP balance %s", ggpBalance.String())
	}

}
