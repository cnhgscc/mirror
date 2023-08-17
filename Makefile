# -- args
export cmdargs="main.go" // path
export output="main"

# -- os
OS=`uname -s | tr 'A-Z' 'a-z'`

# -- env
PWD=`pwd`

BuildMod=`head -n 1 go.mod | cut -d ' ' -f 2`
BuildVersion=`go version | sed -e 's/go version //g' | cut -d ' ' -f 1`
BuildTime=`date '+%Y%m%d%H%M%S'`

# -- git
GITBranch=`git symbolic-ref --short -q HEAD`
GITVersion=`git rev-parse HEAD | cut -c1-8`


# -- build
ifneq ($(cmdargs), main.go)
	output = `echo $(cmdargs) | cut -d '/' -f 2 `
endif

args:
	@echo  $(BuildMod)"("$(GITBranch)")",$(GITVersion), $(BuildVersion)
	@echo

build:args
	@echo go build -o $(PWD)/_dist/$(output)-$(BuildTime)-$(GITVersion) $(PWD)/$(cmdargs)
	@go build -o $(PWD)/_dist/$(output)-$(BuildTime)-$(GITVersion) $(PWD)/$(cmdargs)

protoc:
	@protoc --go_out=.  --go-grpc_out=. ./api/*.proto
