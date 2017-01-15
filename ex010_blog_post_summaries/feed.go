package ex010_blog_post_summaries

// SummariseFeed takes a feed (list) of blog posts and creates a summary for each.
// Returns a map[BlogPost.ID][Summary]
func SummariseFeed(posts []BlogPost, summariser Summariser) map[string]string {
	summaries := make(map[string]string)
	for _, post := range posts {
		summaries[post.ID] = summariser.Summarise(post)
	}
	return summaries
}

// Summariser creates a summary string of a blog post.
type Summariser interface {
	Summarise(BlogPost) string
}
