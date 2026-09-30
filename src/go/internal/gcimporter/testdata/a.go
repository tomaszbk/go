// Input for TestIssue13566

package a

import "encoding/json"

type A struct {
	a    *A
	json json.RawMessage
}
