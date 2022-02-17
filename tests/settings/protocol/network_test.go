package protocol

import (
	"testing"

	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/utils/eth"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
)

func TestNetworkSettings(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Set & get node consensus threshold
	nodeConsensusThreshold := 0.1
	if _, err := protocol.BootstrapNodeConsensusThreshold(ggp, nodeConsensusThreshold, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetNodeConsensusThreshold(ggp, nil); err != nil {
		t.Error(err)
	} else if value != nodeConsensusThreshold {
		t.Error("Incorrect node consensus threshold value")
	}

	// Set & get network balance submissions enabled
	submitBalancesEnabled := false
	if _, err := protocol.BootstrapSubmitBalancesEnabled(ggp, submitBalancesEnabled, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetSubmitBalancesEnabled(ggp, nil); err != nil {
		t.Error(err)
	} else if value != submitBalancesEnabled {
		t.Error("Incorrect network balance submissions enabled value")
	}

	// Set & get network balance submission frequency
	var submitBalancesFrequency uint64 = 10
	if _, err := protocol.BootstrapSubmitBalancesFrequency(ggp, submitBalancesFrequency, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetSubmitBalancesFrequency(ggp, nil); err != nil {
		t.Error(err)
	} else if value != submitBalancesFrequency {
		t.Error("Incorrect network balance submission frequency value")
	}

	// Set & get network price submissions enabled
	submitPricesEnabled := false
	if _, err := protocol.BootstrapSubmitPricesEnabled(ggp, submitPricesEnabled, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetSubmitPricesEnabled(ggp, nil); err != nil {
		t.Error(err)
	} else if value != submitPricesEnabled {
		t.Error("Incorrect network price submissions enabled value")
	}

	// Set & get network price submission frequency
	var submitPricesFrequency uint64 = 10
	if _, err := protocol.BootstrapSubmitPricesFrequency(ggp, submitPricesFrequency, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetSubmitPricesFrequency(ggp, nil); err != nil {
		t.Error(err)
	} else if value != submitPricesFrequency {
		t.Error("Incorrect network price submission frequency value")
	}

	// Set & get minimum node fee
	minimumNodeFee := 0.80
	if _, err := protocol.BootstrapMinimumNodeFee(ggp, minimumNodeFee, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetMinimumNodeFee(ggp, nil); err != nil {
		t.Error(err)
	} else if value != minimumNodeFee {
		t.Error("Incorrect minimum node fee value")
	}

	// Set & get target node fee
	targetNodeFee := 0.85
	if _, err := protocol.BootstrapTargetNodeFee(ggp, targetNodeFee, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetTargetNodeFee(ggp, nil); err != nil {
		t.Error(err)
	} else if value != targetNodeFee {
		t.Error("Incorrect target node fee value")
	}

	// Set & get maximum node fee
	maximumNodeFee := 0.90
	if _, err := protocol.BootstrapMaximumNodeFee(ggp, maximumNodeFee, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetMaximumNodeFee(ggp, nil); err != nil {
		t.Error(err)
	} else if value != maximumNodeFee {
		t.Error("Incorrect maximum node fee value")
	}

	// Set & get node fee demand range
	nodeFeeDemandRange := avax.EthToWei(10)
	if _, err := protocol.BootstrapNodeFeeDemandRange(ggp, nodeFeeDemandRange, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetNodeFeeDemandRange(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(nodeFeeDemandRange) != 0 {
		t.Error("Incorrect node fee demand range value")
	}

	// Set & get target rETH collateral rate
	targetRethCollateralRate := 0.95
	if _, err := protocol.BootstrapTargetRethCollateralRate(ggp, targetRethCollateralRate, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := protocol.GetTargetRethCollateralRate(ggp, nil); err != nil {
		t.Error(err)
	} else if value != targetRethCollateralRate {
		t.Error("Incorrect target rETH collateral rate value")
	}

}
