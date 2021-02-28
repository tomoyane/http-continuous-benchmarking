## HTTP Continuous benchmarking

This repository provides a continuous benchmarking.  
If your project needs benchmark monitoring, it collects performance data by this repository.

|Type|Badge/URL|
|---|---|
|CI/CD|[![ci](https://github.com/tomoyane/http-continuous-benchmarking/actions/workflows/ci.yml/badge.svg)](https://github.com/tomoyane/http-continuous-benchmarking/actions/workflows/ci.yml)|
|Go Report Card|[![Go Report Card](https://goreportcard.com/badge/github.com/tomoyane/http-continuous-benchmarking)](https://goreportcard.com/report/github.com/tomoyane/http-continuous-benchmarking)|
|Coveralls|[![Coverage Status](https://coveralls.io/repos/github/tomoyane/http-continuous-benchmarking/badge.svg?branch=main)](https://coveralls.io/github/tomoyane/http-continuous-benchmarking?branch=main)|
|Coverage report for GitHub Pages|[Coverage report GitHub Pages](https://tomoyane.github.io/http-continuous-benchmarking/#file0)
|Docker Registry|https://hub.docker.com/repository/docker/tomohito/http-continuous-benckmarking|

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
TODO

### Example benchmark result
#### Metrics report HTML
![Screen Shot 2021-02-28 at 11 16 19](https://user-images.githubusercontent.com/9509132/109417530-656e0180-7a07-11eb-922a-e6915d194eb8.png)

#### Log
```bash
$ ./http-continuous-benchmarking
HTTP request pattern according to the ratio = GET GET GET GET GET GET GET GET GET GET
Start time = 1614513241
(Thread-2): Start attack for duration 5 seconds
(Thread-1): Start attack for duration 5 seconds
(Thread-2): End attack
(Thread-1): End attack
Stats info GET request
Latency 99  percentile: 533.000000
Latency 95  percentile: 115.000000
Latency avg percentile: 91.000000
Latency max percentile: 533.000000
Latency min percentile: 66.000000
Request per seconds:    24.000000

(Thread-2): Start attack for duration 5 seconds
(Thread-1): Start attack for duration 5 seconds
(Thread-1): End attack
(Thread-2): End attack
Stats info GET request
Latency 99  percentile: 143.000000
Latency 95  percentile: 127.000000
Latency avg percentile: 88.000000
Latency max percentile: 151.000000
Latency min percentile: 64.000000
Request per seconds:    24.000000

End time = 1614513260
```
