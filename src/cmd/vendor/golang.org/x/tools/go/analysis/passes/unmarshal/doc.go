// The unmarshal package defines an Analyzer that checks for passing
// non-pointer or non-interface types to unmarshal and decode functions.
//
// # Analyzer unmarshal
//
// unmarshal: report passing non-pointer or non-interface values to unmarshal
//
// The unmarshal analysis reports calls to functions such as json.Unmarshal
// in which the argument type is not a pointer or an interface.
package unmarshal
