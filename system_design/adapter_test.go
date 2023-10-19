package system_design

import (
	"fmt"
	"testing"
)

// TriplePin 三孔接口
type TriplePin interface {
	electrify(a, b, c int) error
}

// 两孔接口
type DualPin interface {
	electrify(a, b int) error
}

// TV是一个两孔接口
type TV struct{}

func (T TV) electrify(a, b int) error {
	fmt.Printf("电视的零线与火线, a:[%v], b:[%v]", a, b)
	return nil
}

type AdapterTriple struct {
	adaptee DualPin
}

func (a2 AdapterTriple) electrify(a, b, c int) error {
	a2.adaptee.electrify(a, b)
	return nil
}

func TestAdapter(t *testing.T) {
	// 创建TV
	tv := TV{}

	// 三孔插头
	var triple TriplePin

	// 两孔电视机连不上三孔插头，会报错
	// triple = tv

	// 将电视插在适配器上，然后连上三孔插头
	triple = AdapterTriple{tv}

	// 通电
	triple.electrify(1, 2, 3)
}
