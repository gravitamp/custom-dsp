package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
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
	for i := 1; i < len(csvData); i++ { //len(csvData)
		//parse time
		layouts := "01/02/2006 15:04:05"
		t, err := time.Parse(layouts, csvData[i][2])
		if err != nil {
			fmt.Println(err)
		}
		v, _ := strconv.ParseFloat(csvData[i][6], 64)
		signal.Signal = append(signal.Signal, v)
		date = append(date, float64(t.Unix()))
	}

	return &signal, err
}

func ReadSignalFile2(path string, sampleRate float64) (*Signal, error) {
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
	for i := 1; i < 7200; i++ { //len(csvData)
		//parse time
		layouts := "01/02/2006 15:04:05"
		t, err := time.Parse(layouts, csvData[i][2])
		if err != nil {
			fmt.Println(err)
		}
		v, _ := strconv.ParseFloat(csvData[i][6], 64)
		signal.Signal = append(signal.Signal, v)
		date = append(date, float64(t.Unix()))
	}

	return &signal, err
}
