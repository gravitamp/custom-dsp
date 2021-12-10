package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ReadSignalFile(path string, sampleRate float64) (*Signal, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	signal := Signal{
		SampleRate: sampleRate,
		Signal:     make([]float64, 0),
	}

	csvReader := csv.NewReader(file)
	csvData, _ := csvReader.ReadAll()

	//read without header
	for i := 1; i < len(csvData); i++ {
		v, _ := strconv.ParseFloat(csvData[i][6], 64)
		signal.Signal = append(signal.Signal, v)
	}

	return &signal, err
}
