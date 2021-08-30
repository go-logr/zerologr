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
// equivalent to Zerolog's DebugLevel.  Zerolog's usual "level" field is
// disabled globally and replaced with "v", whose value is a number and is only
// logged on Info(), not Error().
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

// Logger is type alias of logr.Logger
type Logger = logr.Logger

// LogSink implements logr.LogSink and logr.CallDepthLogSink.
type LogSink struct {
	l      *zerolog.Logger
	name   string
	values []interface{}
	depth  int64
	_      int64 // CPU cache line padding
}

// Underlier exposes access to the underlying logging implementation.  Since
// callers only have a logr.Logger, they have to know which implementation is
// in use, so this interface is less of an abstraction and more of way to test
// type conversion.
type Underlier interface {
	GetUnderlying() *zerolog.Logger
}

var (
	_ logr.LogSink          = &LogSink{}
	_ logr.CallDepthLogSink = &LogSink{}
)

// New returns a logr.Logger with logr.LogSink implemented by Zerolog.
func New(l *zerolog.Logger) Logger {
	ls := NewLogSink(l)
	return logr.New(ls)
}

// NewLogSink returns a logr.LogSink implemented by Zerolog.
func NewLogSink(l *zerolog.Logger) *LogSink {
	if zerolog.LevelFieldName == "level" {
		zerolog.LevelFieldName = ""
	}
	return &LogSink{l: l}
}

// Init receives runtime info about the logr library.
func (ls *LogSink) Init(ri logr.RuntimeInfo) {
	ls.depth = int64(ri.CallDepth) + 2
}

// Enabled tests whether this LogSink is enabled at the specified V-level.
func (ls *LogSink) Enabled(level int) bool {
	// Optimization: Info() will check level internally.
	const traceLevel = 1 - int(zerolog.TraceLevel)
	return level <= traceLevel
}

// Info logs a non-error message at specified V-level with the given key/value pairs as context.
func (ls *LogSink) Info(level int, msg string, keysAndValues ...interface{}) {
	e := ls.l.WithLevel(zerolog.Level(1 - level))
	e.Int("v", level)
	ls.msg(e, msg, keysAndValues)
}

// Error logs an error, with the given message and key/value pairs as context.
func (ls *LogSink) Error(err error, msg string, keysAndValues ...interface{}) {
	e := ls.l.Error().Err(err)
	ls.msg(e, msg, keysAndValues)
}

func (ls *LogSink) msg(e *zerolog.Event, msg string, keysAndValues []interface{}) {
	if e == nil {
		return
	}
	if len(ls.values) > 0 {
		e = handleFields(e, ls.values)
	}
	e = handleFields(e, keysAndValues)
	if ls.name != "" {
		e.Str(NameFieldName, ls.name)
	}
	e.CallerSkipFrame(int(ls.depth))
	e.Msg(msg)
}

// WithValues returns a new LogSink with additional key/value pairs.
func (ls LogSink) WithValues(keysAndValues ...interface{}) logr.LogSink {
	n := len(ls.values)
	ls.values = append(ls.values[:n:n], keysAndValues...)
	return &ls
}

// WithName returns a new LogSink with the specified name appended in NameFieldName.
// Name elements are separated by NameSeparator.
func (ls LogSink) WithName(name string) logr.LogSink {
	if ls.name != "" {
		ls.name += NameSeparator + name
	} else {
		ls.name = name
	}
	return &ls
}

// WithCallDepth returns a new LogSink that offsets the call stack by adding specified depths.
func (ls LogSink) WithCallDepth(depth int) logr.LogSink {
	ls.depth += int64(depth)
	return &ls
}

// GetUnderlying returns the zerolog.Logger underneath this logSink.
func (ls *LogSink) GetUnderlying() *zerolog.Logger {
	return ls.l
}

func handleFields(e *zerolog.Event, keysAndValues []interface{}) *zerolog.Event {
	kvLen := len(keysAndValues)
	if kvLen&0x1 == 1 { // odd number
		keysAndValues = append(keysAndValues, "<no-value>")
	}
	for i := 0; i < kvLen; i += 2 {
		key, val := keysAndValues[i], keysAndValues[i+1]
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
