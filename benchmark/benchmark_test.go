/*
Copyright 2021 The logr Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logr_test

import (
	"flag"
	"fmt"
	stdlog "log"
	"os"
	"testing"

	"github.com/bombsimon/logrusr/v2"
	"github.com/go-logr/glogr"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	"github.com/go-logr/stdr"
	"github.com/go-logr/zapr"
	"github.com/hn8/zerologr"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func init() {
	os.Stderr, _ = os.Open("/dev/null")

	// stdr
	stdr.SetVerbosity(1)
	// globr
	flag.Set("v", "1")
	flag.Set("logtostderr", "true")
	// zerologr
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

//go:noinline
func doInfoOneArg(b *testing.B, log logr.Logger) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Info("this is", "a", "string")
	}
}

//go:noinline
func doInfoSeveralArgs(b *testing.B, log logr.Logger) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doV0Info(b *testing.B, log logr.Logger) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.V(0).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doV9Info(b *testing.B, log logr.Logger) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.V(9).Info("multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doError(b *testing.B, log logr.Logger) {
	b.ReportAllocs()
	err := fmt.Errorf("error message")
	for i := 0; i < b.N; i++ {
		log.Error(err, "multi",
			"bool", true, "string", "str", "int", 42,
			"float", 3.14, "struct", struct{ X, Y int }{93, 76})
	}
}

//go:noinline
func doWithValues(b *testing.B, log logr.Logger) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := log.WithValues("k1", "v1", "k2", "v2")
		_ = l
	}
}

//go:noinline
func doWithName(b *testing.B, log logr.Logger) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := log.WithName("name")
		_ = l
	}
}

//go:noinline
func doWithCallDepth(b *testing.B, log logr.Logger) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := log.WithCallDepth(1)
		_ = l
	}
}

// discard

func BenchmarkDiscardInfoOneArg(b *testing.B) {
	var log logr.Logger = logr.Discard()
	doInfoOneArg(b, log)
}

func BenchmarkDiscardInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = logr.Discard()
	doInfoSeveralArgs(b, log)
}

func BenchmarkDiscardV0Info(b *testing.B) {
	var log logr.Logger = logr.Discard()
	doV0Info(b, log)
}

func BenchmarkDiscardV9Info(b *testing.B) {
	var log logr.Logger = logr.Discard()
	doV9Info(b, log)
}

func BenchmarkDiscardError(b *testing.B) {
	var log logr.Logger = logr.Discard()
	doError(b, log)
}

func BenchmarkDiscardWithValues(b *testing.B) {
	var log logr.Logger = logr.Discard()
	doWithValues(b, log)
}

func BenchmarkDiscardWithName(b *testing.B) {
	var log logr.Logger = logr.Discard()
	doWithName(b, log)
}

func BenchmarkDiscardWithCallDepth(b *testing.B) {
	var log logr.Logger = logr.Discard()
	doWithCallDepth(b, log)
}

// funcr

func noop(prefix, args string) {
}

func funcrLogger() logr.Logger {
	return funcr.New(noop, funcr.Options{})
}

func BenchmarkFuncrInfoOneArg(b *testing.B) {
	var log logr.Logger = funcrLogger()
	doInfoOneArg(b, log)
}

func BenchmarkFuncrInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = funcrLogger()
	doInfoSeveralArgs(b, log)
}

func BenchmarkFuncrV0Info(b *testing.B) {
	var log logr.Logger = funcrLogger()
	doV0Info(b, log)
}

func BenchmarkFuncrV9Info(b *testing.B) {
	var log logr.Logger = funcrLogger()
	doV9Info(b, log)
}

func BenchmarkFuncrError(b *testing.B) {
	var log logr.Logger = funcrLogger()
	doError(b, log)
}

func BenchmarkFuncrWithValues(b *testing.B) {
	var log logr.Logger = funcrLogger()
	doWithValues(b, log)
}

func BenchmarkFuncrWithName(b *testing.B) {
	var log logr.Logger = funcrLogger()
	doWithName(b, log)
}

func BenchmarkFuncrWithCallDepth(b *testing.B) {
	var log logr.Logger = funcrLogger()
	doWithCallDepth(b, log)
}

// stdr

func stdrLogger() logr.Logger {
	return stdr.New(stdlog.New(os.Stderr, "", 0))
}

func BenchmarkStdrInfoOneArg(b *testing.B) {
	var log logr.Logger = stdrLogger()
	doInfoOneArg(b, log)
}

func BenchmarkStdrInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = stdrLogger()
	doInfoSeveralArgs(b, log)
}

func BenchmarkStdrV0Info(b *testing.B) {
	var log logr.Logger = stdrLogger()
	doV0Info(b, log)
}

func BenchmarkStdrV9Info(b *testing.B) {
	var log logr.Logger = stdrLogger()
	doV9Info(b, log)
}

func BenchmarkStdrError(b *testing.B) {
	var log logr.Logger = stdrLogger()
	doError(b, log)
}

func BenchmarkStdrWithValues(b *testing.B) {
	var log logr.Logger = stdrLogger()
	doWithValues(b, log)
}

func BenchmarkStdrWithName(b *testing.B) {
	var log logr.Logger = stdrLogger()
	doWithName(b, log)
}

func BenchmarkStdrWithCallDepth(b *testing.B) {
	var log logr.Logger = stdrLogger()
	doWithCallDepth(b, log)
}

// glogr

func BenchmarkGlogrInfoOneArg(b *testing.B) {
	var log logr.Logger = glogr.New()
	doInfoOneArg(b, log)
}

func BenchmarkGlogrInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = glogr.New()
	doInfoSeveralArgs(b, log)
}

func BenchmarkGlogrV0Info(b *testing.B) {
	var log logr.Logger = glogr.New()
	doV0Info(b, log)
}

func BenchmarkGlogrV9Info(b *testing.B) {
	var log logr.Logger = glogr.New()
	doV9Info(b, log)
}

func BenchmarkGlogrError(b *testing.B) {
	var log logr.Logger = glogr.New()
	doError(b, log)
}

func BenchmarkGlogrWithValues(b *testing.B) {
	var log logr.Logger = glogr.New()
	doWithValues(b, log)
}

func BenchmarkGlogrWithName(b *testing.B) {
	var log logr.Logger = glogr.New()
	doWithName(b, log)
}

func BenchmarkGlogrWithCallDepth(b *testing.B) {
	var log logr.Logger = glogr.New()
	doWithCallDepth(b, log)
}

// zapr

func zaprLogger() logr.Logger {
	zc := zap.NewProductionConfig()
	zc.Sampling = nil
	zc.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zc.DisableStacktrace = true
	z, _ := zc.Build()
	return zapr.NewLogger(z)
}

func BenchmarkZaprInfoOneArg(b *testing.B) {
	var log logr.Logger = zaprLogger()
	doInfoOneArg(b, log)
}

func BenchmarkZaprInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = zaprLogger()
	doInfoSeveralArgs(b, log)
}

func BenchmarkZaprV0Info(b *testing.B) {
	var log logr.Logger = zaprLogger()
	doV0Info(b, log)
}

func BenchmarkZaprV9Info(b *testing.B) {
	var log logr.Logger = zaprLogger()
	doV9Info(b, log)
}

func BenchmarkZaprError(b *testing.B) {
	var log logr.Logger = zaprLogger()
	doError(b, log)
}

func BenchmarkZaprWithValues(b *testing.B) {
	var log logr.Logger = zaprLogger()
	doWithValues(b, log)
}

func BenchmarkZaprWithName(b *testing.B) {
	var log logr.Logger = zaprLogger()
	doWithName(b, log)
}

func BenchmarkZaprWithCallDepth(b *testing.B) {
	var log logr.Logger = zaprLogger()
	doWithCallDepth(b, log)
}

// logrusr

func logrusrLogger() logr.Logger {
	logrusLog := logrus.New()
	logrusLog.Level = logrus.DebugLevel
	return logrusr.New(logrusLog)
}

func BenchmarkLogrusrInfoOneArg(b *testing.B) {
	var log logr.Logger = logrusrLogger()
	doInfoOneArg(b, log)
}

func BenchmarkLogrusrInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = logrusrLogger()
	doInfoSeveralArgs(b, log)
}

func BenchmarkLogrusrV0Info(b *testing.B) {
	var log logr.Logger = logrusrLogger()
	doV0Info(b, log)
}

func BenchmarkLogrusrV9Info(b *testing.B) {
	var log logr.Logger = logrusrLogger()
	doV9Info(b, log)
}

func BenchmarkLogrusrError(b *testing.B) {
	var log logr.Logger = logrusrLogger()
	doError(b, log)
}

func BenchmarkLogrusrWithValues(b *testing.B) {
	var log logr.Logger = logrusrLogger()
	doWithValues(b, log)
}

func BenchmarkLogrusrWithName(b *testing.B) {
	var log logr.Logger = logrusrLogger()
	doWithName(b, log)
}

func BenchmarkLogrusrWithCallDepth(b *testing.B) {
	var log logr.Logger = logrusrLogger()
	doWithCallDepth(b, log)
}

// zerologr

func zerologrLogger() logr.Logger {
	zl := zerolog.New(os.Stderr)
	return zerologr.New(&zl)
}

func BenchmarkZerologrInfoOneArg(b *testing.B) {
	var log logr.Logger = zerologrLogger()
	doInfoOneArg(b, log)
}

func BenchmarkZerologrInfoSeveralArgs(b *testing.B) {
	var log logr.Logger = zerologrLogger()
	doInfoSeveralArgs(b, log)
}

func BenchmarkZerologrV0Info(b *testing.B) {
	var log logr.Logger = zerologrLogger()
	doV0Info(b, log)
}

func BenchmarkZerologrV9Info(b *testing.B) {
	var log logr.Logger = zerologrLogger()
	doV9Info(b, log)
}

func BenchmarkZerologrError(b *testing.B) {
	var log logr.Logger = zerologrLogger()
	doError(b, log)
}

func BenchmarkZerologrWithValues(b *testing.B) {
	var log logr.Logger = zerologrLogger()
	doWithValues(b, log)
}

func BenchmarkZerologrWithName(b *testing.B) {
	var log logr.Logger = zerologrLogger()
	doWithName(b, log)
}

func BenchmarkZerologrWithCallDepth(b *testing.B) {
	var log logr.Logger = zerologrLogger()
	doWithCallDepth(b, log)
}
