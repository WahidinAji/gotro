
# Codegen for Web Framework

This framework aims to reduce common task that happened during backend API development which are:

1. generating ORM and migrating model
2. generate for other type of input and output encoding/presentation format (JSON, MsgPack, YAML, JSON5, GraphQL, etc), command line, or different transport (REST, WebSocket, gRPC, CLI, etc)
3. generate API documentation and API clients for different programming languages
4. testable business logic (by untangling `domain/` and decorator/adapter/presentation)

This codegen currently tied to Tarantool (for OLTP use cases), Clickhouse (for OLAP use cases), and Fiber (for web routes) and only tested for generating json using REST/GraphQL transport and plain JS API docs. But you can always create a new codegen for other databases or other web frameworks or other API clients.

## FAQ

Why is it tied to [Tarantool](//www.tarantool.io/en/developers)?
- because it's currently the fastest SQL-OLTP database that also works as in-memory cache, so I can do [integration testing](//kokizzu.blogspot.com/2021/07/mock-vs-fake-and-classical-testing.html) without noticable overhead (can do ~200K writes per second), your first scalability issue would mostly always be the database, unless you're using slow untyped programming language.

Why is it tied to [Clickhouse](//clickhouse.tech)?
- because it's currently the fastest OLAP database (can do ~600K inserts per second).

Why is it tied to [Fiber](//gofiber.io/)?
- because it's currently the [fastest](//www.techempower.com/benchmarks/#section=data-r20&hw=ph&test=update&l=zijocf-sf) minimalist golang framework, ones that are faster but not more painful to use is C#'s [asp.net core](//learn.microsoft.com/en-us/aspnet/core/)

Can I replace the components/database (so instead of using Tarantool/Clickhouse/Fiber)?
- yes, write your own codegen for database and api generator, it's possible

Why using `github.com/goccy/go-json`?
- because I couldn't find any other faster alternative (other than slower `encoding/json`) that can properly parse int64/uint64 (already tried `jsoniter` and `easyjson`, both give wrong result for `{"id":"89388457092187654"}` with `json:"id,string"` tag).

Why tying between business logic and data store (not using Repository pattern)?
- because we are rarely changing database product, like 90% of the time we'll never switch database since it's too costly (time consuming), if read part is the bottleneck we usually cache it (or precalculate or create static snapshot on every update), but if the write is the bottleneck, we usually shard it to different machine (either manual sharding, eg. A to L on node 1, K to Z on node 2, or automatic like with 2 level hash), even when we have to switch database, we usually have to do double-write first, which using `interface{anything}` totally have a very little benefit, compared to the cost (reduced agility when doing source code navigation, since you cannot jump properly to the declaration/implementation/usage if using interface), I still believe that only good time to use interface is for third party (to replace it with fakes for testing 3rd party that usually have a very slow I/O)
- because the data store itself are fast (in-mem) and can be tested using docker-compose or dockertest, no noticable difference between fake/mock and integration test with this database product.
- because we shouldn't separate between memory (data structure) and business logic (algorithm), since they are both the brain of our application, without memory there's nothing to think, without thinking/processing it's just uninformative raw data blob.
- because most entity/repository have 1 to 1 relationship with the persistent/data store adapter, so it would be a overengineering to split when you not yet need it (YAGNI).

How to scale out or load balance?
you can use strategies in this [blog](//kokizzu.blogspot.com/2022/04/automatic-load-balancer.html)

## Usage

See `example/` directory for minimum framework template, or if you want to do codegen manually:

1. create a test file `0_generator_test.go` inside your `domain/` project folder

```go
package domain_test

import (
	"testing"
	`github.com/kokizzu/gotro/W2`
)

//go:generate go test -run=XXX -bench=Benchmark_Generate_WebApiRoutes_CliArgs
//go:generate go test -run=XXX -bench=Benchmark_Generate_SvelteApiDocs

func Benchmark_Generate_WebApiRoutes_CliArgs(b *testing.B) {
	W2.GenerateFiberAndCli(&WW.GeneratorConfig{
		ProjectName: PROJECT_NAME,
	})
	b.SkipNow()
}

func Benchmark_Generate_SvelteApiDocs(b *testing.B) {
	W2.GenerateApiDocs(&WW.GeneratorConfig{
		ProjectName: PROJECT_NAME,
	})
	b.SkipNow()
}

```

2. create a makefile to do the codegen
```Makefile

gen-route:
	cd domain ; rm -f *MSG.GEN.go 
	cd domain ; go test -bench=Benchmark_Generate_WebApiRoutes_CliArgs 0_generator_test.go
	cd domain ; cat *.go | grep '//go:generate ' | cut -d ' ' -f 2- | sh -x
	cd domain ; go test -bench=Benchmark_Generate_SvelteApiDocs 0_generator_test.go

```

3. run `make gen-route`

this would create few generated file:

```
main_cli_args.GEN.go --> cli arguments handler
main_restApi_routes.GEN.go --> used to generating fiber route handlers
svelte/src/pages/api.js --> used for generating API client
```

## Example Generated API Docs

![image](https://user-images.githubusercontent.com/1061610/131266708-44935872-e34a-4538-885a-6056946c9482.png)

## TODO

- censor fields from GraphQL codegen (eg. password, access tokens)
- add gRPC codegen
