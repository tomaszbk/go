// errorcheck


package p

type a struct {
  a int
}

func main() {
  av := a{};
  _ = *a(av); // ERROR "invalid indirect|expected pointer|cannot indirect"
}
