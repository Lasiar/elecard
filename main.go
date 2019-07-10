package main

import (
	"encoding/json"
	"fmt"
	"github.com/Lasiar/elecard/base"
	"github.com/Lasiar/elecard/client"
	"log"
	"math/big"
)

type test []struct {
	Str string
}

type Interface interface {
	print()
	add()
}

func main1() {
	y := big.Int{}
	x := big.Int{}
	x.SetString("10", 10)
	y.SetString("11", 10)
	fmt.Println(new(big.Int).Add(&x, &y).String(), new(big.Int).Sub(&x, &y).String())
}
func main() {
	a := client.New(base.GetConfig().Key)
	a.SetDebug(true)
	tasks, err := a.GetTask()
	if err != nil {
		log.Println(err)
	}
	task := (*tasks)[len(*tasks)-2 ]
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
		//fmt.Printf("string:\t %s \t %s \t %s\n", circle.X.String(), circle.Y.String(), circle.Radius.String())
		rtx := *new(big.Int).Add(&x, &r)
		rty := *new(big.Int).Add(&y, &r)
		lbx := *new(big.Int).Sub(&x, &r)
		lby := *new(big.Int).Sub(&y, &r)
		fmt.Printf("start:\t %s \t %s \t %s \t %s\n", rtx.String(), rty.String(), lbx.String(), lby.String(), )
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
	fmt.Printf("end: \t %s \t %s \t %s \t %s\n", mxrtx.String(), mxrty.String(), mnlbx.String(), mnlby.String())
	response := make([]client.Square, 9, 9)
	response[7] = client.Square{
		LeftBottom: client.CordJson{X: json.Number(mnlbx.String()), Y: json.Number(mnlby.String())},
		RightTop:   client.CordJson{X: json.Number(mxrtx.String()), Y: json.Number(mxrty.String())},
	}

	if _, err := a.CheckResult(response); err != nil {
		log.Println(err)
	}

}

//func main1() {
//	a := client.New("mPEsWPQeOoPPpXEAXw0RiszBUDRh/LHoryBDQ33LMzgxNPK49GOXbRknCBJ5M5BWrAw31QayH6om9QAn2xOF/g==")
//	a.SetDebug(true)
//	tasks, err := a.GetTask()
//	if err != nil {
//		log.Println(err)
//	}
//	response := make([]client.Square, 9, 9)
//
//	task := (*tasks)[len(*tasks)-2 ]
//	//t
//	//for _, task := range *tasks {
//	minLeftBotton := client.CordUint{}
//	maxRightTop := client.CordUint{}
//	for _, circle := range task {
//		s := circle.ToSquare()
//		if minLeftBotton.Y > s.LeftBottom.Y {
//			minLeftBotton.Y = s.LeftBottom.Y
//		}
//		if minLeftBotton.X > s.LeftBottom.X {
//			minLeftBotton.X = s.LeftBottom.X
//		}
//		if maxRightTop.Y < s.RightTop.Y {
//			maxRightTop.Y = s.RightTop.Y
//		}
//		if maxRightTop.X < s.RightTop.X {
//			maxRightTop.X = s.RightTop.X
//		}
//	}
//
//	//*response = append(*response, client.Square{LeftBottom: minLeftBotton, RightTop: maxRightTop})
//	//response[7] = client.Square{LeftBottom: minLeftBotton, RightTop: maxRightTop}
//
//	if _, err := a.CheckResult(response); err != nil {
//		return
//	}
//
//	if err != nil {
//		log.Println(err)
//	}
//
//}
