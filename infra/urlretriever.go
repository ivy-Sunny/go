package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct {

}

func(Retriever) Get(url string) string{
	if resp, err := http.Get(url); err != nil{
		panic(err)
	}else{
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)
		return string(bytes)
	}
}

