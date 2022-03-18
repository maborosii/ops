package apps

import (
	"fmt"
	"testing"
)

func TestNewApp(t *testing.T) {
	s := NewApp(WithAppName("a"), WithAppNewTag("b"))
	fmt.Println(s)
	// t.Errorf("info:", s)
}
