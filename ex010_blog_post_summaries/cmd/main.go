package main

import (
	"fmt"

	app "github.com/dansweeting/go-exercise/ex010_blog_post_summaries"
)

type TitleSummariser struct {}
type ContentSummariser struct {}
type BuzzFeedSummariser struct {}

func (s TitleSummariser) Summarise(p app.BlogPost) string {
	return app.TitleSummarise(p)
}

func (s ContentSummariser) Summarise(p app.BlogPost) string {
	return app.ContentSummarise(p)
}

func (s BuzzFeedSummariser) Summarise(p app.BlogPost) string {
	return app.BuzzFeedSummarise(p)
}

func main() {
	posts := testData()

	titleSummariser := TitleSummariser{}
	contentSummariser := ContentSummariser{}
	buzzFeedSummariser := BuzzFeedSummariser{}

	titleSummaries := app.SummariseFeed(posts, titleSummariser)
	contentSummaries := app.SummariseFeed(posts, contentSummariser)
	buzzFeedSummaries := app.SummariseFeed(posts, buzzFeedSummariser)

	// Print it all out
	fmt.Println("\n\n============= Title Summaries ==============")
	for _, v := range titleSummaries {
		fmt.Println()
		fmt.Println(v)
	}

	fmt.Println("\n\n============= Content Summaries ==============")
	for _, v := range contentSummaries {
		fmt.Println()
		fmt.Println(v)
	}

	fmt.Println("\n\n============= BuzzFeed Summaries ==============")
	for _, v := range buzzFeedSummaries {
		fmt.Println()
		fmt.Println(v)
	}
}

func testData() []app.BlogPost {
	return []app.BlogPost{
		{
			ID:    "d8800ccd605f",
			Title: "WTF Dial: HTTP API",
			Content: `Now that we have our storage layer finished, we need some way for users to access and manipulate that data remotely. The standard approach seems to be JSON over HTTP so that’s what we’ll be using. It’s simple and it’s easy.
			In the Go world there seems to be a lack of consensus on how to write an HTTP API. You can see this in the vast number of HTTP frameworks written for Go. I’m no different and I have my own particular way of writing HTTP API servers in Go.`,
			Author:    "Ben Johnson",
			PermaLink: "https://medium.com/wtf-dial/wtf-dial-http-api-d8800ccd605f#.l2w8fzsds",
		},
		{
			ID:    "186dd7cdea88",
			Title: "Line of sight in code",
			Content: `A good line of sight makes no difference to what your function does, but it does help other humans who might need to read your code.
			The idea is that another programmer (including your future self) can glance down a single column and understand the expected flow of the code.
			If they have to jump around parsing if conditions in their brains moving in and out of code blocks, it makes that task much more difficult.`,
			Author:    "Mat Ryer",
			PermaLink: "https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88#.xfk7em2oo",
		},
		{
			ID:    "911ef4f8bd8e",
			Title: "Modern garbage collection",
			Content: `I’ve seen a bunch of articles lately which promote the Go language’s latest garbage collector in ways that trouble me.
			Some of these articles come from the Go project itself. They make claims that imply a radical breakthrough in GC technology has occurred.`,
			Author:    "Mike Hearn",
			PermaLink: "https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.j8w33br6y",
		},
	}
}
