package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"

	"getme-backend/cmd/utilits/models"
	"getme-backend/internal"
)

var (
	configPath          string
	logLevel            string
	needFile            string
	allFiles            bool
	SearchURL           string
	useServerRepository bool
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
	flag.StringVar(&logLevel, "level", "trace", "skip levels")
	flag.StringVar(&needFile, "name-file", "", "concrete files to print")
	flag.BoolVar(&allFiles, "all", false, "print all logs")
	flag.StringVar(&SearchURL, "search-url", "", "search url")
	flag.BoolVar(&useServerRepository, "server-run", false, "true if it server run, false if it local run")
}

func printLogFromFile(logger *logrus.Logger, fileName string, fileTime time.Time) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer func() {
		if err = file.Close(); err != nil {
			logger.Fatal(err)
		}
	}()

	tmp := time.Now().In(time.UTC)
	diff := tmp.Sub(fileTime)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		lg := &models.Log{}
		err = easyjson.Unmarshal(bytes, lg)
		if err != nil {
			return err
		}

		if SearchURL != "" && SearchURL != lg.Url.String() {
			continue
		}

		level, err := logrus.ParseLevel(lg.Level)
		if err != nil {
			return err
		}

		logger.WithTime(lg.Time.In(time.Now().Location()).Add(diff)).WithFields(logrus.Fields{
			"urls":        lg.Url.String(),
			"method":      lg.Method,
			"remote_addr": lg.Adr,
			"work_time":   lg.WorkTime,
			"req_id":      lg.ReqID,
		}).Log(level, lg.Msg)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getTimeStringFromName(fileName string) string {
	if pos := strings.LastIndex(fileName, "__"); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}

func parseTimeFromFileName(fileName string) (time.Time, error) {
	formatTime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", 2006, 1, 2, 15, 04, 05)
	tmp, err := time.Parse(formatTime, getTimeStringFromName(fileName))
	if err != nil {
		return time.Now(), err
	}
	return tmp, err
}

func main() {
	flag.Parse()

	config := &internal.Config{}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}

	logger := logrus.New()
	logger.SetLevel(level)

	if needFile != "" {
		tmp, err := parseTimeFromFileName(needFile)
		if err != nil {
			log.Printf("error in file %v", err)
		}
		err = printLogFromFile(logger, config.LogAddr+needFile, tmp)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	files, err := os.ReadDir(config.LogAddr)
	if err != nil {
		log.Fatal(err)
	}

	if allFiles {
		for _, file := range files {
			fmt.Printf("Log from : %s\n", file.Name())
			tmp, err := parseTimeFromFileName(file.Name())
			if err != nil {
				log.Printf("error in file %v", err)
			}

			err = printLogFromFile(logger, config.LogAddr+file.Name(), tmp)
			if err != nil {
				log.Printf("error in file %v", err)
			}
		}
		return
	}

	var lastestFile string
	var lastestTime time.Time
	first := true
	for _, file := range files {
		tmp, err := parseTimeFromFileName(file.Name())
		if err == nil && (lastestTime.Before(tmp) || first) {
			lastestTime = tmp
			lastestFile = file.Name()
			first = false
		}
	}

	fmt.Printf("Log from : %s\n", lastestFile)
	tmp, err := parseTimeFromFileName(lastestFile)
	if err != nil {
		log.Printf("error in file %v", err)
	}
	err = printLogFromFile(logger, config.LogAddr+lastestFile, tmp)
	if err != nil {
		log.Printf("error in file %v", err)
	}
}
