package doc

import (
	"io"

	"github.com/gregoryv/draw/shape/design"
	"github.com/gregoryv/fox"
)

func DesignOverview() *design.ClassDiagram {
	var (
		d    = design.NewClassDiagram()
		sync = d.Struct(fox.SyncLog{})
		filt = d.Struct(fox.FilterEmpty{})
		w    = d.Interface((*io.Writer)(nil))
	)

	d.Place(sync).At(20, 20)
	d.Place(filt).RightOf(sync, 70)
	d.HideRealizations()

	d.Place(w).Below(sync)

	lnk := d.Link(sync, w)
	lnk.Head = nil

	d.HAlignCenter(sync, filt)
	d.SetCaption("1. Design of fox package and it's relation to io.Writer")
	return d
}
