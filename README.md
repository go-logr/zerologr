# Zerologr

[![Go Reference](https://pkg.go.dev/badge/github.com/hn8/zerologr.svg)](https://pkg.go.dev/github.com/hn8/zerologr)
![test](https://github.com/hn8/zerologr/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/hn8/zerologr)](https://goreportcard.com/report/github.com/hn8/zerologr)

A [logr](https://github.com/go-logr/logr) implementation using [Zerolog](https://github.com/rs/zerolog).

## Usage

```go
import (
    "os"

    "github.com/go-logr/logr"
    "github.com/hn8/zerologr"
    "github.com/rs/zerolog"
)

func main() {
    zerologr.NameFieldName = "logger"
    zerologr.NameSeparator = "/"

    zl := zerolog.New(os.Stderr)
    var log logr.Logger = zerologr.New(&zl)

    log.Info("Logr in action!", "the answer", 42)
}
```

## Implementation Details

For the most part, concepts in Zerolog correspond directly with those in logr.

Levels in logr correspond to custom debug levels in Zerolog. Any given level
in logr is represents by `zerologLevel = 1 - logrLevel`.

For example `V(2)` is equivalent to Zerolog's `TraceLevel`, while `V(1)` is
equivalent to Zerolog's `DebugLevel`.

## Benchmark

```
go version go1.16.6 linux/amd64
cpu: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz

BenchmarkZerologrInfoOneArg-2        	      2243 ns/op	     176 B/op	       4 allocs/op
BenchmarkZerologrInfoSeveralArgs-2   	      2776 ns/op	     336 B/op	       6 allocs/op
BenchmarkZerologrV0Info-2            	      2833 ns/op	     336 B/op	       6 allocs/op
BenchmarkZerologrV9Info-2            	       124 ns/op	     176 B/op	       2 allocs/op
BenchmarkZerologrError-2             	      2837 ns/op	     336 B/op	       6 allocs/op
BenchmarkZerologrWithValues-2        	       214 ns/op	     192 B/op	       3 allocs/op
BenchmarkZerologrWithName-2          	        58 ns/op	      64 B/op	       1 allocs/op
```

[github.com/go-logr/stdr](https://github.com/go-logr/stdr)

```
BenchmarkStdrInfoOneArg-2            	      1615 ns/op	     232 B/op	       8 allocs/op
BenchmarkStdrInfoSeveralArgs-2       	      2384 ns/op	     552 B/op	      14 allocs/op
BenchmarkStdrV0Info-2                	      2402 ns/op	     552 B/op	      14 allocs/op
BenchmarkStdrV9Info-2                	       121 ns/op	     176 B/op	       2 allocs/op
BenchmarkStdrError-2                 	      2549 ns/op	     600 B/op	      16 allocs/op
BenchmarkStdrWithValues-2            	       230 ns/op	     208 B/op	       3 allocs/op
BenchmarkStdrWithName-2              	        79 ns/op	      80 B/op	       1 allocs/op
```

[github.com/go-logr/glogr](https://github.com/go-logr/glogr)

```
BenchmarkGlogrInfoOneArg-2           	      3188 ns/op	     400 B/op	       9 allocs/op
BenchmarkGlogrInfoSeveralArgs-2      	      4259 ns/op	     664 B/op	      15 allocs/op
BenchmarkGlogrV0Info-2               	      4120 ns/op	     664 B/op	      15 allocs/op
BenchmarkGlogrV9Info-2               	       129 ns/op	     176 B/op	       2 allocs/op
BenchmarkGlogrError-2                	      4699 ns/op	     696 B/op	      17 allocs/op
BenchmarkGlogrWithValues-2           	       215 ns/op	     192 B/op	       3 allocs/op
BenchmarkGlogrWithName-2             	        70 ns/op	      64 B/op	       1 allocs/op
```

[github.com/go-logr/zapr](https://github.com/go-logr/zapr)

```
BenchmarkZaprInfoOneArg-2            	      6894 ns/op	     712 B/op	      15 allocs/op
BenchmarkZaprInfoSeveralArgs-2       	      9028 ns/op	    1192 B/op	      17 allocs/op
BenchmarkZaprV0Info-2                	      8979 ns/op	    1192 B/op	      17 allocs/op
BenchmarkZaprV9Info-2                	       140 ns/op	     176 B/op	       2 allocs/op
BenchmarkZaprError-2                 	      8869 ns/op	    1320 B/op	      18 allocs/op
BenchmarkZaprWithValues-2            	      1199 ns/op	    1560 B/op	      10 allocs/op
BenchmarkZaprWithName-2              	       213 ns/op	     216 B/op	       4 allocs/op
```

[github.com/bombsimon/logrusr/v2](https://github.com/bombsimon/logrusr)

```
BenchmarkLogrusrInfoOneArg-2         	      5812 ns/op	    1752 B/op	      28 allocs/op
BenchmarkLogrusrInfoSeveralArgs-2    	      8973 ns/op	    2088 B/op	      35 allocs/op
BenchmarkLogrusrV0Info-2             	      8754 ns/op	    2088 B/op	      35 allocs/op
BenchmarkLogrusrV9Info-2             	       125 ns/op	     176 B/op	       2 allocs/op
BenchmarkLogrusrError-2              	     10566 ns/op	    2624 B/op	      41 allocs/op
BenchmarkLogrusrWithValues-2         	       875 ns/op	     960 B/op	       8 allocs/op
BenchmarkLogrusrWithName-2           	       646 ns/op	     592 B/op	       7 allocs/op
```
