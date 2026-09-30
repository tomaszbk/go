//go:build !osusergo && darwin

package user

import (
	"internal/syscall/unix"
)

func getGroupList(name *_C_char, userGID _C_gid_t, gids *_C_gid_t, n *_C_int) _C_int {
	err := unix.Getgrouplist(name, userGID, gids, n)
	if err != nil {
		return -1
	}
	return 0
}
