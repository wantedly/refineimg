package refineimg

// #cgo CXXFLAGS: -std=c++11
// #cgo pkg-config: opencv
// #include "image_refinement.h"
import "C"

func Perform(img []byte) []byte {
	return img
}
