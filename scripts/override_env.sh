#!/bin/bash

sed -i'.bak' "s|^$1=.*|$1=$2|g" .env && rm .env.bak
