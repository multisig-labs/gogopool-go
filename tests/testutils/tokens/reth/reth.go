package tokens

import (
	"math/big"

	"github.com/multisig-labs/gogopool-go/deposit"
	"github.com/multisig-labs/gogopool-go/gogopool"
	"github.com/multisig-labs/gogopool-go/tokens"

	"github.com/multisig-labs/gogopool-go/tests/testutils/accounts"
)

// Mint an amount of rETH to an account
func MintRETH(ggp *gogopool.GoGoPool, toAccount *accounts.Account, amount *big.Int) error {

	// Get ETH value of amount
	ethValue, err := tokens.GetETHValueOfRETH(ggp, amount, nil)
	if err != nil {
		return err
	}

	// Deposit from account to mint rETH
	opts := toAccount.GetTransactor()
	opts.Value = ethValue
	_, err = deposit.Deposit(ggp, opts)
	return err

}
