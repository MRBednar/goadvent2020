module github.com/MRBednar/GoAdvent2020/main

go 1.17

replace github.com/MRBednar/GoAdvent2020/s3connect => ./s3connect

require github.com/MRBednar/GoAdvent2020/s3connect v0.0.0-00010101000000-000000000000

require (
	github.com/aws/aws-sdk-go v1.40.27 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)
