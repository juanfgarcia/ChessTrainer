package uci

import (
	"bufio"
	"os/exec"
	"fmt"
	"log"
    "strings"
)

// Result holds the score infromation 	
type Result struct {
    cp1            int       // score centipawns or mate in X if Mate is true
    cp2            int
	mate           int      // whether this move results in forced mate
	variant       string  // best line for this result
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
func (eng *Engine) Go(p0, p1 int) (Result, error) {
    var (
        results Result
        finished bool = false
    )
	_, err := eng.stdin.WriteString(fmt.Sprintf("go depth %d\n", p1))
	if err != nil {
		return Result{},err
    }

    err = eng.stdin.Flush()
	if err != nil {
		return Result{},err
	}

    for !finished{
        line, err := eng.stdout.ReadString('\n')
		if err != nil {
			return Result{},err
		}
        line = strings.Trim(line, "\n")
		s := string(line)

        switch uci_string := ParseLine(s).(type) {
        case Data: {
            if (uci_string.depth==p0){
                results.cp1 = uci_string.cp
            }else if (uci_string.depth==p1){
                results.cp1 = uci_string.cp
                results.variant = uci_string.variant
            }
        }
        case Final: finished=true
        case Neither:
        }
    }
    return results, err
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
