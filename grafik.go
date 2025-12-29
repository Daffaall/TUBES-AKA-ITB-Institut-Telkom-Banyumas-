package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	file, _ := os.Open("data.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	rek := make(plotter.XYs, 0)
	iter := make(plotter.XYs, 0)

	for _, r := range records {
		n, _ := strconv.ParseFloat(r[0], 64)
		tr, _ := strconv.ParseFloat(r[1], 64)
		ti, _ := strconv.ParseFloat(r[2], 64)

		rek = append(rek, plotter.XY{X: n, Y: tr})
		iter = append(iter, plotter.XY{X: n, Y: ti})
	}

	p := plot.New()
	p.Title.Text = "Perbandingan DFS Rekursif vs Iteratif"
	p.X.Label.Text = "Jumlah Node"
	p.Y.Label.Text = "Waktu (µs)"

	l1, _ := plotter.NewLine(rek)
	l2, _ := plotter.NewLine(iter)
	p.Add(l1, l2)
	p.Legend.Add("Rekursif", l1)
	p.Legend.Add("Iteratif", l2)

	p.Save(8*vg.Inch, 5*vg.Inch, "grafik.png")
}