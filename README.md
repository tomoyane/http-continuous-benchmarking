## HTTP Continuous benchmarking

This repository provides a continuous benchmarking.  
If your project needs benchmark monitoring, it collects performance data by this repository.

|Type|Badge/URL|
|---|---|
|CICD|[![ci](https://github.com/tomoyane/http-continuous-benchmarking/actions/workflows/ci.yml/badge.svg)](https://github.com/tomoyane/http-continuous-benchmarking/actions/workflows/ci.yml)|
|Go Report Card|[![Go Report Card](https://goreportcard.com/badge/github.com/tomoyane/http-continuous-benchmarking)](https://goreportcard.com/report/github.com/tomoyane/http-continuous-benchmarking)|
|Coveralls|[![Coverage Status](https://coveralls.io/repos/github/tomoyane/http-continuous-benchmarking/badge.svg?branch=main)](https://coveralls.io/github/tomoyane/http-continuous-benchmarking?branch=main)|
|Coverage report for GitHub Pages|[Coverage report GitHub Pages](https://tomoyane.github.io/http-continuous-benchmarking/#file0)

### Concept
* Simple and lightweight benchmark tool
  * This tool is not recommended for if large-scale performance measurements
* Continuous benchmarking tool
  * Always monitor performance impacts for source code changes
* Lightweight to make various requests
  * The user does not do complicated things
* Create measured reports and notify warnings
  * Create a continuous report and associate it with a commit hash
  * Warn and notice when the threshold is reached

### How to use
