package main

import 	"github.com/valyala/fasthttp"

func main() {

	// pass plain function to fasthttp
	handler := &UrlHolder{
		urlList: map[string]string{},
	}
	fasthttp.ListenAndServe(":8081", handler.HandleFastHTTP)
}
