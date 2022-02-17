package tokens

import (
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	uc "github.com/multisig-labs/gogopool-go/utils/client"

	"github.com/multisig-labs/gogopool-go/tests"
	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
)

var (
	client *uc.EthClientProxy
	ggp    *gogopool.GoGoPool

	ownerAccount       *accounts.Account
	trustedNodeAccount *accounts.Account
	userAccount1       *accounts.Account
	userAccount2       *accounts.Account
	swcAccount         *accounts.Account
)

func TestMain(m *testing.M) {
	var err error

	// Initialize eth client
	client = uc.NewEth1ClientProxy(0, tests.Eth1ProviderAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize contract manager
	ggp, err = gogopool.NewGoGoPool(client, common.HexToAddress(tests.GoGoStorageAddress))
	if err != nil {
		log.Fatal(err)
	}

	// Initialize accounts
	ownerAccount, err = accounts.GetAccount(0)
	if err != nil {
		log.Fatal(err)
	}
	trustedNodeAccount, err = accounts.GetAccount(1)
	if err != nil {
		log.Fatal(err)
	}
	userAccount1, err = accounts.GetAccount(7)
	if err != nil {
		log.Fatal(err)
	}
	userAccount2, err = accounts.GetAccount(8)
	if err != nil {
		log.Fatal(err)
	}
	swcAccount, err = accounts.GetAccount(9)
	if err != nil {
		log.Fatal(err)
	}

	// Run tests
	os.Exit(m.Run())

}
