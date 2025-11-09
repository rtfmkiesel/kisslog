# kisslog

**Keep It Simple <del>Stupid</del> Structured Golang Logger**  

```sh
go get github.com/rtfmkiesel/kisslog@latest
```

```go
import (
    logger "github.com/rtfmkiesel/kisslog"
)

func main(){
    if err := logger.InitDefault("myapp"); err != nil {
        panic(err)
    }

    log := logger.New("main")
    log.Info("beep boop")
}
```

For more, see [examples](./examples).