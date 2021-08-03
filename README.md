# Zerologr

[![Go Reference](https://pkg.go.dev/badge/github.com/hn8/zerologr.svg)](https://pkg.go.dev/github.com/hn8/zerologr)
![test](https://github.com/hn8/zerologr/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/hn8/zerologr)](https://goreportcard.com/report/github.com/hn8/zerologr)

The fastest [logr](https://github.com/go-logr/logr) implementation using
[Zerolog](https://github.com/rs/zerolog).

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
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz

BenchmarkDiscardInfoOneArg-8             1247884               953.4 ns/op           176 B/op          4 allocs/op
BenchmarkDiscardInfoSeveralArgs-8         721035              1484 ns/op             336 B/op          6 allocs/op
BenchmarkDiscardV0Info-8                  770527              1486 ns/op             336 B/op          6 allocs/op
BenchmarkDiscardV9Info-8                 9779283               137.8 ns/op           176 B/op          2 allocs/op
BenchmarkDiscardError-8                   677576              1542 ns/op             336 B/op          6 allocs/op
BenchmarkDiscardWithValues-8             5563468               225.2 ns/op           192 B/op          3 allocs/op
BenchmarkDiscardWithName-8              20706103                58.52 ns/op           64 B/op          1 allocs/op
BenchmarkFuncrInfoOneArg-8               1255844               948.2 ns/op           176 B/op          4 allocs/op
BenchmarkFuncrInfoSeveralArgs-8           802514              1491 ns/op             336 B/op          6 allocs/op
BenchmarkFuncrV0Info-8                    795092              1480 ns/op             336 B/op          6 allocs/op
BenchmarkFuncrV9Info-8                   9467601               140.3 ns/op           176 B/op          2 allocs/op
BenchmarkFuncrError-8                     791719              1528 ns/op             336 B/op          6 allocs/op
BenchmarkFuncrWithValues-8               5428614               222.8 ns/op           192 B/op          3 allocs/op
BenchmarkFuncrWithName-8                18420804                62.07 ns/op           64 B/op          1 allocs/op
```

stdr: https://github.com/go-logr/stdr/pull/14

```
BenchmarkDiscardInfoOneArg-8             1319188               892.8 ns/op           232 B/op          8 allocs/op
BenchmarkDiscardInfoSeveralArgs-8         742125              1580 ns/op             552 B/op         14 allocs/op
BenchmarkDiscardV0Info-8                  755325              1621 ns/op             552 B/op         14 allocs/op
BenchmarkDiscardV9Info-8                 8628398               133.5 ns/op           176 B/op          2 allocs/op
BenchmarkDiscardError-8                   693415              1715 ns/op             600 B/op         16 allocs/op
BenchmarkDiscardWithValues-8            n 4856173              242.6 ns/op           208 B/op          3 allocs/op
BenchmarkDiscardWithName-8              12319874                85.29 ns/op           80 B/op          1 allocs/op
BenchmarkFuncrInfoOneArg-8               1332094               904.2 ns/op           232 B/op          8 allocs/op
BenchmarkFuncrInfoSeveralArgs-8           689158              1561 ns/op             552 B/op         14 allocs/op
BenchmarkFuncrV0Info-8                    769890              1629 ns/op             552 B/op         14 allocs/op
BenchmarkFuncrV9Info-8                   8916390               138.0 ns/op           176 B/op          2 allocs/op
BenchmarkFuncrError-8                     693705              1723 ns/op             600 B/op         16 allocs/op
BenchmarkFuncrWithValues-8               4718301               251.3 ns/op           208 B/op          3 allocs/op
BenchmarkFuncrWithName-8                12068954                91.43 ns/op           80 B/op          1 allocs/op
```

glogr: https://github.com/go-logr/glogr/pull/16

```
BenchmarkDiscardInfoOneArg-8              464660              2199 ns/op             400 B/op          9 allocs/op
BenchmarkDiscardInfoSeveralArgs-8         411957              2908 ns/op             664 B/op         15 allocs/op
BenchmarkDiscardV0Info-8                  382228              2879 ns/op             664 B/op         15 allocs/op
BenchmarkDiscardV9Info-8                 9592717               138.8 ns/op           176 B/op          2 allocs/op
BenchmarkDiscardError-8                   327556              3160 ns/op             696 B/op         17 allocs/op
BenchmarkDiscardWithValues-8             5553832               223.2 ns/op           192 B/op          3 allocs/op
BenchmarkDiscardWithName-8              17091950                71.42 ns/op           64 B/op          1 allocs/op
BenchmarkFuncrInfoOneArg-8                546212              2196 ns/op             400 B/op          9 allocs/op
BenchmarkFuncrInfoSeveralArgs-8           396565              2890 ns/op             664 B/op         15 allocs/op
BenchmarkFuncrV0Info-8                    421482              2872 ns/op             664 B/op         15 allocs/op
BenchmarkFuncrV9Info-8                   8809033               138.8 ns/op           176 B/op          2 allocs/op
BenchmarkFuncrError-8                     383487              3121 ns/op             696 B/op         17 allocs/op
BenchmarkFuncrWithValues-8               5243806               221.5 ns/op           192 B/op          3 allocs/op
BenchmarkFuncrWithName-8                15284324                70.80 ns/op           64 B/op          1 allocs/op
```

zapr: https://github.com/go-logr/zapr/pull/33

```
BenchmarkDiscardInfoOneArg-8              273366              4181 ns/op             713 B/op         15 allocs/op
BenchmarkDiscardInfoSeveralArgs-8         227490              5202 ns/op            1130 B/op         17 allocs/op
BenchmarkDiscardV0Info-8                  238732              5162 ns/op            1129 B/op         17 allocs/op
BenchmarkDiscardV9Info-8                 9315973               141.0 ns/op           176 B/op          2 allocs/op
BenchmarkDiscardError-8                   192272              5543 ns/op            1258 B/op         18 allocs/op
BenchmarkDiscardWithValues-8              831561              1432 ns/op            1544 B/op         10 allocs/op
BenchmarkDiscardWithName-8               5448158               201.5 ns/op           216 B/op          4 allocs/op
BenchmarkFuncrInfoOneArg-8                261658              4174 ns/op             713 B/op         15 allocs/op
BenchmarkFuncrInfoSeveralArgs-8           227588              5172 ns/op            1130 B/op         17 allocs/op
BenchmarkFuncrV0Info-8                    227376              5158 ns/op            1130 B/op         17 allocs/op
BenchmarkFuncrV9Info-8                   8277498               147.4 ns/op           176 B/op          2 allocs/op
BenchmarkFuncrError-8                     205752              5576 ns/op            1258 B/op         18 allocs/op
BenchmarkFuncrWithValues-8                939271              1392 ns/op            1544 B/op         10 allocs/op
BenchmarkFuncrWithName-8                 5572585               208.5 ns/op           216 B/op          4 allocs/op
```
