package connection

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
)

func CreateTable(tableName string) string {

	connectionString := "sqlserver://localhost?database=GOLANG_RESTAPI"

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return "error"
	}

	defer db.Close()

	createSql := fmt.Sprintf("create table %s (id INT, idstr varchar(10))", tableName)
	_, err = db.Exec(createSql)
	if err != nil {
		fmt.Println(err)
		return "error"
	}

	return tableName
}

func DropTable(tableName string) string {

	connectionString := "sqlserver://localhost?database=GOLANG_RESTAPI"

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		// log.Fatal(err)
		strin_err := fmt.Sprintf("%v", err)
		return strin_err
	}

	defer db.Close()

	Sql := fmt.Sprintf("drop table %s", tableName)
	_, err = db.Exec(Sql)
	if err != nil {
		strin_err := fmt.Sprintf("%v", err)
		return strin_err
	}

	return tableName
}

func Insert(tableName string) string {

	connectionString := "sqlserver://localhost?database=GOLANG_RESTAPI"

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		// log.Fatal(err)
		strin_err := fmt.Sprintf("%v", err)
		return strin_err
	}

	defer db.Close()

	cExec := 100

	insertSql := fmt.Sprintf("insert into %s (id, idstr) values (@p1, @p2)", tableName)
	// done := make(chan bool)
	stmt, err := db.Prepare(insertSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Stmt is safe to be used by multiple goroutines
	var wg sync.WaitGroup
	wg.Add(cExec)
	for j := 0; j < cExec; j++ {
		go func(val int) {
			defer wg.Done()
			_, err := stmt.Exec(val, strconv.Itoa(val))
			if err != nil {
				log.Fatal(err)
			}
		}(j)
	}
	wg.Wait()

	return tableName
}

func Select(tableName string) string {
	// db init
	connectionString := "sqlserver://localhost?database=GOLANG_RESTAPI"

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		// log.Fatal(err)
		strin_err := fmt.Sprintf("%v", err)
		return strin_err
	}

	defer db.Close()

	cExec := 100
	done := make(chan bool)
	selectSql := "select idstr from test where id = "
	// DB is safe to be used by multiple goroutines
	for i := 0; i < cExec; i++ {
		go func(key int) {
			rows, err := db.Query(selectSql + strconv.Itoa(key))
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				var id int64
				err := rows.Scan(&id)
				if err != nil {
					log.Fatal(err)
				} else {
					log.Printf("Found %d\n", key)
				}
			}
			done <- true
		}(i)
	}

	for i := 0; i < cExec; i++ {
		<-done
	}
	return tableName
}
