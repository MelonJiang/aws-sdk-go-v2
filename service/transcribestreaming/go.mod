module github.com/aws/aws-sdk-go-v2/service/transcribestreaming

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.16.7
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.3
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.14
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.8
	github.com/aws/smithy-go v1.12.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream => ../../aws/protocol/eventstream/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
