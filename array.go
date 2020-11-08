package main
import "fmt"

func main() {  
    var numbers [3] string  
    numbers[0] = "One"
    numbers[1] = "Two"
    numbers[2] = "Three"
    fmt.Println(numbers[1]) 
    fmt.Println(len(numbers)) 
    fmt.Println(numbers) 

    directions := [...] int {1,2,3,4,5}  
    fmt.Println(directions) 
    fmt.Println(len(directions))
}


PS C:\Users\Sudha\Desktop\golang> go run array.go;
Two
3
[One Two Three]
[1 2 3 4 5]
5




package main
import "fmt"

func main() {  
    var numbers [3] string  
    numbers[0] = "reddy"
    numbers[1] = "sudha"
    numbers[2] = "nayuni"
    fmt.Println(numbers[3]) 
    fmt.Println(len(numbers)) 
    fmt.Println(numbers) 

    directions := [...] int {1,2,3,4,5}  
    fmt.Println(directions) 
    fmt.Println(len(directions))
}
PS C:\Users\Sudha\Desktop\golang>go run len.go;
nayuni
3
[reddy sudha nayuni]
[1 2 3 4 5]
5