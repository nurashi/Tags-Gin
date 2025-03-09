package error_handling

func ErrorPanice(err error) {
	if err != nil {
		panic(err)
	}
}