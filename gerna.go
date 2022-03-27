package samcu

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// var Gerna = Command{
//   "gerna", gerna,
//   "lanli lo gerna be lo'e lojbo",
//   []CommandOption{
//     {"uenzi", "gerna fi ma", nil, StringType},
//   },
//   []CommandOption{
//     {"gerna", "ma gerna tarmi sfaile", [][2]string{
//       {"exp", "cipra"},
//       {"std", "fadni"},
//       {"beta", "cninyrbeta"},
//       {"cbm", "jai du'u ro cmevla cu brivla kei"},
//       {"ckt", "tcekitau"},
//     }, StringType},
//     {"m", "ralte lo'e vlalei tcita", nil, BoolType},
//     {"s", "ralte tu'a lo'e calbu'i", nil, BoolType},
//     {"c", "ralte lo'e selma'o", nil, BoolType},
//     {"t", "ralte lo'e famyma'o", nil, BoolType},
//     {"n", "ralte lo'e ralju gentu'a tcita", nil, BoolType},
//   },
// }

var (
	parserValues = map[string]struct{}{
		"exp":  {},
		"std":  {},
		"beta": {},
		"cbm":  {},
		"ckt":  {},
	}
	flagValues = map[string]struct{}{
		"M": {},
		"S": {},
		"C": {},
		"T": {},
		"N": {},
	}
)

func gerna(args []string) (string, error) {
	parser := "exp"
	flags := ""
	for i, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			args = args[i:]
			break
		}
		arg := arg[1:]
		if _, ok := flagValues[arg]; ok {
			flags += arg
		} else if _, ok := parserValues[arg]; ok {
			parser = arg
		} else {
			return "", fmt.Errorf("na se slabu la'oi %s poi gerna", arg)
		}
	}

	cmdArgs := []string{"ilmentufa/run_camxes.js", "-" + parser}
	if flags != "" {
		cmdArgs = append(cmdArgs, "-m", flags)
	}
	cmdArgs = append(cmdArgs, strings.Join(args, " "))
	cmd := exec.Command("node", cmdArgs...)

	stdout, err := cmd.Output()
	if err != nil {
		log.Printf("camxes failed: %s", err.Error())
		e, notOk := err.(*exec.ExitError)
		if notOk && strings.Contains(string(e.Stderr), "SyntaxError") {
			return "", fmt.Errorf("ki'u di'e na gendra fa'o:\n %s", e.Stderr)
		} else {
			return "", fmt.Errorf("u'u fliba lo ka sazri la camxes")
		}
	}
	return strings.TrimSpace(strings.ReplaceAll(string(stdout), "_", "\\_")), nil
}
