package trustednode

import (
	"testing"

	"github.com/multisig-labs/gogopool-go/settings/trustednode"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	daoutils "github.com/multisig-labs/gogopool-go/tests/testutils/dao"
	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestBootstrapProposalsSettings(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Set & get cooldown
	var cooldown uint64 = 1
	if _, err := trustednode.BootstrapProposalCooldownTime(ggp, cooldown, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalCooldownTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != cooldown {
		t.Error("Incorrect cooldown value")
	}

	// Set & get vote time
	var voteTime uint64 = 10
	if _, err := trustednode.BootstrapProposalVoteTime(ggp, voteTime, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalVoteTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != voteTime {
		t.Error("Incorrect vote time value")
	}

	// Set & get execute time
	var executeTime uint64 = 10
	if _, err := trustednode.BootstrapProposalExecuteTime(ggp, executeTime, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalExecuteTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != executeTime {
		t.Error("Incorrect execute time value")
	}

	// Set & get action time
	var actionTime uint64 = 10
	if _, err := trustednode.BootstrapProposalActionTime(ggp, actionTime, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalActionTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != actionTime {
		t.Error("Incorrect action time value")
	}

	// Set & get vote delay time
	var voteDelayTime uint64 = 1000
	if _, err := trustednode.BootstrapProposalVoteDelayTime(ggp, voteDelayTime, ownerAccount.GetTransactor()); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalVoteDelayTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != voteDelayTime {
		t.Error("Incorrect vote delay time value")
	}

}

func TestProposeProposalsSettings(t *testing.T) {

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

	// Set & get cooldown
	var cooldown uint64 = 1
	if proposalId, _, err := trustednode.ProposeProposalCooldownTime(ggp, cooldown, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalCooldownTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != cooldown {
		t.Error("Incorrect cooldown value")
	}

	// Set & get vote time
	var voteTime uint64 = 10
	if proposalId, _, err := trustednode.ProposeProposalVoteTime(ggp, voteTime, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalVoteTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != voteTime {
		t.Error("Incorrect vote time value")
	}

	// Set & get execute time
	var executeTime uint64 = 10
	if proposalId, _, err := trustednode.ProposeProposalExecuteTime(ggp, executeTime, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalExecuteTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != executeTime {
		t.Error("Incorrect execute time value")
	}

	// Set & get action time
	var actionTime uint64 = 10
	if proposalId, _, err := trustednode.ProposeProposalActionTime(ggp, actionTime, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalActionTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != actionTime {
		t.Error("Incorrect action time value")
	}

	// Set & get vote delay time
	var voteDelayTime uint64 = 1000
	if proposalId, _, err := trustednode.ProposeProposalVoteDelayTime(ggp, voteDelayTime, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Error(err)
	} else if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Error(err)
	} else if value, err := trustednode.GetProposalVoteDelayTime(ggp, nil); err != nil {
		t.Error(err)
	} else if value != voteDelayTime {
		t.Error("Incorrect vote delay time value")
	}

}
