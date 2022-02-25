package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func ReadCmdLine() (string, error) {
	R := bufio.NewReader(os.Stdin)
	Dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	fmt.Printf("GoShell:%v$", Dir)
	input, err := R.ReadString('\n')

	if err != nil {
		return "", err
	}

	return input, nil
}

func Cd(args []string) (string, error) {

	if len(args) == 0 {
		user, err := user.Current()
		if err != nil {
			return "", err
		}
		err = os.Chdir(user.HomeDir)
		if err != nil {
			return "", err
		}

	} else {
		err := os.Chdir(args[0])
		if err != nil {
			return "", err
		}
	}
	return "Directory changed", nil

}

func EchoCmd(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("shell: echo should have some data")
	}
	return args[0], nil
}

func PwdCmd() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return "Current direction :" + path, err
}

func PsCmd() (string, error) {

	var B strings.Builder
	B.WriteString("PID\t|\tCOMMAND\n")
	Prs, err := ps.Processes()
	if err != nil {
		return "", err
	}

	for _, p := range Prs {
		B.WriteString(fmt.Sprintf("%v\t|\t%v\n", p.Pid(), p.Executable()))
	}
	B.WriteString("---------------\n")
	return B.String(), nil
}

func KillCmd(args []string) (string, error) {
	processes, err := process.Processes()
	if err != nil {
		return "", err
	}

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			return "", err
		}
		if name == args[0] {
			err = p.Kill()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return fmt.Sprintf("Process %v successfully killed", args[0]), nil
}

func AllCommands(line string) (string, error) {

	Args := make([]string, 0)
	var Result string
	var err error
	command := strings.Fields(line)
	if len(command) > 1 {
		Args = command[1:]
	}
	switch command[0] {
	case "cd":
		Result, err = Cd(Args)
	case "pwd":
		Result, err = PwdCmd()
	case "echo":
		Result, err = EchoCmd(Args)
	case "kill":
		Result, err = KillCmd(Args)
	case "ps":
		Result, err = PsCmd()
	case "exec":

		CmdExec := exec.Command(command[0], Args...)

		CmdExec.Stdin = os.Stdin
		CmdExec.Stdout = os.Stdout

		err := CmdExec.Run()
		if err != nil {
			return "", err
		}

		Out, err := CmdExec.Output()
		Result = string(Out)

	default:
		fmt.Println("Ошибка! Неизвестная команда")
	}
	return Result, err
}
