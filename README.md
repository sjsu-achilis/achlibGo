# achlibGo

### Use Logger
```
import "github.com/sjsu-achilis/achlibgo/common"

//Log only to stdout/stderr

common.Log().Info("Logging")
common.Log().Debug("Loggging")
common.Log().Warn("Warning")
common.Log().Error("Error message")
common.Log().Panic("panic")

//Log with key and values

common.Log(map[string]interface{}{"key":"value", "key1":10,}).Info("Logging")

//set output file

common.SetLogOutputFile("filename")
common.Log().Info("Logging")

```
### Use dbutil
```
//connect to postgres DB
common.PostgresConnect()

//fetch all from the DB
rr := common.PostgresFetch("select * from users",0)

//limit fetch by 10
rr := common.PostgresFetch("select * from users",10)

//use parameterized statement
rr := common.PostgresFetch("select * from users where name=$1 and email=$2", 0, name, email)

//insert...update and delete statement without parameter
re := common.PostgresInsUp("DELETE FROM users WHERE name='Test Test'")

//insert...update and delete statement with parameter
re := common.PostgresInsUp("DELETE FROM users WHERE name=$1", name)

//close connection
common.Close()

```

### Configs
```
//the default config file is config-default.json
//use this to get value
common.GetFromConfig("postgres.host")

//can use this to add/overrride more configs, the file name should not be config-default.json
common.SetConfigFilePath("file name", "absolute path", "json")
```