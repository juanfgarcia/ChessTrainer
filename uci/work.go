package uci

// create list of jobs
func CreateJobs(amount int) []string {
	var jobs []string

	for i := 0; i < amount/2; i++ {
		jobs = append(jobs,"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	}
    for i := 0; i < amount/2; i++ {
		jobs = append(jobs,"rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2")
	}
	return jobs
}

