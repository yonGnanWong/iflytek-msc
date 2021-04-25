package main

import "msc-test/src"

func main() {
	Msc := src.NewMsc()
	Msc.Prepare()
	Msc.Upload()
	Msc.Merge()
	Msc.GetProgress()
	Msc.GetResult()
}
