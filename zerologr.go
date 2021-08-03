// Package zerologr defines an implementation of the github.com/go-logr/logr
// interfaces built on top of Zerolog (https://github.com/rs/zerolog).
//
// Usage
//
// A new logr.Logger can be constructed from an existing zerolog.Logger using
// the New function:
//
//  log := zerologr.New(someZeroLogger)
//
// Implementation Details
//
// For the most part, concepts in Zerolog correspond directly with those in
// logr.
//
// Levels in logr correspond to custom debug levels in Zerolog.  Any given level
// in logr is represents by `zerologLevel = 1 - logrLevel`.
// For example V(2) is equivalent to Zerolog's TraceLevel, while V(1) is
// equivalent to Zerolog's DebugLevel.
package zerologr

import (
	"encoding/json"
	"net"
	"time"

	"github.com/go-logr/logr"
	"github.com/rs/zerolog"
)

var (
	// NameFieldName is the field key for logr.WithName
	NameFieldName = "logger"
	// NameSeparator separates names for logr.WithName
	NameSeparator = "/"
)

// Logger is the alias of logr.Logger
type Logger = logr.Logger

const (
	infoLevel  = 1 - int(zerolog.InfoLevel)
	debugLevel = 1 - int(zerolog.DebugLevel)
	traceLevel = 1 - int(zerolog.TraceLevel)
)

type logSink struct {
	l      *zerolog.Logger
	name   string
	values []interface{}
	depth  int64
	_      int64 // CPU cache line padding
}

var (
	_ logr.LogSink          = &logSink{}
	_ logr.CallDepthLogSink = &logSink{}
)

// New returns a logr.Logger with logr.LogSink implemented by zerolog.
func New(l *zerolog.Logger) Logger {
	ls := &logSink{l: l}
	return logr.New(ls)
}

func (ls *logSink) Init(ri logr.RuntimeInfo) {
	ls.depth += int64(ri.CallDepth) + 2
}

func (*logSink) Enabled(level int) bool {
	// Info() checks zerolog.GlobalLevel() internally
	return level <= traceLevel
}

func (ls *logSink) Info(level int, msg string, kvList ...interface{}) {
	var e *zerolog.Event
	// small switch: linear search
	switch level {
	case infoLevel:
		e = ls.l.Info()
	case debugLevel:
		e = ls.l.Debug()
	case traceLevel:
		e = ls.l.Trace()
	}
	ls.msg(e, msg, kvList)
}

func (ls *logSink) Error(err error, msg string, kvList ...interface{}) {
	e := ls.l.Error().Err(err)
	ls.msg(e, msg, kvList)
}

func (ls *logSink) msg(e *zerolog.Event, msg string, kvList []interface{}) {
	if e == nil {
		return
	}
	if len(ls.values) > 0 {
		e = handleFields(e, ls.values)
	}
	e = handleFields(e, kvList)
	if len(ls.name) > 0 {
		e.Str(NameFieldName, ls.name)
	}
	e.CallerSkipFrame(int(ls.depth))
	e.Msg(msg)
}

func (ls logSink) WithValues(kvList ...interface{}) logr.LogSink {
	n := len(ls.values)
	ls.values = append(ls.values[:n:n], kvList...)
	return &ls
}

// WithName returns a new logr.Logger with the specified NameFieldName.  zerologr
// uses NameSeparator characters to separate name elements.  Callers should not pass
// NameSeparator in the provided name string, but this library does not actually enforce that.
func (ls logSink) WithName(name string) logr.LogSink {
	if len(ls.name) > 0 {
		ls.name += NameSeparator + name
	} else {
		ls.name = name
	}
	return &ls
}

func (ls logSink) WithCallDepth(depth int) logr.LogSink {
	ls.depth += int64(depth)
	return &ls
}

func handleFields(e *zerolog.Event, kvList []interface{}) *zerolog.Event {
	kvLen := len(kvList)
	if kvLen&0x1 == 1 { // odd number
		kvList = append(kvList, "<no-value>")
	}
	for i := 0; i < kvLen; i += 2 {
		key, val := kvList[i], kvList[i+1]
		k, ok := key.(string)
		if !ok {
			k = "<non-string-key>"
		}
		// concrete type switch: binary search of sorted type hash
		switch v := val.(type) {
		case string:
			e.Str(k, v)
		case []byte:
			e.Bytes(k, v)
		case bool:
			e.Bool(k, v)
		case int:
			e.Int(k, v)
		case int8:
			e.Int8(k, v)
		case int16:
			e.Int16(k, v)
		case int32:
			e.Int32(k, v)
		case int64:
			e.Int64(k, v)
		case uint:
			e.Uint(k, v)
		case uint8:
			e.Uint8(k, v)
		case uint16:
			e.Uint16(k, v)
		case uint32:
			e.Uint32(k, v)
		case uint64:
			e.Uint64(k, v)
		case float32:
			e.Float32(k, v)
		case float64:
			e.Float64(k, v)
		case time.Time:
			e.Time(k, v)
		case time.Duration:
			e.Dur(k, v)
		case []string:
			e.Strs(k, v)
		case []bool:
			e.Bools(k, v)
		case []int:
			e.Ints(k, v)
		case []int8:
			e.Ints8(k, v)
		case []int16:
			e.Ints16(k, v)
		case []int32:
			e.Ints32(k, v)
		case []int64:
			e.Ints64(k, v)
		case []uint:
			e.Uints(k, v)
		case []uint16:
			e.Uints16(k, v)
		case []uint32:
			e.Uints32(k, v)
		case []uint64:
			e.Uints64(k, v)
		case []float32:
			e.Floats32(k, v)
		case []float64:
			e.Floats64(k, v)
		case []time.Time:
			e.Times(k, v)
		case []time.Duration:
			e.Durs(k, v)
		case net.IP:
			e.IPAddr(k, v)
		case net.IPNet:
			e.IPPrefix(k, v)
		case net.HardwareAddr:
			e.MACAddr(k, v)
		case json.RawMessage:
			e.RawJSON(k, v)
		default:
			// interface type switch
			switch v := val.(type) {
			case error:
				e.AnErr(k, v)
			case []error:
				e.Errs(k, v)
			default:
				e.Interface(k, val)
			}
		}
	}
	return e
}
