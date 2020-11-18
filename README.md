# list-github-repos

List active (not archived) repositories for an organization

## Manual

1. Create a token that allows to access the repositories
2. Run the tool with `TOKEN=xxx ORG=yyy go run main.go` or compile it and run the binary by providing the same env.vars

Example output:

```
INFO retrieving a list of repositories, please stand by...
INFO total 2648 repositories
INFO 1925 active repositories
repo-1
repo-2
repo-3
...
INFO 723 archived repositories
repo-4
repo-5
...

```

## License

[MIT](LICENSE.md) © 2020-present Anton Yurchenko
