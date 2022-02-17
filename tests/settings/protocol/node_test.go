package protocol

import (
	"testing"

	"github.com/multisig-labs/gogopool-go/settings/protocol"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
)

func TestNodeSettings(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Set & get node registrations enabled
	nodeRegistrationsEnabled := false
	if _, err := protocol.BootstrapNodeRegistrationEnabled(ggp, nodeRegistrationsEnabled, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetNodeRegistrationEnabled(ggp, nil); err != nil {
		t.Error(err)
	} else if value != nodeRegistrationsEnabled {
		t.Error("Incorrect node registrations enabled value")
	}

	// Set & get node deposits enabled
	nodeDepositsEnabled := false
	if _, err := protocol.BootstrapNodeDepositEnabled(ggp, nodeDepositsEnabled, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetNodeDepositEnabled(ggp, nil); err != nil {
		t.Error(err)
	} else if value != nodeDepositsEnabled {
		t.Error("Incorrect node deposits enabled value")
	}

	// Set & get minimum per minipool GGP stake
	minimumPerMinipoolStake := 1.0
	if _, err := protocol.BootstrapMinimumPerMinipoolStake(ggp, minimumPerMinipoolStake, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetMinimumPerMinipoolStake(ggp, nil); err != nil {
		t.Error(err)
	} else if value != minimumPerMinipoolStake {
		t.Error("Incorrect minimum per minipool stake value")
	}

	// Set & get maximum per minipool GGP stake
	maximumPerMinipoolStake := 10.0
	if _, err := protocol.BootstrapMaximumPerMinipoolStake(ggp, maximumPerMinipoolStake, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetMaximumPerMinipoolStake(ggp, nil); err != nil {
		t.Error(err)
	} else if value != maximumPerMinipoolStake {
		t.Error("Incorrect maximum per minipool stake value")
	}

}
