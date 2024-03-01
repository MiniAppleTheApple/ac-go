# ac-go

"ac" of ac-go stands for "Abstract class".It works just like others language's **abstract class**, but empowered by go's **code generator power**!!! Which people seems to not use that much.Code generator for me, it's like the alternative option of macro provided by go.

## Example project

You can try to build the current project executable and then move to the `./example` directory.
```
go build .
mv ac-go ./example/ac-go // if you are on unix-like system
```

And try to generate the code

```
cd example
go generate ./...
```

And you will see that `ints.go` is generated!!

```go
package main

//go:generate ./ac-go -file=./ints.go -type=Ints -template=./templates/compare.go
type Ints []int

func (xs Ints) Compare(ys Ints) bool {
	if len(xs) != len(ys) {
		return false
	}

	for i := range xs {
		if xs[i] != ys[i] {
			return false
		}
	}

	return true
}
```


