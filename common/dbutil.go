package common

import (
	"database/sql"
	"fmt"
	"math"
	"os"

	//importing to register postgres drivers
	_ "github.com/lib/pq"
)

const (
	host     = "achdb.che00bpg1gs1.us-west-2.rds.amazonaws.com"
	port     = 5432
	user     = "achuser"
	password = "achpassword"
	dbname   = "achdb"
)

var log = NewLogger()

//Db ...
type Db struct {
	Info string
	Con  *sql.DB
}

//NewDb ...
func NewDb() *Db {
	d := new(Db)
	d.SetConnectionInfo()
	return d
}

//SetConnectionInfo ...
func (d *Db) SetConnectionInfo() {
	d.Info = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

//Connect ...
func (d *Db) Connect() {
	con, err := sql.Open("postgres", d.Info)
	if err != nil {
		log.Log().Error(err)
		os.Exit(1)
	}
	d.Con = con
	log.Log().Info("Connection to DB established")
}

//Close ...
func (d *Db) Close() {
	err := d.Con.Close()
	if err != nil {
		log.Log().Warn("no alive connection", err)
		os.Exit(1)
	}
	log.Log().Info("Connection to DB closed")
}

//Fetch ...
func (d *Db) Fetch(query string, limit ...int) []map[string]interface{} {
	l := int(math.Inf(+1))
	if limit != nil {
		l = limit[0]
	}
	rows, err := d.Con.Query(query)
	if err != nil {
		log.Log().Error("error occured while fetching result", err)
		os.Exit(1)
	}
	cols, _ := rows.Columns()
	defer rows.Close()
	i := 0
	var result []map[string]interface{}
	for rows.Next() {
		buff := make([]interface{}, len(cols))
		bufferPointers := make([]interface{}, len(cols))
		buffMap := make(map[string]interface{}, len(cols))
		for i := range buff {
			bufferPointers[i] = &buff[i]
		}
		if err := rows.Scan(bufferPointers...); err != nil {
			log.Log().Error("error while scanning rows, ", err)
			os.Exit(1)
		}
		for i, col := range cols {
			buffMap[col] = *(bufferPointers[i]).(*interface{})
		}
		result = append(result, buffMap)
		i++
		if i == l {
			break
		}
	}
	return result
}
