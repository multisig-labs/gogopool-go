package eth

import (
	"math/big"
	"testing"

	"github.com/multisig-labs/gogopool-go/utils/avax"
)

func TestConversion(t *testing.T) {

	// Equivalent unit amounts
	weiAmount := new(big.Int)
	weiAmount.SetString("999999999999999000000", 0)
	var gweiAmount float64 = 999999999999.999000000
	var ethAmount float64 = 999.999999999999000000

	// Convert wei to eth
	if toEthAmount := avax.WeiToEth(weiAmount); toEthAmount != ethAmount {
		t.Errorf("Incorrect eth amount %f", toEthAmount)
	}

	// Convert eth to wei
	if toWeiAmount := avax.EthToWei(ethAmount); toWeiAmount.Cmp(weiAmount) != 0 {
		t.Errorf("Incorrect wei amount %s", toWeiAmount.String())
	}

	// Convert wei to gigawei
	if toGweiAmount := avax.WeiToGwei(weiAmount); toGweiAmount != gweiAmount {
		t.Errorf("Incorrect gwei amount %f", toGweiAmount)
	}

	// Convert eth to gwei
	if toWeiAmount := avax.GweiToWei(gweiAmount); toWeiAmount.Cmp(weiAmount) != 0 {
		t.Errorf("Incorrect wei amount %s", toWeiAmount.String())
	}

}
