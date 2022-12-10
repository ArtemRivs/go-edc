import "fmt"

func mypanic() {
    var slice []string
    fmt.Println(slice[0])
}

func main() {
    mypanic()
} 
