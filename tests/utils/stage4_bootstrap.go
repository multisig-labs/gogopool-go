package utils

import (
	"time"

	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/settings/protocol"
	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
	"github.com/multisig-labs/gogopool-go/utils/eth"
)

// Bootstrap all of the parameters to mimic Stage 4 so the unit tests work correctly
func Stage4Bootstrap(ggp *gogopool.GoGoPool, ownerAccount *accounts.Account) {

	opts := ownerAccount.GetTransactor()

	protocol.BootstrapDepositEnabled(ggp, true, opts)
	protocol.BootstrapAssignDepositsEnabled(ggp, true, opts)
	protocol.BootstrapMaximumDepositPoolSize(ggp, avax.EthToWei(1000), opts)
	protocol.BootstrapNodeRegistrationEnabled(ggp, true, opts)
	protocol.BootstrapNodeDepositEnabled(ggp, true, opts)
	protocol.BootstrapMinipoolSubmitWithdrawableEnabled(ggp, true, opts)
	protocol.BootstrapMinimumNodeFee(ggp, 0.05, opts)
	protocol.BootstrapTargetNodeFee(ggp, 0.1, opts)
	protocol.BootstrapMaximumNodeFee(ggp, 0.2, opts)
	protocol.BootstrapNodeFeeDemandRange(ggp, avax.EthToWei(1000), opts)
	protocol.BootstrapInflationStartTime(ggp,
		uint64(time.Now().Unix()+(60*60*24*14)), opts)

}
