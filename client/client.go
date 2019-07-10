package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
)

const apiToken = "mPEsWPQeOoPPpXEAXw0RiszBUDRh/LHoryBDQ33LMzgxNPK49GOXbRknCBJ5M5BWrAw31QayH6om9QAn2xOF/g=="
const url = "http://contest.elecard.ru/api"

type API struct {
	Key    string        `json:"key"`
	Method string        `json:"method"`
	Params *[]SquareJSON `json:"params"`
	debug  bool
}

func New(key string) API { return API{Key: key} }

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e errorResponse) error() error {
	if e.Message == "" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("[%d] %s", e.Code, e.Message))
	}
}

type CordUint struct {
	X uint64 `json:"-"`
	Y uint64 `json:"-"`
}

type CordJson struct {
	X json.Number `json:"x"`
	Y json.Number `json:"y"`
}

type number string

func (n number) Less(str string) (bool, error) {
	if strings.Contains(string(str), ".") {

	}
	return false, nil
}

type cordBig struct {
	x Interface
	y Interface
}

func (cb cordBig) Less(bool, error) {

}

type Interface interface {
	Less(str string) (bool, error)
	String() string
}

type Circle struct {
	CordJson
	Radius json.Number `json:"radius"`
}
type Square struct {
	LeftBottom CordJson `json:"left_bottom"`
	RightTop   CordJson `json:"right_top"`
}

type SquareJSON struct {
	LeftBottom CordJson `json:"left_bottom"`
	RightTop   CordJson `json:"right_top"`
}

func (api *API) SetDebug(isDebug bool) {
	api.debug = isDebug
}

func (api API) CheckResult(result []Square) ([]bool, error) {
	api.Method = "CheckResults"
	var js []SquareJSON
	for _, r := range result {
		j := new(SquareJSON)
		j.RightTop.X = json.Number(fmt.Sprint(r.RightTop.X))
		j.RightTop.Y = json.Number(fmt.Sprint(r.RightTop.Y))
		j.LeftBottom.X = json.Number(fmt.Sprint(r.LeftBottom.X))
		j.LeftBottom.Y = json.Number(fmt.Sprint(r.LeftBottom.Y))
		js = append(js, *j)
	}

	api.Params = &js

	resp, err := api.do()
	if err != nil {
		return nil, err
	}
	response := struct {
		Result []bool        `json:"result"`
		Error  errorResponse `json:"error"`
	}{}
	if err := json.NewDecoder(resp).Decode(&response); err != nil {
		return nil, err
	}
	return response.Result, response.Error.error()
}

func (api API) GetTask() (*[][]Circle, error) {
	api.Method = "GetTasks"
	resp, err := api.do()
	if err != nil {
		return nil, err
	}

	response := struct {
		Result *[][]Circle   `json:"result"`
		Error  errorResponse `json:"error"`
	}{}

	if err := json.NewDecoder(resp).Decode(&response); err != nil {
		return nil, err
	}
	return response.Result, response.Error.error()
}

func (api *API) do() (io.Reader, error) {
	query, err := json.Marshal(&api)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", strings.NewReader(string(query)))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("non 200")
	}
	if api.debug {
		fmt.Println(string(query))
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return nil, err
		}
		log.Println(buf.String())
		return &buf, nil
	}
	return resp.Body, nil
}

func (circle *Circle) ToSquare() (square Square) {
	x := big.Int{}
	x.SetString(circle.CordJson.X.String(), 10)
	fmt.Println(x.String(), circle.CordJson.X.String())
	//square.RightTop = CordUint{X: x + uint64(circle.Radius), Y: y + uint64(circle.Radius)}
	//square.LeftBottom = CordUint{X: x - uint64(circle.Radius), Y: y - uint64(circle.Radius)}
	return square
}
