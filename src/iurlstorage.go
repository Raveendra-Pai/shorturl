package main

type IUrlStorage interface {
	Retrieve(key string) (string, error)
	Insert(key string, value string) error
	Init() error
}


