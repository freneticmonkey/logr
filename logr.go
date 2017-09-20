package logr

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

// Writer holds all state required to write a log
type Writer struct {
	level    int
	filepath string
	logFile  *os.File
}

// NewWriter creates a new instance with a default log level of INFO
func NewWriter(level int, file string) Writer {
	if level == 0 {
		level = INFO
	}

	return Writer{
		level:    level,
		filepath: file,
	}
}

// SetLevel of the log writer
func (w *Writer) SetLevel(newLevel int) {
	w.level = newLevel
}

// OpenFile for logging using the filepath set in Writer
func (w *Writer) OpenFile() (func(), error) {

	var (
		err error
	)

	if w.filepath != "" {
		w.logFile, err = os.OpenFile(w.filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}

		log.SetOutput(io.MultiWriter(os.Stderr, w.logFile))

		return func() {
				// TODO: Check if the lambda err variable is copied when func is executed in a caller defer call
				err = w.logFile.Close()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Problem closing the log file: %s\n", err)
				}
			},
			nil
	}

	return func() {
			w.Error("Logging: Invalid Filepath. Defaulting to stdout")
		},
		errors.New("Invalid Filepath")
}

// log sets the message colour and filters by level
func (w Writer) log(out string, attr color.Attribute, level int) {
	if w.level <= level {
		color.Set(attr)
		log.Printf(out)
		color.Unset()
	}
}

// Green writes an Green coloured message with level
func (w Writer) Green(out string, level int) {
	w.log(out, color.FgGreen, level)
}

// White writes an White coloured message with level
func (w Writer) White(out string, level int) {
	w.log(out, color.FgWhite, level)
}

// Magenta writes an Magenta coloured message with level
func (w Writer) Magenta(out string, level int) {
	w.log(out, color.FgMagenta, level)
}

// Yellow writes an Yellow coloured message with level
func (w Writer) Yellow(out string, level int) {
	w.log(out, color.FgYellow, level)
}

// Cyan writes an Cyan coloured message with level
func (w Writer) Cyan(out string, level int) {
	w.log(out, color.FgCyan, level)
}

// Red writes an Red coloured message with level
func (w Writer) Red(out string, level int) {
	w.log(out, color.FgRed, level)
}

// RedBright writes an Bright Red coloured message with level
func (w Writer) RedBright(out string, level int) {
	w.log(out, color.FgHiRed, level)
}

// Info writes an Info message with level INFO
func (w Writer) Info(info ...interface{}) {
	w.White("INFO:  "+fmt.Sprintln(info...), INFO)
}

// Infof writes an Infof message with level INFO
func (w Writer) Infof(format string, info ...interface{}) {
	w.White("INFO:  "+fmt.Sprintf(format, info...), INFO)
}

// Ok writes an Ok message with level DEBUG
func (w Writer) Ok(info ...interface{}) {
	w.Green("OK:    "+fmt.Sprintln(info...), DEBUG)
}

// Okf writes an Okf message with level DEBUG
func (w Writer) Okf(format string, info ...interface{}) {
	w.Green("OK:    "+fmt.Sprintf(format, info...), DEBUG)
}

// Attention writes an Attention message with ATTNlevel
func (w Writer) Attention(info ...interface{}) {
	w.Yellow("ATTN:  "+fmt.Sprintln(info...), ATTN)
}

// Attentionf writes an Attentionf message with levelATTN
func (w Writer) Attentionf(format string, info ...interface{}) {
	w.Yellow("ATTN:  "+fmt.Sprintf(format, info...), ATTN)
}

// Warn writes an Warn message with level WARN
func (w Writer) Warn(warn ...interface{}) {
	w.Magenta("WARN:  "+fmt.Sprintln(warn...), WARN)
}

// Warnf writes an Warnf message with level WARN
func (w Writer) Warnf(format string, warn ...interface{}) {
	w.Magenta("WARN:  "+fmt.Sprintf(format, warn...), WARN)
}

// Alert writes an Alert message with level WARN
func (w Writer) Alert(alert ...interface{}) {
	w.Cyan("ALERT: "+fmt.Sprintln(alert...), WARN)
}

// Alertf writes an Alertf message with level WARN
func (w Writer) Alertf(format string, alert ...interface{}) {
	w.Cyan("ALERT: "+fmt.Sprintf(format, alert...), WARN)
}

// Error writes an Error message with leveERRORl
func (w Writer) Error(err ...interface{}) {
	w.Red("ERROR: "+fmt.Sprintln(err...), ERROR)
}

// Errorf writes an Errorf message with level ERROR
func (w Writer) Errorf(format string, err ...interface{}) {
	w.Red("ERROR: "+fmt.Sprintf(format, err...), ERROR)
}

// Fatal writes an Fatal message with level ERROR
func (w Writer) Fatal(code int, err ...interface{}) {
	w.Red("FATAL: "+fmt.Sprintln(err...), ERROR)
	os.Exit(code)
}
