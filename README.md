# flagday

[![Build Status](https://travis-ci.org/pinzolo/flagday.png)](http://travis-ci.org/pinzolo/flagday)
[![Coverage Status](https://coveralls.io/repos/github/pinzolo/flagday/badge.svg?branch=master)](https://coveralls.io/github/pinzolo/flagday?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/pinzolo/flagday)](https://goreportcard.com/report/github.com/pinzolo/flagday)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/pinzolo/flagday)
[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/pinzolo/flagday/master/LICENSE)

## Description

flagday is library for public holidays in Japan.

## Specification

* It is in compliance with [the law concerning national holidays](http://elaws.e-gov.go.jp/search/elawsSearch/elaws_search/lsg0500/detail?lawId=323AC1000000178&openerCode=1).
* Compatible holidays after 20th July 1948.
* flagday does not use data file, basically it can be used without updating except for the following conditions.
    * Increase and decrease of holidays and migration due to revision of law
    * Equinox days in the distant future (after 2100)
    * Deviation of equinox due to crustal deformation
    * Holidays of Imperial related condolences and auspicious event

## Install

```bash
$ go get github.com/pinzolo/flagday
```

## Contribution

1. Fork ([https://github.com/pinzolo/flagday/fork](https://github.com/pinzolo/flagday/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[pinzolo](https://github.com/pinzolo)
