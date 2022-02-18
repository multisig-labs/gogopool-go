package protocol

import (
	"testing"

	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
)

func TestDepositSettings(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Set & get deposits enabled
	depositEnabled := false
	if _, err := protocol.BootstrapDepositEnabled(ggp, depositEnabled, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetDepositEnabled(ggp, nil); err != nil {
		t.Error(err)
	} else if value != depositEnabled {
		t.Error("Incorrect deposit enabled value")
	}

	// Set & get deposit assignments enabled
	assignDepositsEnabled := false
	if _, err := protocol.BootstrapAssignDepositsEnabled(ggp, assignDepositsEnabled, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetAssignDepositsEnabled(ggp, nil); err != nil {
		t.Error(err)
	} else if value != assignDepositsEnabled {
		t.Error("Incorrect assign deposits enabled value")
	}

	// Set & get minimum deposit amount
	minimumDeposit := avax.EthToWei(1000)
	if _, err := protocol.BootstrapMinimumDeposit(ggp, minimumDeposit, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetMinimumDeposit(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(minimumDeposit) != 0 {
		t.Error("Incorrect minimum deposit value")
	}

	// Set & get maximum deposit pool size
	maximumDepositPoolSize := avax.EthToWei(1)
	if _, err := protocol.BootstrapMaximumDepositPoolSize(ggp, maximumDepositPoolSize, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetMaximumDepositPoolSize(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(maximumDepositPoolSize) != 0 {
		t.Error("Incorrect maximum deposit pool size value")
	}

	// Set & get maximum deposit assignments
	var maximumDepositAssignments uint64 = 50
	if _, err := protocol.BootstrapMaximumDepositAssignments(ggp, maximumDepositAssignments, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetMaximumDepositAssignments(ggp, nil); err != nil {
		t.Error(err)
	} else if value != maximumDepositAssignments {
		t.Error("Incorrect maximum deposit assignments value")
	}

}
