package tokens

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/tokens"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	ggputils "github.com/multisig-labs/gogopool-go/tests/testutils/tokens/ggp"
)

func TestGGPBalances(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Mint GGP
	ggpAmount := avax.EthToWei(100)
	if err := ggputils.MintGGP(ggp, ownerAccount, userAccount1, ggpAmount); err != nil {
		t.Fatal(err)
	}

	// Get & check GGP account balance
	if ggpBalance, err := tokens.GetGGPBalance(ggp, userAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if ggpBalance.Cmp(ggpAmount) != 0 {
		t.Errorf("Incorrect GGP account balance %s", ggpBalance.String())
	}

	// Get & check GGP total supply
	initialTotalSupply := avax.EthToWei(18000000)
	if ggpTotalSupply, err := tokens.GetGGPTotalSupply(ggp, nil); err != nil {
		t.Error(err)
	} else if ggpTotalSupply.Cmp(initialTotalSupply) != 0 {
		t.Errorf("Incorrect GGP total supply %s", ggpTotalSupply.String())
	}

}

func TestTransferGGP(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Mint GGP
	ggpAmount := avax.EthToWei(100)
	if err := ggputils.MintGGP(ggp, ownerAccount, userAccount1, ggpAmount); err != nil {
		t.Fatal(err)
	}

	// Transfer GGP
	toAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	sendAmount := avax.EthToWei(50)
	if _, err := tokens.TransferGGP(ggp, toAddress, sendAmount, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check GGP account balance
	if ggpBalance, err := tokens.GetGGPBalance(ggp, toAddress, nil); err != nil {
		t.Error(err)
	} else if ggpBalance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect GGP account balance %s", ggpBalance.String())
	}

}

func TestTransferFromGGP(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Mint GGP
	ggpAmount := avax.EthToWei(100)
	if err := ggputils.MintGGP(ggp, ownerAccount, userAccount1, ggpAmount); err != nil {
		t.Fatal(err)
	}

	// Approve GGP spender
	sendAmount := avax.EthToWei(50)
	if _, err := tokens.ApproveGGP(ggp, userAccount2.Address, sendAmount, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check spender allowance
	if allowance, err := tokens.GetGGPAllowance(ggp, userAccount1.Address, userAccount2.Address, nil); err != nil {
		t.Error(err)
	} else if allowance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect GGP spender allowance %s", allowance.String())
	}

	// Transfer GGP from account
	toAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	if _, err := tokens.TransferFromGGP(ggp, userAccount1.Address, toAddress, sendAmount, userAccount2.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check GGP account balance
	if ggpBalance, err := tokens.GetGGPBalance(ggp, toAddress, nil); err != nil {
		t.Error(err)
	} else if ggpBalance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect GGP account balance %s", ggpBalance.String())
	}

}

func TestMintInflationGGP(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Constants
	oneDay := 24 * 60 * 60

	// Start GGP inflation
	if _, err := protocol.BootstrapInflationStartTime(ggp, uint64(time.Now().Unix()+3600), ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Increase time until rewards are available
	if err := evm.IncreaseTime(3600 + oneDay); err != nil {
		t.Fatal(err)
	}

	// Get initial total supply
	ggpTotalSupply1, err := tokens.GetGGPTotalSupply(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Mint GGP from inflation
	if _, err := tokens.MintInflationGGP(ggp, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated total supply
	ggpTotalSupply2, err := tokens.GetGGPTotalSupply(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	if ggpTotalSupply2.Cmp(ggpTotalSupply1) != 1 {
		t.Errorf("Incorrect updated GGP total supply %s", ggpTotalSupply2.String())
	}

}

func TestSwapFixedSupplyGGPForGGP(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Mint fixed-supply GGP
	ggpAmount := avax.EthToWei(100)
	if err := ggputils.MintFixedSupplyGGP(ggp, ownerAccount, userAccount1, ggpAmount); err != nil {
		t.Fatal(err)
	}

	// Approve fixed-supply GGP spend
	gogoTokenGGPAddress, err := ggp.GetAddress("rocketTokenGGP")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tokens.ApproveFixedSupplyGGP(ggp, *gogoTokenGGPAddress, ggpAmount, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Swap fixed-supply RP for GGP
	if _, err := tokens.SwapFixedSupplyGGPForGGP(ggp, ggpAmount, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check GGP account balance
	if ggpBalance, err := tokens.GetGGPBalance(ggp, userAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if ggpBalance.Cmp(ggpAmount) != 0 {
		t.Errorf("Incorrect GGP account balance %s", ggpBalance.String())
	}

}
