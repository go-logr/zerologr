module benchmark

go 1.17

require (
	github.com/bombsimon/logrusr/v2 v2.0.1
	github.com/go-logr/glogr v1.1.0
	github.com/go-logr/logr v1.1.0
	github.com/go-logr/stdr v1.1.0
	github.com/go-logr/zapr v1.1.0
	github.com/hn8/zerologr v1.0.1
	github.com/rs/zerolog v1.24.0
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.19.0
)

replace github.com/hn8/zerologr => ./..
