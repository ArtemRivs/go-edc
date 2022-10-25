package solution

const (
	OK = 0
	CANCELLED = 1
	UNKNOWN = 2
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
