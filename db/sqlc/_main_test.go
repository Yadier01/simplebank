package db
var testQueries *Queries

import (
  "database/sql"
  "testing"
  "fmt"
)
const (
  dbDriver = "postgres"
  dbSrouce = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

)
func TestMain(m *testing.M) {
  conn, err := sql.Open(dbDriver, dbSource)
  if err != nil {
    log.Fatal("cannot connect to db:", err)
  }
  
  testQueries = New(conn)
}
