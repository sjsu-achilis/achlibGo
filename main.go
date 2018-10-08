package main

import (
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
	/*
		d := common.NewDb()
		d.Connect()
		re := d.InsUp("DELETE FROM users WHERE name='Test Test'")
		fmt.Println(re)
		rr := d.Fetch("SELECT * FROM users", 0)
		fmt.Println(rr)
		d.Close()
	*/

}
