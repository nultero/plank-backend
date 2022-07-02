#!/bin/env bash

# match
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"jackson":"turbulence"}' \
     http://localhost:9000/auth/logn

# not a match
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"jackson":"doggopants"}' \
     http://localhost:9000/auth/logn
