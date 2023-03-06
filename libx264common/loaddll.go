package libx264common

import (
	"sync"

	"github.com/ying32/dylib"
)

var libx264Dll *dylib.LazyDLL
var libx264DllOnce sync.Once

func GetLibx264Dll() (ans *dylib.LazyDLL) {
	libx264DllOnce.Do(func() {
		libx264Dll = dylib.NewLazyDLL(libx264Path)
	})
	ans = libx264Dll
	return
}

var libx264Path = "libx264.dll"

func SetLibx264Path(path0 string) {
	libx264Path = path0
}
