package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r *Retriever) Get(url string) string {
	if resp, err := http.Get(url);err != nil{
		panic(err)
	}else{

		response, err := httputil.DumpResponse(resp, true)
		resp.Body.Close()
		if err != nil{
			panic(err)
		}
		return string(response)
	}
}