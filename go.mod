module aws-sqs-lambda

go 1.19

require github.com/aws/aws-lambda-go v1.34.1

replace (
	github.com/ShinyTrinkets/overseer => github.com/streamingfast/overseer v0.2.1-0.20210326144022-ee491780e3ef
	github.com/gorilla/rpc => github.com/streamingfast/rpc v1.2.1-0.20201124195002-f9fc01524e38
	github.com/graph-gophers/graphql-go => github.com/streamingfast/graphql-go v0.0.0-20210204202750-0e485a040a3c
)
