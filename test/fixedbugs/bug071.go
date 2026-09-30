// compile


package bug071

type rat struct  {
	den  int;
}

func (u *rat) pr() {
}

type dch struct {
	dat chan  *rat;
}

func dosplit(in *dch){
	dat := <-in.dat;
	_ = dat;
}
