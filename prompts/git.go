package prompts

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func MakeGitBranchComponent() *FunctionalComponent {
	return &FunctionalComponent{
		function: gitBranch,
	}
}

func MakeGitRelativeDirComponent() *FunctionalComponent {
	return &FunctionalComponent{
		function: gitRelativeDir,
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

func gitRepo() string {
	var out bytes.Buffer
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return ""
	}

	return strings.Replace(out.String(), "\n", "", -1)
}

func gitRelativeDir() string {
	repoPath := gitRepo()
	if repoPath == "" {
		s, _ := relativeToHome()
		return s
	}

	repoSplit := strings.Split(repoPath, pathSeparator)
	repoName := repoSplit[len(repoSplit) - 1]

	s, ok := substitutePathPrefix(repoPath, os.Getenv(envPWD), repoName)
	if !ok {
		s, _ = relativeToHome()
	}

	return s
}
