package node

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/multisig-labs/gogopool-go/deposit"
	"github.com/multisig-labs/gogopool-go/minipool"
	"github.com/multisig-labs/gogopool-go/node"
	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/settings/trustednode"
	"github.com/multisig-labs/gogopool-go/tokens"
	"github.com/multisig-labs/gogopool-go/utils/eth"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	minipoolutils "github.com/multisig-labs/gogopool-go/tests/testutils/minipool"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
	ggputils "github.com/multisig-labs/gogopool-go/tests/testutils/tokens/ggp"
)

func TestStakeGGP(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Register node
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get GGP amount required for 2 minipools
	minipoolGgpRequired, err := minipoolutils.GetMinipoolGGPRequired(ggp)
	if err != nil {
		t.Fatal(err)
	}
	ggpAmount := new(big.Int)
	ggpAmount.Mul(minipoolGgpRequired, big.NewInt(2))

	// Mint GGP
	if err := ggputils.MintGGP(ggp, ownerAccount, nodeAccount, ggpAmount); err != nil {
		t.Fatal(err)
	}

	// Approve GGP transfer for staking
	gogoNodeStakingAddress, err := ggp.GetAddress("gogoNodeStaking")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tokens.ApproveGGP(ggp, *gogoNodeStakingAddress, ggpAmount, nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Check initial staking details
	if totalGgpStake, err := node.GetTotalGGPStake(ggp, nil); err != nil {
		t.Error(err)
	} else if totalGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial total GGP stake %s", totalGgpStake.String())
	}
	if totalEffectiveGgpStake, err := node.GetTotalEffectiveGGPStake(ggp, nil); err != nil {
		t.Error(err)
	} else if totalEffectiveGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial total effective GGP stake %s", totalEffectiveGgpStake.String())
	}
	if nodeGgpStake, err := node.GetNodeGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial node GGP stake %s", nodeGgpStake.String())
	}
	if nodeEffectiveGgpStake, err := node.GetNodeEffectiveGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeEffectiveGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial node effective GGP stake %s", nodeEffectiveGgpStake.String())
	}
	if nodeMinimumGgpStake, err := node.GetNodeMinimumGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeMinimumGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial node minimum GGP stake %s", nodeMinimumGgpStake.String())
	}
	if nodeGgpStakedTime, err := node.GetNodeGGPStakedTime(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeGgpStakedTime != 0 {
		t.Errorf("Incorrect initial node GGP staked time %d", nodeGgpStakedTime)
	}
	if nodeMinipoolLimit, err := node.GetNodeMinipoolLimit(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeMinipoolLimit != 0 {
		t.Errorf("Incorrect initial node minipool limit %d", nodeMinipoolLimit)
	}

	// Stake GGP
	if _, err := node.StakeGGP(ggp, ggpAmount, nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Check updated staking details
	if totalGgpStake, err := node.GetTotalGGPStake(ggp, nil); err != nil {
		t.Error(err)
	} else if totalGgpStake.Cmp(ggpAmount) != 0 {
		t.Errorf("Incorrect updated total GGP stake 1 %s", totalGgpStake.String())
	}
	if totalEffectiveGgpStake, err := node.GetTotalEffectiveGGPStake(ggp, nil); err != nil {
		t.Error(err)
	} else if totalEffectiveGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect updated total effective GGP stake 1 %s", totalEffectiveGgpStake.String())
	}
	if nodeGgpStake, err := node.GetNodeGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeGgpStake.Cmp(ggpAmount) != 0 {
		t.Errorf("Incorrect updated node GGP stake 1 %s", nodeGgpStake.String())
	}
	if nodeEffectiveGgpStake, err := node.GetNodeEffectiveGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeEffectiveGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect updated node effective GGP stake 1 %s", nodeEffectiveGgpStake.String())
	}
	if nodeMinimumGgpStake, err := node.GetNodeMinimumGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeMinimumGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect updated node minimum GGP stake 1 %s", nodeMinimumGgpStake.String())
	}
	if nodeGgpStakedTime, err := node.GetNodeGGPStakedTime(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeGgpStakedTime == 0 {
		t.Errorf("Incorrect updated node GGP staked time 1 %d", nodeGgpStakedTime)
	}
	if nodeMinipoolLimit, err := node.GetNodeMinipoolLimit(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeMinipoolLimit != 2 {
		t.Errorf("Incorrect updated node minipool limit 1 %d", nodeMinipoolLimit)
	}

	// Make node deposit to create minipool
	minipoolAddress, _, err := nodeutils.Deposit(t, ggp, nodeAccount, avax.EthToWei(16), 1)
	if err != nil {
		t.Fatal(err)
	}
	mp, err := minipool.NewMinipool(ggp, minipoolAddress)
	if err != nil {
		t.Fatal(err)
	}

	// Make user deposit
	depositOpts := nodeAccount.GetTransactor()
	depositOpts.Value = avax.EthToWei(16)
	if _, err := deposit.Deposit(ggp, depositOpts); err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}

	// Check updated staking details
	if totalEffectiveGgpStake, err := node.GetTotalEffectiveGGPStake(ggp, nil); err != nil {
		t.Error(err)
	} else if totalEffectiveGgpStake.Cmp(ggpAmount) != 0 {
		t.Errorf("Incorrect updated total effective GGP stake 2 %s", totalEffectiveGgpStake.String())
	}
	if nodeEffectiveGgpStake, err := node.GetNodeEffectiveGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeEffectiveGgpStake.Cmp(ggpAmount) != 0 {
		t.Errorf("Incorrect updated node effective GGP stake 2 %s", nodeEffectiveGgpStake.String())
	}
	if nodeMinimumGgpStake, err := node.GetNodeMinimumGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeMinimumGgpStake.Cmp(minipoolGgpRequired) != 0 {
		t.Errorf("Incorrect updated node minimum GGP stake 2 %s", nodeMinimumGgpStake.String())
	}

}

func TestWithdrawGGP(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Register node
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Mint & stake GGP
	ggpAmount := avax.EthToWei(1000)
	if err := nodeutils.StakeGGP(ggp, ownerAccount, nodeAccount, ggpAmount); err != nil {
		t.Fatal(err)
	}

	// Get & set rewards claim interval
	rewardsClaimIntervalTime, err := protocol.GetRewardsClaimIntervalTime(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := protocol.BootstrapRewardsClaimIntervalTime(ggp, 0, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Check initial staking details
	if totalGgpStake, err := node.GetTotalGGPStake(ggp, nil); err != nil {
		t.Error(err)
	} else if totalGgpStake.Cmp(ggpAmount) != 0 {
		t.Errorf("Incorrect initial total GGP stake %s", totalGgpStake.String())
	}
	if nodeGgpStake, err := node.GetNodeGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeGgpStake.Cmp(ggpAmount) != 0 {
		t.Errorf("Incorrect initial node GGP stake %s", nodeGgpStake.String())
	}

	// Withdraw GGP
	if _, err := node.WithdrawGGP(ggp, ggpAmount, nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Check updated staking details
	if totalGgpStake, err := node.GetTotalGGPStake(ggp, nil); err != nil {
		t.Error(err)
	} else if totalGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect updated total GGP stake %s", totalGgpStake.String())
	}
	if nodeGgpStake, err := node.GetNodeGGPStake(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if nodeGgpStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect updated node GGP stake %s", nodeGgpStake.String())
	}

	// Reset rewards claim interval
	if _, err := protocol.BootstrapRewardsClaimIntervalTime(ggp, rewardsClaimIntervalTime, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

}
