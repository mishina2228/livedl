package streamlink

import (
	"../base"
	"fmt"
	"io"
	"os/exec"
)

var cmdList = []string{
	"./bin/streamlink/streamlink",
	"./bin/Streamlink/Streamlink",
	"./bin/streamlink",
	"./bin/Streamlink",
	"./streamlink/streamlink",
	"./Streamlink/Streamlink",
	"./Streamlink",
	"streamlink",
	"Streamlink",
}

func Open(opt ...string) (cmd *exec.Cmd, stdout, stderr io.ReadCloser, err error) {
	cmd, _, stdout, stderr, err = base.Open(&cmdList, false, true, true, false, opt)
	if cmd == nil {
		err = fmt.Errorf("streamlink not found")
		return
	}
	return
}
