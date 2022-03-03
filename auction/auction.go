package auction

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"

	"github.com/multisig-labs/gogopool-go/gogopool"
)

// Settings
const LotDetailsBatchSize = 10

// Lot details
type LotDetails struct {
	Index               uint64   `json:"index"`
	Exists              bool     `json:"exists"`
	StartBlock          uint64   `json:"startBlock"`
	EndBlock            uint64   `json:"endBlock"`
	StartPrice          *big.Int `json:"startPrice"`
	ReservePrice        *big.Int `json:"reservePrice"`
	PriceAtCurrentBlock *big.Int `json:"priceAtCurrentBlock"`
	PriceByTotalBids    *big.Int `json:"priceByTotalBids"`
	CurrentPrice        *big.Int `json:"currentPrice"`
	TotalGGPAmount      *big.Int `json:"totalGgpAmount"`
	ClaimedGGPAmount    *big.Int `json:"claimedGgpAmount"`
	RemainingGGPAmount  *big.Int `json:"remainingGgpAmount"`
	TotalBidAmount      *big.Int `json:"totalBidAmount"`
	AddressBidAmount    *big.Int `json:"addressBidAmount"`
	Cleared             bool     `json:"cleared"`
	GGPRecovered        bool     `json:"ggpRecovered"`
}

// Get all lot details
func GetLots(ggp *gogopool.GoGoPool, opts *bind.CallOpts) ([]LotDetails, error) {

	// Get lot count
	lotCount, err := GetLotCount(ggp, opts)
	if err != nil {
		return []LotDetails{}, err
	}

	// Load lot details in batches
	details := make([]LotDetails, lotCount)
	for bsi := uint64(0); bsi < lotCount; bsi += LotDetailsBatchSize {

		// Get batch start & end index
		lsi := bsi
		lei := bsi + LotDetailsBatchSize
		if lei > lotCount {
			lei = lotCount
		}

		// Load details
		var wg errgroup.Group
		for li := lsi; li < lei; li++ {
			li := li
			wg.Go(func() error {
				lotDetails, err := GetLotDetails(ggp, li, opts)
				if err == nil {
					details[li] = lotDetails
				}
				return err
			})
		}
		if err := wg.Wait(); err != nil {
			return []LotDetails{}, err
		}

	}

	// Return
	return details, nil

}

// Get all lot details with bids from an address
func GetLotsWithBids(ggp *gogopool.GoGoPool, bidder common.Address, opts *bind.CallOpts) ([]LotDetails, error) {

	// Get lot count
	lotCount, err := GetLotCount(ggp, opts)
	if err != nil {
		return []LotDetails{}, err
	}

	// Load lot details in batches
	details := make([]LotDetails, lotCount)
	for bsi := uint64(0); bsi < lotCount; bsi += LotDetailsBatchSize {

		// Get batch start & end index
		lsi := bsi
		lei := bsi + LotDetailsBatchSize
		if lei > lotCount {
			lei = lotCount
		}

		// Load details
		var wg errgroup.Group
		for li := lsi; li < lei; li++ {
			li := li
			wg.Go(func() error {
				lotDetails, err := GetLotDetailsWithBids(ggp, li, bidder, opts)
				if err == nil {
					details[li] = lotDetails
				}
				return err
			})
		}
		if err := wg.Wait(); err != nil {
			return []LotDetails{}, err
		}

	}

	// Return
	return details, nil

}

// Get a lot's details
func GetLotDetails(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (LotDetails, error) {

	// Data
	var wg errgroup.Group
	var exists bool
	var startBlock uint64
	var endBlock uint64
	var startPrice *big.Int
	var reservePrice *big.Int
	var priceAtCurrentBlock *big.Int
	var priceByTotalBids *big.Int
	var currentPrice *big.Int
	var totalGgpAmount *big.Int
	var claimedGgpAmount *big.Int
	var remainingGgpAmount *big.Int
	var totalBidAmount *big.Int
	var cleared bool
	var ggpRecovered bool

	// Load data
	wg.Go(func() error {
		var err error
		exists, err = GetLotExists(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		startBlock, err = GetLotStartBlock(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		endBlock, err = GetLotEndBlock(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		startPrice, err = GetLotStartPrice(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		reservePrice, err = GetLotReservePrice(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		priceAtCurrentBlock, err = GetLotPriceAtCurrentBlock(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		priceByTotalBids, err = GetLotPriceByTotalBids(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		currentPrice, err = GetLotCurrentPrice(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		totalGgpAmount, err = GetLotTotalGGPAmount(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		claimedGgpAmount, err = GetLotClaimedGGPAmount(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		remainingGgpAmount, err = GetLotRemainingGGPAmount(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		totalBidAmount, err = GetLotTotalBidAmount(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		cleared, err = GetLotIsCleared(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		ggpRecovered, err = GetLotGGPRecovered(ggp, lotIndex, opts)
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return LotDetails{}, err
	}

	// Return
	return LotDetails{
		Index:               lotIndex,
		Exists:              exists,
		StartBlock:          startBlock,
		EndBlock:            endBlock,
		StartPrice:          startPrice,
		ReservePrice:        reservePrice,
		PriceAtCurrentBlock: priceAtCurrentBlock,
		PriceByTotalBids:    priceByTotalBids,
		CurrentPrice:        currentPrice,
		TotalGGPAmount:      totalGgpAmount,
		ClaimedGGPAmount:    claimedGgpAmount,
		RemainingGGPAmount:  remainingGgpAmount,
		TotalBidAmount:      totalBidAmount,
		Cleared:             cleared,
		GGPRecovered:        ggpRecovered,
	}, nil

}

// Get a lot's details with address bid amounts
func GetLotDetailsWithBids(ggp *gogopool.GoGoPool, lotIndex uint64, bidder common.Address, opts *bind.CallOpts) (LotDetails, error) {

	// Data
	var wg errgroup.Group
	var details LotDetails
	var addressBidAmount *big.Int

	// Load data
	wg.Go(func() error {
		var err error
		details, err = GetLotDetails(ggp, lotIndex, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		addressBidAmount, err = GetLotAddressBidAmount(ggp, lotIndex, bidder, opts)
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return LotDetails{}, err
	}

	// Return
	details.AddressBidAmount = addressBidAmount
	return details, nil

}

// Get the total GGP balance of the auction contract
func GetTotalGGPBalance(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	totalGgpBalance := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, totalGgpBalance, "getTotalGGPBalance"); err != nil {
		return nil, fmt.Errorf("Could not get auction contract total GGP balance: %w", err)
	}
	return *totalGgpBalance, nil
}

// Get the allotted GGP balance of the auction contract
func GetAllottedGGPBalance(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	allottedGgpBalance := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, allottedGgpBalance, "getAllottedGGPBalance"); err != nil {
		return nil, fmt.Errorf("Could not get auction contract allotted GGP balance: %w", err)
	}
	return *allottedGgpBalance, nil
}

// Get the remaining GGP balance of the auction contract
func GetRemainingGGPBalance(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	remainingGgpBalance := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, remainingGgpBalance, "getRemainingGGPBalance"); err != nil {
		return nil, fmt.Errorf("Could not get auction contract remaining GGP balance: %w", err)
	}
	return *remainingGgpBalance, nil
}

// Get the number of lots for auction
func GetLotCount(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return 0, err
	}
	lotCount := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotCount, "getLotCount"); err != nil {
		return 0, fmt.Errorf("Could not get lot count: %w", err)
	}
	return (*lotCount).Uint64(), nil
}

// Lot details
func GetLotExists(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (bool, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return false, err
	}
	lotExists := new(bool)
	if err := gogoAuctionManager.Call(opts, lotExists, "getLotExists", big.NewInt(int64(lotIndex))); err != nil {
		return false, fmt.Errorf("Could not get lot %d exists status: %w", lotIndex, err)
	}
	return *lotExists, nil
}
func GetLotStartBlock(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (uint64, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return 0, err
	}
	lotStartBlock := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotStartBlock, "getLotStartBlock", big.NewInt(int64(lotIndex))); err != nil {
		return 0, fmt.Errorf("Could not get lot %d start block: %w", lotIndex, err)
	}
	return (*lotStartBlock).Uint64(), nil
}
func GetLotEndBlock(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (uint64, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return 0, err
	}
	lotEndBlock := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotEndBlock, "getLotEndBlock", big.NewInt(int64(lotIndex))); err != nil {
		return 0, fmt.Errorf("Could not get lot %d end block: %w", lotIndex, err)
	}
	return (*lotEndBlock).Uint64(), nil
}
func GetLotStartPrice(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotStartPrice := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotStartPrice, "getLotStartPrice", big.NewInt(int64(lotIndex))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d start price: %w", lotIndex, err)
	}
	return *lotStartPrice, nil
}
func GetLotReservePrice(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotReservePrice := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotReservePrice, "getLotReservePrice", big.NewInt(int64(lotIndex))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d reserve price: %w", lotIndex, err)
	}
	return *lotReservePrice, nil
}
func GetLotTotalGGPAmount(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotTotalGgpAmount := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotTotalGgpAmount, "getLotTotalGGPAmount", big.NewInt(int64(lotIndex))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d total GGP amount: %w", lotIndex, err)
	}
	return *lotTotalGgpAmount, nil
}
func GetLotTotalBidAmount(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotTotalBidAmount := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotTotalBidAmount, "getLotTotalBidAmount", big.NewInt(int64(lotIndex))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d total ETH bid amount: %w", lotIndex, err)
	}
	return *lotTotalBidAmount, nil
}
func GetLotGGPRecovered(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (bool, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return false, err
	}
	lotGgpRecovered := new(bool)
	if err := gogoAuctionManager.Call(opts, lotGgpRecovered, "getLotGGPRecovered", big.NewInt(int64(lotIndex))); err != nil {
		return false, fmt.Errorf("Could not get lot %d GGP recovered status: %w", lotIndex, err)
	}
	return *lotGgpRecovered, nil
}
func GetLotPriceAtCurrentBlock(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotPriceAtCurrentBlock := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotPriceAtCurrentBlock, "getLotPriceAtCurrentBlock", big.NewInt(int64(lotIndex))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d price by current block: %w", lotIndex, err)
	}
	return *lotPriceAtCurrentBlock, nil
}
func GetLotPriceByTotalBids(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotPriceByTotalBids := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotPriceByTotalBids, "getLotPriceByTotalBids", big.NewInt(int64(lotIndex))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d price by total bids: %w", lotIndex, err)
	}
	return *lotPriceByTotalBids, nil
}
func GetLotCurrentPrice(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotCurrentPrice := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotCurrentPrice, "getLotCurrentPrice", big.NewInt(int64(lotIndex))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d current price: %w", lotIndex, err)
	}
	return *lotCurrentPrice, nil
}
func GetLotClaimedGGPAmount(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotClaimedGgpAmount := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotClaimedGgpAmount, "getLotClaimedGGPAmount", big.NewInt(int64(lotIndex))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d claimed GGP amount: %w", lotIndex, err)
	}
	return *lotClaimedGgpAmount, nil
}
func GetLotRemainingGGPAmount(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotRemainingGgpAmount := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotRemainingGgpAmount, "getLotRemainingGGPAmount", big.NewInt(int64(lotIndex))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d remaining GGP amount: %w", lotIndex, err)
	}
	return *lotRemainingGgpAmount, nil
}
func GetLotIsCleared(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.CallOpts) (bool, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return false, err
	}
	lotIsCleared := new(bool)
	if err := gogoAuctionManager.Call(opts, lotIsCleared, "getLotIsCleared", big.NewInt(int64(lotIndex))); err != nil {
		return false, fmt.Errorf("Could not get lot %d cleared status: %w", lotIndex, err)
	}
	return *lotIsCleared, nil
}

// Get the price of a lot at a specific block
func GetLotPriceAtBlock(ggp *gogopool.GoGoPool, lotIndex, blockNumber uint64, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lotPriceAtBlock := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lotPriceAtBlock, "getLotPriceAtBlock", big.NewInt(int64(lotIndex)), big.NewInt(int64(blockNumber))); err != nil {
		return nil, fmt.Errorf("Could not get lot %d price at block: %w", lotIndex, err)
	}
	return *lotPriceAtBlock, nil
}

// Get the ETH amount bid on a lot by an address
func GetLotAddressBidAmount(ggp *gogopool.GoGoPool, lotIndex uint64, bidder common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return nil, err
	}
	lot := new(*big.Int)
	if err := gogoAuctionManager.Call(opts, lot, "getLotAddressBidAmount", big.NewInt(int64(lotIndex)), bidder); err != nil {
		return nil, fmt.Errorf("Could not get lot %d address ETH bid amount: %w", lotIndex, err)
	}
	return *lot, nil
}

// Estimate the gas of CreateLot
func EstimateCreateLotGas(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoAuctionManager.GetTransactionGasInfo(opts, "createLot")
}

// Create a new lot
func CreateLot(ggp *gogopool.GoGoPool, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return 0, common.Hash{}, err
	}
	lotCount, err := GetLotCount(ggp, nil)
	if err != nil {
		return 0, common.Hash{}, err
	}
	hash, err := gogoAuctionManager.Transact(opts, "createLot")
	if err != nil {
		return 0, common.Hash{}, fmt.Errorf("Could not create lot: %w", err)
	}
	return lotCount, hash, nil
}

// Estimate the gas of PlaceBid
func EstimatePlaceBidGas(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoAuctionManager.GetTransactionGasInfo(opts, "placeBid", big.NewInt(int64(lotIndex)))
}

// Place a bid on a lot
func PlaceBid(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.TransactOpts) (common.Hash, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoAuctionManager.Transact(opts, "placeBid", big.NewInt(int64(lotIndex)))
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not place bid on lot %d: %w", lotIndex, err)
	}
	return hash, nil
}

// Estimate the gas of ClaimBid
func EstimateClaimBidGas(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoAuctionManager.GetTransactionGasInfo(opts, "claimBid", big.NewInt(int64(lotIndex)))
}

// Claim GGP from a lot that was bid on
func ClaimBid(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.TransactOpts) (common.Hash, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoAuctionManager.Transact(opts, "claimBid", big.NewInt(int64(lotIndex)))
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not claim bid from lot %d: %w", lotIndex, err)
	}
	return hash, nil
}

// Estimate the gas of RecoverUnclaimedGGP
func EstimateRecoverUnclaimedGGPGas(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoAuctionManager.GetTransactionGasInfo(opts, "recoverUnclaimedGGP", big.NewInt(int64(lotIndex)))
}

// Recover unclaimed GGP from a lot
func RecoverUnclaimedGGP(ggp *gogopool.GoGoPool, lotIndex uint64, opts *bind.TransactOpts) (common.Hash, error) {
	gogoAuctionManager, err := getGoGoAuctionManager(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoAuctionManager.Transact(opts, "recoverUnclaimedGGP", big.NewInt(int64(lotIndex)))
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not recover unclaimed GGP from lot %d: %w", lotIndex, err)
	}
	return hash, nil
}

// Get contracts
var gogoAuctionManagerLock sync.Mutex

func getGoGoAuctionManager(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoAuctionManagerLock.Lock()
	defer gogoAuctionManagerLock.Unlock()
	return ggp.GetContract("rocketAuctionManager")
}
