package iteration

func Repeat(text string, times int) string {
	var repeater string
	for i := 0; i < times; i++ {
		repeater += text
	}
	return repeater
}
