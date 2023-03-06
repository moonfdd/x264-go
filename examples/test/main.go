package main

import (
	"fmt"
	"os"

	"github.com/moonfdd/x264-go/libx264"
	"github.com/moonfdd/x264-go/libx264common"
)

func main() {
	fmt.Println(libx264.X264_POINTVER)
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	libx264common.SetLibx264Path("./lib/libx264-164.dll")
	p := new(libx264.X264ParamT)
	p.X264ParamDefault()
	p.X264ParamDefaultPreset("veryfast", "zerolatency")
}

type A struct {
	B struct{}
}
