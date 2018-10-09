package common

import (
	"database/sql"
	"fmt"
	"math"
	"os"

	//importing to register postgres drivers
	_ "github.com/lib/pq"
)

//Db ...
type Db struct {
	Pg *sql.DB
}

//NewDb ...
func NewDb() *Db {
	d := new(Db)
	return d
}

//PostgresConnect ...
func PostgresConnect() { d.PostgresConnect() }

//PostgresConnect ...
func (d *Db) PostgresConnect() {
	info := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		GetFromConfig("postgres.host"), int(GetFromConfig("postgres.port").(float64)),
		GetFromConfig("postgres.user"), GetFromConfig("postgres.password"),
		GetFromConfig("postgres.dbname"))
	fmt.Println("info: ", info)
	con, err := sql.Open("postgres", info)
	if err != nil {
		Log().Error(err)
		os.Exit(1)
	}
	d.Pg = con
	Log().Info("Connection to DB established")
}

//PostgresClose ...
func PostgresClose() { d.PostgresClose() }

//PostgresClose ...
func (d *Db) PostgresClose() {
	err := d.Pg.Close()
	if err != nil {
		Log().Warn("no alive connection", err)
		os.Exit(1)
	}
	Log().Info("Connection to DB closed")
}

//PostgresFetch ...
func PostgresFetch(query string, limit int, args ...interface{}) []map[string]interface{} {
	return d.PostgresFetch(query, limit, args...)
}

//PostgresFetch ...
func (d *Db) PostgresFetch(query string, limit int, args ...interface{}) []map[string]interface{} {
	var rows *sql.Rows
	var err error
	l := int(math.Inf(+1))
	if limit != 0 {
		l = limit
	}
	if args == nil {
		rows, err = d.Pg.Query(query)
	} else {
		rows, err = d.Pg.Query(query, args...)
	}
	if err != nil {
		Log().Error("error occured while fetching result: ", err)
		os.Exit(1)
	}
	cols, _ := rows.Columns()
	defer rows.Close()
	i := 0
	var result []map[string]interface{}
	for rows.Next() {
		buff := make([]interface{}, len(cols))
		buffPointers := make([]interface{}, len(cols))
		buffMap := make(map[string]interface{}, len(cols))
		for i := range buff {
			buffPointers[i] = &buff[i]
		}
		if err := rows.Scan(buffPointers...); err != nil {
			Log().Error("error while scanning rows, ", err)
			os.Exit(1)
		}
		for i, col := range cols {
			buffMap[col] = *(buffPointers[i]).(*interface{})
		}
		result = append(result, buffMap)
		i++
		if i == l {
			break
		}
	}

	return result
}

//PostgresInsUp ...
func PostgresInsUp(query string, args ...interface{}) int64 { return d.PostgresInsUp(query, args...) }

//PostgresInsUp ...
func (d *Db) PostgresInsUp(query string, args ...interface{}) int64 {
	var result sql.Result
	var err error
	if args == nil {
		result, err = d.Pg.Exec(query)
	} else {
		result, err = d.Pg.Exec(query, args...)
	}
	if err != nil {
		Log().Error("error occured while fetching result ", err)
		os.Exit(1)
	}
	re, _ := result.RowsAffected()
	Log(map[string]interface{}{
		"Rows_Effected ": re,
	}).Info("Querry Successfully Executed")

	return re
}
