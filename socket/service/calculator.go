package service

import (
	"learngo/socket/data"
	"log"
)

type Calculator struct {
}

func (c *Calculator) Add(request *data.CalculatorRequest, respeonse *data.CalculatorRespeonse) error {
	log.Printf("[+]client call add method\n")
	respeonse.Result = request.Left + request.Right
	return nil
}
