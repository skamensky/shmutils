package command

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
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
	cmd              *exec.Cmd
	killRequest      chan error
}

func (c *Command) Pid() int {
	if c.cmd.Process != nil {
		return c.cmd.Process.Pid
	}
	return 0
}

func (c *Command) doesProcessExist() (bool, error) {
	if c.cmd.Process == nil {
		return false, errors.New("command hasn't run yet")
	}
	if c.Pid() == 0 {
		return false, errors.New("command hasn't run yet")
	}
	// from the docs: On Unix systems, FindProcess always succeeds and returns a Process
	// so we can't use os.FindProcess(c.Pid())
	proc, err := os.FindProcess(c.Pid())
	if err != nil {
		return false, fmt.Errorf("could not find process: %w", err)
	}
	// check if process exists
	err = proc.Signal(syscall.Signal(0))
	return err == nil, nil
}

func (c *Command) doKill(timeout int) {
	// send os.Interrupt, if after timeout milliseconds, the proc is still alive, run c.cmd.Process.Kill()
	proc, err := os.FindProcess(c.Pid())
	if err != nil {
		c.killRequest <- fmt.Errorf("could not find proc: %w", err)
		return
	}
	interruptResult := proc.Signal(os.Interrupt)
	if interruptResult != nil {
		c.killRequest <- fmt.Errorf("failed to send interrupt signal to proc: %w", interruptResult)
		return
	}
	start := time.Now()
	ticker := time.NewTicker(100)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			processExists, err := c.doesProcessExist()
			if err != nil {
				c.killRequest <- fmt.Errorf("could not check if process exists: %w", err)
				return
			}
			if !processExists {
				c.killRequest <- nil
				return
			}
		}
		if time.Since(start) > time.Duration(timeout)*time.Millisecond {
			break
		}
	}
	errorMessage := fmt.Sprintf("proc did not exit after %d milliseconds, used Process.Kill to terminate it", timeout)
	killErr := c.cmd.Process.Kill()
	if killErr != nil {
		c.killRequest <- fmt.Errorf(errorMessage+", and then an additional error occured: %w", killErr)
		return
	}
	c.killRequest <- errors.New(errorMessage)
	return
}

/*
Kill sends an SIGINT signal to the process, and if after timeout milliseconds,
the process is still alive, kills the process using SIGKILL.
*/
func (c *Command) Kill(timeout int) error {
	if c.cmd == nil || c.cmd.Process == nil {
		return errors.New("command hasn't run yet")
	}
	processExists, err := c.doesProcessExist()
	if err != nil {
		return fmt.Errorf("could not check if process exists: %w", err)
	}
	if !processExists {
		return nil
	}
	c.killRequest = make(chan error, 1)
	go c.doKill(timeout)
	return <-c.killRequest
}

func New(executable string, args ...string) *Command {
	argsAsArray := make([]string, len(args))
	copy(argsAsArray, args)
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

func (c *Command) WithVerbose(verbose bool) *Command {
	c.verbose = verbose
	return c
}

func (c *Command) WithTreatStderrAsErr(treatStderAsErr bool) *Command {
	c.treatStderrAsErr = treatStderAsErr
	return c
}

func (c *Command) WithDir(dir string) *Command {
	c.dir = dir
	return c
}

func (c *Command) WithAttachToTerminal(attach bool) *Command {
	c.attachToTerminal = attach
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
	c.cmd = exec.Command(c.executable, c.args...)

	if c.attachToTerminal {
		c.cmd.Stdout = os.Stdout
		// what should we do with stderr?
		//cmd.Stderr = os.Stderr
		c.cmd.Stdin = os.Stdin
	} else {
		c.cmd.Stdout = &c.Result.Stdout
		c.cmd.Stderr = &c.Result.Stderr
	}

	if c.dir != "" {
		c.cmd.Dir = c.dir
	}

	ticker := time.NewTicker(time.Duration(500) * time.Millisecond)
	defer ticker.Stop()
	done := make(chan int, 1)
	go func() {
		c.Result.Err = c.cmd.Run()
		done <- 0
	}()
	// wait for the command to finish, or for the kill request to be fulfilled
	for {
		select {
		case <-ticker.C:
			if c.killRequest != nil && len(c.killRequest) == 1 {
				break
			}
		}
		if len(done) == 1 {
			break
		}
	}

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
