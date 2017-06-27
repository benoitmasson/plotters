# Pie chart

Pie chart plotter for [github.com/gonum/plot](https://github.com/gonum/plot) package.

**Install** with `go get github.com/benoitmasson/plotters/piechart`, then **import** `github.com/benoitmasson/plotters/piechart` in your Go file, and finally **create** a new chart with `piechart.NewPieChart(plotter.Values{...})`.

See [documentation](https://godoc.org/github.com/benoitmasson/plotters/piechart) for a complete overview of the functionalities.

### Examples

Such plots need to call `p.HideAxes()`, since axes are not relevant for this kind of chart.

#### Basic usage
```go
    // Initialise chart
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.HideAxes()

	// Setup pie chart
	pie, err := piechart.NewPieChart(plotter.Values{1, 2, 3, 2, 4})
	if err != nil {
		panic(err)
	}
	pie.Color = color.RGBA{255, 0, 0, 255}
	p.Add(pie)
```
![basic](https://github.com/benoitmasson/plotters/blob/master/piechart/examples/basic.png)

#### Advanced usage
```go
	// Initialise chart
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Legend.Top = true
	p.HideAxes()

	// Setup pie charts
	pie1, err := piechart.NewPieChart(plotter.Values{1, 2})
	if err != nil {
		panic(err)
	}
	pie1.Color = color.RGBA{255, 0, 0, 255}
	pie1.Total = 12
	pie1.Labels.Nominal = []string{"one", "two"}
	pie1.Labels.Values.Show = true
	pie1.Labels.Values.Percentage = true
	p.Add(pie1)
	p.Legend.Add("sample 1", pie1)

	pie2, err := piechart.NewPieChart(plotter.Values{3, 2})
	if err != nil {
		panic(err)
	}
	pie2.Color = color.RGBA{0, 255, 0, 255}
	pie2.Offset.Value = 3
	pie2.Total = 12
	pie2.Labels.Nominal = []string{"three", "four"}
	pie2.Labels.Values.Show = true
	pie2.Labels.Values.Percentage = true
	p.Add(pie2)
	p.Legend.Add("sample 2", pie2)

	pie3, err := piechart.NewPieChart(plotter.Values{4})
	if err != nil {
		panic(err)
	}
	pie3.Color = color.RGBA{0, 0, 255, 255}
	pie3.Offset.Value = 8
	pie3.Total = 12
	pie3.Offset.X = vg.Length(10)
	pie3.Offset.Y = vg.Length(-15)
	pie3.Labels.Position = 1.1
	pie3.Labels.Nominal = []string{"five"}
	pie3.Labels.Values.Show = true
	pie3.Labels.Values.Percentage = true
	p.Add(pie3)
	p.Legend.Add("sample 3", pie3)
```
![advanced](https://github.com/benoitmasson/plotters/blob/master/piechart/examples/advanced.png)
