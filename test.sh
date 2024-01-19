#!/bin/bash

API_ENDPOINT="http://localhost:3322/course-all"  # Replace with your actual API endpoint
CONCURRENT_REQUESTS=10000  # Adjust the number of concurrent requests
BEARER_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFobWFkaXNhYWRyaWFudXNAZ21haWwuY29tIiwiZXhwIjoxNzA1NjcyNDA0LCJyb2xlX25hbWUiOiJhZG1pbmlzdHJhdG9yIn0.h6BsdGHK8ovnW89eZL0cBWNwhnPXKmWcJtyTnMYOMVY"  # Replace with your actual Bearer token

for ((i=1; i<=$CONCURRENT_REQUESTS; i++)); do
    curl -s -H "Authorization: Bearer $BEARER_TOKEN" "$API_ENDPOINT" >/dev/null &
done

wait

