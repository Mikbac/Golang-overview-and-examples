# Data structures

## Array

Fixed length list.

```go
var sampleArray1 [10]int = int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sampleArray2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

e.g.:

* [demoArrays.go](../samples/002-demo-app/demoArrays.go)

## Slice

Array that can grow and shrink.

```go
var sampleSlice1 []string = []string{"aaa", "bbb"}
sampleSlice2 := []string{"aaa", "bbb"}
```

e.g.:

* [demoSlices.go](../samples/002-demo-app/demoSlices.go)

## Map

```go
    values1 := map[string]int{
"aaa": 1,
"bbb": 2,
}
```

e.g.:

* [demoMap.go](../samples/002-demo-app/demoMap.go)
