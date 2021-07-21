package uci

import (
    "regexp"
    "strconv"
    "strings"
)

var reInfo = regexp.MustCompile(`info depth (?P<depth>\d+) seldepth [0-9]+ multipv \d+ score cp (?P<cp>\d+) nodes \d+ nps \d+ tbhits \d+ time \d+ pv (?P<variant>[a-zA-Z0-9 ]+)`)
var groups = reInfo.SubexpNames()

type UCIString interface{
    isExpr()
}

type Final bool
func isFinal(s string) bool{
    if strings.HasPrefix(s, "bestmove"){
        return true
    }
    return false
}
func (f Final) isExpr() {}


type Data struct{
    depth   int
    cp      int
    variant string
}
func (d Data) isExpr() {}

func isData(s string) (bool,Data){
    match := reInfo.FindStringSubmatch(s)
    if len(match)==0{
        return false,Data{}
    }
    m := map[string]string{}
    for i, n := range(match) {
        m[groups[i]] = n
    }
    depth,_ := strconv.Atoi(m["depth"])
    cp,_    := strconv.Atoi(m["cp"])
    variant := m["variant"]
    return true,Data{depth,cp,variant}
}

type Neither struct{}
func (n Neither) isExpr(){}

func ParseLine(line string) UCIString {
    if (isFinal(line)) {
        return Final(true)
    } else if isData,data := isData(line); isData==true{
        return data
    }
    return Neither(struct{}{})
}
