package trustednode

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/multisig-labs/gogopool-go/dao"
	trustednodedao "github.com/multisig-labs/gogopool-go/dao/trustednode"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/node"
	trustednodesettings "github.com/multisig-labs/gogopool-go/settings/trustednode"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	daoutils "github.com/multisig-labs/gogopool-go/tests/testutils/dao"
	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestProposeInviteMember(t *testing.T) {

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

	// Submit, pass & execute invite member proposal
	proposalMemberAddress := nodeAccount.Address
	proposalMemberId := "coolguy"
	proposalMemberEmail := "coolguy@gogopool.net"
	proposalId, _, err := trustednodedao.ProposeInviteMember(ggp, "invite coolguy", proposalMemberAddress, proposalMemberId, proposalMemberEmail, trustedNodeAccount1.GetTransactor())
	if err != nil {
		t.Fatal(err)
	}
	if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Fatal(err)
	}

	// Get & check initial member exists status
	if exists, err := trustednodedao.GetMemberExists(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if exists {
		t.Error("Incorrect initial member exists status")
	}

	// Mint trusted node GGP bond & join trusted node DAO
	if err := nodeutils.MintTrustedNodeBond(ggp, ownerAccount, nodeAccount); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodedao.Join(ggp, nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated member exists status
	if exists, err := trustednodedao.GetMemberExists(ggp, nodeAccount.Address, nil); err != nil {
		t.Error(err)
	} else if !exists {
		t.Error("Incorrect updated member exists status")
	}

	// Get & check proposal payload string
	if payloadStr, err := dao.GetProposalPayloadStr(ggp, proposalId, nil); err != nil {
		t.Error(err)
	} else if payloadStr != fmt.Sprintf("proposalInvite(%s,%s,%s)", proposalMemberId, proposalMemberEmail, proposalMemberAddress.Hex()) {
		t.Errorf("Incorrect proposal payload string %s", payloadStr)
	}

	// Get & check member invite executed block
	if inviteExecutedTime, err := trustednodedao.GetMemberInviteProposalExecutedTime(ggp, proposalMemberAddress, nil); err != nil {
		t.Error(err)
	} else if inviteExecutedTime == 0 {
		t.Errorf("Incorrect member invite proposal executed time %d", inviteExecutedTime)
	}

}

func TestProposeMemberLeave(t *testing.T) {

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
	if _, err := trustednodesettings.BootstrapProposalCooldownTime(ggp, 0, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodesettings.BootstrapProposalVoteDelayTime(ggp, 5, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Register nodes
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount1); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount2); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount3); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount4); err != nil {
		t.Fatal(err)
	}

	// Submit, pass & execute member leave proposal
	proposalMemberAddress := trustedNodeAccount1.Address
	proposalId, _, err := trustednodedao.ProposeMemberLeave(ggp, "node 1 leave", proposalMemberAddress, trustedNodeAccount1.GetTransactor())
	if err != nil {
		t.Fatal(err)
	}
	if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{
		trustedNodeAccount1,
		trustedNodeAccount2,
		trustedNodeAccount3,
		trustedNodeAccount4,
	}); err != nil {
		t.Fatal(err)
	}

	// Get & check member leave executed time
	if leaveExecutedTime, err := trustednodedao.GetMemberLeaveProposalExecutedTime(ggp, proposalMemberAddress, nil); err != nil {
		t.Error(err)
	} else if leaveExecutedTime == 0 {
		t.Errorf("Incorrect member leave proposal executed time %d", leaveExecutedTime)
	}

	// Get & check initial member exists status
	if exists, err := trustednodedao.GetMemberExists(ggp, trustedNodeAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if !exists {
		t.Error("Incorrect initial member exists status")
	}

	// Leave trusted node DAO
	if _, err := trustednodedao.Leave(ggp, trustedNodeAccount1.Address, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated member exists status
	if exists, err := trustednodedao.GetMemberExists(ggp, trustedNodeAccount1.Address, nil); err != nil {
		t.Error(err)
	} else if exists {
		t.Error("Incorrect updated member exists status")
	}

	// Get & check proposal payload string
	if payloadStr, err := dao.GetProposalPayloadStr(ggp, proposalId, nil); err != nil {
		t.Error(err)
	} else if payloadStr != fmt.Sprintf("proposalLeave(%s)", proposalMemberAddress.Hex()) {
		t.Errorf("Incorrect proposal payload string %s", payloadStr)
	}

}

func TestProposeKickMember(t *testing.T) {

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
	if _, err := trustednodesettings.BootstrapProposalCooldownTime(ggp, 0, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodesettings.BootstrapProposalVoteDelayTime(ggp, 5, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Register nodes
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount1); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount2); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(ggp, ownerAccount, trustedNodeAccount3); err != nil {
		t.Fatal(err)
	}

	// Get & check initial member exists status
	if exists, err := trustednodedao.GetMemberExists(ggp, trustedNodeAccount2.Address, nil); err != nil {
		t.Error(err)
	} else if !exists {
		t.Error("Incorrect initial member exists status")
	}

	// Submit, pass & execute kick member proposal
	proposalMemberAddress := trustedNodeAccount2.Address
	proposalFineAmount := avax.EthToWei(1000)
	proposalId, _, err := trustednodedao.ProposeKickMember(ggp, "kick node 2", proposalMemberAddress, proposalFineAmount, trustedNodeAccount1.GetTransactor())
	if err != nil {
		t.Fatal(err)
	}
	if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Fatal(err)
	}

	// Get & check updated member exists status
	if exists, err := trustednodedao.GetMemberExists(ggp, trustedNodeAccount2.Address, nil); err != nil {
		t.Error(err)
	} else if exists {
		t.Error("Incorrect updated member exists status")
	}

	// Get & check proposal payload string
	if payloadStr, err := dao.GetProposalPayloadStr(ggp, proposalId, nil); err != nil {
		t.Error(err)
	} else if payloadStr != fmt.Sprintf("proposalKick(%s,%s)", proposalMemberAddress.Hex(), proposalFineAmount.String()) {
		t.Errorf("Incorrect proposal payload string %s", payloadStr)
	}

}

func TestProposeUpgradeContract(t *testing.T) {

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
	if _, err := trustednodesettings.BootstrapProposalCooldownTime(ggp, 0, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodesettings.BootstrapProposalVoteDelayTime(ggp, 5, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

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

	// Submit, pass & execute upgrade contract proposal
	proposalUpgradeType := "upgradeContract"
	proposalContractName := "rocketDepositPool"
	proposalContractAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	proposalContractAbi := "[{\"name\":\"foo\",\"type\":\"function\",\"inputs\":[],\"outputs\":[]}]"
	proposalId, _, err := trustednodedao.ProposeUpgradeContract(ggp, "upgrade gogoDepositPool", proposalUpgradeType, proposalContractName, proposalContractAbi, proposalContractAddress, trustedNodeAccount1.GetTransactor())
	if err != nil {
		t.Fatal(err)
	}
	if err := daoutils.PassAndExecuteProposal(ggp, proposalId, []*accounts.Account{trustedNodeAccount1, trustedNodeAccount2}); err != nil {
		t.Fatal(err)
	}

	// Get & check updated contract details
	if contractAddress, err := ggp.GetAddress(proposalContractName); err != nil {
		t.Error(err)
	} else if !bytes.Equal(contractAddress.Bytes(), proposalContractAddress.Bytes()) {
		t.Errorf("Incorrect updated contract address %s", contractAddress.Hex())
	}
	if contractAbi, err := ggp.GetABI(proposalContractName); err != nil {
		t.Error(err)
	} else if _, ok := contractAbi.Methods["foo"]; !ok {
		t.Errorf("Incorrect updated contract ABI")
	}

	// Get & check proposal payload string
	if payloadStr, err := dao.GetProposalPayloadStr(ggp, proposalId, nil); err != nil {
		t.Error(err)
	} else if encodedAbi, err := gogopool.EncodeAbiStr(proposalContractAbi); err != nil {
		t.Error(err)
	} else if payloadStr != fmt.Sprintf("proposalUpgrade(%s,%s,%s,%s)", proposalUpgradeType, proposalContractName, encodedAbi, proposalContractAddress.Hex()) {
		t.Errorf("Incorrect proposal payload string %s", payloadStr)
	}

}
