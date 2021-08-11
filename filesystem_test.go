package phpgo

import (
	"log"
	"testing"
)

func TestBasename(t *testing.T) {
	filename := "/etc/sudoers.d"
	suffix := ".d"
	expect := "sudoers"
	res := Basename(filename, suffix)
	if res != expect {
		log.Printf("filename:%s, suffix:%s, expect:%s, result:%s", filename, suffix, expect, res)
		t.Fail()
	}

	filename = "/etc/sudoers.d"
	suffix = ""
	expect = "sudoers.d"
	res = Basename(filename, suffix)
	log.Printf("filename:%s, suffix:%s, expect:%s, result:%s", filename, suffix, expect, res)
	if res != expect {
		log.Printf("filename:%s, suffix:%s, expect:%s, result:%s", filename, suffix, expect, res)
		t.Fail()
	}

	filename = "/etc/passwd"
	suffix = ""
	expect = "passwd"
	res = Basename(filename, suffix)
	log.Printf("filename:%s, suffix:%s, expect:%s, result:%s", filename, suffix, expect, res)
	if res != expect {
		log.Printf("filename:%s, suffix:%s, expect:%s, result:%s", filename, suffix, expect, res)
		t.Fail()
	}

	filename = "/etc/"
	suffix = ""
	expect = "etc"
	res = Basename(filename, suffix)
	log.Printf("filename:%s, suffix:%s, expect:%s, result:%s", filename, suffix, expect, res)
	if res != expect {
		log.Printf("filename:%s, suffix:%s, expect:%s, result:%s", filename, suffix, expect, res)
		t.Fail()
	}
	filename = "."
	suffix = ""
	expect = "."
	res = Basename(filename, suffix)
	log.Printf("filename:%s, suffix:%s, expect:%s, result:%s", filename, suffix, expect, res)
	if res != expect {
		log.Printf("filename:%s, suffix:%s, expect:%s, result:%s", filename, suffix, expect, res)
		t.Fail()
	}

}

func TestDirname(t *testing.T) {
	filename := "/etc/passwd"
	expect := "/etc"
	res := Dirname(filename)
	if res != expect {
		log.Printf("filename:%s,  expect:%s, result:%s", filename, expect, res)
		t.Fail()
	}

	filename = "/etc"
	expect = "/"
	res = Dirname(filename)
	if res != expect {
		log.Printf("filename:%s,  expect:%s, result:%s", filename, expect, res)
		t.Fail()
	}
}

func TestReadDir(t *testing.T) {
	dir := "c:\\"
	set,_ := ReadDir(dir)
	for _, item := range set {
		log.Println(item.Name())
		log.Println(item.Info())
	}
}