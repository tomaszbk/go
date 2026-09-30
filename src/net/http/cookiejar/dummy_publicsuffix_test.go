package cookiejar_test

import "net/http/cookiejar"

type dummypsl struct {
	List cookiejar.PublicSuffixList
}

func (dummypsl) PublicSuffix(domain string) string {
	return domain
}

func (dummypsl) String() string {
	return "dummy"
}

var publicsuffix = dummypsl{}
