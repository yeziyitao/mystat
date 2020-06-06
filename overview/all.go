package overview

import (
        "fmt"
        "time"

	"github.com/yezi/mystat/system"
)

//show w,df,free,top,iostat
func ShowAll() {
        //w
        result, _ := system.ExeCmd("w")
        fmt.Println(result)

        //df
        result, _ = system.ExeCmd("df -hT|head")
        fmt.Println(result)

        //free
        result, _ = system.ExeCmd("free -mt")
        fmt.Println(result)

        //top
        result, _ = system.ExeCmd("top -n 1 |head -5")
        fmt.Println(result)

        //iostat
        result, _ = system.ExeCmd("iostat -xm 1 1")
        fmt.Println(result)
        time.Sleep(time.Duration(1)*time.Second)
}
