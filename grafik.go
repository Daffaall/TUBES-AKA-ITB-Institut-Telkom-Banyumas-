package main

import (
	"encoding/csv"
	"image/color"
	"os"
	"sort"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	rek := make(plotter.XYs, 0)
	iter := make(plotter.XYs, 0)

	for _, r := range records {
		if len(r) < 3 {
			continue
		}

		n, err1 := strconv.ParseFloat(r[0], 64)
		tr, err2 := strconv.ParseFloat(r[1], 64)
		ti, err3 := strconv.ParseFloat(r[2], 64)
		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}

		tr = tr / 1e6
		ti = ti / 1e6

		rek = append(rek, plotter.XY{X: n, Y: tr})
		iter = append(iter, plotter.XY{X: n, Y: ti})
	}

	sort.Slice(rek, func(i, j int) bool { return rek[i].X < rek[j].X })
	sort.Slice(iter, func(i, j int) bool { return iter[i].X < iter[j].X })

	p := plot.New()
	p.Title.Text = "Perbandingan Waktu DFS Rekursif dan DFS Iteratif"
	p.X.Label.Text = "Jumlah paket (n)"
	p.Y.Label.Text = "Waktu eksekusi (detik)"

	p.Add(plotter.NewGrid())

	l1, _ := plotter.NewLine(rek)
	l2, _ := plotter.NewLine(iter)

	s1, _ := plotter.NewScatter(rek)
	s2, _ := plotter.NewScatter(iter)

	blue := color.RGBA{R: 31, G: 119, B: 180, A: 255}
	orange := color.RGBA{R: 255, G: 127, B: 14, A: 255}

	l1.Color = blue
	l2.Color = orange
	s1.Color = blue
	s2.Color = orange

	s1.GlyphStyle.Shape = draw.CircleGlyph{}
	s2.GlyphStyle.Shape = draw.SquareGlyph{}
	s1.GlyphStyle.Radius = vg.Points(3)
	s2.GlyphStyle.Radius = vg.Points(3)

	l1.Width = vg.Points(1.5)
	l2.Width = vg.Points(1.5)

	p.Legend.Top = true
	p.Legend.Left = true

	p.Add(l1, s1, l2, s2)
	p.Legend.Add("DFS Rekursif", l1, s1)
	p.Legend.Add("DFS Iteratif", l2, s2)

	if err := p.Save(8*vg.Inch, 5*vg.Inch, "grafik.png"); err != nil {
		panic(err)
	}
}