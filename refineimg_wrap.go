package refineimg

// #cgo pkg-config: opencv
import "C"

func Perform(img []byte) []byte {
	bv := NewByteVector()
	bv.Reserve(int64(len(img)))
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
