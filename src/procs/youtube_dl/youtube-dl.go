package youtube_dl

import (
	"../base"
	"fmt"
	"io"
	"os/exec"
)

var cmdList = []string{
	"./bin/youtube-dl/youtube-dl",
	"./bin/youtube-dl",
	"./youtube-dl/youtube-dl",
	"./youtube-dl",
	"youtube-dl",
}

func Open(opt ...string) (cmd *exec.Cmd, stdout, stderr io.ReadCloser, err error) {
	cmd, _, stdout, stderr, err = base.Open(&cmdList, false, true, true, false, opt)
	if cmd == nil {
		err = fmt.Errorf("youtube-dl not found")
		return
	}
	return
}
