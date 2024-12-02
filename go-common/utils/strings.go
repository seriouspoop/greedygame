package utils

type Stringer[V any] interface {
	String() string
}

func StringSlice[V Stringer[V]](s []V) (strSlice []string) {
	for _, val := range s {
		strSlice = append(strSlice, val.String())
	}
	return
}
