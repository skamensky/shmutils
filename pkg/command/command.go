package command

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Result struct {
	Stdout bytes.Buffer
	Stderr bytes.Buffer
	Err    error
}

type Command struct {
	*Result
	executable       string
	args             []string
	verbose          bool
	treatStderrAsErr bool
	dir              string
	attachToTerminal bool
}

func NewCommand(executable string, args ...string) *Command {
	argsAsArray := make([]string, len(args))
	for i, arg := range args {
		argsAsArray[i] = arg
	}
	c := &Command{
		executable:       executable,
		args:             argsAsArray,
		verbose:          true,
		treatStderrAsErr: false,
		attachToTerminal: false,
	}
	c.Result = &Result{
		Stdout: bytes.Buffer{},
		Stderr: bytes.Buffer{},
		Err:    nil,
	}
	return c
}

func (c *Command) WithVerbose() *Command {
	c.verbose = true
	return c
}

func (c *Command) WithTreatStderrAsErr() *Command {
	c.treatStderrAsErr = true
	return c
}

func (c *Command) WithDir(dir string) *Command {
	c.dir = dir
	return c
}

func (c *Command) WithAttachToTerminal() *Command {
	c.attachToTerminal = true
	return c
}

func (c *Command) Run() *Result {
	if c.verbose {
		if c.dir == "" {
			fmt.Printf("Running command: %v %v\n", c.executable, strings.Join(c.args, " "))
		} else {
			fmt.Printf("Running command (cwd=%s) : %v %v\n", c.dir, c.executable, strings.Join(c.args, " "))
		}
	}
	cmd := exec.Command(c.executable, c.args...)
	if c.attachToTerminal {
		cmd.Stdout = os.Stdout
		// what should we do with stderr?
		//cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	} else {
		cmd.Stdout = &c.Result.Stdout
		cmd.Stderr = &c.Result.Stderr
	}

	if c.dir != "" {
		cmd.Dir = c.dir
	}

	c.Result.Err = cmd.Run()
	if c.treatStderrAsErr && len(c.Result.Stderr.Bytes()) != 0 {
		if c.Result.Err != nil {
			c.Result.Err = errors.New(fmt.Sprintf("running the command returned an error: {%s}. Additionally, Stderr is not empty: %s", c.Result.Err.Error(), string(c.Stderr.Bytes())))
		} else {
			c.Result.Err = c.stderrAsErr(c.Stderr)
		}
	}
	return c.Result
}

func (c *Command) stderrAsErr(stderr bytes.Buffer) error {
	return errors.New(fmt.Sprintf("Stderr from command is not empty: %v", string(stderr.Bytes())))
}
