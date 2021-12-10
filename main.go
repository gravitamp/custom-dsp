package main

import (
	"fmt"
)

func main() {
	// Read a signal sampled at 31hz
	signal1, _ := ReadSignalFile("CS2_34_12_08_10.csv", 2)
	fmt.Println(signal1)

	// // Read a signal sampled at 100Hz
	// signal2, _ := ReadSignalFile("example_signal_100_hz.txt", 100)
	// fmt.Println(signal2)

	// // Get a 10 second sample of the signal
	// signal10s := signal1.Sample(10 * time.Second)

	// fmt.Println(signal10s)

	// // Normalize the signal between -1 and 1
	// normalized, _ := signal1.Normalize()

	// // Calculate the frequency spectrum of the signal (FFT + massage of the numbers)
	// spectrum, _ := normalized.FrequencySpectrum()

	// fmt.Println(spectrum)

	// // Run some filters on the signal
	// _, _ = signal10s.LowPassFilter(3)
	// _, _ = signal10s.HighPassFilter(10)
	// _, _ = signal10s.BandPassFilter(3, 10)
}
