//go:build android

package user

import "errors"

func lookupUser(string) (*User, error) {
	return nil, errors.New("user: Lookup not implemented on android")
}

func lookupUserId(string) (*User, error) {
	return nil, errors.New("user: LookupId not implemented on android")
}

func lookupGroup(string) (*Group, error) {
	return nil, errors.New("user: LookupGroup not implemented on android")
}

func lookupGroupId(string) (*Group, error) {
	return nil, errors.New("user: LookupGroupId not implemented on android")
}
