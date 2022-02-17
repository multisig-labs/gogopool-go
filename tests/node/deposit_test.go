package node

import (
	"testing"

	"github.com/multisig-labs/gogopool-go/minipool"
	"github.com/multisig-labs/gogopool-go/node"
	"github.com/multisig-labs/gogopool-go/utils/eth"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	minipoolutils "github.com/multisig-labs/gogopool-go/tests/testutils/minipool"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestDeposit(t *testing.T) {

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

	// Get initial node minipool count
	minipoolCount1, err := minipool.GetNodeMinipoolCount(ggp, nodeAccount.Address, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Mint & stake GGP required for mininpool
	ggpRequired, err := minipoolutils.GetMinipoolGGPRequired(ggp)
	if err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.StakeGGP(ggp, ownerAccount, nodeAccount, ggpRequired); err != nil {
		t.Fatal(err)
	}

	// Deposit
	if _, _, err := nodeutils.Deposit(t, ggp, nodeAccount, avax.EthToWei(16), 1); err != nil {
		t.Fatal(err)
	}

	// Get & check updated node minipool count
	minipoolCount2, err := minipool.GetNodeMinipoolCount(ggp, nodeAccount.Address, nil)
	if err != nil {
		t.Fatal(err)
	} else if minipoolCount2 != minipoolCount1+1 {
		t.Error("Incorrect node minipool count")
	}

}
