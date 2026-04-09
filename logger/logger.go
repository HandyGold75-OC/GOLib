package logger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

type (
	Color string

	Logger struct {
		// When logging to file this file will be used.
		//
		// If `logger.DynamicFileName` is not nil this becomes the used path (.log Suffix is trimmed).
		FilePath string
		// Append the return to `logger.FilePath` (.log Suffix is trimmed from `logger.FilePath`).
		DynamicFileName func() string
		// Mapping of Verbosities to set allowed verbosities and their priority.
		Verbosities map[string]int
		// Mapping of Verbosities Colors to set their cli color.
		VerbositiesColors map[string]Color
		// Minimal verbose priority to log message to CLI.
		VerboseToCLI int
		// Minimal verbose priority to log message to file.
		VerboseToFile int
		// Prepend logs with the date time.
		AppendDateTime bool
		// Prepend logs with the verbosity.
		AppendVerbosity bool
		// Prepend logs with this.
		PrependCLI string
		// Called after every log to CLI.
		MessageCLIHook func(msg string)
		// Minimal char count a log part will take up.
		CharCountPerPart int
		// Minimal char space the verbosity part will take up (AppendVerbosity must be true to take effect).
		CharCountVerbosity int
		// When true RecordSeparator and EORSeparator are used when logging to file, otherwise log the raw message.
		UseSeparators bool
		// Separator string between message parts when logging to file (Logged message can not contain this string).
		RecordSeparator string
		// End of record string after a message when logging to file (Logged message can not contain this string).
		EORSeparator string
	}
)

const (
	Reset Color = "\033[0m"

	Bold            Color = "\033[1m"
	Faint           Color = "\033[2m"
	Italic          Color = "\033[3m"
	Underline       Color = "\033[4m"
	StrikeTrough    Color = "\033[9m"
	DoubleUnderline Color = "\033[21m"

	Black   Color = "\033[30m"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	White   Color = "\033[37m"

	BGBlack   Color = "\033[40m"
	BGRed     Color = "\033[41m"
	BGGreen   Color = "\033[42m"
	BGYellow  Color = "\033[43m"
	BGBlue    Color = "\033[44m"
	BGMagenta Color = "\033[45m"
	BGCyan    Color = "\033[46m"
	BGWhite   Color = "\033[47m"

	BrightBlack   Color = "\033[90m"
	BrightRed     Color = "\033[91m"
	BrightGreen   Color = "\033[92m"
	BrightYellow  Color = "\033[93m"
	BrightBlue    Color = "\033[94m"
	BrightMagenta Color = "\033[95m"
	BrightCyan    Color = "\033[96m"
	BrightWhite   Color = "\033[97m"

	BGBrightBlack   Color = "\033[100m"
	BGBrightRed     Color = "\033[101m"
	BGBrightGreen   Color = "\033[102m"
	BGBrightYellow  Color = "\033[103m"
	BGBrightBlue    Color = "\033[104m"
	BGBrightMagenta Color = "\033[105m"
	BGBrightCyan    Color = "\033[106m"
	BGBrightWhite   Color = "\033[107m"
)

var (
	// Default value for `logger.DynamicFileName`, does not effect existing loggers.
	DynamicFileName func() string = nil
	// Default value for `logger.Verbosities`, does not effect existing loggers.
	Verbosities = map[string]int{"high": 3, "medium": 2, "low": 1}
	// Default value for `logger.VerbositiesColors`, does not effect existing loggers.
	VerbositiesColors = map[string]Color{"high": Red, "low": BrightBlack}
	// Default value for `logger.VerboseToCLI`, does not effect existing loggers.
	VerboseToCLI = 1
	// Default value for `logger.VerboseToFile`, does not effect existing loggers.
	VerboseToFile = 2
	// Default value for `logger.AppendDateTime`, does not effect existing loggers.
	AppendDateTime = true
	// Default value for `logger.AppendVerbosity`, does not effect existing loggers.
	AppendVerbosity = true
	// Default value for `logger.PrependCLI`, does not effect existing loggers.
	PrependCLI = ""
	// Default value for `logger.MessageCLIHook`, does not effect existing loggers.
	MessageCLIHook func(msg string) = nil
	// Default value for `logger.CharCountPerPart`, does not effect existing loggers.
	CharCountPerPart = 32
	// Default value for `logger.CharCountVerbosity`, does not effect existing loggers.
	CharCountVerbosity = 7
	// Default value for `logger.UseSeparators`, does not effect existing loggers.
	UseSeparators = true
	// Default value for `logger.RecordSeparator`, does not effect existing loggers.
	RecordSeparator = "<SEP>"
	// Default value for `logger.EORSeparator`, does not effect existing loggers.
	EORSeparator = "<EOR>\n"
)

// Create new logger instance.
//
// If log file is not present then it tries creating it.
//
// Log file is stored in `./golib/<name>.log` relative to `os.UserConfigDir`.
func New(name string) (*Logger, error) {
	file, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	return NewAbs(file + "/golib/" + name + ".log"), nil
}

// Create new logger instance.
//
// If log file is not present then it tries creating it.
//
// Log file is stored in `./<name>.log` relative to `os.Executable`.
func NewRel(name string) (*Logger, error) {
	file, err := os.Executable()
	if err != nil {
		return nil, err
	}
	fileSplit := strings.Split(strings.ReplaceAll(file, "\\", "/"), "/")
	return NewAbs(strings.Join(fileSplit[:len(fileSplit)-1], "/") + "/" + name + ".log"), nil
}

// Create new logger instance.
//
// If log file is not present then it tries creating it.
func NewAbs(file string) *Logger {
	return &Logger{
		FilePath:           file,
		DynamicFileName:    DynamicFileName,
		Verbosities:        Verbosities,
		VerbositiesColors:  VerbositiesColors,
		VerboseToCLI:       VerboseToCLI,
		VerboseToFile:      VerboseToFile,
		AppendDateTime:     AppendDateTime,
		AppendVerbosity:    AppendVerbosity,
		PrependCLI:         PrependCLI,
		MessageCLIHook:     MessageCLIHook,
		CharCountPerPart:   CharCountPerPart,
		CharCountVerbosity: CharCountVerbosity,
		UseSeparators:      UseSeparators,
		RecordSeparator:    RecordSeparator,
		EORSeparator:       EORSeparator,
	}
}

func (logger Logger) logToCLI(verbosity string, msgs ...any) {
	width, _, _ := term.GetSize(0)
	msg := fmt.Sprintf(strings.Repeat("%-"+strconv.Itoa(min(logger.CharCountPerPart, int(float64(width)/float64(len(msgs)))))+"v", len(msgs)), msgs...)

	if logger.AppendVerbosity {
		msg = fmt.Sprintf("%-"+strconv.Itoa(logger.CharCountVerbosity)+"v ", verbosity) + msg
	}
	if logger.AppendDateTime {
		msg = "[" + time.Now().Format(time.DateTime) + "] " + msg
	}
	msg = logger.PrependCLI + msg

	col, ok := logger.VerbositiesColors[verbosity]
	if ok {
		msg = string(col) + msg + string(Reset)
	}

	if len([]rune(msg)) > width {
		fmt.Printf("%."+strconv.Itoa(width-3)+"s...\n", msg)
	} else {
		fmt.Printf("%."+strconv.Itoa(width)+"s\n", msg)
	}

	if logger.MessageCLIHook != nil {
		logger.MessageCLIHook(msg)
	}
}

func (logger Logger) logToFile(verbosity string, msgs ...any) {
	var msg string

	if logger.UseSeparators {
		msg = fmt.Sprintf(strings.Repeat("%v"+logger.RecordSeparator, len(msgs)), msgs...)

		if logger.AppendVerbosity {
			msg = verbosity + logger.RecordSeparator + msg
		}
		if logger.AppendDateTime {
			msg = time.Now().Format(time.RFC3339Nano) + logger.RecordSeparator + msg
		}

		i := strings.LastIndex(msg, logger.RecordSeparator)
		msg = msg[:i] + strings.Replace(msg[i:], logger.RecordSeparator, "", 1) + logger.EORSeparator
	} else {
		msg = fmt.Sprintf(strings.Repeat("%-"+strconv.Itoa(logger.CharCountPerPart)+"v", len(msgs))+"\n", msgs...)

		if logger.AppendVerbosity {
			msg = fmt.Sprintf("%-"+strconv.Itoa(logger.CharCountVerbosity)+"v ", verbosity) + msg
		}
		if logger.AppendDateTime {
			msg = "[" + time.Now().Format(time.DateTime) + "] " + msg
		}
	}

	fp := logger.FilePath
	if logger.DynamicFileName != nil {
		fp = strings.TrimSuffix(fp, ".log") + "/" + logger.DynamicFileName()
	}

	if _, err := os.Stat(fp); os.IsNotExist(err) {
		fileSplit := strings.Split(strings.ReplaceAll(fp, "\\", "/"), "/")
		err := os.MkdirAll(strings.Join(fileSplit[:len(fileSplit)-1], "/"), os.ModePerm)
		if err != nil {
			logger.logToCLI("ERROR", "Failed creating logpath", err)
			return
		}
		if err := os.WriteFile(fp, []byte(msg), 0640); err != nil {
			logger.logToCLI("ERROR", "Failed creating logfile", err)
		}
		return
	}

	logFile, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.logToCLI("ERROR", "Failed opening logfile", err)
	}
	if _, err := logFile.Write([]byte(msg)); err != nil {
		logger.logToCLI("ERROR", "Failed writing to logfile", err)
	}
	if err := logFile.Close(); err != nil {
		logger.logToCLI("ERROR", "Failed closing logfile", err)
	}
}

// Log an message.
//
// If `verbosity` is not present in `logger.Verbosities` then it is set to `{ "ERROR": 99 }`
//
// Message is logged to CLI if `verbosity >= logger.VerboseToCLI`.
//
// Message is logged to file if `verbosity >= logger.VerboseToFile`.
func (logger Logger) Log(verbosity string, msgs ...any) {
	verboseLevel, ok := logger.Verbosities[verbosity]
	if !ok {
		verbosity, verboseLevel = "ERROR", 99
	}

	if verboseLevel >= logger.VerboseToFile {
		for _, msg := range msgs {
			if strings.Contains(fmt.Sprintf("%v", msg), logger.RecordSeparator) {
				logger.logToCLI("ERROR", "Msg contains "+logger.RecordSeparator, msg)
				return
			}
			if strings.Contains(fmt.Sprintf("%v", msg), logger.EORSeparator) {
				logger.logToCLI("ERROR", "Msg contains "+strings.ReplaceAll(logger.EORSeparator, "\n", "\\n"), strings.ReplaceAll(msg.(string), "\n", "\\n"))
				return
			}
		}
		logger.logToFile(verbosity, msgs...)
	}
	if verboseLevel >= logger.VerboseToCLI {
		logger.logToCLI(verbosity, msgs...)
	}
}
