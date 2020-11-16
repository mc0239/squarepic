package main

import (
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

type generationParams struct {
	help         bool
	size         int
	squaresCount int
}

func extractQueryParams(params url.Values) generationParams {
	genParams := generationParams{
		help:         false,
		size:         globalConfig.defaultSize,
		squaresCount: globalConfig.defaultSquaresCount,
	}

	if s := params.Get("help"); len(s) > 0 {
		genParams.help = true
	}

	// extract square count or use default
	if s := params.Get("squares"); len(s) > 0 {
		squaresCount, err := strconv.Atoi(s)
		if err == nil {
			genParams.squaresCount = squaresCount
		}
	}

	// extract size or use default
	if s := params.Get("size"); len(s) > 0 {
		size, err := strconv.Atoi(s)
		if err == nil {
			genParams.size = size
		}
	}

	// limit size between min and max
	if genParams.size < globalConfig.minSize {
		genParams.size = globalConfig.minSize
	}
	if genParams.size > globalConfig.maxSize {
		genParams.size = globalConfig.maxSize
	}

	// limit squares count to size
	if genParams.squaresCount > genParams.size {
		genParams.squaresCount = genParams.size
	}

	// floor size to squares count
	sizeRemain := genParams.size % genParams.squaresCount
	genParams.size = genParams.size - sizeRemain

	return genParams
}

func getGeneratedImageFilename(hashedPath int64, params generationParams) string {
	return strconv.FormatUint(uint64(hashedPath), 16) +
		"_" + strconv.Itoa(params.size) +
		"_" + strconv.Itoa(params.squaresCount) +
		".png"
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	logger.Verbose("%d for %s %s", 404, r.Method, r.URL.Path)
	w.WriteHeader(404)
}

func handleGenerateImageRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	params := extractQueryParams(r.URL.Query())

	if params.help == true {
		io.WriteString(w, helpText)
		return
	}

	// generate hashed value and a random number generator with seed from hashed
	// value
	reqPath := r.URL.Path
	hashedReqPath, err := hashInput(reqPath)

	if err != nil {
		logger.Error("Hashing failed with error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	randGen := rand.New(rand.NewSource(hashedReqPath))

	// unique part of file name is hashed path in hex
	filename := getGeneratedImageFilename(hashedReqPath, params)
	// full file name also includes parameters
	filepath := path.Join(globalConfig.imagesFolder, filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// if image doesn't exist yet, generate it
		if err := generateImage(filepath, params, randGen); err != nil {
			logger.Error("Image generation failed with error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	imageFile, err := os.Open(filepath)
	defer imageFile.Close()

	if err != nil {
		logger.Error("Error opening image file: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "image/png")

	_, err = io.Copy(w, imageFile)
	if err != nil {
		logger.Error("Error writing response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	elapsed := time.Now().Sub(start)
	logger.Verbose("Request to %s took %s", r.URL, elapsed)
}
