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
func ReadJobs(backends ...BackEnd) (jobs []Job, err error) {
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

func MarshalJobs(b *BackEnd, jobs *[]Job) error {
	if b.Type == protoFileBackend {
		buf := make([]byte, 1)
		for _, j := range *jobs {
			buf = append(buf, []byte(j.Msg.String())...)
			buf = append(buf, '\n')
		}
		file_err := ioutil.WriteFile(b.URL, buf, 0750)
		if file_err != nil {
			return file_err
		}
	}
	return nil
}

func (job *Job) Persist(backend, backend_type string) error {
	log.Debug("Persisting jobs")
	b := BackEnd{}
	b.URL = backend
	switch backend_type {
	case jsonFileBackend:
		log.Debug("JSON backend")
		b.Type = jsonFileBackend
	case protoFileBackend:
		log.Debug("Protobuf backend")
		b.Type = protoFileBackend
	default:
		log.Error("backend_type '%s' invalid", backend_type)
		return fmt.Errorf("backend_type invalid")
	}
	all_jobs, err := ReadJobs(b)
	if err != nil {
		log.Warning("Error reading existing jobs.")
		return err
	}
	jobs := append(all_jobs, *job)
	MarshalJobs(&b, &jobs)

	return nil
}
