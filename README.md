# p7
Open source microservice platform

# Example
./p7-cli create lambda --payload '{"name": "test", "code": "exports.handler = function(event, context, cb) {\n  console.log(process.execPath)\n  console.log(process.execArgv)\n  console.log(event)\n  console.log(context)\n  context.callbackWaitsForEmptyEventLoop = false\n  console.log(context.getRemainingTimeInMillis())\n  cb()\n}"}

./p7-cli list lambda
http :8888/p7/lambdas

./p7-cli code lambda --lambda_id 1
:8888/p7/lambdas/1/actions/code

curl -X POST --header 'Content-Type: application/json' --header 'Accept: text/plain' -d '{ \
   "codePath": "http://localhost:8888/p7/lambdas/1/actions/code", \
   "provider": "aws", \
   "tag": "p7/testimage" \
 }' 'http://localhost:8890/image/images'
 http POST :8890/image/images codePath="http://localhost:8000/test.js" provider="aws" tag="p7/testimage:0.0.1"

 # Debug
 docker run --rm -it --entrypoint=/bin/bash --user root p7/testimage
