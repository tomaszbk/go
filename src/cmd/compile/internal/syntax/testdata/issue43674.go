package p

func _(... /* ERROR [.][.][.] is missing type */ )
func _(... /* ERROR [.][.][.] is missing type */ , int)

func _(a, b ... /* ERROR [.][.][.] is missing type */ )
func _(a, b ... /* ERROR [.][.][.] is missing type */ , x int)

func _()(... /* ERROR [.][.][.] is missing type */ )
