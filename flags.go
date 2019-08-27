package main

import "flag"

var flagP = flag.String("p", "", "Flag p")
var flagD = flag.Bool("d", false, "Enable debug")

func init() {
	flag.Parse()
}
