package s3util_test

import (
	"io"
	"os"

	"github.com/kr/s3"
	"github.com/kr/s3/s3util"
)

func ExampleCreate() {
	keys := &s3.StaticKeys{
		AccessKeyValue: "...access key...",
		SecretKeyValue: "...secret key...",
	}
	s3util.DefaultConfig.Keys = keys
	r, _ := os.Open("/dev/stdin")
	w, _ := s3util.Create("https://mybucket.s3.amazonaws.com/log.txt", nil, nil)
	io.Copy(w, r)
	w.Close()
}

func ExampleOpen() {
	keys := &s3.StaticKeys{
		AccessKeyValue: "...access key...",
		SecretKeyValue: "...secret key...",
	}
	s3util.DefaultConfig.Keys = keys
	r, _ := s3util.Open("https://mybucket.s3.amazonaws.com/log.txt", nil)
	w, _ := os.Create("out.txt")
	io.Copy(w, r)
	w.Close()
}
