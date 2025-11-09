package kisslog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	FlagDebug           = false // Enables logger.Debug()
	FlagTime            = true  // Timestamp prefix
	FlagColor           = true  // Colored levels via github.com/fatih/color
	FlagSilenceAll      = false // Mute all output
	currentGlobalLogger *globalLogger
)

type level int

const (
	levelDebug level = iota
	levelInfo
	levelWarning
	levelError
	levelFatal
)

// Returns the (optionally colored) string for the level.
func (l level) String() string {
	var text string
	var colorFunc func(string, ...any) string

	switch l {
	case levelDebug:
		text, colorFunc = "DEBG", color.BlueString
	case levelInfo:
		text, colorFunc = "INFO", color.WhiteString
	case levelWarning:
		text, colorFunc = "WARN", color.YellowString
	case levelError:
		text, colorFunc = "ERRO", color.RedString
	case levelFatal:
		text, colorFunc = "FATL", color.RedString
	default:
		text, colorFunc = "????", color.MagentaString
	}

	if FlagColor {
		return colorFunc(text)
	}

	return text
}

type Config struct {
	Base    string    // Log base, e.g. "github.com/myuser/mytool"
	TimeStr string    // Time format string to be used when FlagTime is set to true, default: time.RFC3339
	Delim   byte      // Byte to use for separating log parts, default '|'
	Output  io.Writer // Where log messages will be written to, default: github.com/fatih/color.Error
}

type globalLogger struct {
	config *Config
}

// Sets a baseline config if the logger is called without Init().
func init() {
	_ = InitDefault("kisslog")
}

// Initializes the logger based on config.
func Init(config *Config) error {
	if config == nil || strings.TrimSpace(config.Base) == "" {
		return fmt.Errorf("invalid config")
	}

	// Set defaults
	if config.Delim == 0 {
		config.Delim = '|'
	}
	if config.Output == nil {
		config.Output = color.Error
	}
	if config.TimeStr == "" {
		config.TimeStr = time.RFC3339
	}

	currentGlobalLogger = &globalLogger{config: config}
	return nil
}

// Initialized the logger to base with the default config.
func InitDefault(base string) error {
	return Init(&Config{Base: base})
}

// Formats message of lvl from module and then prints it to the configured output.
func (gl *globalLogger) write(module string, lvl level, message string) {
	if FlagSilenceAll {
		return
	}

	var b strings.Builder

	if FlagTime {
		b.WriteString(time.Now().Format(gl.config.TimeStr))
		b.WriteByte(gl.config.Delim)
	}

	b.WriteString(lvl.String())
	b.WriteByte(gl.config.Delim)
	b.WriteString(gl.config.Base)
	b.WriteByte(gl.config.Delim)
	b.WriteString(module)
	b.WriteByte(gl.config.Delim)
	b.WriteString(message)

	if !strings.HasSuffix(message, "\n") {
		b.WriteByte('\n')
	}

	_, _ = fmt.Fprint(gl.config.Output, b.String())
}

// A logger to be used inside a module.
type Logger struct {
	module string
}

// Returns a new logger which can be used inside a module.
func New(moduleName string) *Logger {
	return &Logger{module: moduleName}
}

// Prints a format string with the level debug. Only effective if ShowDebug() is set to true.
func (l *Logger) Debug(s string, args ...any) {
	if FlagDebug {
		currentGlobalLogger.write(l.module, levelDebug, fmt.Sprintf(s, args...))
	}
}

// Prints a format string with the level info.
func (l *Logger) Info(s string, args ...any) {
	currentGlobalLogger.write(l.module, levelInfo, fmt.Sprintf(s, args...))
}

// Prints a format string with the level warning.
func (l *Logger) Warning(s string, args ...any) {
	currentGlobalLogger.write(l.module, levelWarning, fmt.Sprintf(s, args...))
}

// Prints an error or a format string with the level error.
func (l *Logger) Error(v any, args ...any) {
	var msg string
	switch val := v.(type) {
	case error:
		msg = val.Error()
	case string:
		msg = fmt.Sprintf(val, args...)
	default:
		msg = fmt.Sprint(v)
	}

	currentGlobalLogger.write(l.module, levelError, msg)
}

// Prints an error or a format string  with the level fatal, then quits with os.Exit(1).
func (l *Logger) Fatal(v any, args ...any) {
	var msg string
	switch val := v.(type) {
	case error:
		msg = val.Error()
	case string:
		msg = fmt.Sprintf(val, args...)
	default:
		msg = fmt.Sprint(v)
	}

	currentGlobalLogger.write(l.module, levelFatal, msg)

	os.Exit(1)
}
