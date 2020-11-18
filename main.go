package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetReportCaller(false)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	if os.Getenv("TOKEN") == "" || os.Getenv("ORG") == "" {
		fmt.Println("missing env.var (TOKEN or ORG)")
		os.Exit(1)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("TOKEN")},
	)

	client := github.NewClient(oauth2.NewClient(ctx, ts))

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	log.Info("retrieving a list of repositories, please stand by...")

	var repos []*github.Repository
	for {
		result, resp, err := client.Repositories.ListByOrg(ctx, os.Getenv("ORG"), opt)
		if err != nil {
			log.Fatal(err)
		}
		repos = append(repos, result...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	log.Infof("total %v repositories", len(repos))

	var archived []string
	var active []string
	for _, r := range repos {
		if !*r.Archived {
			active = append(active, *r.Name)
		} else {
			archived = append(archived, *r.Name)
		}
	}

	log.Infof("%v active repositories", len(active))
	fmt.Println(strings.Join(active, "\n"))
	log.Infof("%v archived repositories", len(archived))
	fmt.Println(strings.Join(archived, "\n"))
}
