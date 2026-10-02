package git

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/johanmcquillan/prompts"
	"github.com/johanmcquillan/prompts/env"
)

func MakeGitBranchComponent() *prompts.FunctionalComponent {
	return &prompts.FunctionalComponent{
		Function: gitBranch,
	}
}

type GitRelativeDirComponent struct {
	prompts.Formatter
	aliases map[string]string
}

func MakeGitRelativeDirComponent() *GitRelativeDirComponent {
	return &GitRelativeDirComponent{
		aliases: map[string]string{},
	}
}

func (c *GitRelativeDirComponent) WithFormatter(formatter prompts.Formatter) *GitRelativeDirComponent {
	c.Formatter = formatter
	return c
}

// WithRepoAlias shows the repo at repoPath as alias instead of its directory
// name, e.g. WithRepoAlias("~/core3/src", "core3"). A leading ~ in repoPath is
// expanded to $HOME.
func (c *GitRelativeDirComponent) WithRepoAlias(repoPath, alias string) *GitRelativeDirComponent {
	c.aliases[expandRepoPath(repoPath)] = alias
	return c
}

func (c *GitRelativeDirComponent) MakeElement() prompts.Element {
	rawValue := c.gitRelativeDir()

	if c.Formatter == nil {
		return prompts.Element{
			Output: rawValue,
			Length: len(rawValue),
		}
	}

	return prompts.Element{
		Output: c.Format(rawValue),
		Length: len(rawValue),
	}
}

func expandRepoPath(repoPath string) string {
	if repoPath == "~" || strings.HasPrefix(repoPath, "~"+env.PathSeparator) {
		repoPath = filepath.Join(os.Getenv(env.EnvHome), repoPath[1:])
	}
	// git reports the repo root with symlinks resolved, so match that.
	if resolved, err := filepath.EvalSymlinks(repoPath); err == nil {
		return resolved
	}
	return filepath.Clean(repoPath)
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

func (c *GitRelativeDirComponent) gitRelativeDir() string {
	repoPath := gitRepo()
	if repoPath == "" {
		s, _ := env.RelativeToHome()
		return s
	}

	repoSplit := strings.Split(repoPath, env.PathSeparator)
	repoName := repoSplit[len(repoSplit)-1]
	if alias, ok := c.aliases[repoPath]; ok {
		repoName = alias
	}

	s, ok := env.SubstitutePathPrefix(repoPath, os.Getenv(env.EnvPWD), repoName)
	if !ok {
		s, _ = env.RelativeToHome()
	}

	return s
}
