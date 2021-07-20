package uci

import (
	"bufio"
	"os/exec"
	"fmt"
	"log"
    "strings"
)

// Result holds the score infromation 	
type Results struct {
    cp1            int       // score centipawns or mate in X if Mate is true
    cp2            int
	Mate           int      // whether this move results in forced mate
	BestMove       string  // best line for this result
}

// Engine holds the information needed to communicate with
// a chess engine executable. Engines should be created with
// a call to NewEngine(/path/to/executable)
type Engine struct {
	cmd    *exec.Cmd
	stdout *bufio.Reader
	stdin  *bufio.Writer
}


// NewEngine returns an Engine it has spun up
// and connected communication to
func NewEngine(path string) (*Engine, error) {
	eng := Engine{}
	eng.cmd = exec.Command(path)
	stdin, err := eng.cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := eng.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := eng.cmd.Start(); err != nil {
		return nil, err
	}
	eng.stdin = bufio.NewWriter(stdin)
	eng.stdout = bufio.NewReader(stdout)

    return &eng, nil
}

// SetFEN takes a FEN string and tells the engine to set the position
func (eng *Engine) SetFEN(fen string) error {
	_, err := eng.stdin.WriteString(fmt.Sprintf("position fen %s\n", fen))
	if err != nil {
		return err
	}
	err = eng.stdin.Flush()
    if err != nil {
		return err
	}
	return err
}


// Go recieves to depth p0 and p1 and return the score at 
// each depth, p0 < p1 is assumed.
func (eng *Engine) Go(p0, p1 int) (error) {
	_, err := eng.stdin.WriteString(fmt.Sprintf("go depth %d\n", p1))
	if err != nil {
		return err
    }

    err = eng.stdin.Flush()
	if err != nil {
		return err
	}

    for {
        line, err := eng.stdout.ReadString('\n')
		if err != nil {
			return err
		}
        line = strings.Trim(line, "\n")
		s := string(line)
        if strings.HasPrefix(line, "bestmove") {
			break
		}
        match := reInfo.FindStringSubmatch(s)
        m := map[string]string{}
        for i, n := range(match) {
            m[groups[i]] = n
        }
        fmt.Println(m["depth"])
        fmt.Println(m["cp"])
    }
    return err
}

func (eng *Engine) Close() {
	_, err := eng.stdin.WriteString("stop\n")
	if err != nil {
		log.Println("failed to stop engine:", err)
	}
	eng.stdin.Flush()
	err = eng.cmd.Process.Kill()
	if err != nil {
		log.Println("failed to kill engine:", err)
	}
	eng.cmd.Wait()
}

//var reInfo = regexp.MustCompile(`info depth (?P<depth>\d+) seldepth [0-9]+ multipv \d+ score cp (?P<cp>\d+)`)
//var groups = reInfo.SubexpNames()
//
//type Expr interface{
//    isExpr()
//
//}
//
//type FinalLine bool
//
//    final bool,
//    struct content {
//        depth   int,
//        cp      int,
//        variant string,
//    }
//}
//
//func parseLine(line string) uciLine {
//    match := reInfo.FindStringSubmatch(line)
//    m := map[string]string{}
//    for i, n := range(match) {
//        m[groups[i]] = n
//    }
//    depth,_ = strconv.Atoi(m["depth"])
//    cp,_    = strconv.Atoi(m["cp"])
//
//    return
//}
