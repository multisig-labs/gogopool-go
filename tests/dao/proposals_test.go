package dao

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/multisig-labs/gogopool-go/dao"
	trustednodedao "github.com/multisig-labs/gogopool-go/dao/trustednode"
	"github.com/multisig-labs/gogopool-go/node"
	trustednodesettings "github.com/multisig-labs/gogopool-go/settings/trustednode"
	ggptypes "github.com/multisig-labs/gogopool-go/types"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestProposalDetails(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// The DAO to check for proposals under
	proposalDaoName := "gogoDAONodeTrustedProposals"

	// Set proposal cooldown
	if _, err := trustednodesettings.BootstrapProposalCooldownTime(ggp, 0, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodesettings.BootstrapProposalVoteDelayTime(ggp, 5, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Register nodes
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount1); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount2); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount3); err != nil {
		t.Fatal(err)
	}

	// Get & check initial proposal details
	if proposals, err := dao.GetProposals(ggp, nil); err != nil {
		t.Error(err)
	} else if len(proposals) != 0 {
		t.Error("Incorrect initial proposal count")
	}
	if proposals, err := dao.GetProposalsWithMember(ggp, trustedNodeAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if len(proposals) != 0 {
		t.Error("Incorrect initial proposal count")
	}
	if daoProposals, err := dao.GetDAOProposals(ggp, proposalDaoName, nil); err != nil {
		t.Error(err)
	} else if len(daoProposals) != 0 {
		t.Error("Incorrect initial DAO proposal count")
	}
	if daoProposals, err := dao.GetDAOProposalsWithMember(ggp, proposalDaoName, trustedNodeAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if len(daoProposals) != 0 {
		t.Error("Incorrect initial DAO proposal count")
	}

	// Submit invite member proposal
	proposalMessage := "invite coolguy"
	proposalMemberAddress := nodeAccount.Address
	proposalMemberId := "coolguy"
	proposalMemberEmail := "coolguy@gogopool.net"
	proposalId, _, err := trustednodedao.ProposeInviteMember(ggp, proposalMessage, proposalMemberAddress, proposalMemberId, proposalMemberEmail, trustedNodeAccount1.GetTransactor())
	if err != nil {
		t.Fatal(err)
	}

	// Increase time until proposal voting delay has passed
	voteDelayTime, err := trustednodesettings.GetProposalVoteDelayTime(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := evm.IncreaseTime(int(voteDelayTime)); err != nil {
		t.Fatal(err)
	}

	// Vote on & execute proposal
	if _, err := trustednodedao.VoteOnProposal(ggp, proposalId, true, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodedao.VoteOnProposal(ggp, proposalId, true, trustedNodeAccount2.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodedao.ExecuteProposal(ggp, proposalId, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Submit invite member proposal & cancel it
	cancelledProposalId, _, err := trustednodedao.ProposeInviteMember(ggp, "cancel this", nodeAccount.Address, "cancel", "cancel@gogopool.net", trustedNodeAccount1.GetTransactor())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodedao.CancelProposal(ggp, cancelledProposalId, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated proposal details
	if proposals, err := dao.GetProposals(ggp, nil); err != nil {
		t.Error(err)
	} else if len(proposals) != 2 {
		t.Error("Incorrect updated proposal count")
	} else if proposals[0].ID != proposalId || proposals[1].ID != cancelledProposalId {
		t.Error("Incorrect proposal indexes")
	}
	if proposals, err := dao.GetProposalsWithMember(ggp, trustedNodeAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if len(proposals) != 2 {
		t.Error("Incorrect updated proposal count")
	} else {

		// Passed proposal
		proposal := proposals[0]
		if proposal.ID != proposalId {
			t.Errorf("Incorrect proposal ID %d", proposal.ID)
		}
		if proposal.DAO != proposalDaoName {
			t.Errorf("Incorrect proposal DAO %s", proposal.DAO)
		}
		if !bytes.Equal(proposal.ProposerAddress.Bytes(), trustedNodeAccount1.Address.Bytes()) {
			t.Errorf("Incorrect proposal proposer address %s", proposal.ProposerAddress.Hex())
		}
		if proposal.Message != proposalMessage {
			t.Errorf("Incorrect proposal message %s", proposal.Message)
		}
		if proposal.CreatedTime == 0 {
			t.Errorf("Incorrect proposal created time %d", proposal.CreatedTime)
		}
		if proposal.StartTime <= proposal.CreatedTime {
			t.Errorf("Incorrect proposal start time %d", proposal.StartTime)
		}
		if proposal.EndTime <= proposal.StartTime {
			t.Errorf("Incorrect proposal end time %d", proposal.EndTime)
		}
		if proposal.ExpiryTime <= proposal.EndTime {
			t.Errorf("Incorrect proposal expiry time %d", proposal.ExpiryTime)
		}
		if proposal.VotesRequired == 0.0 {
			t.Errorf("Incorrect proposal required votes %f", proposal.VotesRequired)
		}
		if proposal.VotesFor != 2.0 {
			t.Errorf("Incorrect proposal votes for %f", proposal.VotesFor)
		}
		if proposal.VotesAgainst != 0.0 {
			t.Errorf("Incorrect proposal votes against %f", proposal.VotesAgainst)
		}
		if !proposal.MemberVoted {
			t.Error("Incorrect proposal member voted status")
		}
		if !proposal.MemberSupported {
			t.Error("Incorrect proposal member supported status")
		}
		if proposal.IsCancelled {
			t.Error("Incorrect proposal cancelled status")
		}
		if !proposal.IsExecuted {
			t.Error("Incorrect proposal executed status")
		}
		if proposal.PayloadStr != fmt.Sprintf("proposalInvite(%s,%s,%s)", proposalMemberId, proposalMemberEmail, proposalMemberAddress.Hex()) {
			t.Errorf("Incorrect proposal payload string %s", proposal.PayloadStr)
		}
		if proposal.State != ggptypes.Executed {
			t.Errorf("Incorrect proposal state %s", proposal.State.String())
		}

		// Cancelled proposal
		cancelledProposal := proposals[1]
		if cancelledProposal.ID != cancelledProposalId {
			t.Errorf("Incorrect cancelled proposal ID %d", cancelledProposal.ID)
		}
		if !cancelledProposal.IsCancelled {
			t.Error("Incorrect cancelled proposal cancelled status")
		}

	}
	if daoProposals, err := dao.GetDAOProposals(ggp, proposalDaoName, nil); err != nil {
		t.Error(err)
	} else if len(daoProposals) != 2 {
		t.Error("Incorrect updated DAO proposal count")
	} else if daoProposals[0].ID != proposalId || daoProposals[1].ID != cancelledProposalId {
		t.Error("Incorrect DAO proposal indexes")
	}
	if daoProposals, err := dao.GetDAOProposalsWithMember(ggp, proposalDaoName, trustedNodeAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if len(daoProposals) != 2 {
		t.Error("Incorrect updated DAO proposal count")
	} else if daoProposals[0].ID != proposalId || daoProposals[1].ID != cancelledProposalId {
		t.Error("Incorrect DAO proposal indexes")
	}

}
