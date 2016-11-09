package refineimg

// #cgo CXXFLAGS: -std=c++11
// #cgo pkg-config: opencv
import "C"

func Perform(img []byte) []byte {
	bv := NewByteVector()
	for _, b := range img {
		bv.Add(b)
	}
	outBv := Refine(bv)
	n := int(outBv.Size())
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = outBv.Get(i)
	}
	return out
}
