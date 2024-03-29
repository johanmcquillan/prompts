package git

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/johanmcquillan/prompts"
	"github.com/johanmcquillan/prompts/env"
)

func MakeGitBranchComponent() *prompts.FunctionalComponent {
	return &prompts.FunctionalComponent{
		Function: gitBranch,
	}
}

func MakeGitRelativeDirComponent() *prompts.FunctionalComponent {
	return &prompts.FunctionalComponent{
		Function: gitRelativeDir,
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
		s, _ := env.RelativeToHome()
		return s
	}

	repoSplit := strings.Split(repoPath, env.PathSeparator)
	repoName := repoSplit[len(repoSplit)-1]

	s, ok := env.SubstitutePathPrefix(repoPath, os.Getenv(env.EnvPWD), repoName)
	if !ok {
		s, _ = env.RelativeToHome()
	}

	return s
}
