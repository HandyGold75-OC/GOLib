# GOLib

A collection of Go utility packages for building CLI applications and tools.

## Packages

### `argp`

Argument parser with struct tag-based configuration. Supports switches, positional arguments, and automatic help generation.

```go
import "github.com/HandyGold75/GOLib/argp"

args := argp.ParseArgs(struct {
    Help   bool     `switch:"h,-help" opts:"help" help:"Show this help message"`
    Name   string   `switch:"n,-name" opts:"required" help:"Your name"`
    Count  int      `switch:"c,-count" default:"1" help:"Repeat count"`
    Verbose bool    `switch:"v,-verbose" help:"Enable verbose output"`
})
```

### `cfg`

JSON configuration file management. Load and save configs with automatic file creation.

```go
import "github.com/HandyGold75/GOLib/cfg"

type Config struct {
    Host string
    Port int
}

var config Config
if err := cfg.Load("myapp", &config); err != nil {
    log.Fatal(err)
}
```

### `gapo`

TP-Link Tapo smart plug control. Control Tapo devices over the local network.

```go
import "github.com/HandyGold75/GOLib/gapo"

device, err := gapo.NewTapo("192.168.1.100", "email@example.com", "password")
if err != nil {
    log.Fatal(err)
}

// Turn device on
device.On()

// Get device info
info, _ := device.GetDeviceInfo()
fmt.Printf("Device: %s, On: %v\n", info.Nickname, info.DeviceOn)
```

### `keyboard`

Linux keyboard event input handling. Read and send keyboard events via `/dev/input/event*`.

```go
import "github.com/HandyGold75/GOLib/keyboard"

kb, err := keyboard.NewKeyboard("")
if err != nil {
    log.Fatal(err)
}

events := kb.Read()
for event := range events {
    if event.IsPress() {
        fmt.Printf("Key pressed: %s\n", event.String())
    }
}
```

### `logger`

Configurable logging with terminal colors, file output, and verbosity levels.

```go
import "github.com/HandyGold75/GOLib/logger"

log, err := logger.New("myapp")
if err != nil {
    log.Fatal(err)
}

log.Log("info", "Application started")
log.Log("error", "Something went wrong")
```

### `pbar`

Progress bar for CLI applications with multiple verbosity modes.

```go
import "github.com/HandyGold75/GOLib/pbar"

pbar.Total = 100
for i := 0; i < 100; i++ {
    pbar.Next("", "")
    time.Sleep(10 * time.Millisecond)
}
```

### `scheduler`

Cron-like scheduling utilities with time-based wake/sleep functions.

```go
import "github.com/HandyGold75/GOLib/scheduler"

schedule := scheduler.Schedule{
    Months:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
    Weeks:   []int{1, 2, 3, 4, 5},
    Days:    []int{0, 1, 2, 3, 4, 5, 6},
    Hours:   []int{0, 6, 12, 18},
    Minutes: []int{0},
}

nextRun := time.Now()
scheduler.SetNextTime(&nextRun, schedule)
fmt.Printf("Next run: %v\n", nextRun)
```

### `tui`

Terminal User Interface with interactive menus. Supports basic and bulky renderers.

```go
import "github.com/HandyGold75/GOLib/tui"

menu := tui.NewMenuBasic("Main Menu")

menu.Menu.NewAction("Say Hello", func() {
    fmt.Println("Hello!")
})

textItem := menu.Menu.NewText("Username", tui.GeneralCharSet, "")
menu.Menu.NewAction("Greet", func() {
    fmt.Printf("Hello, %s!\n", textItem.Value())
})

if err := menu.Run(); err != nil {
    log.Fatal(err)
}
```

### `yts`

YouTube search client without requiring API authentication.

```go
import "github.com/HandyGold75/GOLib/yts"

client := yts.NewSearch("golang tutorial", yts.FilterVideo, yts.OrderRelevance)
results, err := client.Next()
if err != nil {
    log.Fatal(err)
}

for _, video := range results.Videos {
    fmt.Printf("%s: %s\n", video.Title, video.URL)
}
```


