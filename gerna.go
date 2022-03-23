package samcu

import (
	"os/exec"
	"strings"
  "fmt"
  "log"
)

var Gerna = Command{
  "gerna", gerna,
  "lanli lo gerna be lo'e lojbo",
  []CommandOption{
    {"uenzi", "gerna fi ma", nil, StringType},
  },
  []CommandOption{
    {"gerna", "ma gerna tarmi sfaile", [][2]string{
      {"exp", "cipra"},
      {"std", "fadni"},
      {"beta", "cninyrbeta"},
      {"cbm", "jai du'u ro cmevla cu brivla kei"},
      {"ckt", "tcekitau"},
    }, StringType},
    {"m", "ralte lo'e vlalei tcita", nil, BoolType},
    {"s", "ralte tu'a lo'e calbu'i", nil, BoolType},
    {"c", "ralte lo'e selma'o", nil, BoolType},
    {"t", "ralte lo'e famyma'o", nil, BoolType},
    {"n", "ralte lo'e ralju gentu'a tcita", nil, BoolType},
  },
}

func gerna(args map[string]any) (string, error) {
  uenzi := args["uenzi"].(string)
  parser := WithDefault(args, "gerna", "exp")
  flags := ""
  for key, val := range args {
    switch val.(type) {
    case bool:
      if val == true {
        flags += string(key[0] + 'A' - 'a')
      }
    }
  }

  cmdArgs := []string{"ilmentufa/run_camxes.js", fmt.Sprintf("-%s", parser)}
  if flags != "" {
    cmdArgs = append(cmdArgs, "-m", flags)
  }
  cmdArgs = append(cmdArgs, uenzi)
  cmd := exec.Command("node", cmdArgs...)

	stdout, err := cmd.Output()
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok && strings.Contains(string(e.Stderr), "SyntaxError") {
      log.Printf("camxes failed: %v %s", err, e.Stderr)
      return "", fmt.Errorf("i u'u ki'u di'e na gendra fa'o:\n %s", e.Stderr)
		} else {
      log.Printf("camxes failed: %v", err)
      return "", fmt.Errorf("i u'u fliba lo ka sazri la camxes")
    }
	}
	return strings.TrimSpace(string(stdout)), nil
}
