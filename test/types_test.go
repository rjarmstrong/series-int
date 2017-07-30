package test

import (
	"fmt"
	"github.com/bmizerany/assert"
	"github.com/rjarmstrong/series-int/series"
	"math/big"
	"testing"
)

func Test_SeriesInt(t *testing.T) {
	i := series.NewInt16(60)
	i.Set(0, 512)
	i.Set(1, 255)
	i.Set(30, 200)
	assert.Equal(t, uint16(512), i.Val(0))
	fmt.Println(i)
	j := series.NewInt16(60)
	j.Set(0, 666)
	j.Set(1, 1)
	j.Set(30, 111)
	j.Set(40, 65535)
	j.Set(59, 123)
	fmt.Println(j)
	i.Add(&j)
	fmt.Println(i)
	assert.Equal(t, uint16(512+666), i.Val(0))
	i.Incr(0)
	i.Incr(0)
	i.Incr(30)
	i.Incr(31)
	fmt.Println(i)
	i.AddRange(1, 10, 513)
	fmt.Println(i)
}

func BenchmarkSeriesInt16_Add(b *testing.B) {
	i := series.NewInt16(60)
	i.Set(0, 512)
	j := series.NewInt16(60)
	j.Set(0, 666)
	for n := 0; n < b.N; n++ {
		i.Add(&j)
	}
}

func MonthsAsBigNum() *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(60*4), nil)
}

func BenchmarkBigNum_Add(b *testing.B) {
	n1 := MonthsAsBigNum()
	base := MonthsAsBigNum()
	n1.Add(n1, new(big.Int).Exp(big.NewInt(512), big.NewInt(59*4), nil))
	n2 := MonthsAsBigNum()
	n2.Add(n2, new(big.Int).Exp(big.NewInt(666), big.NewInt(59*4), nil))
	for n := 0; n < b.N; n++ {
		n1.Add(n1, n2)
		n1.Sub(n1, base)
	}
}

func BenchmarkSeriesInt16_Get(b *testing.B) {
	i := series.NewInt16(60)
	i.Set(0, 512)
	for n := 0; n < b.N; n++ {
		i.Val(0)
	}
}

func BenchmarkBigNum_Get(b *testing.B) {
	n1 := MonthsAsBigNum()
	n1.Add(n1, new(big.Int).Exp(big.NewInt(512), big.NewInt(59*4), nil))
	for n := 0; n < b.N; n++ {
		n1.String()
	}
}
