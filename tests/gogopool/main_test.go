package gogopool

import (
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/gogopool"
	uc "github.com/multisig-labs/gogopool-go/utils/client"

	"github.com/multisig-labs/gogopool-go/tests"
)

var (
	client *uc.EthClientProxy
	ggp    *gogopool.GoGoPool
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

	// Run tests
	os.Exit(m.Run())

}
