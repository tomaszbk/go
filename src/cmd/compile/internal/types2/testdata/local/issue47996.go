package p

// don't crash
func T /* ERROR "missing" */ [P] /* ERROR "missing" */ m /* ERROR "unexpected" */ () /* ERROR ")" */ { /* ERROR "{" */ } /* ERROR "}" */
