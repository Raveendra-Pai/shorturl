package main

func main() {

	// pass plain function to fasthttp
	handler := &UrlHolder{
		urlList:       map[string]string{},
		configuration: Config{},
	}

	if nil == handler.Initialize() {
		handler.Start()
	}

	//	fasthttp.ListenAndServe(":8090", handler.HandleFastHTTP)

}
