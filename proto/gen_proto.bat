@echo off

for /D %%f in (*) do (
		cd %%f
		protoc -I . -I ../common --go_out=%GOPATH%/src *.proto
		cd ..)
