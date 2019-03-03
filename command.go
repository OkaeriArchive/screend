package main

import (
	"github.com/pkg/errors"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

type UserInfo struct {
	name string
	gid  uint32
	uid  uint32
}

func getUserInfoByName(userName string) (UserInfo, error) {

	screenUser, err := user.Lookup(userName)
	if err != nil {
		return UserInfo{}, err
	}

	gid64, _ := strconv.ParseUint(screenUser.Gid, 10, 32)
	gid := uint32(gid64)

	uid64, _ := strconv.ParseUint(screenUser.Uid, 10, 32)
	uid := uint32(uid64)

	return UserInfo{userName, gid, uid}, nil
}

func executeCommand(userInfo UserInfo, ignoreErrors bool, runDirectory string, args ...string) (string, error) {

	baseCmd := args[0]
	cmdArgs := args[1:]

	cmd := exec.Command(baseCmd, cmdArgs...)
	cmd.Dir = runDirectory
	currentUser, _ := user.Current()
	root := currentUser.Gid == "0"

	if userInfo.name != currentUser.Username && !root {
		return "", errors.New("Screen daemon requires root privileges to switch between users")
	}

	if userInfo.name != currentUser.Username && root {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.SysProcAttr.Credential = &syscall.Credential{Uid: userInfo.uid, Gid: userInfo.gid}
	}

	out, err := cmd.Output()
	if err != nil && !ignoreErrors {
		return string(out), err
	}

	return string(out), nil
}
