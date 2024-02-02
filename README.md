# go-projects

## 1. Build A Simple Web Server With Go
 - link http://github.com/kojoluh/goprojects
## 2. Build A CRUD API With Go
 - link http://github.com/kojoluh/goprojects
## 3. Go With MySQL Book Management System 
 - link http://github.com/kojoluh/goprojects
## 4. Simple SlackBot To Calculate Age 
 - link http://github.com/kojoluh/goprojects
## 5. Go Slackbot for File Uploading 
 - link http://github.com/kojoluh/goprojects
## 6. Email Verifier Tool With Go 
 - link http://github.com/kojoluh/goprojects
## 7. AWS Lambda With Go 
 - link http://github.com/kojoluh/goprojects

 1. #create a role from cli for the lambda
 - $ aws iam create-role --role-name go-lambda-ex --assume-role-policy-document '
{"Version": "2012-10-17","Statement":[{"Effect": "Allow", "Principal":"{"Service":"lambda.amazonaws.com"}, "Action": "sts:AssumeRole"}]}' 
- alternatively you can add this json policy part to a file trust-policy.json and link the file 
with command as below (ensure you run the command from the same path as file):
- $ aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json 

- next, run the following command to setup the lmbda role arn
$ aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

- build your main.go file
$ go build main.go

- zip the main executable
$ zip function.zip main

- create the lambda by attaching the zip file
$ aws lambda create-function --function-name go-lambda-function \
--zip-file fileb://function-zip --handler main --runtime go1.x \
--role arn:aws:iam::xxxxxx:role/lambda-ex

- invoke the lambda function by running the comman below:
$ aws lambda invoke --function-name go-lambda-function --cli-binary-format raw-in-base64-out --payload '{"What is your name": "Foster Kojo", "How old are you?": 33}' output.txt



## 8. CRM with Go Fiber 
 - link http://github.com/kojoluh/goprojects
## 9. HRMS with Go Fiber 
 - link http://github.com/kojoluh/goprojects
## 10. Complete Serverless Stack 
 - link http://github.com/kojoluh/goprojects
## 11. A.I Bot with Wolfram, Wit.ai and Go
 - link http://github.com/kojoluh/goprojects

------------------
Pointers
Concurrency
Goroutines
mutex
Structs and Struct methods
