## HTTP Continuous benchmarking

This repository provides a continuous benchmarking.  
If your project needs benchmark monitoring, it collects performance data by this repository.

|Type|Badge/URL|
|---|---|
|CI|[![ci](https://github.com/tomoyane/http-continuous-benchmarking/actions/workflows/ci.yml/badge.svg)](https://github.com/tomoyane/http-continuous-benchmarking/actions/workflows/ci.yml)|
|Release|[![release](https://github.com/tomoyane/http-continuous-benchmarking/actions/workflows/release.yml/badge.svg)](https://github.com/tomoyane/http-continuous-benchmarking/actions/workflows/release.yml)|
|Go Report Card|[![Go Report Card](https://goreportcard.com/badge/github.com/tomoyane/http-continuous-benchmarking)](https://goreportcard.com/report/github.com/tomoyane/http-continuous-benchmarking)|
|Coveralls|[![Coverage Status](https://coveralls.io/repos/github/tomoyane/http-continuous-benchmarking/badge.svg?branch=main)](https://coveralls.io/github/tomoyane/http-continuous-benchmarking?branch=main)|
|Coverage report for GitHub Pages|[Coverage report GitHub Pages](https://tomoyane.github.io/http-continuous-benchmarking/#file0)
|Docker Registry|[http-continuous-benchmarking](https://hub.docker.com/r/tomohito/http-continuous-benckmarking)|

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

### How to use GitHub Actions
With input variable.  
**(※)** is required.

|With input name|Description|Example|
|---|---|---|
|target_url **(※)**|Request destination URL.|http(s)://xxxxxxx.com/api/v1/users|
|http_headers **(※)**|Request HTTP Headers. `{}` format.|{"Authorization": "Bearer xxx", "Content-Type": "application/json"}|
|thread_num **(※)**|Client thread num.|5|
|trial_num **(※)**|Benchmark trial number while 5seconds. If its 5times, the benchmark try 5times * 5seconds.|5 <br>(Ex: Case of API 100rps, 100(rps) * 5(seconds) * 5(times))|
|req_http_method_ratio **(※)**|HTTP method percentage of request.`{}` format.|{"POST": 4, "GET": 6}|
|req_body|HTTP Request Body. If you use PUT or PATCH or POST, its required.`{}` format.|{"email": "xx@gmail.com"}|

Sample GitHub actions workflow yaml.
```yaml
on: [push]

jobs:
  benchmarking:
    runs-on: ubuntu-latest
    name: Attack
    steps:
      - name: Benchmarking
        id: benchmarking
        uses: tomoyane/http-continuous-benchmarking@1.0.0
        with:
          target_url: 'https://example.com'
          http_headers: '{"Content-Type":"application/json"}'
          thread_num: '1'
          trial_num: '1'
          req_http_method_ratio: '{"GET": 10}'
      - name: Completed
        run: echo "Completed benchmarking"
```

### How to use simple application
[This repository application usage](https://github.com/tomoyane/http-continuous-benchmarking)

Application basic usage.
```bash
# Set required environment variable before execution
$ git clone https://github.com/tomoyane/http-continuous-benchmarking.git
$ cd http-continuous-benchmarking; go build
$ export INPUT_TARGET_URL='https://example.com' \
         INPUT_REQ_HTTP_METHOD_RATIO='{"GET":10}' \
         INPUT_HTTP_HEADERS='{"Content-Type":"application/json"}' \
         INPUT_THREAD_NUM=2 \
         INPUT_TRIAL_NUM=2
$ ./http-continuous-benchmarking
```

Docker image usage.
```bash
$ docker pull tomohito/http-continuous-benckmarking
$ docker run \
  --env INPUT_TARGET_URL='https://example.com' \
  --env INPUT_REQ_HTTP_METHOD_RATIO='{"GET":10}' \
  --env INPUT_HTTP_HEADERS='{"Content-Type":"application/json"}' \
  --env INPUT_THREAD_NUM=2 \
  --env INPUT_TRIAL_NUM=2 \
  -i tomohito/http-continuous-benckmarking
```

### Example benchmark result
#### Metrics report HTML
![Screen Shot 2021-02-28 at 11 16 19](https://user-images.githubusercontent.com/9509132/109417530-656e0180-7a07-11eb-922a-e6915d194eb8.png)

#### Log
```bash
$ ./http-continuous-benchmarking
HTTP request pattern according to the ratio = GET GET GET GET GET GET GET GET GET GET
Start warnmup for duration 5 seconds
End warmup

Start time = 1614653873
(Thread-3): Start attack for duration 5 seconds
(Thread-1): Start attack for duration 5 seconds
(Thread-2): Start attack for duration 5 seconds
(Thread-1): End attack
(Thread-2): End attack
(Thread-3): End attack

GET request stats information
Latency 99  percentile: 190 milliseconds
Latency 95  percentile: 172 milliseconds
Latency avg percentile: 115 milliseconds
Latency max percentile: 210 milliseconds
Latency min percentile: 72 milliseconds
Request per seconds:    30

(Thread-1): Start attack for duration 5 seconds
(Thread-3): Start attack for duration 5 seconds
(Thread-2): Start attack for duration 5 seconds
(Thread-1): End attack
(Thread-2): End attack
(Thread-3): End attack

GET request stats information
Latency 99  percentile: 153 milliseconds
Latency 95  percentile: 129 milliseconds
Latency avg percentile: 94 milliseconds
Latency max percentile: 192 milliseconds
Latency min percentile: 65 milliseconds
Request per seconds:    36

End time = 1614653894
```
