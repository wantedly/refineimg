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

### Explicitly generate swig code

Usually, `go build` internally call swig.
You can see this process by calling `go build -x`.

The swig command to generate wrapper code is:

```
swig -go -cgo -c++ -intgosize 64 refineimg.swigcxx
```
