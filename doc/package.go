package doc

import (
	"github.com/gregoryv/draw/shape/design"
	"github.com/gregoryv/fox"
)

func DesignOverview() *design.ClassDiagram {
	var (
		d    = design.NewClassDiagram()
		sync = d.Struct(fox.SyncLog{})
	)
	d.Place(sync).At(20, 20)
	return d
}
