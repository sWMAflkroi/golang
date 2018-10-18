package rationalNumbers

type Rational struct {
	num int
	den int
}

func (x Rational) Multiply(y Rational) Rational {
	return NewRational(x.num*y.den, x.den*y.num)
}

func NewRational(num int, den int) Rational {
	if den == 0 {
		panic("divison by zero is not allowed!")
	}
	r := Rational{}
	divisor := gcd(num, den)
	r.num = num / divisor
	r.den = den / divisor
	return r
}
