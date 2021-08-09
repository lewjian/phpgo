package phpgo

import (
	"fmt"
	"log"
	"testing"
)

func TestExplode(t *testing.T) {
	s := "hello,world,love,peace"
	result := []string{"hello", "world", "love", "peace"}
	cr := Explode(",", s)
	for i, item := range cr {
		if item != result[i] {
			t.Fail()
		}
	}
}

func TestMD5(t *testing.T) {
	s := "hello,world!"
	sign := "c0e84e870874dd37ed0d164c7986f03a"
	if sign != MD5([]byte(s)) {
		t.Fail()
	}
}

func TestSha1(t *testing.T) {
	s := "hello,world!"
	sign := "4518135c05e0706c0a34168996517bb3f28d94b5"
	if sign != Sha1([]byte(s)) {
		t.Fail()
	}
}

func TestStrIReplace(t *testing.T) {
	s := "AbC,i O U"
	searchStr := "abc"
	repl := "OOO"
	res, err := StrIReplace(searchStr, repl, s)
	log.Println(res, err)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	if res != "OOO,i O U" {
		t.Fail()
	}
}

func TestStrPad(t *testing.T) {
	s := "hello"
	padLen := 10
	padStr := "-"
	expect := "-----hello"
	if StrPad(s, padLen, padStr, StrPadLeft) != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, StrPad(s, padLen, padStr, StrPadLeft))
		t.Fail()
	}
	expect = "hello-----"
	if StrPad(s, padLen, padStr, StrPadRight) != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, StrPad(s, padLen, padStr, StrPadRight))
		t.Fail()
	}
	expect = "--hello---"
	if StrPad(s, padLen, padStr, StrPadBoth) != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, StrPad(s, padLen, padStr, StrPadBoth))
		t.Fail()
	}
	padStr = "ab"
	expect = "ababahello"
	if StrPad(s, padLen, padStr, StrPadLeft) != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, StrPad(s, padLen, padStr, StrPadLeft))
		t.Fail()
	}
	expect = "helloababa"
	if StrPad(s, padLen, padStr, StrPadRight) != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, StrPad(s, padLen, padStr, StrPadRight))
		t.Fail()
	}
	expect = "abhelloaba"
	if StrPad(s, padLen, padStr, StrPadBoth) != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, StrPad(s, padLen, padStr, StrPadBoth))
		t.Fail()
	}

}

func TestStrShuffle(t *testing.T) {
	s := "hello"
	log.Println(s, StrShuffle(s))
}

func BenchmarkStrShuffle(b *testing.B) {
	s := "hello"
	for i := 0; i < b.N; i++ {
		StrShuffle(s)
	}
}

func TestStrSplit(t *testing.T) {
	s := "我爱天安门，i iove tiananmen"
	log.Println("原始长度:", len([]rune(s)))
	res := StrSplit(s, 1)
	log.Printf("len=%d, v=%v",len(res), res)
	res = StrSplit(s, 2)
	log.Printf("len=%d, v=%v",len(res), res)
	res = StrSplit(s, 3)
	log.Printf("len=%d, v=%v",len(res), res)
	res = StrSplit(s, 10)
	log.Printf("len=%d, v=%v",len(res), res)
	res = StrSplit(s, 50)
	log.Printf("len=%d, v=%v",len(res), res)
}

func TestStrRev(t *testing.T) {
	s := "我爱天安门"
	expect := "门安天爱我"
	res := StrRev(s)
	log.Println(s, res)
	if res != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, res)
		t.Fail()
	}
	s = "我爱天安门呢"
	expect = "呢门安天爱我"
	res = StrRev(s)
	log.Println(s, res)
	if res != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, res)
		t.Fail()
	}
	s = "我爱天安门呢yes"
	expect = "sey呢门安天爱我"
	res = StrRev(s)
	log.Println(s, res)
	if res != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, res)
		t.Fail()
	}
	s = "我爱天安门呢yess"
	expect = "ssey呢门安天爱我"
	res = StrRev(s)
	log.Println(s, res)
	if res != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, res)
		t.Fail()
	}
}

func TestStrStr(t *testing.T) {
	s := "name@example.com"
	substr := "@"
	expect := "name"
	res := StrStr(s, substr, true)
	log.Println(s, res)
	if res != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, res)
		t.Fail()
	}
	expect = "@example.com"
	res = StrStr(s, substr, false)
	log.Println(s, res)
	if res != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, res)
		t.Fail()
	}
}

func TestSubStr(t *testing.T) {
	s := "name@example.com"
	start := 2
	length := 5
	expect := "me@ex"
	res := SubStr(s, start, length)
	log.Println(s, res)
	if res != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, res)
		t.Fail()
	}
}
func TestSubStrRune(t *testing.T) {
	s := "我爱你"
	start := 1
	length := 2
	expect := "爱你"
	res := SubStrRune(s, start, length)
	log.Println(s, res)
	if res != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, res)
		t.Fail()
	}
}

func TestUCWords(t *testing.T) {
	s := "hello, world"
	expect := "Hello, World"
	res := UCWords(s)
	log.Println(s, res)
	if res != expect {
		fmt.Sprintf("s:%s, expect:%s, get:%s", s, expect, res)
		t.Fail()
	}
}