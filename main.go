package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var date []float64

func main() {
	// Read a signal sampled at 31hz
	// signal1, _ := ReadSignalFile("6_26_13_tricklecharge.csv", 2)
	signal1, _ := ReadSignalFile("CS2_34_11_11_10.csv", 2)

	fmt.Println(signal1)

	//add plot
	p := plot.New()
	p.Title.Text = "Traffic Volume Forecasting"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	for i := 0; i < len(signal1.Signal); i++ {
		date = append(date, float64(i)/2)
	}

	err := plotutil.AddLinePoints(p,
		"Train", makePoints(signal1.Signal, date))
	if err != nil {
		panic(err)
	}
	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}

	// // Get a 10 second sample of the signal
	// signal10s := signal1.Sample(10 * time.Second)

	// fmt.Println(signal10s)

	// Normalize the signal between -1 and 1
	normalized, _ := signal1.Normalize()

	// Calculate the frequency spectrum of the signal (FFT + massage of the numbers)
	spectrum, _ := normalized.FrequencySpectrum()

	fmt.Println(spectrum)

	//add plot
	p2 := plot.New()
	p2.Title.Text = "Traffic Volume Forecasting"
	p2.X.Label.Text = "X"
	p2.Y.Label.Text = "Y"

	err2 := plotutil.AddLinePoints(p2,
		"Train", makePoints(spectrum.Spectrum, date))
	if err2 != nil {
		panic(err2)
	}
	// Save the plot to a PNG file.
	if err3 := p2.Save(10*vg.Inch, 4*vg.Inch, "plot2.png"); err3 != nil {
		panic(err)
	}

	// // Run some filters on the signal
	// _, _ = signal10s.LowPassFilter(3)
	// _, _ = signal10s.HighPassFilter(10)
	// _, _ = signal10s.BandPassFilter(3, 10)
}

func makePoints(data []float64, date []float64) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = date[i]
		pts[i].Y = data[i]
	}
	return pts
}
