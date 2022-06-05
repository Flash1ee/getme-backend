package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"
)

// Запускать из папки app!
var (
	yamlPath string
	jsonPath string
)

func init() {
	flag.StringVar(&yamlPath, "yaml-path", "./api/swagger.yaml", "path yaml file")
	flag.StringVar(&jsonPath, "json-path", "./api/swagger.json", "path to json output")

}
func main() {
	f, err := os.Open(yamlPath)
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Error(err)
		}
	}(f)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	buf := make([]byte, 1024)
	data := make([]byte, 0)
	for {
		n, err := r.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		data = append(data, buf[:n]...)
	}

	y, err := yaml.YAMLToJSON(data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	var resJson bytes.Buffer
	err = json.Indent(&resJson, y, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	res, err := os.Create(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = res.Write(resJson.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	defer func(res *os.File) {
		err := res.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res)
}
