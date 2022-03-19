package iris

import (
	"go4ml.xyz/base/fu/lazy"
	"go4ml.xyz/base/model"
	"go4ml.xyz/base/tables"
	"go4ml.xyz/base/tables/csv"
	"go4ml.xyz/iokit"
)

func source(x string) iokit.Input {
	const base = "https://datahub.io/machine-learning/iris/r/"
	return iokit.Url(base+x, iokit.Cache("go-ml/dataset/iris/"+x))
}

var dataset = source("iris.csv")
var Features = []string{"Feature1", "Feature2", "Feature3", "Feature4"}

var Data tables.Lazy = func() lazy.Stream {
	var cls = tables.Enumset{}
	return csv.Source(dataset,
		csv.Float32("sepallength").As("Feature1"),
		csv.Float32("sepalwidth").As("Feature2"),
		csv.Float32("petallength").As("Feature3"),
		csv.Float32("petalwidth").As("Feature4"),
		csv.Meta(cls.Integer(), "class").As(model.LabelCol))()
}
