#!/bin/bash

ENABLE_MFA=$1
sed -i'.original' "s#ENABLE_MFA?=.*#ENABLE_MFA?=$ENABLE_MFA#" Makefile
rm -f Makefile.original