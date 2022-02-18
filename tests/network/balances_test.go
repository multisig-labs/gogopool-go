package network

import (
	"testing"

	"github.com/multisig-labs/gogopool-go/network"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestSubmitBalances(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Register trusted node
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount); err != nil {
		t.Fatal(err)
	}

	// Submit balances
	var balancesBlock uint64 = 100
	totalEth := avax.EthToWei(100)
	stakingEth := avax.EthToWei(80)
	rethSupply := avax.EthToWei(70)
	if _, err := network.SubmitBalances(ggp, balancesBlock, totalEth, stakingEth, rethSupply, trustedNodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check network balances block
	if networkBalancesBlock, err := network.GetBalancesBlock(ggp, nil); err != nil {
		t.Error(err)
	} else if networkBalancesBlock != balancesBlock {
		t.Errorf("Incorrect network balances block %d", networkBalancesBlock)
	}

	// Get & check network total ETH
	if networkTotalEth, err := network.GetTotalETHBalance(ggp, nil); err != nil {
		t.Error(err)
	} else if networkTotalEth.Cmp(totalEth) != 0 {
		t.Errorf("Incorrect network total ETH balance %s", networkTotalEth.String())
	}

	// Get & check network staking ETH
	if networkStakingEth, err := network.GetStakingETHBalance(ggp, nil); err != nil {
		t.Error(err)
	} else if networkStakingEth.Cmp(stakingEth) != 0 {
		t.Errorf("Incorrect network staking ETH balance %s", networkStakingEth.String())
	}

	// Get & check network rETH supply
	if networkRethSupply, err := network.GetTotalRETHSupply(ggp, nil); err != nil {
		t.Error(err)
	} else if networkRethSupply.Cmp(rethSupply) != 0 {
		t.Errorf("Incorrect network total rETH supply %s", networkRethSupply.String())
	}

	// Get & check ETH utilization rate
	if ethUtilizationRate, err := network.GetETHUtilizationRate(ggp, nil); err != nil {
		t.Error(err)
	} else if ethUtilizationRate != avax.WeiToEth(stakingEth)/avax.WeiToEth(totalEth) {
		t.Errorf("Incorrect network ETH utilization rate %f", ethUtilizationRate)
	}

}
