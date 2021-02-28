## HTTP Continuous benchmarking

This repository provides a continuous benchmarking.  
If your project needs benchmark monitoring, it collects performance data by this repository.

|Workflow|Badge|
|---|---|
|CI||

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
