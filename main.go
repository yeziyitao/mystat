package main

import (
        "fmt"
        "time"
        "flag"
        "os"
        "encoding/json"

	"github.com/yezi/mystat/overview"
	"github.com/yezi/mystat/disk"
)

var (
        //VERSION tool version
        VERSION = "1.1.0"

        //h help
        h bool
        //v,V version
        v, V bool

	e = "Mystat"
	//for all overview
	a bool
)

/*
struct MSG
markdown msg
*/
type MSG struct {
        // msg type
        Msgtype string `json:"msgtype"`

        // markdown msg content
        Info Markdown `json:"markdown"`
}

/*
struct Markdown
markdown content
*/
type Markdown struct {
        // content detail info ,markdown type
        Content string `json:"content"`
}

/*
@ConvertJson
Convert msg to json format with markdown
send msg to wx
*/
func ConvertJson(content string) (jsonStr []byte) {
        mymsg := MSG{
                Msgtype: "json",
                Info: Markdown{
                        Content: content,
                },
        }
        // convert json to []byte
        myjsonStr, _ := json.Marshal(mymsg)
        return myjsonStr
}

/*
@usage
edit help msg
*/
func usage() {
        fmt.Println(os.Args[0], `version:`, VERSION, `
Usage: `, os.Args[0], `[OPTIONS]

Options:
        `)
        flag.PrintDefaults()
}

/*
@init
init flag
*/
func init() {
        // set help and version
        flag.BoolVar(&h, "h", false, "this help")
        flag.BoolVar(&v, "v", false, "show version and exit")
        flag.BoolVar(&V, "V", false, "show version and configure options then exit")

        // wx msg webhook addr
        flag.StringVar(&e, "e", "", "")

        //all
        flag.BoolVar(&a, "a", false, "show all overview,about cpu,men,disk")

        //set defalut Usage to usage
        flag.Usage = usage

}

func main() {
        flag.Parse()
        if h || v || V {
                flag.Usage()
                os.Exit(0)
        }

        if a {
		overview.ShowAll()
                os.Exit(0)
        }
        var cstZone = time.FixedZone("CST", 8*3600) // 8 timezone
        timestamp := time.Now().In(cstZone).Format("2006-01-02 15:04:05")

        info := timestamp + "=" + e
        fmt.Println(e, string(ConvertJson(info)))

	disk.IoStat()

}
