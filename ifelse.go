package main
import "fmt"

func main() {  
    var x = 50
    if x < 10 {
        fmt.Println("x is less than 10")
    } else {
        fmt.Println("x is greater than or equals 10")
    }
}



PS C:\Users\Sudha\Desktop\golang> go run ifelse.go;
x is greater than or equals 10



