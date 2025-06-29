#!/run/current-system/sw/bin/bash

# strip the response from the http_req_dump_gemini.log to get the json only

count=0
while IFS= read -r line; do
  ((count++))
  [ $count -eq 11 ] && echo "$line" && break
done < ../http_req_dump_gemini.log

# run the test and preetify the json output
./http_req_test.sh | jq
