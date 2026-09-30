// errorcheckdir


// The gccgo compiler had a bug: if one file in a package did a dot
// import, then an earlier file in the package would incorrectly
// resolve to the imported names rather than reporting undefined
// errors.

package ignored
