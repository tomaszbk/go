// errorcheck


// Here for reference, but hard to test automatically
// because the BOM muddles the
// processing done by ../run.

package main

func main() {
	﻿// There's a bom here.	// ERROR "BOM"
	//﻿ And here.	// ERROR "BOM"
	/*﻿ And here.*/	// ERROR "BOM"
	println("hi﻿ there") // and here	// ERROR "BOM"
}
