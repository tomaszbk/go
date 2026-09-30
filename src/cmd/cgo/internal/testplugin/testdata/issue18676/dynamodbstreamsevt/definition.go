package dynamodbstreamsevt

import "encoding/json"

var foo json.RawMessage

type Event struct{}

func (e *Event) Dummy() {}
