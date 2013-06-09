package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"code.google.com/p/gcfg"
)

func readConfig(loc string) (Config, error) {
	raw, err := readJSON(loc)
	if err != nil {
		fmt.Println("Oh crap, stuff is broken")
	}
	return parseStructFromJSON(raw)
}


func parseStructFromJSON(raw []byte) (Config, error) {
	var parsed_config Config
	err := json.Unmarshal(raw, &parsed_config)
	if err != nil {
		return Config{}, err
	}
	return parsed_config, err
}

func cleanConfig(loc string) ([]byte, error) {
	config, err := readConfig(loc)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	jsonout, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return nil, err
	}
	return jsonout, nil
}

func makeConfig() ([]byte)

// Returns a []byte of raw JSON with comments removed.
func readJSON(filein string) ([]byte, error) {
	file, err := ioutil.ReadFile(filein)
	if err != nil {
		return nil, err
	}

	raw, err := stripComments(file)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

// Replaces all C-style comments (prefixed with "//" and inside "/* */") with empty strings. This is necessary in parsing JSON files that contain them.
// Returns b without comments. Credit to SashaCrofter, thanks!
func stripComments(b []byte) ([]byte, error) {
	regComment, err := regexp.Compile("(?s)//.*?\n|/\\*.*?\\*/")
	if err != nil {
		return nil, err
	}
	out := regComment.ReplaceAllLiteral(b, nil)
	return out, nil
}
