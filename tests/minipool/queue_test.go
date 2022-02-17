package minipool

import (
	"testing"

	trustednodesettings "github.com/multisig-labs/gogopool-go/settings/trustednode"

	"github.com/multisig-labs/gogopool-go/minipool"
	"github.com/multisig-labs/gogopool-go/node"
	"github.com/multisig-labs/gogopool-go/utils/eth"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	minipoolutils "github.com/multisig-labs/gogopool-go/tests/testutils/minipool"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestQueueLengths(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Register nodes
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount); err != nil {
		t.Fatal(err)
	}

	// Disable min commission rate for unbonded pools
	if _, err := trustednodesettings.BootstrapMinipoolUnbondedMinFee(ggp, uint64(0), ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check queue lengths
	if queueLengths, err := minipool.GetQueueLengths(ggp, nil); err != nil {
		t.Error(err)
	} else {
		if queueLengths.Total != 0 {
			t.Errorf("Incorrect total queue length 1 %d", queueLengths.Total)
		}
		if queueLengths.FullDeposit != 0 {
			t.Errorf("Incorrect full deposit queue length 1 %d", queueLengths.FullDeposit)
		}
		if queueLengths.HalfDeposit != 0 {
			t.Errorf("Incorrect half deposit queue length 1 %d", queueLengths.HalfDeposit)
		}
		if queueLengths.EmptyDeposit != 0 {
			t.Errorf("Incorrect empty deposit queue length 1 %d", queueLengths.EmptyDeposit)
		}
	}

	// Create full deposit minipool
	if _, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, nodeAccount, avax.EthToWei(32), 1); err != nil {
		t.Fatal(err)
	}

	// Get & check queue lengths
	if queueLengths, err := minipool.GetQueueLengths(ggp, nil); err != nil {
		t.Error(err)
	} else {
		if queueLengths.Total != 1 {
			t.Errorf("Incorrect total queue length 2 %d", queueLengths.Total)
		}
		if queueLengths.FullDeposit != 1 {
			t.Errorf("Incorrect full deposit queue length 2 %d", queueLengths.FullDeposit)
		}
		if queueLengths.HalfDeposit != 0 {
			t.Errorf("Incorrect half deposit queue length 2 %d", queueLengths.HalfDeposit)
		}
		if queueLengths.EmptyDeposit != 0 {
			t.Errorf("Incorrect empty deposit queue length 2 %d", queueLengths.EmptyDeposit)
		}
	}

	// Create half deposit minipool
	if _, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, nodeAccount, avax.EthToWei(16), 2); err != nil {
		t.Fatal(err)
	}

	// Get & check queue lengths
	if queueLengths, err := minipool.GetQueueLengths(ggp, nil); err != nil {
		t.Error(err)
	} else {
		if queueLengths.Total != 2 {
			t.Errorf("Incorrect total queue length 3 %d", queueLengths.Total)
		}
		if queueLengths.FullDeposit != 1 {
			t.Errorf("Incorrect full deposit queue length 3 %d", queueLengths.FullDeposit)
		}
		if queueLengths.HalfDeposit != 1 {
			t.Errorf("Incorrect half deposit queue length 3 %d", queueLengths.HalfDeposit)
		}
		if queueLengths.EmptyDeposit != 0 {
			t.Errorf("Incorrect empty deposit queue length 3 %d", queueLengths.EmptyDeposit)
		}
	}

	// Create empty deposit minipool
	//if _, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, trustedNodeAccount, eth.EthToWei(0), 3); err != nil { t.Fatal(err) }

	// Get & check queue lengths
	if queueLengths, err := minipool.GetQueueLengths(ggp, nil); err != nil {
		t.Error(err)
	} else {
		if queueLengths.Total != 2 {
			t.Errorf("Incorrect total queue length 4 %d", queueLengths.Total)
		}
		if queueLengths.FullDeposit != 1 {
			t.Errorf("Incorrect full deposit queue length 4 %d", queueLengths.FullDeposit)
		}
		if queueLengths.HalfDeposit != 1 {
			t.Errorf("Incorrect half deposit queue length 4 %d", queueLengths.HalfDeposit)
		}
		//if queueLengths.EmptyDeposit != 1 {
		//    t.Errorf("Incorrect empty deposit queue length 4 %d", queueLengths.EmptyDeposit)
		//}
	}

}

func TestQueueCapacity(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Register nodes
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount); err != nil {
		t.Fatal(err)
	}

	// Disable min commission rate for unbonded pools
	if _, err := trustednodesettings.BootstrapMinipoolUnbondedMinFee(ggp, uint64(0), ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check queue capacity
	if queueCapacity, err := minipool.GetQueueCapacity(ggp, nil); err != nil {
		t.Error(err)
	} else {
		if queueCapacity.Total.Cmp(avax.EthToWei(0)) != 0 {
			t.Errorf("Incorrect queue total capacity 1 %s", queueCapacity.Total.String())
		}
		if queueCapacity.Effective.Cmp(avax.EthToWei(0)) != 0 {
			t.Errorf("Incorrect queue effective capacity 1 %s", queueCapacity.Effective.String())
		}
		if queueCapacity.NextMinipool.Cmp(avax.EthToWei(0)) != 0 {
			t.Errorf("Incorrect queue next minipool capacity 1 %s", queueCapacity.NextMinipool.String())
		}
	}

	/* TODO: Unbonded minipools are temporarily disabled
	   // Create empty deposit minipool
	   if _, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, trustedNodeAccount, eth.EthToWei(0)); err != nil { t.Fatal(err) }

	   // Get & check queue capacity
	   if queueCapacity, err := minipool.GetQueueCapacity(ggp, nil); err != nil {
	       t.Error(err)
	   } else {
	       if queueCapacity.Total.Cmp(eth.EthToWei(32)) != 0 {
	           t.Errorf("Incorrect queue total capacity 2 %s", queueCapacity.Total.String())
	       }
	       if queueCapacity.Effective.Cmp(eth.EthToWei(0)) != 0 {
	           t.Errorf("Incorrect queue effective capacity 2 %s", queueCapacity.Effective.String())
	       }
	       if queueCapacity.NextMinipool.Cmp(eth.EthToWei(32)) != 0 {
	           t.Errorf("Incorrect queue next minipool capacity 2 %s", queueCapacity.NextMinipool.String())
	       }
	   }
	*/

	// Create half deposit minipool
	if _, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, nodeAccount, avax.EthToWei(16), 1); err != nil {
		t.Fatal(err)
	}

	// Get & check queue capacity
	if queueCapacity, err := minipool.GetQueueCapacity(ggp, nil); err != nil {
		t.Error(err)
	} else {
		if queueCapacity.Total.Cmp(avax.EthToWei(16)) != 0 {
			t.Errorf("Incorrect queue total capacity 3 %s", queueCapacity.Total.String())
		}
		if queueCapacity.Effective.Cmp(avax.EthToWei(16)) != 0 {
			t.Errorf("Incorrect queue effective capacity 3 %s", queueCapacity.Effective.String())
		}
		if queueCapacity.NextMinipool.Cmp(avax.EthToWei(16)) != 0 {
			t.Errorf("Incorrect queue next minipool capacity 3 %s", queueCapacity.NextMinipool.String())
		}
	}

	// Create full deposit minipool
	if _, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, nodeAccount, avax.EthToWei(32), 2); err != nil {
		t.Fatal(err)
	}

	// Get & check queue capacity
	if queueCapacity, err := minipool.GetQueueCapacity(ggp, nil); err != nil {
		t.Error(err)
	} else {
		if queueCapacity.Total.Cmp(avax.EthToWei(32)) != 0 {
			t.Errorf("Incorrect queue total capacity 4 %s", queueCapacity.Total.String())
		}
		if queueCapacity.Effective.Cmp(avax.EthToWei(32)) != 0 {
			t.Errorf("Incorrect queue effective capacity 4 %s", queueCapacity.Effective.String())
		}
		if queueCapacity.NextMinipool.Cmp(avax.EthToWei(16)) != 0 {
			t.Errorf("Incorrect queue next minipool capacity 4 %s", queueCapacity.NextMinipool.String())
		}
	}

}
