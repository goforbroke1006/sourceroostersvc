package cvs

import "os/exec"

func ExecGitRemoteShowList {
	exec.Command("git", "remote show list")
	// todo: implement
}