# achlibGo

### Use Logger
```import "github.com/sjsu-achilis/achlibgo/common"</br>
var log = common.NewLogger()</br>
const logfile = "test.log"
func init() {
	log.SetOutputFile("test.log")
}
func main(){
    log.Log(map[string]interface{}{
				"attempt": 1,
			}).Info("This is awesome!!")

    log.Log().info("This is awesome!")
}
```
### Use db wrappers
```d := common.NewDb()
d.Connect()
//fetch all
rr := d.Fetch("select * from users")
//limit by 10
rr := d.Fetch("select * from users", 10)
d.Close()```