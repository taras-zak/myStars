package main

import (
	"context"
	"github.com/google/go-github/github"
	"log"
	"os"
	"text/template"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	repos, err := GetStarredRepos("taras-zak")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %s stars", len(repos))
	//PrintRepos(repos)
	MakeMarkdown(repos)
}

func GetStarredRepos(user string) ([]*github.Repository, error) {
	var starredRepos []*github.Repository
	opt := github.ActivityListStarredOptions{ListOptions: github.ListOptions{Page: 0, PerPage: 1000}}
	client := github.NewClient(nil)
	for {
		repos, res, err := client.Activity.ListStarred(context.Background(), user, &opt)
		if err != nil {
			return starredRepos, err
		}
		for _, repo := range repos {
			starredRepos = append(starredRepos, repo.Repository)
		}
		opt.ListOptions.Page = res.NextPage
		if res.NextPage == 0 {
			log.Println(res)
			break
		}
	}
	return starredRepos, nil
}

func PrintRepos(repos []*github.Repository) {
	for _, repo := range repos {
		log.Println("Repo: ", repo.GetFullName())
		log.Println("  Description ", repo.GetDescription())
		log.Println("  Homepage ", repo.GetHomepage())
		log.Println("  HTMLURL ", repo.GetHTMLURL())
		log.Println("  Language ", repo.GetLanguage())
		log.Println("  ForksCount ", repo.GetForksCount())
		log.Println("  StargazersCount ", repo.GetStargazersCount())
		log.Println("  OpenIssuesCount ", repo.GetOpenIssuesCount())
	}
}

func MakeMarkdown(repos []*github.Repository) {
	t := template.Must(template.New("README.md").Parse(ReadmeTemplate))
	err := t.Execute(os.Stdout, repos)
	if err != nil {
		log.Fatalf("executing template:", err)
	}
}
