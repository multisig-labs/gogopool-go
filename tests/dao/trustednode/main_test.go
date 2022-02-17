package trustednode

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

	ownerAccount        *accounts.Account
	trustedNodeAccount1 *accounts.Account
	trustedNodeAccount2 *accounts.Account
	trustedNodeAccount3 *accounts.Account
	trustedNodeAccount4 *accounts.Account
	nodeAccount         *accounts.Account
)

func TestMain(m *testing.M) {
	var err error

	// Initialize eth client
	client = uc.NewEth1ClientProxy(0, tests.Eth1ProviderAddress)

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
	trustedNodeAccount1, err = accounts.GetAccount(1)
	if err != nil {
		log.Fatal(err)
	}
	trustedNodeAccount2, err = accounts.GetAccount(2)
	if err != nil {
		log.Fatal(err)
	}
	trustedNodeAccount3, err = accounts.GetAccount(3)
	if err != nil {
		log.Fatal(err)
	}
	trustedNodeAccount4, err = accounts.GetAccount(4)
	if err != nil {
		log.Fatal(err)
	}
	nodeAccount, err = accounts.GetAccount(5)
	if err != nil {
		log.Fatal(err)
	}

	// Run tests
	os.Exit(m.Run())

}
