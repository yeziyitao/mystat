package disk 

import (
        "fmt"
        "time"

	"github.com/yezi/mystat/system"
)

//show w,df,free,top,iostat
func IoStat() {
        //iostat
        result, _ := system.ExeCmd("iostat -xm 1 1")
        fmt.Println(result)
        time.Sleep(time.Duration(1)*time.Second)
}
