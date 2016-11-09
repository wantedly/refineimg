package refineimg

// #cgo pkg-config: opencv
// #include "image_refinement.h"
import "C"

func Perform(img []byte) []byte {
	return img
}
