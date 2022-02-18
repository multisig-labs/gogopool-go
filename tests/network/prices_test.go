package network

import (
	"testing"

	"github.com/multisig-labs/gogopool-go/network"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestSubmitPrices(t *testing.T) {

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

	// Submit prices
	var pricesBlock uint64 = 100
	ggpPrice := avax.EthToWei(1000)
	effectiveGgpStake := avax.EthToWei(24000)
	if _, err := network.SubmitPrices(ggp, pricesBlock, ggpPrice, effectiveGgpStake, trustedNodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check network prices block
	if networkPricesBlock, err := network.GetPricesBlock(ggp, nil); err != nil {
		t.Error(err)
	} else if networkPricesBlock != pricesBlock {
		t.Errorf("Incorrect network prices block %d", networkPricesBlock)
	}

	// Get & check network GGP price
	if networkGgpPrice, err := network.GetGGPPrice(ggp, nil); err != nil {
		t.Error(err)
	} else if networkGgpPrice.Cmp(ggpPrice) != 0 {
		t.Errorf("Incorrect network GGP price %s", networkGgpPrice.String())
	}

}
