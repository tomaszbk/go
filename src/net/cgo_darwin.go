package net

import "internal/syscall/unix"

const cgoAddrInfoFlags = (unix.AI_CANONNAME | unix.AI_V4MAPPED | unix.AI_ALL) & unix.AI_MASK
