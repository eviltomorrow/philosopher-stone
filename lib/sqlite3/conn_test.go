package sqlite3

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	DSN = "/home/shepard/tmp/test.db"

	if err := Build(); err != nil {
		log.Fatal(err)
	}
}

func TestConn(t *testing.T) {
	rows, err := DB.Query("select * from user")
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%d-%s-%d\r\n", id, name, age)
	}

}
