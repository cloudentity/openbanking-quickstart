#!/bin/bash

SCAN_RESULT=${1:-/tmp/scanresult.json}

jq -r '.alerts[].issues[] | [.severity, .impacted_artifacts[].infected_files[].display_name, .cve] | @tsv' \
    $SCAN_RESULT |  column -ts $'\t'

rm $SCAN_RESULT
