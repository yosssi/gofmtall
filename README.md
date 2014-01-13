# gofmtall - Format tool in Go

[![Build Status](https://drone.io/github.com/yosssi/gofmtall/status.png)](https://drone.io/github.com/yosssi/gofmtall/latest)

gofmtall is a tool for formatting the go files. By using this tool, you can format all the go files in the specific directory and its subdirectories.

## Installation

	$ go get github.com/yosssi/gofmtall

## Execution

Change the current directory to the your go package path.

	$ cd your-go-package-path

Execute the `gofmtall` binary file.

	$ your-go-workspace-path/bin/gofmtall

Or, if you added the your go workspace's bin directory on PATH, you can execute simply the following command.

	$ gofmtall

If three are go files that can be formatted, the following message would be shown.

	The following go files will be formatted:

	./file1.go
	./file2.go
	./sub1/file3.go
	./sub1/sub2/file4.go

	Are you sure to format them? [y/n]

If you input `y` and press Enter key, the files will be formatted and the following message would be shown.

	Done.
