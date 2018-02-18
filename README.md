# Contacts API
## Service for handling contacts calls

## Setup
This service runs on go.

Install and setup go
- On OSX run `brew install go glide`.
- Make sure that the executable `go` is in your shell's path.
```
GOPATH=<workspace_dir>
export GOPATH
PATH="${PATH}:${GOPATH}/bin"
export PATH
```

### Get source code
```shell
```

### Setup
copy application.yml.sample to application.yml and update the config
```shell
go run main.go

```

###Vegata
vegeta attack -targets=./target.txt -body=./contacts_request.json -rate=100 -duration=120s | vegeta report -reporter=plot -output=report.html
