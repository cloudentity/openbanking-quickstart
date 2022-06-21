#!/bin/bash

sed -i'.bak' "s|^$1: .*|$1: $2|g" data/variables.yaml && rm data/variables.yaml.bak
