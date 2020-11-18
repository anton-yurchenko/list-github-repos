# list-github-repos

List GitHub repositories for an organization

## Manual

1. Create a token that allows to access the repositories
2. Download latest release
3. Export previously created GitHub token: `export TOKEN=xxx`
4. Run the tool: `./list-github-repos`

Example output:

```log
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
