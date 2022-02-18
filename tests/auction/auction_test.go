package auction

import (
	"math/big"
	"testing"

	"github.com/multisig-labs/gogopool-go/settings/trustednode"

	"github.com/multisig-labs/gogopool-go/auction"
	"github.com/multisig-labs/gogopool-go/network"
	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/tokens"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	auctionutils "github.com/multisig-labs/gogopool-go/tests/testutils/auction"
	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestAuctionDetails(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Register node
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount1); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount2); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount3); err != nil {
		t.Fatal(err)
	}

	// Disable min commission rate for unbonded pools
	if _, err := trustednode.BootstrapMinipoolUnbondedMinFee(ggp, uint64(0), ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check initial GGP balances
	totalBalance1, err := auction.GetTotalGGPBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	allottedBalance1, err := auction.GetAllottedGGPBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	remainingBalance1, err := auction.GetRemainingGGPBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	if totalBalance1.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial auction contract total GGP balance %s", totalBalance1.String())
	}
	if allottedBalance1.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial auction contract allotted GGP balance %s", allottedBalance1.String())
	}
	if remainingBalance1.Cmp(totalBalance1) != 0 {
		t.Errorf("Incorrect initial auction contract remaining GGP balance %s", remainingBalance1.String())
	}

	// Mint slashed GGP to auction contract
	if err := auctionutils.CreateSlashedGGP(t, ggp, ownerAccount, trustedNodeAccount1, trustedNodeAccount2, userAccount1); err != nil {
		t.Fatal(err)
	}

	// Get & check updated GGP balances
	totalBalance2, err := auction.GetTotalGGPBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	allottedBalance2, err := auction.GetAllottedGGPBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	remainingBalance2, err := auction.GetRemainingGGPBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	if totalBalance2.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Incorrect updated auction contract total GGP balance 1 %s", totalBalance2.String())
	}
	if allottedBalance2.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect updated auction contract allotted GGP balance 1 %s", allottedBalance2.String())
	}
	if remainingBalance2.Cmp(totalBalance2) != 0 {
		t.Errorf("Incorrect updated auction contract remaining GGP balance 1 %s", remainingBalance2.String())
	}

	// Create a new lot
	if _, _, err := auction.CreateLot(ggp, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated GGP balances
	totalBalance3, err := auction.GetTotalGGPBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	allottedBalance3, err := auction.GetAllottedGGPBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	remainingBalance3, err := auction.GetRemainingGGPBalance(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	var expectedRemainingBalance big.Int
	expectedRemainingBalance.Sub(totalBalance3, allottedBalance3)
	if allottedBalance3.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Incorrect updated auction contract allotted GGP balance 2 %s", allottedBalance3.String())
	}
	if remainingBalance3.Cmp(&expectedRemainingBalance) != 0 {
		t.Errorf("Incorrect updated auction contract remaining GGP balance 2 %s", remainingBalance3.String())
	}

}

func TestLotDetails(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Register node
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount1); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount2); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount3); err != nil {
		t.Fatal(err)
	}

	// Disable min commission rate for unbonded pools
	if _, err := trustednode.BootstrapMinipoolUnbondedMinFee(ggp, uint64(0), ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Set network parameters
	if _, err := network.SubmitPrices(ggp, 1, avax.EthToWei(1), avax.EthToWei(24), trustedNodeAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := network.SubmitPrices(ggp, 1, avax.EthToWei(1), avax.EthToWei(24), trustedNodeAccount2.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := protocol.BootstrapLotStartingPriceRatio(ggp, 1.0, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := protocol.BootstrapLotReservePriceRatio(ggp, 0.5, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := protocol.BootstrapLotMaximumEthValue(ggp, avax.EthToWei(10), ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := protocol.BootstrapLotDuration(ggp, 5, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Mint slashed GGP to auction contract
	if err := auctionutils.CreateSlashedGGP(t, ggp, ownerAccount, trustedNodeAccount1, trustedNodeAccount2, userAccount1); err != nil {
		t.Fatal(err)
	}

	// Get & check initial lot details
	if lots, err := auction.GetLots(ggp, nil); err != nil {
		t.Error(err)
	} else if len(lots) != 0 {
		t.Error("Incorrect initial lot count")
	}
	if lots, err := auction.GetLotsWithBids(ggp, userAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if len(lots) != 0 {
		t.Error("Incorrect initial lot count")
	}

	// Create lots
	lot1Index, _, err := auction.CreateLot(ggp, userAccount1.GetTransactor())
	if err != nil {
		t.Fatal(err)
	}
	lot2Index, _, err := auction.CreateLot(ggp, userAccount1.GetTransactor())
	if err != nil {
		t.Fatal(err)
	}

	// Place bid on lot 1
	bidAmount := avax.EthToWei(1)
	bid1Opts := userAccount1.GetTransactor()
	bid1Opts.Value = bidAmount
	if _, err := auction.PlaceBid(ggp, lot1Index, bid1Opts); err != nil {
		t.Fatal(err)
	}

	// Place another bid on lot 1 to clear it
	bid2Opts := userAccount2.GetTransactor()
	bid2Opts.Value = avax.EthToWei(1000)
	if _, err := auction.PlaceBid(ggp, lot1Index, bid2Opts); err != nil {
		t.Fatal(err)
	}

	// Mine blocks until lot 2 hits reserve price & recover unclaimed GGP from it
	if err := evm.MineBlocks(5); err != nil {
		t.Fatal(err)
	}
	if _, err := auction.RecoverUnclaimedGGP(ggp, lot2Index, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated lot details
	if lots, err := auction.GetLots(ggp, nil); err != nil {
		t.Error(err)
	} else if len(lots) != 2 {
		t.Error("Incorrect updated lot count")
	} else if lots[0].Index != lot1Index || lots[1].Index != lot2Index {
		t.Error("Incorrect lot indexes")
	}
	if lots, err := auction.GetLotsWithBids(ggp, userAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if len(lots) != 2 {
		t.Error("Incorrect updated lot count")
	} else {
		lot1 := lots[0]
		lot2 := lots[1]

		// Lot 1
		if lot1.Index != lot1Index {
			t.Errorf("Incorrect lot index %d", lot1.Index)
		}
		if !lot1.Exists {
			t.Error("Incorrect lot exists status")
		}
		if lot1.StartBlock == 0 {
			t.Errorf("Incorrect lot start block %d", lot1.StartBlock)
		}
		if lot1.EndBlock <= lot1.StartBlock {
			t.Errorf("Incorrect lot end block %d", lot1.EndBlock)
		}
		if lot1.StartPrice.Cmp(avax.EthToWei(1)) != 0 {
			t.Errorf("Incorrect lot start price %s", lot1.StartPrice.String())
		}
		if lot1.ReservePrice.Cmp(avax.EthToWei(0.5)) != 0 {
			t.Errorf("Incorrect lot reserve price %s", lot1.ReservePrice.String())
		}
		if lot1.PriceAtCurrentBlock.Cmp(lot1.StartPrice) == 1 || lot1.PriceAtCurrentBlock.Cmp(lot1.ReservePrice) == -1 {
			t.Errorf("Incorrect lot price at current block %s", lot1.PriceAtCurrentBlock.String())
		}
		if lot1.PriceByTotalBids.Cmp(lot1.StartPrice) == 1 || lot1.PriceByTotalBids.Cmp(lot1.ReservePrice) == -1 {
			t.Errorf("Incorrect lot price at current block %s", lot1.PriceByTotalBids.String())
		}
		if lot1.CurrentPrice.Cmp(lot1.StartPrice) == 1 || lot1.CurrentPrice.Cmp(lot1.ReservePrice) == -1 {
			t.Errorf("Incorrect lot price at current block %s", lot1.CurrentPrice.String())
		}
		if lot1.TotalGGPAmount.Cmp(avax.EthToWei(10)) != 0 {
			t.Errorf("Incorrect lot total GGP amount %s", lot1.TotalGGPAmount.String())
		}
		if lot1.ClaimedGGPAmount.Cmp(avax.EthToWei(10)) != 0 {
			t.Errorf("Incorrect lot claimed GGP amount %s", lot1.ClaimedGGPAmount.String())
		}
		if lot1.RemainingGGPAmount.Cmp(big.NewInt(0)) != 0 {
			t.Errorf("Incorrect lot remaining GGP amount %s", lot1.RemainingGGPAmount.String())
		}
		if lot1.TotalBidAmount.Cmp(bidAmount) != 1 {
			t.Errorf("Incorrect lot total bid amount %s", lot1.TotalBidAmount.String())
		}
		if lot1.AddressBidAmount.Cmp(bidAmount) != 0 {
			t.Errorf("Incorrect lot address bid amount %s", lot1.AddressBidAmount.String())
		}
		if !lot1.Cleared {
			t.Error("Incorrect lot cleared status")
		}
		if lot1.GGPRecovered {
			t.Error("Incorrect lot GGP recovered status")
		}

		// Lot 1 prices at blocks
		if priceAtBlock, err := auction.GetLotPriceAtBlock(ggp, lot1Index, 0, nil); err != nil {
			t.Error(err)
		} else if priceAtBlock.Cmp(lot1.StartPrice) != 0 {
			t.Errorf("Incorrect lot price at block 1 %s", priceAtBlock.String())
		}
		if priceAtBlock, err := auction.GetLotPriceAtBlock(ggp, lot1Index, 1000000, nil); err != nil {
			t.Error(err)
		} else if priceAtBlock.Cmp(lot1.ReservePrice) != 0 {
			t.Errorf("Incorrect lot price at block 2 %s", priceAtBlock.String())
		}

		// Lot 2
		if lot2.Index != lot2Index {
			t.Errorf("Incorrect lot index %d", lot2.Index)
		}
		if !lot2.GGPRecovered {
			t.Error("Incorrect lot GGP recovered status")
		}

	}

	// Get & check initial bidder GGP balance
	if ggpBalance, err := tokens.GetGGPBalance(ggp, userAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if ggpBalance.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Incorrect initial bidder GGP balance %s", ggpBalance.String())
	}

	// Claim bid on lot 1
	if _, err := auction.ClaimBid(ggp, lot1Index, userAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated bidder GGP balance
	if ggpBalance, err := tokens.GetGGPBalance(ggp, userAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if ggpBalance.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Incorrect updated bidder GGP balance %s", ggpBalance.String())
	}

}
