// errorcheck


package main

func main() {
	var x map[string]string{"a":"b"}		// ERROR "unexpected { at end of statement|expected ';' or '}' or newline"
