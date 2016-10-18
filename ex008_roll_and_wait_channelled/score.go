package main

type Score struct {
	Name  string
	Score int
}

type ByScore []Score

func (items ByScore) Len() int {
	return len(items)
}

func (items ByScore) Less(x, y int) bool {
	return items[x].Score < items[y].Score
}

func (s ByScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
