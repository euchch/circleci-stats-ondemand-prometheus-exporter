package polingqueue

type RepoBranch struct {
	Repo   string `json:"Name"`
	Branch string `json:"Branch"`
}

type PolingQueue struct {
	queue []RepoBranch
}

func (rb *RepoBranch) equals(repo string, branch string) bool {
	return repo == rb.Repo && branch == rb.Branch
}

func (p *PolingQueue) EnqueueUniqe(repo string, branch string) {
	if (repo == "") || (branch == "") {
		return
	}

	for _, pr := range p.queue {
		if pr.equals(repo, branch) {
			return
		}
	}

	p.queue = append(p.queue, *NewRepoBranch(repo, branch))
}

func (p *PolingQueue) Dequeue() *RepoBranch {
	if len(p.queue) == 0 {
		return nil
	}

	dequeuedRP := NewRepoBranch(p.queue[0].Repo, p.queue[0].Branch)
	p.queue = p.queue[1:]
	return dequeuedRP
}

func NewRepoBranch(repo string, branch string) *RepoBranch {
	return &RepoBranch{
		Repo:   repo,
		Branch: branch,
	}
}

func NewPolingQueue() *PolingQueue {
	return &PolingQueue{}
}
