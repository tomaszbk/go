package noresultvalues

func x() { return nil } // want `no result values expected|too many return values`

func y() { return nil, "hello" } // want `no result values expected|too many return values`
