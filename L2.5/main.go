package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	customErr := test()

	if customErr != nil { // проверим нашу переменную на nil
		var err error = customErr
		println(err.Error())
		return
	}

	println("ok")
}
