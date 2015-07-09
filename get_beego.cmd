@pushd %~dp0

@set GOPATH=%~dp0
go get github.com/astaxie/beego
go get github.com/beego/bee

@pause

@popd