// errorcheck


package p

func _ () {
	if {} // ERROR "missing condition in if statement"

	if
	{} // ERROR "missing condition in if statement"

	if ; {} // ERROR "missing condition in if statement"

	if foo; {} // ERROR "missing condition in if statement"

	if foo; // ERROR "missing condition in if statement"
	{}

	if foo {}

	if ; foo {}

	if foo // ERROR "unexpected newline, expected { after if clause"
	{}
}
