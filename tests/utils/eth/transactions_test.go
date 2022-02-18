package eth

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	uc "github.com/multisig-labs/gogopool-go/utils/client"

	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests"
	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	"github.com/multisig-labs/gogopool-go/utils"
)

func TestSendTransaction(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Initialize eth client
	client := uc.NewEth1ClientProxy(0, tests.Eth1ProviderAddress)

	// Initialize accounts
	userAccount, err := accounts.GetAccount(9)
	if err != nil {
		t.Fatal(err)
	}

	// Transaction parameters
	toAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	sendAmount := avax.EthToWei(50)

	// Send transaction
	opts := userAccount.GetTransactor()
	opts.Value = sendAmount
	hash, err := avax.SendTransaction(client, toAddress, big.NewInt(1337), opts) // Ganache's default chain ID is 1337
	if err != nil {
		t.Fatal(err)
	}
	if _, err := utils.WaitForTransaction(client, hash); err != nil {
		t.Fatal(err)
	}

	// Get & check to address balance
	if balance, err := client.BalanceAt(context.Background(), toAddress, nil); err != nil {
		t.Error(err)
	} else if balance.Cmp(sendAmount) != 0 {
		t.Errorf("Incorrect to address balance %s", balance.String())
	}

}
