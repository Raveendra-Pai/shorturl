package main

import "errors"

type MemoryMapStorage struct {
	urlMap map[string]string
}

func (m *MemoryMapStorage) Retrieve(key string) (string, error) {

	if val, ok := m.urlMap[key]; ok {
		return val, nil
	}

	return "", errors.New(key + " Not found in the map")

}

func (m *MemoryMapStorage) Insert(key string, value string) error {
	m.urlMap[key] = value
	return nil
}

func (m *MemoryMapStorage) Init(conf Config) error {
	m.urlMap = make(map[string]string)
	return nil
}
