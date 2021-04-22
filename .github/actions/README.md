## HTTP Continuous benchmarking

This repository provides a continuous benchmarking for GitHub Actions.  
If your project needs benchmark monitoring, it collects performance data by this repository.

## Concept
* Simple and lightweight benchmark tool
  * This tool is not recommended for if large-scale performance measurements
* Continuous benchmarking tool
  * Always monitor performance impacts for source code changes
* Lightweight to make various requests
  * The user does not do complicated things
* Create measured reports and notify warnings
  * Create a continuous report and associate it with a commit hash
  * Warn and notice when the threshold is reached
    
## Screenshots
![Screen Shot 2021-02-28 at 11 16 19](https://user-images.githubusercontent.com/9509132/109417530-656e0180-7a07-11eb-922a-e6915d194eb8.png)

## Input params
|input param name|description|example|
|---|---|---|
|target_url (※)|Request destination URL.|http(s)://xxxxxxx.com/api/v1/users|
|http_headers (※)|Request HTTP Headers.|{"Authorization": "Bearer xxx", "Content-Type": "application/json"}|
|thread_num (※)|Client thread num.|5|
|trial_num (※)|Number of trials to apply load per load duration.|5|
|req_http_method_ratio (※)|HTTP method ratio of request.|{"POST": 4, "GET": 6}|

## Run sample
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
