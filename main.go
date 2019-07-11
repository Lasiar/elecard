package main

import (
	"log"

	"github.com/Lasiar/elecard/base"
	"github.com/Lasiar/elecard/client"
	"github.com/Lasiar/elecard/square"
)

func main() {
	a := client.New(base.GetConfig().Key)
	a.SetDebug(base.GetConfig().Debug)
	tasks, err := a.GetTask()
	if err != nil {
		log.Fatal(err)
	}
	result := new([]square.Square)
	for _, task := range *tasks {
		if square.IsFloat(task) {
			*result = append(*result, square.CalcF(task))
			continue
		}
		*result = append(*result, square.Calc(task))
	}
	if _, err := a.CheckResult(*result); err != nil {
		log.Println(err)
	}
}
