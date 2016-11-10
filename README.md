# refineimg

Refine Image package

## How to use

```
package refineimg

func Perform(img []byte) []byte
```

- input: original webp binary image (type: []byte)
- output: refined webp binary image (type: []byte)

### Example

Refer the [test code](refineimg_test.go)

## How to develop

### modify swig and generate code

The swig command to generate wrapper code is:

```
swig -go -cgo -c++ -intgosize 64 refineimg.i
```

Then `refineimg_wrap.cxx` and `refineimg.go` are generated
