package deposit

import (
	"testing"

	"github.com/multisig-labs/gogopool-go/deposit"
	"github.com/multisig-labs/gogopool-go/node"
	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/utils/eth"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	minipoolutils "github.com/multisig-labs/gogopool-go/tests/testutils/minipool"
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

	// Make deposit
	opts := userAccount.GetTransactor()
	opts.Value = avax.EthToWei(10)
	if _, err := deposit.Deposit(ggp, opts); err != nil {
		t.Fatal(err)
	}

	// Get & check deposit pool balance
	if balance, err := deposit.GetBalance(ggp, nil); err != nil {
		t.Error(err)
	} else if balance.Cmp(opts.Value) != 0 {
		t.Error("Incorrect deposit pool balance")
	}

	// Get & check deposit pool excess balance
	if excessBalance, err := deposit.GetExcessBalance(ggp, nil); err != nil {
		t.Error(err)
	} else if excessBalance.Cmp(opts.Value) != 0 {
		t.Error("Incorrect deposit pool excess balance")
	}

}

func TestAssignDeposits(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Disable deposit assignments
	if _, err := protocol.BootstrapAssignDepositsEnabled(ggp, false, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Make user deposit
	userDepositOpts := userAccount.GetTransactor()
	userDepositOpts.Value = avax.EthToWei(32)
	if _, err := deposit.Deposit(ggp, userDepositOpts); err != nil {
		t.Fatal(err)
	}

	// Register node & create minipool
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, nodeAccount, avax.EthToWei(16), 1); err != nil {
		t.Fatal(err)
	}

	// Re-enable deposit assignments
	if _, err := protocol.BootstrapAssignDepositsEnabled(ggp, true, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get initial deposit pool balance
	balance1, err := deposit.GetBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Assign deposits
	if _, err := deposit.AssignDeposits(ggp, userAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated deposit pool balance
	balance2, err := deposit.GetBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	} else if balance2.Cmp(balance1) != -1 {
		t.Error("Deposit pool balance did not decrease after assigning deposits")
	}

}
