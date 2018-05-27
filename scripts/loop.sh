#!/bin/bash
for i in {1..1000}
do
  curl http://demo.fnproject.io
  sleep 0.15
done
