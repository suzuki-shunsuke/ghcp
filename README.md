# ghcp

Fork of [int128/ghcp](https://github.com/int128/ghcp).

This is a release engineering tool for GitHub.
It depends on GitHub APIs and works without git installation.

It provides the following features:

- Commit files to a repository
- Create an empty commit
- Fork a repository and commit files to the forked repository
- Create a pull request
- Upload files to GitHub Releases

## Install

There are several ways to install ghcp.

1. [Homebrew](https://brew.sh/)

```sh
brew install suzuki-shunsuke/ghcp/ghcp
```

2. [Scoop](https://scoop.sh/)

```sh
scoop bucket add suzuki-shunsuke https://github.com/suzuki-shunsuke/scoop-bucket
scoop install ghcp
```

3. [aqua](https://aquaproj.github.io/)

```sh
aqua g -i suzuki-shunsuke/ghcp
```

4. Download a prebuilt binary from [GitHub Releases](https://github.com/suzuki-shunsuke/ghcp/releases) and unarchive it and install the executable file into `$PATH`

## GitHub Access Token

You need to set a GitHub Access token to the environment variable `GITHUB_TOKEN`.

```sh
export GITHUB_TOKEN=xxx
```

Or you can also pass the access token with `--token` option.

e.g.

```sh
ghcp --token xxx commit -r OWNER/REPO -m MESSAGE file1 file2
```

## Getting Started

### Commit files to a branch

To commit files to the default branch:

```sh
ghcp commit -r OWNER/REPO -m MESSAGE file1 file2
```

To commit files to `feature` branch:

```sh
ghcp commit -r OWNER/REPO -b feature -m MESSAGE file1 file2
```

If `feature` branch does not exist, ghcp will create it from the default branch.

To create `feature` branch from `develop` branch:

```sh
ghcp commit -r OWNER/REPO -b feature --parent=develop -m MESSAGE file1 file2
```

If `feature` branch already exists, ghcp will fail.
Currently only fast-forward is supported.

ghcp performs a commit operation as follows:

- An author and committer of a commit are set to the login user (depending on the token).
- If the branch has same files, do not create a new commit. It prevents an empty commit.
- It excludes `.git` directories.
- It does not support `.gitconfig`.

You can set the following options.

```
Flags:
      --author-email string      Author email (default: login email)
      --author-name string       Author name (default: login name)
  -b, --branch string            Name of the branch to create or update (default: the default branch of repository)
      --committer-email string   Committer email (default: login email)
      --committer-name string    Committer name (default: login name)
      --dry-run                  Upload files but do not update the branch actually
  -h, --help                     help for commit
  -m, --message string           Commit message (mandatory)
      --no-file-mode             Ignore executable bit of file and treat as 0644
      --no-parent                Create a commit without a parent
  -u, --owner string             Repository owner
      --parent string            Create a commit from the parent branch/tag (default: fast-forward)
  -r, --repo string              Repository name, either -r OWNER/REPO or -u OWNER -r REPO (mandatory)
```

### Create an empty commit to a branch

To create an empty commit to the default branch:

```sh
ghcp empty-commit -r OWNER/REPO -m MESSAGE
```

To create an empty commit to the branch:

```sh
ghcp empty-commit -r OWNER/REPO -b BRANCH -m MESSAGE
```

If the branch does not exist, ghcp creates a branch from the default branch.
It the branch exists, ghcp updates the branch by fast-forward.

To create an empty commit to a new branch from the parent branch:

```sh
ghcp empty-commit -r OWNER/REPO -b BRANCH --parent PARENT -m MESSAGE
```

If the branch exists, it will fail.

You can set the following options.

```
Flags:
      --author-email string      Author email (default: login email)
      --author-name string       Author name (default: login name)
  -b, --branch string            Name of the branch to create or update (default: the default branch of repository)
      --committer-email string   Committer email (default: login email)
      --committer-name string    Committer name (default: login name)
      --dry-run                  Do not update the branch actually
  -h, --help                     help for empty-commit
  -m, --message string           Commit message (mandatory)
  -u, --owner string             Repository owner
      --parent string            Create a commit from the parent branch/tag (default: fast-forward)
  -r, --repo string              Repository name, either -r OWNER/REPO or -u OWNER -r REPO (mandatory)
```

### Fork the repository and commit files to a new branch

To fork repository `UPSTREAM/REPO` and create `feature` branch from the default branch:

```sh
ghcp fork-commit -u UPSTREAM/REPO -b feature -m MESSAGE file1 file2
```

To fork repository `UPSTREAM/REPO` and create `feature` branch from `develop` branch of the upstream:

```sh
ghcp fork-commit -u UPSTREAM/REPO -b feature --parent develop -m MESSAGE file1 file2
```

If the branch already exists, ghcp will fail.
Currently only fast-forward is supported.

You can set the following options.

```
Flags:
      --author-email string      Author email (default: login email)
      --author-name string       Author name (default: login name)
  -b, --branch string            Name of the branch to create (mandatory)
      --committer-email string   Committer email (default: login email)
      --committer-name string    Committer name (default: login name)
      --dry-run                  Upload files but do not update the branch actually
  -h, --help                     help for fork-commit
  -m, --message string           Commit message (mandatory)
      --no-file-mode             Ignore executable bit of file and treat as 0644
  -u, --owner string             Upstream repository owner
      --parent string            Upstream branch name (default: the default branch of the upstream repository)
  -r, --repo string              Upstream repository name, either -r OWNER/REPO or -u OWNER -r REPO (mandatory)
```

### Create a pull request

To create a pull request from `feature` branch to the default branch:

```sh
ghcp pull-request -r OWNER/REPO -b feature --title TITLE --body BODY
```

To create a pull request from `feature` branch to the `develop` branch:

```sh
ghcp pull-request -r OWNER/REPO -b feature --base develop --title TITLE --body BODY
```

To create a pull request from `feature` branch of `OWNER/REPO` repository to the default branch of `UPSTREAM/REPO` repository:

```sh
ghcp pull-request -r OWNER/REPO -b feature --base-repo UPSTREAM/REPO --title TITLE --body BODY
```

To create a pull request from `feature` branch of `OWNER/REPO` repository to the default branch of `UPSTREAM/REPO` repository:

```sh
ghcp pull-request -r OWNER/REPO -b feature --base-repo UPSTREAM/REPO --base feature --title TITLE --body BODY
```

If a pull request already exists, ghcp do nothing.

You can set the following options.

```
Flags:
      --base string         Base branch name (default: default branch of base repository)
      --base-owner string   Base repository owner (default: head)
      --base-repo string    Base repository name, either --base-repo OWNER/REPO or --base-owner OWNER --base-repo REPO (default: head)
      --body string         Body of a pull request
      --draft               If set, mark as a draft
  -b, --head string         Head branch name (mandatory)
  -u, --head-owner string   Head repository owner
  -r, --head-repo string    Head repository name, either -r OWNER/REPO or -u OWNER -r REPO (mandatory)
  -h, --help                help for pull-request
      --reviewer string     If set, request a review
      --title string        Title of a pull request (mandatory)
```

### Release assets

To upload files to the release associated to tag `v1.0.0`:

```sh
ghcp release -r OWNER/REPO -t v1.0.0 dist/
```

If the release does not exist, it will create a release.
If the tag does not exist, it will create a tag from the default branch and create a release.

To create a tag and release on commit `COMMIT_SHA` and upload files to the release:

```sh
ghcp release -r OWNER/REPO -t v1.0.0 --target COMMIT_SHA dist/
```

If the tag already exists, it ignores the target commit.
If the release already exist, it only uploads the files.

You can set the following options.

```
Flags:
      --dry-run         Do not create a release and assets actually
  -h, --help            help for release
  -u, --owner string    Repository owner
  -r, --repo string     Repository name, either -r OWNER/REPO or -u OWNER -r REPO (mandatory)
  -t, --tag string      Tag name (mandatory)
      --target string   Branch name or commit SHA of a tag. Unused if the Git tag already exists (default: the default branch)
```

## Usage

### Global options

You can set the following options.

```
Global Flags:
      --api string         GitHub API v3 URL (v4 will be inferred) [$GITHUB_API]
      --debug              Show debug logs
  -C, --directory string   Change to directory before operation
      --token string       GitHub API token [$GITHUB_TOKEN]
```

### GitHub Enterprise

You can set a GitHub API v3 URL by `GITHUB_API` environment variable or `--api` option.

```sh
export GITHUB_API=https://github.example.com/api/v3/
```

GitHub API v4 URL will be automatically inferred from the v3 URL by resolving the relative path `../graphql`.

## Contributions

Please see [the guide](https://github.com/suzuki-shunsuke/oss-contribution-guide).

## LICENSE

[Apache License 2.0](LICENSE)

This is a fork of [int128/ghcp](https://github.com/int128/ghcp) ([base commit](https://github.com/int128/ghcp/commit/6de98923653dac599f0cc0e53dc805a9c963669e)).
The original author is [Hidetake Iwata](https://github.com/int128) and the original license is also [Apache License 2.0](https://github.com/int128/ghcp/blob/6de98923653dac599f0cc0e53dc805a9c963669e/LICENSE).
