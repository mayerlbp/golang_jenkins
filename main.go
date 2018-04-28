package main
import "fmt"
import (
  "testing"
  "flag"
)

func printInfo() {
    fmt.Println("Hello World")
}

func main() {
    printInfo() 
}

var systemTest *bool
func init() {
  systemTest = flag.Bool("systemTest", false, "Set to true when running system tests")
}

func TestSystem(t *testing.T) {
  if *systemTest {
     main()
  }
}



