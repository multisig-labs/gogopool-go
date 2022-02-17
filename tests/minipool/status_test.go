package minipool

import (
	"fmt"
	"testing"

	"github.com/multisig-labs/gogopool-go/settings/trustednode"
	"github.com/multisig-labs/gogopool-go/types"

	"github.com/multisig-labs/gogopool-go/minipool"
	"github.com/multisig-labs/gogopool-go/node"
	"github.com/multisig-labs/gogopool-go/utils/eth"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	minipoolutils "github.com/multisig-labs/gogopool-go/tests/testutils/minipool"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestSubmitMinipoolWithdrawable(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Register nodes
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount); err != nil {
		t.Fatal(err)
	}

	// Create & stake minipool
	mp, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, nodeAccount, avax.EthToWei(32), 1)
	if err != nil {
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

	if err := minipoolutils.StakeMinipool(ggp, mp, nodeAccount); err != nil {
		t.Fatal(err)
	}

	// Get & check initial minipool withdrawable status
	if status, err := mp.GetStatus(nil); err != nil {
		t.Error(err)
	} else if status == types.Withdrawable {
		t.Error("Incorrect initial minipool withdrawable status")
	}

	// Submit minipool withdrawable status
	if _, err := minipool.SubmitMinipoolWithdrawable(ggp, mp.Address, trustedNodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated minipool withdrawable status
	if status, err := mp.GetStatus(nil); err != nil {
		t.Error(err)
	} else if status != types.Withdrawable {
		t.Error("Incorrect updated minipool withdrawable status")
	}

}
