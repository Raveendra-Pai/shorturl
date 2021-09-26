package main

var Applog FileLogger

func main() {

	handler := &UrlHolder{
		urlstorage:    nil,
		configuration: Config{},
	}

	handler.Init()
	Applog.Info("Succesfully initialized shorturl handler")
	handler.Start()

}
