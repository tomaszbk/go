// compile -p=p

package v4

var sink interface{}

//go:noinline
func Do(result, body interface{}) {
	sink = &result
}

func DataAction(result DataActionResponse, body DataActionRequest) {
	Do(&result, body)
}

type DataActionRequest struct {
	Action *interface{}
}

type DataActionResponse struct {
	ValidationErrors *ValidationError
}

type ValidationError struct {
}
