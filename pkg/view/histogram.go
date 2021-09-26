package view

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Histogram struct {
	Plot   *plot.Plot
	Values []float64
}

func NewHistogram(title string, values []float64, normalize bool) (*Histogram, error) {
	p := plot.New()
	p.Title.Text = title
	v := make(plotter.Values, len(values))
	for i := range v {
		v[i] = values[i]
	}
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		return nil, err
	}
	if normalize {
		h.Normalize(1)
	}
	p.Add(h)

	return &Histogram{
		Plot:   p,
		Values: values,
	}, nil
}

func (h *Histogram) SaveImage(path string) error {
	return h.Plot.Save(4*vg.Inch, 4*vg.Inch, path)
}
