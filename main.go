package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mc0239/logm"
)

const configFilePath = "squares.config"

var logger = logm.New("root")

var globalConfig = struct {
	port                int
	imagesFolder        string
	defaultSquaresCount int
	minSize             int
	maxSize             int
	defaultSize         int
}{9001, "images/", 5, 5, 5000, 250}

func initImagesFolder() {
	imagesFolder := globalConfig.imagesFolder

	if _, err := os.Stat(imagesFolder); os.IsNotExist(err) {
		logger.Info("Creating folder '%s'...", imagesFolder)
		if err := os.MkdirAll(imagesFolder, os.ModeDir); err != nil {
			logger.Error("Failed to create dir '%s'", imagesFolder)
			log.Fatal(err)
		} else {
			logger.Info("Created folder '%s'", imagesFolder)
		}
	} else {
		logger.Info("Folder '%s' already exists", imagesFolder)
		logger.Warning("Existing images will not be regenerated, delete them if you want to regenerate images with current config")
	}
}

func initConfigFile() {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		logger.Info("Creating config file '%s'...", configFilePath)

		configFile, err := os.Create(configFilePath)
		defer configFile.Close()
		if err != nil {
			logger.Error("Failed to create file '%s'", configFilePath)
			log.Fatal(err)
		}

		configFile.WriteString("port=9001\nimages_folder=images/\ndefault_squares_count=5\nmin_size=5\nmax_size=5000\ndefault_size=250")
		logger.Log(logm.LvlNotice, "Config file '%s' created with default values", configFilePath)
		logger.Log(logm.LvlNotice, "Edit config file if needed and rerun program")
		os.Exit(0)
	} else {
		readGlobalConfig()
	}
}

func readGlobalConfig() {
	logger.Info("Reading config file '%s'...", configFilePath)
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		logger.Error("Failed to read file '%s'", configFilePath)
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}
		keyval := strings.Split(line, "=")
		if len(keyval) != 2 {
			logger.Warning("Invalid config line %s", line)
			continue
		}

		switch keyval[0] {
		case "images_folder":
			globalConfig.imagesFolder = strings.TrimSpace(keyval[1])
		case "min_size":
			globalConfig.minSize, err = strconv.Atoi(keyval[1])
			if err != nil {
				logger.Warning("Invalid config value %s", keyval[1])
				globalConfig.minSize = 5
			}
		case "max_size":
			globalConfig.maxSize, err = strconv.Atoi(keyval[1])
			if err != nil {
				logger.Warning("Invalid config value %s", keyval[1])
				globalConfig.maxSize = 5000
			}
		case "default_squares_count":
			globalConfig.defaultSquaresCount, err = strconv.Atoi(keyval[1])
			if err != nil {
				logger.Warning("Invalid config value %s", keyval[1])
				globalConfig.defaultSquaresCount = 5
			}
		case "default_size":
			globalConfig.defaultSize, err = strconv.Atoi(keyval[1])
			if err != nil {
				logger.Warning("Invalid config value %s", keyval[1])
				globalConfig.defaultSize = 250
			}
		case "port":
			globalConfig.port, err = strconv.Atoi(keyval[1])
			if err != nil {
				logger.Warning("Invalid config value %s", keyval[1])
				globalConfig.port = 9001
			}
		default:
			logger.Warning("Invalid config line %s", line)
		}
	}
}

func initRequestHandler() {
	// handle everything, but only make valid response to GET requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGenerateImageRequest(w, r)
		} else {
			handleNotFound(w, r)
		}
	})
}

func main() {
	initConfigFile()
	initImagesFolder()
	initRequestHandler()

	logger.Info("Started serving on port %d", globalConfig.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", globalConfig.port), nil))
}
