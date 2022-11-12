package solution

const (
	OK = 0
	CANCELLED = 1
	UNKNOWN = 2
)

const (
    Black = iota
    Gray
    White
)

// счётчик обнуляется
const (
    Yellow = iota
    Red
    Green = iota // это присваивание не обнулит iota
    Blue
)

const (
    one = iota*2 + 1 // укажите здесь формулу с iota
    three
    five
    seven
    nine
    eleven
)


func ErrorMessageToCode(msg string) int {
	var resCode = UNKNOWN
	switch(msg) {
		case "OK":
			resCode = OK
		case "CANCELLED":
			resCode = CANCELLED
	}

	return resCode

}


func main() {
	fmt.Println(Black, Gray, White) 
	fmt.Println(Yellow, Red, Green, Blue)
	fmt.Println(one, three, five, seven, nine, eleven)
} 
