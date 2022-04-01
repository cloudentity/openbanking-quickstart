docker-compose -f docker-compose.acp.local.yaml exec -T runner sh -c \
    "go run ./scripts/go/clean_saas.go \
        -spec=$1 \
        -tenant=$2 \
        -cid=$3 \
        -csec=$4"
