package trustednode

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	trustednodedao "github.com/multisig-labs/gogopool-go/dao/trustednode"
	"github.com/multisig-labs/gogopool-go/node"
	trustednodesettings "github.com/multisig-labs/gogopool-go/settings/trustednode"
	"github.com/multisig-labs/gogopool-go/utils/avax"

	"github.com/multisig-labs/gogopool-go/tests/testutils/evm"
	minipoolutils "github.com/multisig-labs/gogopool-go/tests/testutils/minipool"
	nodeutils "github.com/multisig-labs/gogopool-go/tests/testutils/node"
)

func TestMemberDetails(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Disable min commission rate for unbonded pools
	if _, err := trustednodesettings.BootstrapMinipoolUnbondedMinFee(ggp, uint64(0), ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check minimum member count
	if minMemberCount, err := trustednodedao.GetMinimumMemberCount(ggp, nil); err != nil {
		t.Error(err)
	} else if minMemberCount == 0 {
		t.Error("Incorrect trusted node DAO minimum member count")
	}

	// Get & check initial member details
	if members, err := trustednodedao.GetMembers(ggp, nil); err != nil {
		t.Error(err)
	} else if len(members) != 0 {
		t.Error("Incorrect initial trusted node DAO member count")
	}

	// Set proposal cooldown
	if _, err := trustednodesettings.BootstrapProposalCooldownTime(ggp, 0, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Register nodes
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", trustedNodeAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", trustedNodeAccount2.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", trustedNodeAccount3.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := node.RegisterNode(ggp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Bootstrap trusted node DAO member
	memberId := "coolguy"
	memberEmail := "coolguy@gogopool.net"
	if _, err := trustednodedao.BootstrapMember(ggp, memberId, memberEmail, trustedNodeAccount1.Address, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodedao.BootstrapMember(ggp, memberId, memberEmail, trustedNodeAccount2.Address, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodedao.BootstrapMember(ggp, memberId, memberEmail, trustedNodeAccount3.Address, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get GGP bond amount
	ggpBondAmount, err := trustednodesettings.GetGGPBond(ggp, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Mint trusted node GGP bond & join trusted node DAO
	if err := nodeutils.MintTrustedNodeBond(ggp, ownerAccount, trustedNodeAccount1); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.MintTrustedNodeBond(ggp, ownerAccount, trustedNodeAccount2); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.MintTrustedNodeBond(ggp, ownerAccount, trustedNodeAccount3); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodedao.Join(ggp, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodedao.Join(ggp, trustedNodeAccount2.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if _, err := trustednodedao.Join(ggp, trustedNodeAccount3.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Submit a proposal
	if _, _, err := trustednodedao.ProposeMemberLeave(ggp, "bye", trustedNodeAccount1.Address, trustedNodeAccount1.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Create an unbonded minipool
	if _, err := minipoolutils.CreateMinipool(t, ggp, ownerAccount, trustedNodeAccount1, avax.EthToWei(16), 1); err != nil {
		t.Fatal(err)
	}

	// Get & check updated member details
	if members, err := trustednodedao.GetMembers(ggp, nil); err != nil {
		t.Error(err)
	} else if len(members) != 3 {
		t.Error("Incorrect updated trusted node DAO member count")
	} else {
		member := members[0]
		if !bytes.Equal(member.Address.Bytes(), trustedNodeAccount1.Address.Bytes()) {
			t.Errorf("Incorrect member address %s", member.Address.Hex())
		}
		if !member.Exists {
			t.Error("Incorrect member exists status")
		}
		if member.ID != memberId {
			t.Errorf("Incorrect member ID %s", member.ID)
		}
		if member.Url != memberEmail {
			t.Errorf("Incorrect member email %s", member.Url)
		}
		if member.JoinedTime == 0 {
			t.Errorf("Incorrect member joined time %d", member.JoinedTime)
		}
		if member.LastProposalTime == 0 {
			t.Errorf("Incorrect member last proposal time %d", member.LastProposalTime)
		}
		if member.GGPBondAmount.Cmp(ggpBondAmount) != 0 {
			t.Errorf("Incorrect member GGP bond amount %s", member.GGPBondAmount.String())
		}
		/* TEMPORARILY DISABLED
		   if member.UnbondedValidatorCount != 1 {
		       t.Errorf("Incorrect member unbonded validator count %d", member.UnbondedValidatorCount)
		   }
		*/
	}

}

func TestUpgradeContract(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Upgrade contract
	contractName := "gogoDepositPool"
	contractNewAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	contractNewAbi := "[{\"name\":\"foo\",\"type\":\"function\",\"inputs\":[],\"outputs\":[]}]"
	if _, err := trustednodedao.BootstrapUpgrade(ggp, "upgradeContract", contractName, contractNewAbi, contractNewAddress, ownerAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated contract details
	if contractAddress, err := ggp.GetAddress(contractName); err != nil {
		t.Error(err)
	} else if !bytes.Equal(contractAddress.Bytes(), contractNewAddress.Bytes()) {
		t.Errorf("Incorrect updated contract address %s", contractAddress.Hex())
	}
	if contractAbi, err := ggp.GetABI(contractName); err != nil {
		t.Error(err)
	} else if _, ok := contractAbi.Methods["foo"]; !ok {
		t.Errorf("Incorrect updated contract ABI")
	}

}
