package ex010_blog_post_summaries

import "fmt"

func TitleSummarise(post BlogPost) string {
	return fmt.Sprintf("[%s](%s)", post.Title, post.PermaLink)
}

func ContentSummarise(post BlogPost) string {
	return fmt.Sprintf("%s\n%s ...", TitleSummarise(post), post.Content[:100])
}

func BuzzFeedSummarise(post BlogPost) string {
	return fmt.Sprintf("[You'll never believe what %s just posted on his blog!](%s)", post.Author, post.PermaLink)
}
