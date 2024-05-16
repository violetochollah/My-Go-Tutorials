package main
import (
"strings"
"testing"
)

//ReverseWords takes a string as input and returns the string with each word ReverseWords func ReverseWords(input string) string {
words :=strings.Fields(input)
reversed := make([]string,len(words))
for _, char := range word  {
reverseWord =string(char) +reverseWord
}
reversed[i] =reverseWord
}
return strings.join(reversed,"
")
}
//Test ReverseWords tests the ReverseWords function
func TestReverseWords(t*testing.T) {
Testcase :=[]struct {
input        string
expectedOutput  string
}{
{"Hello, world!", "olleH,! dlrow"},
{"Go programming", "OGgnimmargorp"},
{"Ilove GO,!," "I  evol !OG"},
}
For _, tc:= range testcases {output :+
ReverseWords(tc.input) if output != tc.expectedOutput
{
t.Errorf("ReverseWords(%s) =%s, expected %s",tc.input,output)
   }
  }
}
