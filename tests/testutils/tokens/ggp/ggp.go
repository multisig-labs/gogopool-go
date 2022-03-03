package ggp

import (
	"fmt"
	"math/big"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/tokens"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
)

// Mint an amount of GGP to an account
func MintGGP(ggp *gogopool.GoGoPool, ownerAccount *accounts.Account, toAccount *accounts.Account, amount *big.Int) error {

	// Get GGP token contract address
	gogoTokenGGPAddress, err := ggp.GetAddress("rocketTokenGGP")
	if err != nil {
		return err
	}

	// Mint, approve & swap fixed-supply GGP
	if err := MintFixedSupplyGGP(ggp, ownerAccount, toAccount, amount); err != nil {
		return err
	}
	if _, err := tokens.ApproveFixedSupplyGGP(ggp, *gogoTokenGGPAddress, amount, toAccount.GetTransactor()); err != nil {
		return err
	}
	if _, err := tokens.SwapFixedSupplyGGPForGGP(ggp, amount, toAccount.GetTransactor()); err != nil {
		return err
	}

	// Return
	return nil

}

// Mint an amount of fixed-supply GGP to an account
func MintFixedSupplyGGP(ggp *gogopool.GoGoPool, ownerAccount *accounts.Account, toAccount *accounts.Account, amount *big.Int) error {
	gogoTokenFixedSupplyGGP, err := ggp.GetContract("rocketTokenGGPFixedSupply")
	if err != nil {
		return err
	}
	if _, err := gogoTokenFixedSupplyGGP.Transact(ownerAccount.GetTransactor(), "mint", toAccount.Address, amount); err != nil {
		return fmt.Errorf("Could not mint fixed-supply GGP tokens to %s: %w", toAccount.Address.Hex(), err)
	}
	return nil
}
