package Util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"

	"github.com/0xAX/notificator"
	"github.com/Splizard/go-espeak/espeak"
)

func RunCommandOnBash(script string) {
	cmd := exec.Command("/bin/bash", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func RunCommandOnBashReturnResult(script string) string {
	var cmdOut []byte
	var err error
	if cmdOut, err = exec.Command("/bin/bash", "-c", script).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	return clearEnterKeyInLast(cmdOut)
}

// clearEnterKeyInLast : clearing enter key in out's end.
func clearEnterKeyInLast(out []byte) string {
	outLength := len(out)
	if outLength < 0 {
		return string(out)
	}
	if out[outLength-1] == 10 {
		return clearEnterKeyInLast(out[:outLength-1])
	}
	return string(out)
}

// IsRoot : Checking user is root!
func IsRoot() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if user.Uid != "0" {
		fmt.Println("Please, You must be root!")
		os.Exit(1)
	}
}

func Notification(message string) {
	var notify *notificator.Notificator

	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "My test App",
	})

	notify.Push("title", message, "", notificator.UR_CRITICAL)
	//RunCommandOnBash("notify-send " + message + " -u critical")
}

func Espeak(message string) {
	if err := espeak.Init(); err == -1 {
		return
	}
	espeak.Say(message)
	espeak.Sync()
	//RunCommandOnBash("espeak " + message)
}
