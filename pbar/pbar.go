package pbar

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/term"
)

// PBar is a progress bar instance.
type PBar struct {
	// Completed actions.
	Done int

	// Total actions.
	Total int

	// Size multiplier (0.0 - 1.0).
	Size float64

	// Verbosity.
	Verbose int
}

// Deprecated: Use NewPBar() instead.
var (
	// Completed actions.
	Done = 0

	// Total actions.
	Total = 0

	// Size multiplier (0.0 - 1.0).
	Size = 0.25

	// Verbosity.
	Verbose = 0
)

// NewPBar creates a new progress bar instance with default values.
//
// Default values:
//   - Done: 0
//   - Total: 0
//   - Size: 0.25
//   - Verbose: 0
func NewPBar() *PBar {
	return &PBar{
		Done:    0,
		Total:   0,
		Size:    0.25,
		Verbose: 0,
	}
}

// Log progress bar.
func (pb *PBar) Log() {
	prog := 0.0
	if pb.Total != 0 {
		prog = float64(pb.Done) / float64(max(0, pb.Total))
	}

	width, _, _ := term.GetSize(0)
	progLen := float64(prog) * float64(width) * min(1.0, max(0.0, pb.Size))

	lastChar := "█"
	switch i := progLen - float64(int(progLen)); {
	case i < 0.25 && prog != 1:
		lastChar += " "
	case i < 0.5 && prog != 1:
		lastChar += "▄"
	case i < 0.75 && prog != 1:
		lastChar += "▀"
	}

	msg := fmt.Sprintf("\r|%--"+strconv.Itoa(int(float64(width)/4)+1)+"v| %.1f%%", strings.Repeat("█", int(progLen))+lastChar, prog*100)

	fmt.Printf("\r%"+strconv.Itoa(width)+"v", "")
	if len([]rune(msg)) > width {
		fmt.Printf("%."+strconv.Itoa(width-3)+"s...", msg)
	} else {
		fmt.Printf("%."+strconv.Itoa(width)+"s", msg)
	}
}

// Next increments Done by 1 and logs the progress bar.
func (pb *PBar) Next(msg string, format string, v ...any) {
	pb.Done += 1
	pb.Log()
}

// Back decrements Done by 1 and logs the progress bar.
func (pb *PBar) Back() {
	pb.Done -= 1
	pb.Log()
}

// LogMsg logs progress bar with message capabilities.
//
// `msg` and `msgLong` may be ignored based on `pb.Verbose`.
//
// Verbosities:
//
//	<= 0: Plain progress bar.
//	== 1: `msg` is appended to the progress bar.
//	>= 2: Progress bar is discarded, `msgLong` is forwarded to `fmt.Print`.
func (pb *PBar) LogMsg(msg string, msgLong string) {
	if pb.Verbose >= 2 {
		fmt.Print(msgLong)
		return
	}

	prog := 0.0
	if pb.Total != 0 {
		prog = float64(pb.Done) / float64(max(0, pb.Total))
	}

	width, _, _ := term.GetSize(0)
	progLen := float64(prog) * float64(width) * min(1.0, max(0.0, pb.Size))

	lastChar := "█"
	switch i := progLen - float64(int(progLen)); {
	case i < 0.25 && prog != 1:
		lastChar += " "
	case i < 0.5 && prog != 1:
		lastChar += "▄"
	case i < 0.75 && prog != 1:
		lastChar += "▀"
	}

	if pb.Verbose >= 1 {
		msg = fmt.Sprintf("\r|%--"+strconv.Itoa(int(float64(width)/4)+1)+"v| %.1f%% (%v/%v) -> %v", strings.Repeat("█", int(progLen))+lastChar, prog*100, pb.Done, pb.Total, msg)
	} else {
		msg = fmt.Sprintf("\r|%--"+strconv.Itoa(int(float64(width)/4)+1)+"v| %.1f%%", strings.Repeat("█", int(progLen))+lastChar, prog*100)
	}

	fmt.Printf("\r%"+strconv.Itoa(width)+"v", "")
	if len([]rune(msg)) > width {
		fmt.Printf("%."+strconv.Itoa(width-3)+"s...", msg)
	} else {
		fmt.Printf("%."+strconv.Itoa(width)+"s", msg)
	}
}

// NextMsg increments Done by 1 and logs the progress bar with message.
func (pb *PBar) NextMsg(msg string, msgLong string) {
	pb.Done += 1
	pb.LogMsg(msg, msgLong)
}

// BackMsg decrements Done by 1 and logs the progress bar with message.
func (pb *PBar) BackMsg(msg string, msgLong string) {
	pb.Done -= 1
	pb.LogMsg(msg, msgLong)
}

// Deprecated: Use NewPBar().Log() instead.
func Log() {
	prog := 0.0
	if Total != 0 {
		prog = float64(Done) / float64(max(0, Total))
	}

	width, _, _ := term.GetSize(0)
	progLen := float64(prog) * float64(width) * min(1.0, max(0.0, Size))

	lastChar := "█"
	switch i := progLen - float64(int(progLen)); {
	case i < 0.25 && prog != 1:
		lastChar += " "
	case i < 0.5 && prog != 1:
		lastChar += "▄"
	case i < 0.75 && prog != 1:
		lastChar += "▀"
	}

	msg := fmt.Sprintf("\r|%--"+strconv.Itoa(int(float64(width)/4)+1)+"v| %.1f%%", strings.Repeat("█", int(progLen))+lastChar, prog*100)

	fmt.Printf("\r%"+strconv.Itoa(width)+"v", "")
	if len([]rune(msg)) > width {
		fmt.Printf("%."+strconv.Itoa(width-3)+"s...", msg)
	} else {
		fmt.Printf("%."+strconv.Itoa(width)+"s", msg)
	}
}

// Deprecated: Use NewPBar().Next() instead.
func Next(msg string, format string, v ...any) { Done += 1; Log() }

// Deprecated: Use NewPBar().Back() instead.
func Back() { Done -= 1; Log() }

// Deprecated: Use NewPBar().LogMsg() instead.
func LogMsg(msg string, msgLong string) {
	if Verbose >= 2 {
		fmt.Print(msgLong)
		return
	}

	prog := 0.0
	if Total != 0 {
		prog = float64(Done) / float64(max(0, Total))
	}

	width, _, _ := term.GetSize(0)
	progLen := float64(prog) * float64(width) * min(1.0, max(0.0, Size))

	lastChar := "█"
	switch i := progLen - float64(int(progLen)); {
	case i < 0.25 && prog != 1:
		lastChar += " "
	case i < 0.5 && prog != 1:
		lastChar += "▄"
	case i < 0.75 && prog != 1:
		lastChar += "▀"
	}

	if Verbose >= 1 {
		msg = fmt.Sprintf("\r|%--"+strconv.Itoa(int(float64(width)/4)+1)+"v| %.1f%% (%v/%v) -> %v", strings.Repeat("█", int(progLen))+lastChar, prog*100, Done, Total, msg)
	} else {
		msg = fmt.Sprintf("\r|%--"+strconv.Itoa(int(float64(width)/4)+1)+"v| %.1f%%", strings.Repeat("█", int(progLen))+lastChar, prog*100)
	}

	fmt.Printf("\r%"+strconv.Itoa(width)+"v", "")
	if len([]rune(msg)) > width {
		fmt.Printf("%."+strconv.Itoa(width-3)+"s...", msg)
	} else {
		fmt.Printf("%."+strconv.Itoa(width)+"s", msg)
	}
}

// Deprecated: Use NewPBar().NextMsg() instead.
func NextMsg(msg string, msgLong string) { Done += 1; LogMsg(msg, msgLong) }

// Deprecated: Use NewPBar().BackMsg() instead.
func BackMsg(msg string, msgLong string) { Done -= 1; LogMsg(msg, msgLong) }
