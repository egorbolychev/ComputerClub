package models

// Struct for sorting Tasks
type TaskSlice []*Task

func (t TaskSlice) Len() int { return len(t) }

func (t TaskSlice) Less(i, j int) bool { return t[i].Time.Before(t[j].Time) }

func (t TaskSlice) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
