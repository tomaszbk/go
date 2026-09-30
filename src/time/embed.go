// This file is used with build tag timetzdata to embed tzdata into
// the binary.

//go:build timetzdata

package time

import _ "time/tzdata"
