package pkg1

var x = make(chan interface{})

func Do() int {
	return (<-x).(int)
}
