package greetings
import (
	"testing"
	"regexp"
)

//check Greetings function
func TestGreetings(t *testing.T){
name:="Rohan"
want:=regexp.MustCompile(`\b`+name+`\b`);
msg,err := Greetings(name)
if !want.MatchString(msg) || err!=nil {
	t.Errorf(`Hello ("Rohan") = %q %v, want match for %#q,nil`,msg,err,want)
}
}