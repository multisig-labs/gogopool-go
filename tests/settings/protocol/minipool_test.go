package protocol

import (
	"testing"
	"time"

	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/utils/eth"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
)

func TestMinipoolSettings(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Get & check launch balance and deposit amounts
	fullMinipoolBalance := avax.EthToWei(32)
	halfMinipoolBalance := avax.EthToWei(16)
	emptyMinipoolBalance := avax.EthToWei(0)
	if value, err := protocol.GetMinipoolLaunchBalance(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(fullMinipoolBalance) != 0 {
		t.Error("Incorrect minipool launch balance")
	}
	if value, err := protocol.GetMinipoolFullDepositNodeAmount(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(fullMinipoolBalance) != 0 {
		t.Error("Incorrect minipool full deposit node amount")
	}
	if value, err := protocol.GetMinipoolHalfDepositNodeAmount(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(halfMinipoolBalance) != 0 {
		t.Error("Incorrect minipool half deposit node amount")
	}
	if value, err := protocol.GetMinipoolEmptyDepositNodeAmount(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(emptyMinipoolBalance) != 0 {
		t.Error("Incorrect minipool empty deposit node amount")
	}
	if value, err := protocol.GetMinipoolFullDepositUserAmount(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(halfMinipoolBalance) != 0 {
		t.Error("Incorrect minipool full deposit user amount")
	}
	if value, err := protocol.GetMinipoolHalfDepositUserAmount(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(halfMinipoolBalance) != 0 {
		t.Error("Incorrect minipool half deposit user amount")
	}
	if value, err := protocol.GetMinipoolEmptyDepositUserAmount(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(fullMinipoolBalance) != 0 {
		t.Error("Incorrect minipool empty deposit user amount")
	}

	// Set & get submit withdrawable enabled
	submitWithdrawableEnabled := false
	if _, err := protocol.BootstrapMinipoolSubmitWithdrawableEnabled(ggp, submitWithdrawableEnabled, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetMinipoolSubmitWithdrawableEnabled(ggp, nil); err != nil {
		t.Error(err)
	} else if value != submitWithdrawableEnabled {
		t.Error("Incorrect minipool withdrawable submissions enabled value")
	}

	// Set & get minipool launch timeout
	var minipoolLaunchTimeout time.Duration = 5 * time.Second
	if _, err := protocol.BootstrapMinipoolLaunchTimeout(ggp, minipoolLaunchTimeout, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetMinipoolLaunchTimeout(ggp, nil); err != nil {
		t.Error(err)
	} else if value != minipoolLaunchTimeout {
		t.Error("Incorrect minipool launch timeout value")
	}
}
