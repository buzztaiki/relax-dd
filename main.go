package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s <src> <dst> [<dd-args>...]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	src := getStat(flag.Arg(0))
	dst := getStat(flag.Arg(1))
	extArgs := flag.Args()[2:]

	if err := run(src, dst, extArgs); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

func run(src, dst *stat, extArgs []string) error {
	fmt.Println("src:")
	if err := src.inspect(); err != nil {
		return fmt.Errorf("inspect error: src=%s: %w", src.Path, err)
	}
	fmt.Println("dst:")
	if err := dst.inspect(); err != nil {
		return fmt.Errorf("inspect error: dst=%s: %w", dst.Path, err)
	}

	ddArgs := makeDDArgs(src, dst, extArgs)
	dd, err := exec.LookPath("dd")
	if err != nil {
		return err
	}

	fmt.Printf("execute the follwing command:\n")
	fmt.Printf("  %s %s\n", dd, strings.Join(ddArgs, " "))
	if !askOk("ok? (yes/NO) ", "yes", os.Stdin) {
		return nil
	}

	syscall.Exec(dd, append([]string{dd}, ddArgs...), nil)
	return nil
}

type stat struct {
	Path        string
	Exists      bool
	BlockDevice bool
}

func getStat(path string) *stat {
	st := syscall.Stat_t{}
	if err := syscall.Stat(path, &st); err != nil {
		return &stat{path, false, false}
	} else {
		return &stat{path, true, st.Mode&syscall.S_IFBLK != 0}
	}
}

func (st *stat) inspect() error {
	if !st.Exists {
		fmt.Printf("%s: not exists\n", st.Path)
	} else if st.BlockDevice {
		if err := runCmd(exec.Command("lsblk", "--fs", st.Path), os.Stdout); err != nil {
			return fmt.Errorf("lsblk error: %w", err)
		}
	} else {
		if err := runCmd(exec.Command("file", st.Path), os.Stdout); err != nil {
			return fmt.Errorf("file error: %w", err)
		}
	}
	fmt.Println()
	return nil
}

func runCmd(cmd *exec.Cmd, w io.Writer) error {
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}

func makeDDArgs(src, dst *stat, extArgs []string) []string {
	ddArgs := []string{fmt.Sprintf("if=%s", src.Path), fmt.Sprintf("of=%s", dst.Path), "bs=4M", "status=progress"}
	if dst.BlockDevice {
		ddArgs = append(ddArgs, "conv=fsync", "oflag=direct")
	}
	return append(ddArgs, extArgs...)
}

func askOk(prompt string, ok string, r io.Reader) bool {
	fmt.Print(prompt)
	line, err := bufio.NewReader(r).ReadString('\n')
	if err != nil {
		return false
	}
	return strings.TrimSpace(line) == ok
}
