agenda
======

[![Build Status](https://drone.io/github.com/ttysteale/agenda/status.png)](https://drone.io/github.com/ttysteale/agenda/latest)

A simple todo (task) list built in Go and GTK (by way of the Gotk3 bindings: http://github.com/conformal/gotk3)

## Installation
agenda depends on GTK+3.6.8 or greater.

1. Install Go (http://golang.org/doc/install)
2. Get the project: `$ go get github.com/ttysteale/agenda` or `$ git clone` it
3. Install it
  1. If you have GTK >= 3.10 installed: `$ go install github.com/ttysteale/agenda`
  2. If you have GTK >= 3.8 < 3.10 installed: `$ go install -tags=gtk_3_8 github.com/ttysteale/agenda`
  3. If you have GTK >= 3.6.8 < 3.8 installed: `$ go install -tags=gtk_3_6 github.com/ttysteale/agenda`
4. Run it: `$ agenda`
