package dao

import (
	"github.com/multisig-labs/gogopool-go/dao"
	trustednodedao "github.com/multisig-labs/gogopool-go/dao/trustednode"
	"github.com/multisig-labs/gogopool-go/gogopool"
	trustednodesettings "github.com/multisig-labs/gogopool-go/settings/trustednode"
	ggptypes "github.com/multisig-labs/gogopool-go/types"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
)

// Pass and execute a proposal
func PassAndExecuteProposal(ggp *gogopool.GoGoPool, proposalId uint64, trustedNodeAccounts []*accounts.Account) error {

	// Get proposal voting delay
	voteDelayTime, err := trustednodesettings.GetProposalVoteDelayTime(ggp, nil)
	if err != nil {
		return err
	}

	// Increase time until proposal voting delay has passed
	if err := evm.IncreaseTime(int(voteDelayTime)); err != nil {
		return err
	}

	// Vote on proposal until passed
	for _, account := range trustedNodeAccounts {
		if state, err := dao.GetProposalState(ggp, proposalId, nil); err != nil {
			return err
		} else if state == ggptypes.Succeeded {
			break
		}
		if _, err := trustednodedao.VoteOnProposal(ggp, proposalId, true, account.GetTransactor()); err != nil {
			return err
		}
	}

	// Execute proposal
	if _, err := trustednodedao.ExecuteProposal(ggp, proposalId, trustedNodeAccounts[0].GetTransactor()); err != nil {
		return err
	}

	// Return
	return nil

}
