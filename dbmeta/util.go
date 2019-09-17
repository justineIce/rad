package dbmeta

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"math/big"
	math "math/rand"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

const (
	// Set of characters to use for generating random strings
	Alphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Numerals     = "1234567890"
	Alphanumeric = Alphabet + Numerals
	Ascii        = Alphanumeric + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)

// fmtFieldName formats a string as a struct key
//
// Example:
// 	fmtFieldName("foo_id")
// Output: FooID
func FmtFieldName(s string) string {
	name := lintFieldName(s)
	runes := []rune(name)
	for i, c := range runes {
		ok := unicode.IsLetter(c) || unicode.IsDigit(c)
		if i == 0 {
			ok = unicode.IsLetter(c)
		}
		if !ok {
			runes[i] = '_'
		}
	}
	return string(runes)
}

func FmtFieldName2(s string) string {
	str := FmtFieldName(s)
	return strings.ToLower(str[:1]) + str[1:]
}

func lintFieldName(name string) string {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return name
	}

	for len(name) > 0 && name[0] == '_' {
		name = name[1:]
	}

	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		runes := []rune(name)
		if u := strings.ToUpper(name); commonInitialisms[u] {
			copy(runes[0:], []rune(u))
		} else {
			runes[0] = unicode.ToUpper(runes[0])
		}
		return string(runes)
	}

	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word

		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))

		} else if strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// convert first character ints to strings
func stringifyFirstChar(str string) string {
	first := str[:1]

	i, err := strconv.ParseInt(first, 10, 8)

	if err != nil {
		return str
	}

	return intToWordMap[i] + "_" + str[1:]
}

func Copy(dst interface{}, src interface{}) error {
	dstV := reflect.Indirect(reflect.ValueOf(dst))
	srcV := reflect.Indirect(reflect.ValueOf(src))

	if !dstV.CanAddr() {
		return errors.New("copy to value is unaddressable")
	}

	if srcV.Type() != dstV.Type() {
		return errors.New("different types can be copied")
	}

	for i := 0; i < dstV.NumField(); i++ {
		f := srcV.Field(i)
		if !isZeroOfUnderlyingType(f.Interface()) {
			dstV.Field(i).Set(f)
		}
	}

	return nil
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// global value that performs all random number operations
var globalSource = math.New(&source{})

// math/rand Source using entropy from crypto/rand
type source struct {
	data [8]byte // 64 bits
	mtx  sync.Mutex
}

// Float64 returns a pseudo-random number in [0.0,1.0)
func Float64() float64 { return globalSource.Float64() }

// needed by math/rand.Source, but we don't require seeding
func (src *source) Seed(_ int64) {}

// needed by math/rand.Source
func (src *source) Int63() int64 {
	src.mtx.Lock()
	defer src.mtx.Unlock()

	data := src.data[:]
	n, err := rand.Read(data)
	if err != nil {
		panic("crypto.Read failed: " + err.Error())
	}
	if n != 8 {
		panic("read too few random bytes")
	}
	x := binary.BigEndian.Uint64(data)
	return int64(x >> 1) // need 63 bit number
}

// Intn returns, as an int, a non-negative pseudo-random number in
// [0,n). It panics if n <= 0.
func Intn(n int) int {
	return globalSource.Intn(n)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an
// int64.
func Int63() int64 {
	return globalSource.Int63()
}

// Int63n returns, as an int64, a non-negative pseudo-random number in
// [0,n). It panics if n <= 0.
func Int63n(n int64) int64 {
	return globalSource.Int63n(n)
}

// String returns a random string n characters long, composed of entities
// from charset.
func String(n int, charset string) (string, error) {
	randstr := make([]byte, n) // Random string to return
	charlen := big.NewInt(int64(len(charset)))
	for i := 0; i < n; i++ {
		b, err := rand.Int(rand.Reader, charlen)
		if err != nil {
			return "", err
		}
		r := int(b.Int64())
		randstr[i] = charset[r]
	}
	return string(randstr), nil
}

// AlphaString returns a random alphanumeric string n characters long.
func GetRandomString(n int) (s string) {
	s, _ = String(n, Alphanumeric)
	return
}

func GetRandom(length int) int {
	if length <= 0 {
		return 0
	}
	var arr []string
	for i := 0; i < length; i++ {
		arr = append(arr, strconv.Itoa(Intn(10)))
	}
	return StrToInt(strings.Join(arr, ""))
}

func StrToInt(str string) (i int) {
	i, _ = strconv.Atoi(str)
	if i == 0 {
		i = 1
	}
	if i > 100 {
		return 100
	}
	return
}

func IDString() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return strings.Replace(fmt.Sprintf("%s", id), "-", "", -1)
}
