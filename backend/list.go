package backend

import (
	"fmt"
	"io/ioutil"
	"os"
)

type BackEndType string

const (
	jsonFileBackend  = "jsonfile"
	protoFileBackend = "protofile"
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
func JobReader(backends ...BackEnd) (jobs []Job, err error) {
	log.Debug("Received %d backends to read jobs from", len(backends))
	for _, back := range backends {
		switch back.Type {
		case jsonFileBackend:
			readJSONInto(&jobs, back.URL)
			//do stuff
			log.Debug("Reading JSON file")
		case protoFileBackend:
			log.Error("Protobuf backend unimplemented")
		default:
			log.Error("Backend type '%s' unknown with URL '%s'",
				back.Type, back.URL)
		}
		log.Debug("Setting up backend %s", back.URL)
	}
	return nil, nil
}

func readJSONInto(jobs *[]Job, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	src, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	_ = src
	// read []byte as JSON using protobuf
	return nil
}
