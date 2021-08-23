package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("sql: database is closed")
	var errDBClosed = errors.New("sql: database is closed")
	if err.Error() == errDBClosed.Error() {
		fmt.Println(err.Error())
	}
}
