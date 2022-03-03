package node

import (
	"fmt"

	trustednodedao "github.com/multisig-labs/gogopool-go/dao/trustednode"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/node"
	trustednodesettings "github.com/multisig-labs/gogopool-go/settings/trustednode"
	"github.com/multisig-labs/gogopool-go/tokens"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	ggputils "github.com/multisig-labs/gogopool-go/tests/testutils/tokens/ggp"
)

// Trusted node counter
var trustedNodeIndex = 0

// Register a trusted node
func RegisterTrustedNode(ggp *gogopool.GoGoPool, ownerAccount *accounts.Account, trustedNodeAccount *accounts.Account) error {

	// Register node
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", trustedNodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Bootstrap trusted node DAO member
	if _, err := trustednodedao.BootstrapMember(ggp, fmt.Sprintf("tn%d", trustedNodeIndex), fmt.Sprintf("tn%d@gogopool.net", trustedNodeIndex), trustedNodeAccount.Address, ownerAccount.GetTransactor()); err != nil {
		return err
	}

	// Mint trusted node GGP bond
	if err := MintTrustedNodeBond(ggp, ownerAccount, trustedNodeAccount); err != nil {
		return err
	}

	// Join trusted node DAO
	if _, err := trustednodedao.Join(ggp, trustedNodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Increment trusted node counter & return
	trustedNodeIndex++
	return nil

}

// Mint trusted node DAO GGP bond to a node account and approve it for spending
func MintTrustedNodeBond(ggp *gogopool.GoGoPool, ownerAccount *accounts.Account, trustedNodeAccount *accounts.Account) error {

	// Get GGP bond amount
	ggpBondAmount, err := trustednodesettings.GetGGPBond(ggp, nil)
	if err != nil {
		return err
	}

	// Get GoGoDAONodeTrustedActions contract address
	gogoDAONodeTrustedActionsAddress, err := ggp.GetAddress("rocketDAONodeTrustedActions")
	if err != nil {
		return err
	}

	// Mint GGP to node & allow trusted node DAO contract to spend it
	if err := ggputils.MintGGP(ggp, ownerAccount, trustedNodeAccount, ggpBondAmount); err != nil {
		return err
	}
	if _, err := tokens.ApproveGGP(ggp, *gogoDAONodeTrustedActionsAddress, ggpBondAmount, trustedNodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Return
	return nil

}
