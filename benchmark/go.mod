module benchmark

go 1.17

require (
	github.com/bombsimon/logrusr/v2 v2.0.1
	github.com/go-logr/glogr v1.1.0
	github.com/go-logr/logr v1.1.0
	github.com/go-logr/stdr v1.1.0
	github.com/go-logr/zapr v1.1.0
	github.com/hn8/zerologr v0.0.0
	github.com/rs/zerolog v1.24.0
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.19.0
)

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20210608053332-aa57babbf139 // indirect
)

replace github.com/hn8/zerologr => ./..
