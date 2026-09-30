//go:build android

package user

import (
	"errors"
)

func init() {
	groupListImplemented = false
}

func listGroups(*User) ([]string, error) {
	return nil, errors.New("user: list groups not implemented")
}
