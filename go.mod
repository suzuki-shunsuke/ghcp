module github.com/suzuki-shunsuke/ghcp

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.6.0
	github.com/google/go-github/v60 v60.0.0
	github.com/google/wire v0.6.0
	github.com/shurcooL/githubv4 v0.0.0-20221229060216-a8d4a561cc93
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	golang.org/x/oauth2 v0.18.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f // indirect
	golang.org/x/net v0.22.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

go 1.21

toolchain go1.22.1
