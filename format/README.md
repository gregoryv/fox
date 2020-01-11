[fox/format](https://godoc.org/github.com/gregoryv/fox/format) - package provides basic formating funcs

## Quick start

    go get github.com/gregoryv/fox/format

and use it with

    import . "github.com/gregoryv/fox/format"

    Log := fox.NewSyncLog(os.Stdout).Log
	Log("something")
	Log(Debug("will show dir/file.go:fileno: ..."))
	Log(Debugf("also but with %s", "formated value"))
