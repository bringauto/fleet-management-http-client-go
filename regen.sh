#!/bin/bash

npx @openapitools/openapi-generator-cli generate -i api/openapi.yaml \
 --generator-name go \
 --output .