package routebsd

import (
	"syscall"
	"testing"
)

func TestFetchRIBMessagesOnDarwin(t *testing.T) {
	for _, typ := range []int{syscall.NET_RT_FLAGS, syscall.NET_RT_DUMP2, syscall.NET_RT_IFLIST2} {
		ms, err := FetchRIBMessages(typ, 0)
		if err != nil {
			t.Error(typ, err)
			continue
		}
		ss, err := msgs(ms).validate()
		if err != nil {
			t.Error(typ, err)
			continue
		}
		for _, s := range ss {
			t.Log(s)
		}
	}
}
