package store

// Manage clients who are waiting for empty table.
// Data struct: queue, stores usernames
type QueueManager struct {
	queu []string
}

func (q *QueueManager) Add(user string) {
	q.queu = append(q.queu, user)
}

func (q *QueueManager) Get() string {
	res := q.queu[0]
	q.queu = q.queu[1:]
	return res
}

func (q *QueueManager) Len() int {
	return len(q.queu)
}
