package main
import (
 "bufio"
 "fmt"
 "os"
 //"database/sql" //non-functional MySQL imports
 // _ "/mysql"
)
func main(){

  con, err := sql.Open("mysql", store.user+":"+store.password+"@/"+store.database)
  defer con.Close()
  if err != nil {
    fmt.Println("OH NO!", err)
  }
  fmt.Println("read line:",)
}

func MySQLQuery (input string) string{
        //this will return the MySQL query string, right now it doesn't
        con, err := sql.Open("mysql", store.user+":"+store.password+"@/"+store.database)
        defer con.Close()
}

