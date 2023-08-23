package main

import (
	"context"
	"fmt"
	"mockery/article"
	"mockery/author"
)

func main() {

	return
}

func Get(ctx context.Context, articleRepository article.ArticleRepository, authorRepository author.AuthorRepository) {
	res1, _ := articleRepository.GetByID(ctx, 1)
	res2, _ := authorRepository.GetByID(ctx, 2)
	fmt.Print(res1, res2)
}
