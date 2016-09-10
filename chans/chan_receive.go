package chans

import "github.com/willfaught/goo"

type ChanReceive interface {
	Cap() int

	Len() int

	Make(c int) Chan

	Receive() interface{}

	ReceiveCheck() (interface{}, bool)
}

// TODO: Cycle, Iterate, Repeat

func Equal(c, d ChanReceive) bool {
	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if vc != vd {
			return false
		}
	}

	return true
}

func EqualDeep(c, d ChanReceive) bool {
	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if !goo.Equal(vc, vd) {
			return false
		}
	}

	return true
}

func Iterator(c ChanReceive) goo.Iterator {
	return &iterator{c: c}
}

type iterator struct {
	c ChanReceive

	checked bool

	ok bool

	v interface{}
}

func (i *iterator) More() bool {
	if i.checked {
		return i.ok
	}

	i.v, i.ok = i.c.ReceiveCheck()
	i.checked = true

	return i.ok
}

func (i *iterator) Next() interface{} {
	if !i.checked {
		i.More()
	}

	i.checked = false

	return i.v
}
