package converter

import "math/big"

const DEFAULT = 1000000000000000000

type MinterConverter struct {
}

func ConvertValue(num *big.Int, to string) *big.Int {

	if to == "pip" {
		return big.NewInt(0).Mul(big.NewInt(DEFAULT), num)
	}

	if to == "bip" {
		return big.NewInt(0).Div(num, big.NewInt(DEFAULT))
	}

	return nil
}
