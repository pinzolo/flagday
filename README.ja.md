# flagday

[![Test](https://github.com/pinzolo/flagday/actions/workflows/test.yml/badge.svg)](https://github.com/pinzolo/flagday/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/pinzolo/flagday/badge.svg?branch=master)](https://coveralls.io/github/pinzolo/flagday?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/pinzolo/flagday)](https://goreportcard.com/report/github.com/pinzolo/flagday)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/pinzolo/flagday)
[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/pinzolo/flagday/master/LICENSE)

## 概要

flagday は日本の祝日を扱うライブラリです。

## 特徴

* [国民の休日に関する法律](http://elaws.e-gov.go.jp/search/elawsSearch/elaws_search/lsg0500/detail?lawId=323AC1000000178&openerCode=1)に準じています。
* 1948年7月20日以降の祝日に対応しています。
* データファイルを使用していないので、下記の条件を除けば基本的には更新無しで使用できます。
    * 法律の改正による祝日の増減および移動
    * 遠い未来(2100年以降)の春分・秋分の日
    * 地殻変動などによる春分・秋分の日のずれ
    * 皇室関連の慶弔時による当年限りの休日

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
