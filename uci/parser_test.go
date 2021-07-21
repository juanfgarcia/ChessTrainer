package uci

import (
	"testing"
)

func TestParseLine(t *testing.T){

    tests := []struct{
        name    string
        got     string
        expr    UCIString
    }{
        {
            name: "Initial line",
            got:  "info string NNUE evaluation using nn-82215d0fd0df.nnue enabled",
            expr:  Neither(struct{}{}),
        },
        {
            name: "First line",
            got: "info depth 1 seldepth 1 multipv 1 score cp 3 nodes 20 nps 10000 tbhits 0 time 2 pv d2d4",
            expr: Data{depth: 1, cp: 3,variant: "d2d4"},
        },
        {
            name: "Last line",
            got:"info depth 7 seldepth 7 multipv 1 score cp 37 nodes 1680 nps 105000 tbhits 0 time 16 pv c2c4 e7e5 e2e3 g8f6 g1f3",
            expr: Data{depth: 7, cp: 37, variant: "c2c4 e7e5 e2e3 g8f6 g1f3"},
        },
        {
            name: "Final line",
            got: "bestmove c2c4 ponder e7e5",
            expr: Final(true),
        },
    }

    for _, tt := range tests {
        got := ParseLine(tt.got)
        t.Run(tt.name, func (t *testing.T){
            switch got.(type) {
                case Final:{
                    if (got != tt.expr){
                        t.Errorf("Wanted %b but got %b for string %s",tt.expr, got, tt.got)
                    }
                }
                case Data:{
                    if (got.(Data).cp != tt.expr.(Data).cp){
                        t.Errorf("Wanted %d but got %d for string %s",tt.expr.(Data).cp, got.(Data).cp, tt.got)
                    }
                    if (got.(Data).depth != tt.expr.(Data).depth){
                        t.Errorf("Wanted %d but got %d for string %s",tt.expr.(Data).depth, got.(Data).depth, tt.got)
                    }
                    if (got.(Data).variant != tt.expr.(Data).variant){
                        t.Errorf("Wanted %s but got %s for string %s",tt.expr.(Data).variant, got.(Data).variant, tt.got)
                    }
                }
            }
        })
    }
}
