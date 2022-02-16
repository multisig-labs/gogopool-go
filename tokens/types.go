package tokens

import "math/big"

// Token balances
type Balances struct {
    AVAX *big.Int            `json:"avax"`
    GAVAX *big.Int           `json:"gavax"`
    GGP *big.Int            `json:"ggp"`
    FixedSupplyGGP *big.Int `json:"fixedSupplyGgp"`
}

