docker-compose exec runner sh -c  \
    "swagger generate spec \
        --include-tag=$1 \
        -m \
        -o api/internal/bank-$1.yaml \
         -w ./apps/bank"