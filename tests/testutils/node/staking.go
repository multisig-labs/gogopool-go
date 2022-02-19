package node

import (
	"math/big"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/node"
	"github.com/multisig-labs/gogopool-go/tokens"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	ggputils "github.com/multisig-labs/gogopool-go/tests/testutils/tokens/ggp"
)

// Mint & stake an amount of GGP against a node
func StakeGGP(ggp *gogopool.GoGoPool, ownerAccount, nodeAccount *accounts.Account, amount *big.Int) error {

	// Get GoGoNodeStaking contract address
	gogoNodeStakingAddress, err := ggp.GetAddress("gogoNodeStaking")
	if err != nil {
		return err
	}

	// Mint, approve & stake GGP
	if err := ggputils.MintGGP(ggp, ownerAccount, nodeAccount, amount); err != nil {
		return err
	}
	if _, err := tokens.ApproveGGP(ggp, *gogoNodeStakingAddress, amount, nodeAccount.GetTransactor()); err != nil {
		return err
	}
	if _, err := node.StakeGGP(ggp, amount, nodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Return
	return nil

}
