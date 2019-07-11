package square

import (
	"encoding/json"
	"math/big"
)

func Calc(task []Circle) Square {
	mxrtx := big.Int{}
	mxrty := big.Int{}
	mnlbx := big.Int{}
	mnlby := big.Int{}
	for _, circle := range task {
		x := big.Int{}
		y := big.Int{}
		r := big.Int{}
		x.SetString(circle.X.String(), 10)
		y.SetString(circle.Y.String(), 10)
		r.SetString(circle.Radius.String(), 10)
		rtx := *new(big.Int).Add(&x, &r)
		rty := *new(big.Int).Add(&y, &r)
		lbx := *new(big.Int).Sub(&x, &r)
		lby := *new(big.Int).Sub(&y, &r)
		if cmp := mxrtx.Cmp(&rtx); cmp == -1 {
			mxrtx = rtx
		}
		if cmp := mxrty.Cmp(&rty); cmp == -1 {
			mxrty = rty
		}
		if cmp := mnlbx.Cmp(&lbx); cmp == 1 {
			mnlbx = lbx
		}
		if cmp := mnlby.Cmp(&lby); cmp == 1 {
			mnlby = lby
		}
	}
	return Square{
		LeftBottom: Cord{X: json.Number(mnlbx.String()), Y: json.Number(mnlby.String())},
		RightTop:   Cord{X: json.Number(mxrtx.String()), Y: json.Number(mxrty.String())},
	}
}
