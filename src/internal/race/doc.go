/*
Package race contains helper functions for manually instrumenting code for the race detector.

The runtime package intentionally exports these functions only in the race build;
this package exports them unconditionally but without the "race" build tag they are no-ops.
*/
package race
