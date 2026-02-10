#!/bin/bash

npx @openapitools/openapi-generator-cli@2.28.3 generate -i api/openapi.yaml \
 --generator-name go \
 --git-user-id bringauto \
 --git-repo-id fleet-management-http-client-go \
 --output .