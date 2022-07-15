package tis620_test

import (
	"os"
	"testing"

	"github.com/tupkung/tis620"
)

func TestUTF8(t *testing.T) {
	ut8 := []rune(`ทดสอบภาษาไทย and English 
ทดสอบขึ้นบรรทัดใหม่ด้วย|ทดสอบคั่นหน้า,พวกนี้ แล้วก็ ##!!@##$...`)
	t.Logf("%d", len(ut8))

	t620 := tis620.ToTIS620(ut8)

	os.WriteFile("tis-620-text.txt", t620, 0600)
}

func TestTIS620(t *testing.T) {
	b, err := os.ReadFile("tis-620-text.txt")
	if err != nil {
		t.Fatal(err)
	}

	ut8 := tis620.ToUTF8(b)

	os.WriteFile("utf-8-text.txt", ut8, 0600)
}
