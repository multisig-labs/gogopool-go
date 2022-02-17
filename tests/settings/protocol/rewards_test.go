package protocol

import (
	"testing"

	protocoldao "github.com/multisig-labs/gogopool-go/dao/protocol"
	protocolsettings "github.com/multisig-labs/gogopool-go/settings/protocol"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
)

func TestRewardsSettings(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Bootstrap a claimer & get claimer settings
	claimerPerc := 0.1
	if _, err := protocoldao.BootstrapClaimer(ggp, "gogoClaimNode", claimerPerc, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else {
		if value, err := protocolsettings.GetRewardsClaimerPerc(ggp, "gogoClaimNode", nil); err != nil {
			t.Error(err)
		} else if value != claimerPerc {
			t.Errorf("Incorrect rewards claimer percent %f", value)
		}
		if value, err := protocolsettings.GetRewardsClaimerPercTimeUpdated(ggp, "gogoClaimNode", nil); err != nil {
			t.Error(err)
		} else if value == 0 {
			t.Errorf("Incorrect rewards claimer percent time updated %d", value)
		}
		if value, err := protocolsettings.GetRewardsClaimersPercTotal(ggp, nil); err != nil {
			t.Error(err)
		} else if value == 0 {
			t.Errorf("Incorrect rewards claimers total percent %f", value)
		}
	}

	// Set & get rewards claim interval time
	var rewardsClaimIntervalTime uint64 = 1
	if _, err := protocolsettings.BootstrapRewardsClaimIntervalTime(ggp, rewardsClaimIntervalTime, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocolsettings.GetRewardsClaimIntervalTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != rewardsClaimIntervalTime {
		t.Error("Incorrect rewards claim interval time value")
	}

}
