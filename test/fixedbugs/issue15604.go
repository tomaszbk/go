// compile

package bug

import "os"

func f(err error) {
	var ok bool
	if err, ok = err.(*os.PathError); ok {
		if err == os.ErrNotExist {
		}
	}
}
