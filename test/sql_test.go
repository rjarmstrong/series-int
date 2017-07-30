package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rjarmstrong/series-int/series"
	"testing"
)

func Test_SerializeDeserialize(t *testing.T) {
	db, err := sql.Open("mysql", "root:rootPassword@tcp(localhost:3306)/informatics?parseTime=true&multiStatements=true")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	s := series.NewInt16(60)
	s.SetRange(0, 29, 666)
	_, err = db.Exec("insert into series (active_months) values (?)", []byte(s))
	if err != nil {
		t.Fatal(err)
	}

	rows, err := db.Query("select active_months from series")
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		p := series.NewInt16(60)
		var d []byte
		err := rows.Scan(&d)
		p = series.Int16(d)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(len(p), p)
	}
}
