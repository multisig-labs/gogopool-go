package trustednode

import (
	"testing"

	"github.com/multisig-labs/gogopool-go/settings/trustednode"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	daoutils "github.com/multisig-labs/gogopool-go/tests/testutils/dao"
	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestBootstrapMembersSettings(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Set & get quorum
	quorum := 0.1
	if _, err := trustednode.BootstrapQuorum(ggp, quorum, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetQuorum(ggp, nil); err != nil {
		t.Error(err)
	} else if value != quorum {
		t.Error("Incorrect quorum value")
	}

	// Set & get ggp bond
	ggpBond := avax.EthToWei(1)
	if _, err := trustednode.BootstrapGGPBond(ggp, ggpBond, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetGGPBond(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(ggpBond) != 0 {
		t.Error("Incorrect ggp bond value")
	}

	// Set & get maximum unbonded minipools
	var minipoolUnbondedMax uint64 = 1
	if _, err := trustednode.BootstrapMinipoolUnbondedMax(ggp, minipoolUnbondedMax, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetMinipoolUnbondedMax(ggp, nil); err != nil {
		t.Error(err)
	} else if value != minipoolUnbondedMax {
		t.Error("Incorrect maximum unbonded minipools value")
	}

}

func TestProposeMembersSettings(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Set proposal cooldown
	if _, err := trustednode.BootstrapProposalCooldownTime(ggp, 0, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednode.BootstrapProposalVoteDelayTime(ggp, 5, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Register trusted node
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount1); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount2); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount3); err != nil {
		t.Fatal(err)
	}

	// Set & get quorum
	quorum := 0.1
	if proposalId, _, err := trustednode.ProposeQuorum(ggp, quorum, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetQuorum(ggp, nil); err != nil {
		t.Error(err)
	} else if value != quorum {
		t.Error("Incorrect quorum value")
	}

	// Set & get ggp bond
	ggpBond := avax.EthToWei(1)
	if proposalId, _, err := trustednode.ProposeGGPBond(ggp, ggpBond, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetGGPBond(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(ggpBond) != 0 {
		t.Error("Incorrect ggp bond value")
	}

	// Set & get maximum unbonded minipools
	var minipoolUnbondedMax uint64 = 1
	if proposalId, _, err := trustednode.ProposeMinipoolUnbondedMax(ggp, minipoolUnbondedMax, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetMinipoolUnbondedMax(ggp, nil); err != nil {
		t.Error(err)
	} else if value != minipoolUnbondedMax {
		t.Error("Incorrect maximum unbonded minipools value")
	}

	// Set & get member challenge cooldown period
	var memberChallengeCooldown uint64 = 1
	if proposalId, _, err := trustednode.ProposeChallengeCooldown(ggp, memberChallengeCooldown, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetChallengeCooldown(ggp, nil); err != nil {
		t.Error(err)
	} else if value != memberChallengeCooldown {
		t.Error("Incorrect member challenge cooldown value")
	}

	// Set & get member challenge window period
	var memberChallengeWindow uint64 = 1
	if proposalId, _, err := trustednode.ProposeChallengeWindow(ggp, memberChallengeWindow, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetChallengeWindow(ggp, nil); err != nil {
		t.Error(err)
	} else if value != memberChallengeWindow {
		t.Error("Incorrect member challenge window value")
	}

	// Set & get member challenge cost amount
	challengeCost := avax.EthToWei(1)
	if proposalId, _, err := trustednode.ProposeChallengeCost(ggp, challengeCost, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetChallengeCost(ggp, nil); err != nil {
		t.Error(err)
	} else if value.Cmp(challengeCost) != 0 {
		t.Error("Incorrect member challenge cost value")
	}

}
