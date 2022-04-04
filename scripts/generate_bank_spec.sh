docker-compose -f docker-compose.acp.local.yaml exec runner sh -c  \
    "export SWAGGER_GENERATE_EXTENSION=false && swagger generate spec \
        --include-tag=$1 \
        -m \
        -o api/internal/bank-$1.yaml \
         -w ./apps/bank"