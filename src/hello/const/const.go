package solution

const (
	OK = 0
	CANCELLED = 1
	UNKNOWN = 2
)

func ErrorMessageToCode(msg string) int {
	var resCode int
	switch(msg) {
		case "OK":
			resCode = OK
		case "CANCELLED":
			resCode = CANCELLED
		default:
			resCode = UNKNOWN
	}

	return resCode

}
