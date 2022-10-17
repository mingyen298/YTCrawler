package test

import (
	"fmt"
	"testing"
	"time"
	"yt_crawler/decrypt"
)

func Speed(f func()) {
	s := time.Now()
	f()
	fmt.Println(time.Since(s))
}
func TestProccess1(t *testing.T) {
	yt := decrypt.VideoDecrypter{}

	raw := "Z=gXbsqCBBs3awYhAFD=9crBYVtSZGC0-YJQZtAvnV2abCQICYQ-u5ATrZMDuITNylZ9_7LqA_hCTQMx1mO1hhadDnNIgIQRw8JQ0AObAOb"
	out := yt.Proccess(&raw)
	if out != "" {
		t.Log(out)
	} else {
		t.Error(out)
	}

}
