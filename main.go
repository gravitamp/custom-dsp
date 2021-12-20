package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var date []float64

func main() {
	f, _ := os.OpenFile("fft.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	w := csv.NewWriter(f)
	// Read a signal
	signal1, _ := ReadSignalFile("11_05_2015_SP20-2_DST_80SOC.csv", 1)

	fmt.Println(signal1)

	//add plot
	p := plot.New()
	p.Title.Text = "Traffic Volume Forecasting"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// for i := 0; i < len(signal1.Signal); i++ {
	// 	date = append(date, float64(i)/2)
	// }

	err := plotutil.AddLinePoints(p,
		"Train", makePoints(signal1.Signal, date))
	if err != nil {
		panic(err)
	}
	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "plot1.png"); err != nil {
		panic(err)
	}

	// Normalize the signal between -1 and 1
	normalized, _ := signal1.Normalize()

	// Calculate the frequency spectrum of the signal (FFT + massage of the numbers)
	spectrum, _ := normalized.FrequencySpectrum()

	fmt.Println(spectrum)

	//add plot
	p2 := plot.New()
	p2.Title.Text = "Traffic Volume Forecasting"
	p2.X.Label.Text = "Frequency"
	p2.Y.Label.Text = "Spectrum"

	err2 := plotutil.AddLinePoints(p2,
		"Train", makePoints(spectrum.Spectrum, spectrum.Frequencies))
	if err2 != nil {
		panic(err2)
	}
	// Save the plot to a PNG file.
	if err3 := p2.Save(10*vg.Inch, 4*vg.Inch, "plot2.png"); err3 != nil {
		panic(err)
	}
	fmt.Println(len(spectrum.Frequencies), len(signal1.Signal))
	// for j := 0; j < len(spectrum.Frequencies); j++ {
	// 	fmt.Println("spectrum :", spectrum.Spectrum[j], "-", "frequency", spectrum.Frequencies[j])
	// }
	for j := 0; j < len(spectrum.Frequencies); j++ {
		w.Write([]string{"1", strconv.FormatFloat(spectrum.Spectrum[j], 'E', -1, 64),
			strconv.FormatFloat(spectrum.Frequencies[j], 'E', -1, 64)})
		w.Flush()
	}
}

func makePoints(data []float64, date []float64) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = date[i]
		pts[i].Y = data[i]
	}
	fmt.Println(len(pts)) //sesuai frekuensi
	return pts
}
