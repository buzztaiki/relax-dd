package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetStat(t *testing.T) {
	blockDevice := func() string {
		xs, err := os.ReadDir("/dev/block")
		if err != nil {
			t.Fatal(err)
		}
		return filepath.Join("/dev/block", xs[0].Name())
	}

	cases := []struct {
		name string
		path string
		want *stat
	}{
		{"exists", "main.go", &stat{"main.go", true, false}},
		{"not exists", "not_exists", &stat{"not_exists", false, false}},
		{"block device", blockDevice(), &stat{blockDevice(), true, true}},
	}

	for _, c := range cases {
		got := getStat(c.path)
		t.Run(c.name, func(t *testing.T) {
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("-want +got:\n%s", diff)
			}
		})
	}
}

func TestRunCmd(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	if err := runCmd(exec.Command("echo", "-n", "moo"), buf); err != nil {
		t.Fatal(err)
	}
	if buf.String() != "moo" {
		t.Errorf("got %q, want %q", buf.String(), "moo")
	}
}

func TestAskOk(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"match", "ok", true},
		{"not match", "ng", false},
		{"empty", "", false},
		{"trim", " ok ", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := askOk("", "ok", bytes.NewReader([]byte(c.input+"\n")))
			if got != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestMakeDDArgs(t *testing.T) {
	cases := []struct {
		name     string
		src, dst *stat
		extArgs  []string
		want     []string
	}{
		{"normal", &stat{"src", true, false}, &stat{"dst", true, false}, nil,
			[]string{"if=src", "of=dst", "bs=4M", "status=progress"}},
		{"block", &stat{"src", true, false}, &stat{"dst", true, true}, nil,
			[]string{"if=src", "of=dst", "bs=4M", "status=progress", "conv=fsync", "oflag=direct"}},
		{"extArgs", &stat{"src", true, false}, &stat{"dst", true, false}, []string{"iflag=sync"},
			[]string{"if=src", "of=dst", "bs=4M", "status=progress", "iflag=sync"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := makeDDArgs(c.src, c.dst, c.extArgs)
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("-got +want:\n%s", diff)
			}
		})
	}
}
