package exec_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tenntenn/exec"
)

func TestEnv(t *testing.T) {
	t.Parallel()

	type S = []string
	cases := map[string]struct {
		cmds   string
		want   string
		hasErr bool
	}{
		"single":              {"go list fmt", "fmt\n", false},
		"multiple":            {"go list fmt;go list io", "fmt\nio\n", false},
		"single with error":   {"go list _err", "", true},
		"multiple with error": {"go list _err; go list fmt", "", true},
		"before error": {"go list fmt; go list _err", "fmt\n", true},
		"with dir": {"go mod init tmpmod; go list -m", "tmpmod\n", false},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var stdout, stderr bytes.Buffer
			env := &exec.Env{
				Dir:    t.TempDir(),
				Stdout: &stdout,
				Stderr: &stderr,
			}
			cmds := strings.Split(tt.cmds, ";")
			for _, cmd := range cmds {
				cmd = strings.TrimSpace(cmd)
				args := strings.Split(cmd, " ")
				env.Run(args[0], args[1:]...)
			}
			switch {
			case env.Err() != nil && !tt.hasErr:
				t.Fatalf("unexpected error:%v\n\tstderr:%s\tstdout:%s", env.Err(), &stderr, &stdout)
			case env.Err() == nil && tt.hasErr:
				t.Error("expected error have not occured")
			}
			if got := stdout.String(); got != tt.want {
				t.Errorf("want %v but got %v", tt.want, got)
			}
		})
	}
}
