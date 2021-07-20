package uci

import (
    "regexp"
    "strconv"
)

var reInfo = regexp.MustCompile(`info depth (?P<depth>\d+) seldepth [0-9]+ multipv \d+ score cp (?P<cp>\d+)`)
var groups = reInfo.SubexpNames()

type Expr interface{
    isExpr()
}

type Final bool
func (f Final) isExpr() {}

type Data struct{
    depth   int
    cp      int
    variant string
}
func (d Data) isExpr() {}


func parseLine(line string) Expr {
    match := reInfo.FindStringSubmatch(line)
    m := map[string]string{}
    for i, n := range(match) {
        m[groups[i]] = n
    }
    depth,_ := strconv.Atoi(m["depth"])
    cp,_    := strconv.Atoi(m["cp"])

    return nil
}
