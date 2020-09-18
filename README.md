# Z Tests
Golang tests with containers without friction

[![Go Report Card](https://goreportcard.com/badge/github.com/alexrios/ztests)](https://goreportcard.com/report/github.com/alexrios/ztests)

## Getting started
The purpose is to make testing docker containers easier without having to worry about managing them. 
The main goal of this library is to be as transparent as possible when it comes to a simple suite of unit tests.

To achieve this results this project uses [Test Containers Go](https://github.com/testcontainers/testcontainers-go) and
[Canned Containers](https://github.com/alexrios/canned-containers)

#### Example
This example has 3 easy steps:
* Call `Setup()` to receive a `db connection` and a `termination function`
* `defer teardown()` to handle all the docker parts
* The test itself 

```go
import (
	"context"
	"github.com/alexrios/ztests/postgres"
	"testing"
)

func Test(t *testing.T) {
	ctx := context.Background()
	testEnv, err := postgres.Setup()
	defer testEnv.Teardown()
	conn, err := testEnv.DB.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var twelve int
	row := conn.QueryRowContext(ctx, "select $1", 12)
	err = row.Scan(&twelve)
	if err != nil {
		t.Fatal(err)
	}
	if twelve != 12 {
		t.FailNow()
	}
}

```

### TODO
[x] support for Postgres canned container 

[ ] support for Redis canned container 

[ ] support for Elastic Search canned container

[ ] support for MySQL Search canned container

[ ] support for nginx Search canned container

[ ] and many more...

## Contributing
If you have any questions or feedback regarding ztests, bring it!
Your feedback is always welcome.

#### Getting help?
Hit me on twitter [@alextrending](https://twitter.com/alextrending)!

