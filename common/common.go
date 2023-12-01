package common

func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
