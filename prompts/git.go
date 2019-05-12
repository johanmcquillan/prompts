package prompts

import (
	"bytes"
	"os/exec"
	"strings"
)

func MakeGitBranchComponent() *FunctionalComponent {
	return &FunctionalComponent{
		function: gitBranch,
	}
}

func gitBranch() string {
	var out bytes.Buffer
	cmd := exec.Command("git", "branch")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return ""
	}

	var branchString string
	branches := strings.Split(out.String(), "\n")
	for i := 0; i < len(branches) && branchString == ""; i++ {
		if strings.Contains(branches[i], "*") {
			branchString = branches[i][2:]
		}
	}

	return branchString
}

//func gitRepo() string {
//	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
//	var out bytes.Buffer
//	cmd.Stdout = &out
//	if err := cmd.Run(); err != nil {
//		return ""
//	}
//	return path[len(path) - 1]
//}
