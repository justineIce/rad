package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Random random
type Random struct{}

// NewRandom new random
func NewRandom() *Random {
	return &Random{}
}

// Number random number
func (rd Random) Number(length float64) int {
	i := int(math.Pow(10, length-1))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(9*i) + i
}

// String random string
func (rd Random) String(length int) string {
	b := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	d := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		d = append(d, b[r.Intn(62)])
	}
	return string(d)
}

func (rd Random) NumberByInt(start, end int) int {
	r := rand.New(rand.NewSource(ID()))
	return start + r.Intn(end-start+1)
}

//随机数
func GetRand() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%04v", rnd.Int31n(10000))
}
