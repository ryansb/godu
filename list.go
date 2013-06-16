package main

import (
	"fmt"
)

type BackEndType int

const (
	_               = iota
	JSONFileBackend = 1 + iota
	ProtoFileBackend
)

type BackEnd struct {
	URL  string
	Type BackEndType
}

func ListExecs() (execs []string, err error) {
	fmt.Println("listExecs unimplemented")
	return nil, nil
}

// JobReader can take a list of backends (only files currently)
func JobReader(backends []BackEnd) (jobs []Job) {
	log.Debug("")
	for _, back := range backends {
		fmt.Println(back.URL)
	}
	return make([]Job, 3)
}
