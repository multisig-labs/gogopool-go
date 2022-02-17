package tokens

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/tokens"
	"github.com/multisig-labs/gogopool-go/utils/eth"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	ggputils "github.com/multisig-labs/gogopool-go/tests/testutils/tokens/ggp"
)

func TestFixedSupplyGGPBalances(t *testing.T) {

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
	fixedGgpAmount := avax.EthToWei(100)
	if err := ggputils.MintFixedSupplyGGP(ggp, ownerAccount, userAccount1, fixedGgpAmount); err != nil {
		t.Fatal(err)
	}

	// Get & check fixed-supply GGP total supply
	if fixedGgpTotalSupply, err := tokens.GetFixedSupplyGGPTotalSupply(ggp, nil); err != nil {
		t.Error(err)
	} else if fixedGgpTotalSupply.Cmp(fixedGgpAmount) != 0 {
		t.Errorf("Incorrect fixed-supply GGP total supply %s", fixedGgpTotalSupply.String())
	}

	// Get & check fixed-supply GGP account balance
	if fixedGgpBalance, err := tokens.GetFixedSupplyGGPBalance(ggp, userAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if fixedGgpBalance.Cmp(fixedGgpAmount) != 0 {
		t.Errorf("Incorrect fixed-supply GGP account balance %s", fixedGgpBalance.String())
	}

}

func TestTransferFixedSupplyGGP(t *testing.T) {

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
	fixedGgpAmount := avax.EthToWei(100)
	if err := ggputils.MintFixedSupplyGGP(ggp, ownerAccount, userAccount1, fixedGgpAmount); err != nil {
		t.Fatal(err)
	}

	// Transfer fixed-supply GGP
	toAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	sendAmount := avax.EthToWei(50)
	if _, err := tokens.TransferFixedSupplyGGP(ggp, toAddress, sendAmount, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check fixed-supply GGP account balance
	if fixedGgpBalance, err := tokens.GetFixedSupplyGGPBalance(ggp, toAddress, nil); err != nil {
		t.Error(err)
	} else if fixedGgpBalance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect fixed-supply GGP account balance %s", fixedGgpBalance.String())
	}

}

func TestTransferFromFixedSupplyGGP(t *testing.T) {

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
	fixedGgpAmount := avax.EthToWei(100)
	if err := ggputils.MintFixedSupplyGGP(ggp, ownerAccount, userAccount1, fixedGgpAmount); err != nil {
		t.Fatal(err)
	}

	// Approve fixed-supply GGP spender
	sendAmount := avax.EthToWei(50)
	if _, err := tokens.ApproveFixedSupplyGGP(ggp, userAccount2.Address, sendAmount, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check spender allowance
	if allowance, err := tokens.GetFixedSupplyGGPAllowance(ggp, userAccount1.Address, userAccount2.Address, nil); err != nil {
		t.Error(err)
	} else if allowance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect fixed-supply GGP spender allowance %s", allowance.String())
	}

	// Transfer fixed-supply GGP from account
	toAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	if _, err := tokens.TransferFromFixedSupplyGGP(ggp, userAccount1.Address, toAddress, sendAmount, userAccount2.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check fixed-supply GGP account balance
	if fixedGgpBalance, err := tokens.GetFixedSupplyGGPBalance(ggp, toAddress, nil); err != nil {
		t.Error(err)
	} else if fixedGgpBalance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect fixed-supply GGP account balance %s", fixedGgpBalance.String())
	}

}
