# achlibGo

### Use Logger
```
import "github.com/sjsu-achilis/achlibgo/common"</br>
var log = common.NewLogger()
const logfile = "test.log"
func init() {
	log.SetOutputFile("test.log")
}
func main(){
    //log with kay and values
    log.Log(map[string]interface{}{
				"attempt": 1,
			}).Info("This is awesome!!")
    // log without keys and values
    log.Log().info("This is awesome!")
}
```
### Use db wrappers
```
d := common.NewDb()
d.Connect()
//fetch all
rr := d.Fetch("select * from users",0)
//limit by 10
rr := d.Fetch("select * from users", 10)
//use parameterized statement
rr := d.Fetch("select * from users where name=$1 and email=$2", 0, name, email)
//insert...update and delete statement without parameter
re := d.InsUp("DELETE FROM users WHERE name='Test Test'")
//insert...update and delete statement with parameter
re := d.InsUp("DELETE FROM users WHERE name=$1", name)
d.Close()```