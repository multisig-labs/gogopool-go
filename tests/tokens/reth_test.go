package tokens

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/network"
	"github.com/multisig-labs/gogopool-go/tokens"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
	rethutils "github.com/multisig-labs/gogopool-go/tests/testutils/tokens/reth"
)

// GetRETHContractETHBalance test under minipool.TestWithdrawValidatorBalance
// GetRETHTotalCollateral test under minipool.TestWithdrawValidatorBalance
// GetRETHCollateralRate test under minipool.TestWithdrawValidatorBalance

func TestRETHBalances(t *testing.T) {

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
	rethAmount := avax.EthToWei(100)
	if err := rethutils.MintRETH(ggp, userAccount1, rethAmount); err != nil {
		t.Fatal(err)
	}

	// Get & check rETH total supply
	if rethTotalSupply, err := tokens.GetRETHTotalSupply(ggp, nil); err != nil {
		t.Error(err)
	} else if rethTotalSupply.Cmp(rethAmount) != 0 {
		t.Errorf("Incorrect rETH total supply %s", rethTotalSupply.String())
	}

	// Get & check rETH account balance
	if gavaxBalance, err := tokens.GetRETHBalance(ggp, userAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if gavaxBalance.Cmp(rethAmount) != 0 {
		t.Errorf("Incorrect rETH account balance %s", gavaxBalance.String())
	}

}

func TestTransferRETH(t *testing.T) {

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
	rethAmount := avax.EthToWei(100)
	if err := rethutils.MintRETH(ggp, userAccount1, rethAmount); err != nil {
		t.Fatal(err)
	}

	// Mine pre-requisite 5760 blocks before being able to transfer
	if err := evm.MineBlocks(5760); err != nil {
		t.Fatal(err)
	}

	// Transfer rETH
	toAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	sendAmount := avax.EthToWei(50)
	if _, err := tokens.TransferRETH(ggp, toAddress, sendAmount, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check rETH account balance
	if gavaxBalance, err := tokens.GetRETHBalance(ggp, toAddress, nil); err != nil {
		t.Error(err)
	} else if gavaxBalance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect rETH account balance %s", gavaxBalance.String())
	}

}

func TestTransferFromRETH(t *testing.T) {

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
	rethAmount := avax.EthToWei(100)
	if err := rethutils.MintRETH(ggp, userAccount1, rethAmount); err != nil {
		t.Fatal(err)
	}

	// Approve rETH spender
	sendAmount := avax.EthToWei(50)
	if _, err := tokens.ApproveRETH(ggp, userAccount2.Address, sendAmount, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check spender allowance
	if allowance, err := tokens.GetRETHAllowance(ggp, userAccount1.Address, userAccount2.Address, nil); err != nil {
		t.Error(err)
	} else if allowance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect rETH spender allowance %s", allowance.String())
	}

	// Mine pre-requisite 5760 blocks before being able to transfer
	if err := evm.MineBlocks(5760); err != nil {
		t.Fatal(err)
	}

	// Transfer rETH from account
	toAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	if _, err := tokens.TransferFromRETH(ggp, userAccount1.Address, toAddress, sendAmount, userAccount2.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check rETH account balance
	if gavaxBalance, err := tokens.GetRETHBalance(ggp, toAddress, nil); err != nil {
		t.Error(err)
	} else if gavaxBalance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect rETH account balance %s", gavaxBalance.String())
	}

}

func TestRETHExchangeRate(t *testing.T) {

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

	// Submit network balances
	if _, err := network.SubmitBalances(ggp, 1, avax.EthToWei(100), avax.EthToWei(100), avax.EthToWei(50), trustedNodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check ETH value of rETH amount
	rethAmount := avax.EthToWei(1)
	if ethValue, err := tokens.GetETHValueOfRETH(ggp, rethAmount, nil); err != nil {
		t.Error(err)
	} else if ethValue.Cmp(avax.EthToWei(2)) != 0 {
		t.Errorf("Incorrect ETH value %s of rETH amount %s", ethValue.String(), rethAmount.String())
	}

	// Get & check rETH value of ETH amount
	ethAmount := avax.EthToWei(2)
	if rethValue, err := tokens.GetRETHValueOfETH(ggp, ethAmount, nil); err != nil {
		t.Error(err)
	} else if rethValue.Cmp(avax.EthToWei(1)) != 0 {
		t.Errorf("Incorrect rETH value %s of ETH amount %s", rethValue.String(), ethAmount.String())
	}

	// Get & check ETH : rETH exchange rate
	if exchangeRate, err := tokens.GetRETHExchangeRate(ggp, nil); err != nil {
		t.Error(err)
	} else if exchangeRate != 2 {
		t.Errorf("Incorrect ETH : rETH exchange rate %f : 1", exchangeRate)
	}

}

func TestBurnRETH(t *testing.T) {

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
	rethAmount := avax.EthToWei(100)
	if err := rethutils.MintRETH(ggp, userAccount1, rethAmount); err != nil {
		t.Fatal(err)
	}

	// Get initial balances
	balances1, err := tokens.GetBalances(ggp, userAccount1.Address, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Mine pre-requisite 5760 blocks before being able to burn
	if err := evm.MineBlocks(5760); err != nil {
		t.Fatal(err)
	}

	// Burn rETH
	burnAmount := avax.EthToWei(50)
	if _, err := tokens.BurnRETH(ggp, burnAmount, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated balances
	balances2, err := tokens.GetBalances(ggp, userAccount1.Address, nil)
	if err != nil {
		t.Fatal(err)
	} else {
		if balances2.GAVAX.Cmp(balances1.GAVAX) != -1 {
			t.Error("rETH balance did not decrease after burning rETH")
		}
		if balances2.AVAX.Cmp(balances1.AVAX) != 1 {
			t.Error("ETH balance did not increase after burning rETH")
		}
	}

}
