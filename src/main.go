package main

var Applog FileLogger

func main() {

	// pass plain function to fasthttp
	handler := &UrlHolder{
		urlstorage: nil,
		configuration: Config{},
	}

	handler.Init()
	Applog.Info("Succesfully initialized shorturl handler")
	handler.Start()

	//	fasthttp.ListenAndServe(":8090", handler.HandleFastHTTP)

}
