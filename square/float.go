package square

import (
	"encoding/json"
	"strconv"
)

// CalcFloat calc with circle have float
func CalcFloat(task []Circle) Square {
	var (
		mxrtx,
		mxrty,
		mnlbx,
		mnlby float64
	)
	for _, circle := range task {
		x, _ := strconv.ParseFloat(circle.X.String(), 64)
		y, _ := strconv.ParseFloat(circle.Y.String(), 64)
		r, _ := strconv.ParseFloat(circle.Radius.String(), 64)

		rtx := x + r
		rty := y + r
		lbx := x - r
		lby := y - r

		if mxrtx < rtx {
			mxrtx = rtx
		}
		if mxrty < rty {
			mxrty = rty
		}
		if mnlbx > lbx {
			mnlbx = lbx
		}
		if mnlby > lby {
			mnlby = lby
		}
	}
	return Square{
		LeftBottom: Cord{X: json.Number(strconv.FormatFloat(mnlbx, 'f', -1, 64)), Y: json.Number(strconv.FormatFloat(mnlby, 'f', -1, 64))},
		RightTop:   Cord{X: json.Number(strconv.FormatFloat(mxrtx, 'f', -1, 64)), Y: json.Number(strconv.FormatFloat(mxrty, 'f', -1, 64))},
	}
}
