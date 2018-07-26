# Graylog hook for [zerolog](https://github.com/rs/zerolog)

This is a hook to send log to graylog, support GELF.

## Usage

Sample code using graylog hook

```golang
import (
    "github.com/halink0803/zerolog-graylog-hook/graylog"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    // set default logger to print to stdout and prettify    
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})  
    // udp://127.0.0.1:12201 is log udp url 
	hook, err := graylog.NewGraylogHook("udp://127.0.0.1:12201")
	if err != nil {
		panic(err)
    }
    //Set global logger with graylog hook
    log.Logger = log.Hook(hook)
    
    // Print and send sample log
    for range time.Tick(time.Millisecond * 200) {
		for range time.Tick(time.Millisecond * 200) {
			log.Info.Msg("info log")
            log.Warn.Msg("warning log")
            log.Debug.Msg("debug log")
		}
	}
}
```

*Note: When you first install Graylog, remember to go to /system/inputs to enable GELF input*