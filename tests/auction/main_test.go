package auction

import (
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	uc "github.com/multisig-labs/gogopool-go/utils/client"

	"github.com/multisig-labs/gogopool-go/tests"
	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	"github.com/multisig-labs/gogopool-go/tests/utils"
)

var (
	client *uc.EthClientProxy
	ggp    *gogopool.GoGoPool

	ownerAccount        *accounts.Account
	trustedNodeAccount1 *accounts.Account
	trustedNodeAccount2 *accounts.Account
	trustedNodeAccount3 *accounts.Account
	userAccount1        *accounts.Account
	userAccount2        *accounts.Account
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
	userAccount1, err = accounts.GetAccount(8)
	if err != nil {
		log.Fatal(err)
	}
	userAccount2, err = accounts.GetAccount(9)
	if err != nil {
		log.Fatal(err)
	}

	// Do the bootstrap settings
	utils.Stage4Bootstrap(ggp, ownerAccount)

	// Run tests
	os.Exit(m.Run())

}
