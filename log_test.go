package log

import "testing"

func init()  {
	SetPrefix("test")
	SetLogLevel(DEBUG)
}

func TestInfo(t *testing.T) {
	Info("%s","test")
}