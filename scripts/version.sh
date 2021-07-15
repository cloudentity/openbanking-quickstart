#!/bin/bash

type=$(git branch | grep "\\*" | cut -d ' ' -f2 | cut -d '/' -f1)

if [ "$type" = "release" ] || [ "$type" = "hotfix" ] || [ "$type" = "feature" ] || [ "$type" = "bugfix" ] ; then
  git branch | grep "\\*" | cut -d ' ' -f2 | cut -d '/' -f2
else
  echo "latest"
fi
