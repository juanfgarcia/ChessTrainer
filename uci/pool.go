package uci

import (
    "sync"
)

type Work struct {
    Fen string
    P0 int
    P1 int
}

type Worker struct {
    eng *Engine
    m *sync.Mutex
	WorkerChannel chan chan Work // used to communicate between dispatcher and workers
	Channel chan Work
	End chan bool
}

// start worker
func (w *Worker) Start() {
    eng, _ := NewEngine("../ChessTrainer/uci/stockfish12")
    w.eng = eng

	go func() {
		for {
			w.WorkerChannel <-w.Channel // when the worker is available place channel in queue
			select {
			case job := <-w.Channel: // worker has received job
				w.DoWork(job.Fen,job.P0,job.P1) // do work
			case <-w.End:
				return
			}
		}
	}()
}

func (w *Worker) DoWork(fen string, p0, p1 int){
    w.eng.SetFEN(fen)
    _, _ = w.eng.Go(p0,p1)
}

// end worker
func (w *Worker) Stop() {
	w.End <- true
}
