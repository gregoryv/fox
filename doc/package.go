package doc

import (
	"io"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
	"github.com/gregoryv/fox"
)

func DesignOverview() *design.ClassDiagram {
	var (
		d    = design.NewClassDiagram()
		sync = d.Struct(fox.SyncLog{})
		filt = d.Struct(fox.FilterEmpty{})
		w    = d.Interface((*io.Writer)(nil))
		log  = d.Interface((*fox.Logger)(nil))
	)

	d.Place(log).At(20, 20)
	d.Place(sync).Below(log)
	d.Place(filt).RightOf(sync, 70)
	d.VAlignCenter(log, sync, w)
	d.HideRealizations()
	shape.Move(log, 80, 0)
	d.Place(w).Below(sync)

	lnk := d.Link(sync, w)
	lnk.Head = nil

	d.HAlignCenter(sync, filt)
	d.SetCaption("1. Design of fox package and it's relation to io.Writer")
	return d
}
