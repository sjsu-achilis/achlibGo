package main

import (
	"fmt"

	"github.com/sjsu-achilis/achlibgo/common"
)

const logfile = "test.log"

var log = common.NewLogger()

/*
func init() {
	log.SetOutputFile("test.log")
}
*/

func main() {
	/*
		i := 0
		for {
			log.Log(map[string]interface{}{
				"attempt": 1,
			}).Info("This is awesome!!")
			time.Sleep(5 * time.Second)
			i++
		}
	*/
	d := common.NewDb()
	d.Connect()
	rr := d.Fetch("select * from users")
	for _, r := range rr {
		fmt.Println(r["organization"])
	}
	d.Close()

}
