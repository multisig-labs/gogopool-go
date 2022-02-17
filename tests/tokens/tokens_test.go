package tokens

import (
	"math/big"
	"testing"

	"github.com/multisig-labs/gogopool-go/tokens"
	"github.com/multisig-labs/gogopool-go/utils/eth"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	ggputils "github.com/multisig-labs/gogopool-go/tests/testutils/tokens/ggp"
	rethutils "github.com/multisig-labs/gogopool-go/tests/testutils/tokens/reth"
)

func TestTokenBalances(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Mint rETH
	rethAmount := avax.EthToWei(102)
	if err := rethutils.MintRETH(ggp, userAccount1, rethAmount); err != nil {
		t.Fatal(err)
	}

	// Mint GGP
	ggpAmount := avax.EthToWei(103)
	if err := ggputils.MintGGP(ggp, ownerAccount, userAccount1, ggpAmount); err != nil {
		t.Fatal(err)
	}

	// Mint fixed-supply GGP
	fixedGgpAmount := avax.EthToWei(104)
	if err := ggputils.MintFixedSupplyGGP(ggp, ownerAccount, userAccount1, fixedGgpAmount); err != nil {
		t.Fatal(err)
	}

	// Get & check token balances
	if balances, err := tokens.GetBalances(ggp, userAccount1.Address, nil); err != nil {
		t.Error(err)
	} else {
		if balances.ETH.Cmp(big.NewInt(0)) != 1 {
			t.Errorf("Incorrect ETH balance %s", balances.ETH.String())
		}
		if balances.RETH.Cmp(rethAmount) != 0 {
			t.Errorf("Incorrect rETH balance %s", balances.RETH.String())
		}
		if balances.GGP.Cmp(ggpAmount) != 0 {
			t.Errorf("Incorrect GGP balance %s", balances.GGP.String())
		}
		if balances.FixedSupplyGGP.Cmp(fixedGgpAmount) != 0 {
			t.Errorf("Incorrect fixed-supply GGP balance %s", balances.FixedSupplyGGP.String())
		}
	}

}
