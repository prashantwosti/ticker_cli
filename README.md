# ticker_cli
![Go](https://github.com/prashantwosti/ticker_cli/workflows/Go/badge.svg?branch=master)

Shows you the company price from the ticker symbol.
` $ ticker_cli <exchange> <symbol>`

e.g 
```
$ ticker_cli US TSLA


$€£¥$€£¥$€£¥$€£¥$€£¥$€£¥
Tesla, Inc.
Price: USD 883.09
Last change:  0.26%
$€£¥$€£¥$€£¥$€£¥$€£¥$€£¥


$ ticker_cli ASX APT


$€£¥$€£¥$€£¥$€£¥$€£¥$€£¥
Afterpay Limited
Price: AUD 146.00
Last change:  1.77%
$€£¥$€£¥$€£¥$€£¥$€£¥$€£¥


```

## How to build an executable:
Navigate inside the project dir.

Assuming you have [Go](https://golang.org/) installed.

Using [Gox](https://github.com/mitchellh/gox) to make a binary file.

1. `$ go get github.com/mitchellh/gox`

2. Execute the following command as per your platform & arch.
   - for linux: `gox -osarch="linux/amd64"`
   - for mac: `gox -osarch="darwin/amd64"`
   - for windows: `gox -osarch="windows/amd64"`
   - [Other than above](https://github.com/mitchellh/gox#usage)?

You'll find the binary file in the root dir of the project.
