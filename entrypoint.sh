#!/bin/sh
echo $INPUT_TARGET_URL
export INPUT_TARGET_URL=${INPUT_TARGET_URL}
/tomoyane/http-continuous-benchmarking
