package db

import (
	"fmt"
	"testing"
)

func TestSqlStore(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	fmt.Println(">> before:", user1.FirstName, user2.FirstName)
}
