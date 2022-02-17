package protocol

import (
	"testing"
	"time"

	"github.com/multisig-labs/gogopool-go/settings/protocol"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
)

func TestInflationSettings(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Set & get inflation interval rate
	inflationIntervalRate := 0.5
	if _, err := protocol.BootstrapInflationIntervalRate(ggp, inflationIntervalRate, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetInflationIntervalRate(ggp, nil); err != nil {
		t.Error(err)
	} else if value != inflationIntervalRate {
		t.Error("Incorrect inflation interval rate value")
	}

	// Set & get inflation start block
	inflationStartTime := uint64(time.Now().Unix()) + 3600
	if _, err := protocol.BootstrapInflationStartTime(ggp, inflationStartTime, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetInflationStartTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != inflationStartTime {
		t.Error("Incorrect inflation start time value")
	}

}
