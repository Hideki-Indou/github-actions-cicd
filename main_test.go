package main
import "tesing"
func TestGreeting(t *tesing.T) {
	want:="Hello"
	got:=Greeting()
	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}