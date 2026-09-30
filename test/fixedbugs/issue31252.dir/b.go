package b

import "fmt"

type IndexController struct{}

func (this *IndexController) Index(m *string) {
	fmt.Println(m)
}
