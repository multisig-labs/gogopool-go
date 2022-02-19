package trustednode

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/utils/strings"
)

// Settings
const (
	MemberAddressBatchSize = 50
	MemberDetailsBatchSize = 20
)

// Proposal details
type MemberDetails struct {
	Address                common.Address `json:"address"`
	Exists                 bool           `json:"exists"`
	ID                     string         `json:"id"`
	Url                    string         `json:"url"`
	JoinedTime             uint64         `json:"joinedTime"`
	LastProposalTime       uint64         `json:"lastProposalTime"`
	GGPBondAmount          *big.Int       `json:"ggpBondAmount"`
	UnbondedValidatorCount uint64         `json:"unbondedValidatorCount"`
}

// Get all member details
func GetMembers(ggp *gogopool.GoGoPool, opts *bind.CallOpts) ([]MemberDetails, error) {

	// Get member addresses
	memberAddresses, err := GetMemberAddresses(ggp, opts)
	if err != nil {
		return []MemberDetails{}, err
	}

	// Load member details in batches
	details := make([]MemberDetails, len(memberAddresses))
	for bsi := 0; bsi < len(memberAddresses); bsi += MemberDetailsBatchSize {

		// Get batch start & end index
		msi := bsi
		mei := bsi + MemberDetailsBatchSize
		if mei > len(memberAddresses) {
			mei = len(memberAddresses)
		}

		// Load details
		var wg errgroup.Group
		for mi := msi; mi < mei; mi++ {
			mi := mi
			wg.Go(func() error {
				memberAddress := memberAddresses[mi]
				memberDetails, err := GetMemberDetails(ggp, memberAddress, opts)
				if err == nil {
					details[mi] = memberDetails
				}
				return err
			})
		}
		if err := wg.Wait(); err != nil {
			return []MemberDetails{}, err
		}

	}

	// Return
	return details, nil

}

// Get all member addresses
func GetMemberAddresses(ggp *gogopool.GoGoPool, opts *bind.CallOpts) ([]common.Address, error) {

	// Get member count
	memberCount, err := GetMemberCount(ggp, opts)
	if err != nil {
		return []common.Address{}, err
	}

	// Load member addresses in batches
	addresses := make([]common.Address, memberCount)
	for bsi := uint64(0); bsi < memberCount; bsi += MemberAddressBatchSize {

		// Get batch start & end index
		msi := bsi
		mei := bsi + MemberAddressBatchSize
		if mei > memberCount {
			mei = memberCount
		}

		// Load addresses
		var wg errgroup.Group
		for mi := msi; mi < mei; mi++ {
			mi := mi
			wg.Go(func() error {
				address, err := GetMemberAt(ggp, mi, opts)
				if err == nil {
					addresses[mi] = address
				}
				return err
			})
		}
		if err := wg.Wait(); err != nil {
			return []common.Address{}, err
		}

	}

	// Return
	return addresses, nil

}

// Get a member's details
func GetMemberDetails(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (MemberDetails, error) {

	// Data
	var wg errgroup.Group
	var exists bool
	var id string
	var url string
	var joinedTime uint64
	var lastProposalTime uint64
	var ggpBondAmount *big.Int
	var unbondedValidatorCount uint64

	// Load data
	wg.Go(func() error {
		var err error
		exists, err = GetMemberExists(ggp, memberAddress, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		id, err = GetMemberID(ggp, memberAddress, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		url, err = GetMemberUrl(ggp, memberAddress, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		joinedTime, err = GetMemberJoinedTime(ggp, memberAddress, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		lastProposalTime, err = GetMemberLastProposalTime(ggp, memberAddress, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		ggpBondAmount, err = GetMemberGGPBondAmount(ggp, memberAddress, opts)
		return err
	})
	wg.Go(func() error {
		var err error
		unbondedValidatorCount, err = GetMemberUnbondedValidatorCount(ggp, memberAddress, opts)
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return MemberDetails{}, err
	}

	// Return
	return MemberDetails{
		Address:                memberAddress,
		Exists:                 exists,
		ID:                     id,
		Url:                    url,
		JoinedTime:             joinedTime,
		LastProposalTime:       lastProposalTime,
		GGPBondAmount:          ggpBondAmount,
		UnbondedValidatorCount: unbondedValidatorCount,
	}, nil

}

// Get the minimum member count
func GetMinimumMemberCount(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return 0, err
	}
	minMemberCount := new(*big.Int)
	if err := gogoDAONodeTrusted.Call(opts, minMemberCount, "getMemberMinRequired"); err != nil {
		return 0, fmt.Errorf("Could not get trusted node DAO minimum member count: %w", err)
	}
	return (*minMemberCount).Uint64(), nil
}

// Get the member count
func GetMemberCount(ggp *gogopool.GoGoPool, opts *bind.CallOpts) (uint64, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return 0, err
	}
	memberCount := new(*big.Int)
	if err := gogoDAONodeTrusted.Call(opts, memberCount, "getMemberCount"); err != nil {
		return 0, fmt.Errorf("Could not get trusted node DAO member count: %w", err)
	}
	return (*memberCount).Uint64(), nil
}

// Get a member address by index
func GetMemberAt(ggp *gogopool.GoGoPool, index uint64, opts *bind.CallOpts) (common.Address, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return common.Address{}, err
	}
	memberAddress := new(common.Address)
	if err := gogoDAONodeTrusted.Call(opts, memberAddress, "getMemberAt", big.NewInt(int64(index))); err != nil {
		return common.Address{}, fmt.Errorf("Could not get trusted node DAO member %d address: %w", index, err)
	}
	return *memberAddress, nil
}

// Member details
func GetMemberExists(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (bool, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return false, err
	}
	exists := new(bool)
	if err := gogoDAONodeTrusted.Call(opts, exists, "getMemberIsValid", memberAddress); err != nil {
		return false, fmt.Errorf("Could not get trusted node DAO member %s exists status: %w", memberAddress.Hex(), err)
	}
	return *exists, nil
}
func GetMemberID(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (string, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return "", err
	}
	id := new(string)
	if err := gogoDAONodeTrusted.Call(opts, id, "getMemberID", memberAddress); err != nil {
		return "", fmt.Errorf("Could not get trusted node DAO member %s ID: %w", memberAddress.Hex(), err)
	}
	return strings.Sanitize(*id), nil
}
func GetMemberUrl(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (string, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return "", err
	}
	url := new(string)
	if err := gogoDAONodeTrusted.Call(opts, url, "getMemberUrl", memberAddress); err != nil {
		return "", fmt.Errorf("Could not get trusted node DAO member %s URL: %w", memberAddress.Hex(), err)
	}
	return strings.Sanitize(*url), nil
}
func GetMemberJoinedTime(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (uint64, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return 0, err
	}
	joinedTime := new(*big.Int)
	if err := gogoDAONodeTrusted.Call(opts, joinedTime, "getMemberJoinedTime", memberAddress); err != nil {
		return 0, fmt.Errorf("Could not get trusted node DAO member %s joined time: %w", memberAddress.Hex(), err)
	}
	return (*joinedTime).Uint64(), nil
}
func GetMemberLastProposalTime(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (uint64, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return 0, err
	}
	lastProposalTime := new(*big.Int)
	if err := gogoDAONodeTrusted.Call(opts, lastProposalTime, "getMemberLastProposalTime", memberAddress); err != nil {
		return 0, fmt.Errorf("Could not get trusted node DAO member %s last proposal time: %w", memberAddress.Hex(), err)
	}
	return (*lastProposalTime).Uint64(), nil
}
func GetMemberGGPBondAmount(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return nil, err
	}
	ggpBondAmount := new(*big.Int)
	if err := gogoDAONodeTrusted.Call(opts, ggpBondAmount, "getMemberGGPBondAmount", memberAddress); err != nil {
		return nil, fmt.Errorf("Could not get trusted node DAO member %s GGP bond amount: %w", memberAddress.Hex(), err)
	}
	return *ggpBondAmount, nil
}
func GetMemberUnbondedValidatorCount(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (uint64, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return 0, err
	}
	unbondedValidatorCount := new(*big.Int)
	if err := gogoDAONodeTrusted.Call(opts, unbondedValidatorCount, "getMemberUnbondedValidatorCount", memberAddress); err != nil {
		return 0, fmt.Errorf("Could not get trusted node DAO member %s unbonded validator count: %w", memberAddress.Hex(), err)
	}
	return (*unbondedValidatorCount).Uint64(), nil
}

// Get the time that a proposal for a member was executed at
func GetMemberInviteProposalExecutedTime(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (uint64, error) {
	return GetMemberProposalExecutedTime(ggp, "invited", memberAddress, opts)
}
func GetMemberLeaveProposalExecutedTime(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (uint64, error) {
	return GetMemberProposalExecutedTime(ggp, "leave", memberAddress, opts)
}
func GetMemberReplaceProposalExecutedTime(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (uint64, error) {
	return GetMemberProposalExecutedTime(ggp, "replace", memberAddress, opts)
}
func GetMemberProposalExecutedTime(ggp *gogopool.GoGoPool, proposalType string, memberAddress common.Address, opts *bind.CallOpts) (uint64, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return 0, err
	}
	proposalExecutedTime := new(*big.Int)
	if err := gogoDAONodeTrusted.Call(opts, proposalExecutedTime, "getMemberProposalExecutedTime", proposalType, memberAddress); err != nil {
		return 0, fmt.Errorf("Could not get trusted node DAO %s proposal executed time for member %s: %w", proposalType, memberAddress.Hex(), err)
	}
	return (*proposalExecutedTime).Uint64(), nil
}

// Get a member's replacement address if being replaced
func GetMemberReplacementAddress(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (common.Address, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return common.Address{}, err
	}
	replacementAddress := new(common.Address)
	if err := gogoDAONodeTrusted.Call(opts, replacementAddress, "getMemberReplacedAddress", "new", memberAddress); err != nil {
		return common.Address{}, fmt.Errorf("Could not get trusted node DAO member %s replacement address: %w", memberAddress.Hex(), err)
	}
	return *replacementAddress, nil
}

// Get whether a member has an active challenge against them
func GetMemberIsChallenged(ggp *gogopool.GoGoPool, memberAddress common.Address, opts *bind.CallOpts) (bool, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return false, err
	}
	isChallenged := new(bool)
	if err := gogoDAONodeTrusted.Call(opts, isChallenged, "getMemberIsChallenged", memberAddress); err != nil {
		return false, fmt.Errorf("Could not get trusted node DAO member %s is challenged status: %w", memberAddress.Hex(), err)
	}
	return *isChallenged, nil
}

// Estimate the gas of BootstrapBool
func EstimateBootstrapBoolGas(ggp *gogopool.GoGoPool, contractName, settingPath string, value bool, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrusted.GetTransactionGasInfo(opts, "bootstrapSettingBool", contractName, settingPath, value)
}

// Bootstrap a bool setting
func BootstrapBool(ggp *gogopool.GoGoPool, contractName, settingPath string, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrusted.Transact(opts, "bootstrapSettingBool", contractName, settingPath, value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap trusted node setting %s.%s: %w", contractName, settingPath, err)
	}
	return hash, nil
}

// Estimate the gas of BootstrapUint
func EstimateBootstrapUintGas(ggp *gogopool.GoGoPool, contractName, settingPath string, value *big.Int, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrusted.GetTransactionGasInfo(opts, "bootstrapSettingUint", contractName, settingPath, value)
}

// Bootstrap a uint256 setting
func BootstrapUint(ggp *gogopool.GoGoPool, contractName, settingPath string, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrusted.Transact(opts, "bootstrapSettingUint", contractName, settingPath, value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap trusted node setting %s.%s: %w", contractName, settingPath, err)
	}
	return hash, nil
}

// Estimate the gas of BootstrapMember
func EstimateBootstrapMemberGas(ggp *gogopool.GoGoPool, id, url string, nodeAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	url = strings.Sanitize(url)
	return gogoDAONodeTrusted.GetTransactionGasInfo(opts, "bootstrapMember", id, url, nodeAddress)
}

// Bootstrap a DAO member
func BootstrapMember(ggp *gogopool.GoGoPool, id, url string, nodeAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	url = strings.Sanitize(url)
	hash, err := gogoDAONodeTrusted.Transact(opts, "bootstrapMember", id, url, nodeAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap trusted node member %s: %w", id, err)
	}
	return hash, nil
}

// Estimate the gas of BootstrapUpgrade
func EstimateBootstrapUpgradeGas(ggp *gogopool.GoGoPool, upgradeType, contractName, contractAbi string, contractAddress common.Address, opts *bind.TransactOpts) (gogopool.GasInfo, error) {
	compressedAbi, err := gogopool.EncodeAbiStr(contractAbi)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return gogopool.GasInfo{}, err
	}
	return gogoDAONodeTrusted.GetTransactionGasInfo(opts, "bootstrapUpgrade", upgradeType, contractName, compressedAbi, contractAddress)
}

// Bootstrap a contract upgrade
func BootstrapUpgrade(ggp *gogopool.GoGoPool, upgradeType, contractName, contractAbi string, contractAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	compressedAbi, err := gogopool.EncodeAbiStr(contractAbi)
	if err != nil {
		return common.Hash{}, err
	}
	gogoDAONodeTrusted, err := getGoGoDAONodeTrusted(ggp)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := gogoDAONodeTrusted.Transact(opts, "bootstrapUpgrade", upgradeType, contractName, compressedAbi, contractAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap contract '%s' upgrade (%s): %w", contractName, upgradeType, err)
	}
	return hash, nil
}

// Get contracts
var gogoDAONodeTrustedLock sync.Mutex

func getGoGoDAONodeTrusted(ggp *gogopool.GoGoPool) (*gogopool.Contract, error) {
	gogoDAONodeTrustedLock.Lock()
	defer gogoDAONodeTrustedLock.Unlock()
	return ggp.GetContract("gogoDAONodeTrusted")
}
