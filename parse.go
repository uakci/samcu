package samcu

import (
	"os/exec"
	"strings"
)

func parse(_ string, args []string) string {
	if len(args) == 0 {
		return ".i mo?"
	}

	parser := "-exp"
	mode := "X"
	parsers := map[string]bool{
		"-std":    true,
		"-beta":   true,
		"-cbm":    true,
		"-ckt":    true,
		"-exp":    true,
		"-morpho": true,
	}

	// Parse arguments:
	for len(args) >= 1 && strings.HasPrefix(args[0], "-") {
		if parsers[args[0]] {
			parser = args[0]
			args = args[1:]
		} else if len(args) >= 2 && args[0] == "-m" {
			mode = args[1]
			args = args[2:]
		} else {
			break
		}
	}

	// Run camxes:
	text := strings.Join(args, " ")
	cmd := exec.Command("node", "ilmentufa/run_camxes.js", parser, "-m", mode, text)
	stdout, err := cmd.Output()
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok && strings.Contains(string(e.Stderr), "SyntaxError") {
			return ".i .u'u na gendra"
		}
		return "failed to run camxes"
	}
	return strings.TrimSpace(string(stdout))
}
