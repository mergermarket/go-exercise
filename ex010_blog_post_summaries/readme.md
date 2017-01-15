Blog Post Summaries
===================

Our goal is to show summaries of a feed of blog posts.

Everything we need is already in place - look at [feed.go](feed.go) and [summarise.go](summarise.go)

We just need to wire things together in [main.go](cmd/main.go)

There are two tasks, solve them in separate git branches.

### Task 1
Create three Summarisers (for each behaviour defined in [summarise.go](summarise.go)) and call `SummariseFeed` with each one.

### Task 2
Again, call `SummariseFeed` once for each behaviour in  [summarise.go](summarise.go). However, avoid creating the three Summariser empty structs. Create an adapter to use the summarise functions directly.

