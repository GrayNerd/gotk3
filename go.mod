module github.com/graynerd/gtk

go 1.13

replace github.com/gotk3/gotk3 => /home/ken/gostuff/src/github.com/graynerd/gotk3

replace github.com/gotk3/gotk3/gtk => /home/ken/gostuff/src/github.com/graynerd/gotk3/gtk

replace github.com/gotk3/gotk3/gdk => /home/ken/gostuff/src/github.com/graynerd/gotk3/gdk

replace github.com/gotk3/gotk3/glib => /home/ken/gostuff/src/github.com/graynerd/gotk3/glib

replace github.com/gotk3/gotk3/cairo => /home/ken/gostuff/src/github.com/graynerd/gotk3/cairo

require (
	github.com/gotk3/gotk3 v0.0.0-20191027191019-60cba67d4ea4
	github.com/gotk3/gotk3/cairo v0.0.0-00010101000000-000000000000
	github.com/gotk3/gotk3/gdk v0.0.0-00010101000000-000000000000
	github.com/gotk3/gotk3/glib v0.0.0-00010101000000-000000000000
)
