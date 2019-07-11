package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Lasiar/elecard/base"
	"github.com/Lasiar/elecard/square"
)

// API struct for work with API
type API struct {
	Key    string           `json:"key"`
	Method string           `json:"method"`
	Params *[]square.Square `json:"params"`
	debug  bool
	logger *log.Logger
}

// New Get api
func New(key string) API { return API{Key: key, logger: log.New(os.Stderr, "[DEBUG]", 0)} }

// errorResponse handing error api
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e errorResponse) error() error {
	if e.Message == "" {
		return nil
	}
	return fmt.Errorf("[%d] %s", e.Code, e.Message)
}

// SetDebug print request and response on Stderr
func (api *API) SetDebug(isDebug bool) {
	api.debug = isDebug
}

// CheckResult sends data to check
func (api API) CheckResult(result []square.Square) ([]bool, error) {
	api.Method = "CheckResults"
	var js []square.Square
	for _, r := range result {
		j := new(square.Square)
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

// GetTask get task from server
func (api API) GetTask() (*[][]square.Circle, error) {
	api.Method = "GetTasks"
	resp, err := api.do()
	if err != nil {
		return nil, err
	}

	response := struct {
		Result *[][]square.Circle `json:"result"`
		Error  errorResponse      `json:"error"`
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
	resp, err := http.Post(base.GetConfig().URL, "application/json", strings.NewReader(string(query)))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("non 200")
	}
	if api.debug {
		api.logger.Printf("request: %s", string(query))
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return nil, err
		}
		api.logger.Println("response: $s", buf.String())
		return &buf, nil
	}
	return resp.Body, nil
}
