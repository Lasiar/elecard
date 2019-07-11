package main

import (
	"log"

	"github.com/Lasiar/elecard/base"
	"github.com/Lasiar/elecard/client"
	"github.com/Lasiar/elecard/square"
)

func main() {
	api := client.New(base.GetConfig().Key)
	api.SetDebug(base.GetConfig().Debug)
	tasks, err := api.GetTask()
	if err != nil {
		log.Fatal(err)
	}
	result := new([]square.Square)
	for _, task := range *tasks {
		if square.IsFloat(task) {
			*result = append(*result, square.CalcFloat(task))
			continue
		}
		*result = append(*result, square.CalcBig(task))
	}
	if _, err := api.CheckResult(*result); err != nil {
		log.Println(err)
	}
}
