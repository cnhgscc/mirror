# -- args
export cmdargs="main.go" // path

# -- os
OS=`uname -s | tr 'A-Z' 'a-z'`

# -- env
PWD=`pwd`

BuildMod=`head -n 1 go.mod | cut -d ' ' -f 2`
BuildVersion=`go version | sed -e 's/go version //g' | cut -d ' ' -f 1`

# -- git
GITBranch=`git symbolic-ref --short -q HEAD`
GITVersion=``


args:
	@echo  $(BuildMod)"("$(GITBranch)")", $(BuildVersion)
	@echo

build:args
	go build -o $(PWD)/_dist/ $(PWD)/$(cmdargs)
