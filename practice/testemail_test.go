package practice

import(
	"testing"
)

func TestIsEmail(t *testing.T){
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}
}