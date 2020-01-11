[fox](https://godoc.org/github.com/gregoryv/fox) - package provides an alternate Logger design

![fox logo](doc/logo-small.png)

## Quick start

    go get github.com/gregoryv/fox

    Log := fox.NewSyncLog(os.Stdout).Log
	Log("something")

	func Test_thing(t *testing.T) {
	   thingWith.Logger := t
	   ...
	}

## Design

Based on the principle that interfaces should be kept small this package
provides a Logger interface with one func only

    Log(v ...interface{})

![design overview](doc/design_overview.svg)

The design focuses on separation between writing and formating

- SyncLog only writes messages to the output ensuring each one ends with a new line
