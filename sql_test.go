package appsayhello

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO tb_customer (name) VALUES ( 'monyet')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM tb_customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id :", id)
		fmt.Println("name:", name)
	}

	defer rows.Close()
}

func TestGetById(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM tb_customer WHERE id = 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id disini", id, "name:", name)
	}

	defer rows.Close()

}
