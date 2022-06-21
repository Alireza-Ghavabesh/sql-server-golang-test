package main

import (
	// "database/sql"
	// "log"
	// "strconv"
	// "sync"
	"github.com/alireza-ghavabesh/sql-server-test/server"
)

func main() {
	server.StartServer()
}

// // number of executions for go threads
// cExec := 100

// // dropSql := "drop table test"
// // db.Exec(dropSql)

// createSql := "create table test (id INT, idstr varchar(10))"
// _, err = db.Exec(createSql)
// if err != nil {
// 	log.Fatal(err)
// }

// insertSql := "insert into test (id, idstr) values (@p1, @p2)"
// done := make(chan bool)
// stmt, err := db.Prepare(insertSql)
// if err != nil {
// 	log.Fatal(err)
// }
// defer stmt.Close()

// // Stmt is safe to be used by multiple goroutines
// var wg sync.WaitGroup
// wg.Add(cExec)
// for j := 0; j < cExec; j++ {
// 	go func(val int) {
// 		defer wg.Done()
// 		_, err := stmt.Exec(val, strconv.Itoa(val))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}(j)
// }
// wg.Wait()

// selectSql := "select idstr from test where id = "
// // DB is safe to be used by multiple goroutines
// for i := 0; i < cExec; i++ {
// 	go func(key int) {
// 		rows, err := db.Query(selectSql + strconv.Itoa(key))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer rows.Close()
// 		for rows.Next() {
// 			var id int64
// 			err := rows.Scan(&id)
// 			if err != nil {
// 				log.Fatal(err)
// 			} else {
// 				log.Printf("Found %d\n", key)
// 			}
// 		}
// 		done <- true
// 	}(i)
// }

// for i := 0; i < cExec; i++ {
// 	<-done
// }
