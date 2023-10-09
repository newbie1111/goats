package utils

/*
Orderable is an interface defines orderable type
*/
type Orderable interface {
	Real | ~string
}

/*
Real is an interface defines Real number (approximated on a computer).
*/
type Real interface {
	Zahl | Float
}

/*
Float is an interface defines decimal fractions (approximated on a computer).
*/
type Float interface {
	~float32 | ~float64
}

/*
Zahl is an interface defines integer (approximated on a computer).
*/
type Zahl interface {
	Integer | NonNegativeInteger
}

/*
Integer is an interface defines signed integer (approximated on a computer).
*/
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

/*
NonNegativeInteger is an interface defines non-negative integer (approximated on a computer).
*/
type NonNegativeInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
