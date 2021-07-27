[![Build Status](https://travis-ci.org/anjulapaulus/rflush.svg?branch=master)](https://travis-ci.com/anjulapaulus/rflush)
[![codecov](https://codecov.io/gh/anjulapaulus/rflush/branch/master/graph/badge.svg)](https://codecov.io/gh/anjulapaulus/rflush)
### Go H3

Go utilities for H3 library function support.

### Support
- H3 Radius Lookup

### Usage

H3 Index

```go
func TestH3Index(t *testing.T) {
	result := "87283472bffffff"
	index := H3Go.H3Index(37.3615593, -122.0553238, 7)
}
```

KRing

```go
func TestKRing(t *testing.T) {
    resolution := 8
    radiusInKM := 4
	origin := H3Go.H3Index(37.3615593, -122.0553238, resolution)
	resultArray := H3Go.KRing(origin, resolution, radiusInKM)
}
```

### Performance Benchmarks

```
Benchmark Done by HP ProBook 450 G5 8GB Ram machine

+-------------------------------+------------+-----------------------------------+
| Method                                    | Ops           | Average time ns/op |
+-------------------------------+------------+-----------------------------------+
| BenchmarkH3Index                          | 1000000       | 1154 ns/op         |
| BenchmarkKRing                            | 528979        | 2381 ns/op         |
+-------------------------------+------------+-----------------------------------+

```
