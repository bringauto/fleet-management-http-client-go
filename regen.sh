#!/bin/bash

npx @openapitools/openapi-generator-cli generate -i api/openapi.yaml \
 --generator-name go \
 --git-user-id bringauto \
 --git-repo-id fleet-management-http-client-go \
 --output .