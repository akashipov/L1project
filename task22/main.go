package main

import "fmt"

type LargeInt struct {
	Values [3]byte
}

func (v LargeInt) Add(other LargeInt) *LargeInt {
	k := 0
	if other.Values[0]>>7 == 1 && v.Values[0]>>7 == 1 {
		other.Values[0] &^= (1 << 7)
		v.Values[0] &^= (1 << 7)
		k := v.Add(other)
		k.Values[0] |= (1 << 7)
		return k
	} else if other.Values[0]>>7 == 1 && v.Values[0]>>7 == 0 {
		other.Values[0] &^= (1 << 7)
		return v.Minus(other)
	} else if other.Values[0]>>7 == 0 && v.Values[0]>>7 == 0 {
		for i := len(v.Values) - 1; i >= 0; i-- {
			k = int(v.Values[i]) + int(other.Values[i]) + int(k)
			var d int
			if i == 0 {
				d = 128
			} else {
				d = 256
			}
			g := k % d
			v.Values[i] = byte(g)
			k = (k - g) / d
		}
		if k != 0 {
			fmt.Println("Too large numeric")
			return nil
		}
		return &v
	} else {
		return other.Add(v)
	}
}

func (v LargeInt) Less(other LargeInt) bool {
	if other.Values[0]>>7 == 0 && v.Values[0]>>7 == 1 {
		return true
	} else if other.Values[0]>>7 == 1 && v.Values[0]>>7 == 0 {
		return false
	} else if other.Values[0]>>7 == 1 && v.Values[0]>>7 == 1 {
		other.Values[0] &^= 1 << 7
		v.Values[0] &^= 1 << 7
		return !v.Less(other)
	} else {
		for i := 0; i < len(v.Values); i++ {
			if v.Values[i] < other.Values[i] {
				return true
			} else if v.Values[i] > other.Values[i] {
				return false
			}
		}
		// fmt.Println("There are equal")
		return false
	}
}

func (v LargeInt) Minus(other LargeInt) *LargeInt {
	if v.Less(other) {
		o := other.Minus(v)
		if o == nil {
			return nil
		}
		o.Values[0] |= (1 << 7)
		return o
	}
	if other.Values[0]>>7 == 0 && v.Values[0]>>7 == 1 {
		other.Values[0] |= (1 << 7)
		return v.Add(other)
	} else if other.Values[0]>>7 == 1 && v.Values[0]>>7 == 0 {
		other.Values[0] &^= (1 << 7)
		return v.Add(other)
	} else if other.Values[0]>>7 == 0 && v.Values[0]>>7 == 0 {
		k := 0
		for i := len(v.Values) - 1; i >= 0; i-- {
			g := int(v.Values[i]) - int(other.Values[i]) - k
			if g < 0 {
				g += 256
				k = 1
			} else {
				k = 0
			}
			v.Values[i] = byte(g)
		}
		return &v
	} else {
		return other.Minus(v)
	}
}

func (v LargeInt) MultiplyBy2Degree(degree int) *LargeInt {
	wholeCountBytes := degree / 8
	// leftPart=5 11111111 -> {00011111} {11100000}
	leftPart := degree % 8
	var newK, nextK int8
	var newI, nextI int
	var l int16
	for i := 0; i < len(v.Values); i++ {
		l = int16(v.Values[i]) << leftPart
		newK = int8(l & 255)
		nextK = int8((l >> 8) & 255)
		newI = i - wholeCountBytes
		nextI = newI - 1
		if (newI < 0 && newK > 0) || (nextI < 0 && nextK > 0) {
			fmt.Println("Too much to multiply")
			return nil
		}
		v.Values[i] = 0
		if nextI >= 0 {
			v.Values[nextI] |= byte(nextK)
		}
		if newI >= 0 {
			v.Values[newI] |= byte(newK)
		}
	}
	return &v
}

func (v LargeInt) DivideBy2Degree(degree int) *LargeInt {
	wholeCountBytes := degree / 8
	// leftPart=5 11111111 -> {00011111} {11100000}
	leftPart := degree % 8
	var newK, nextK int8
	var newI, nextI int
	var l int16
	for i := len(v.Values) - 1; i >= 0; i-- {
		l = int16(v.Values[i]) << (8 - leftPart)
		newK = int8((l >> 8) & 255)
		nextK = int8(l & 255)
		newI = i + wholeCountBytes
		nextI = newI + 1
		v.Values[i] = 0
		if nextI >= 0 && nextI <= len(v.Values)-1 {
			v.Values[nextI] |= byte(nextK)
		}
		if newI >= 0 && newI <= len(v.Values)-1 {
			v.Values[newI] |= byte(newK)
		}
	}
	return &v
}

func (v LargeInt) Multiply(other LargeInt) *LargeInt {
	var s LargeInt
	var degree int
	for i := 0; i < len(other.Values); i++ {
		for j := 7; j >= 0; j-- {
			g := (int(other.Values[i]) >> j) & 1
			if g == 1 {
				degree = (len(other.Values)-i-1)*8 + j
			} else {
				continue
			}
			fmt.Println("degree:", degree)
			l := v.MultiplyBy2Degree(degree)
			if l == nil {
				return nil
			}
			o := s.Add(*l)
			if o == nil {
				return nil
			}
			s = *o
		}
	}
	return &s
}

func (v LargeInt) Divide(other LargeInt) *LargeInt {
	r := LargeInt{}
	a := LargeInt{}
	b := [len(v.Values)]byte{}
	b[len(b)-1] = 1
	one := LargeInt{b}
	for i := 0; i < (len(v.Values) * 8); i++ {
		r = *r.MultiplyBy2Degree(1)
		a = *a.MultiplyBy2Degree(1)
		if (v.DivideBy2Degree(len(v.Values)*8 - i - 1).Values[len(v.Values)-1] & 1) == 1 {
			a = *a.Add(one)
		}
		if a.Less(other) {
			continue
		}
		a = *a.Minus(other)
		r = *r.Add(one)
	}
	return &r
}

func main() {
	a := LargeInt{[3]byte{0, 1, 36}}
	fmt.Printf("%08b\n", a.Values)
	// b := LargeInt{[3]byte{0, 0, 3}}
	l := a.Divide(LargeInt{[3]byte{0, 0, 7}})
	if l == nil {
		return
	}
	fmt.Printf("%08b\n", (*l).Values)
	// fmt.Printf("%08b", a.Values)
}
