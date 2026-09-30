// Ensure there is no "imported and not used" error
// if a package wasn't imported in the first place.

package p

import . "/foo" // ERROR "could not import /foo"
