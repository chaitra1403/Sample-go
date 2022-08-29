##To get started initialize your local project by running the following Go command.
go mod init aws_s3

##The following commands used to retrieve the standard set of SDK modules.
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3

##in your aws create S3 bucket named awsgolang, make the bucket acl enabled create the bucket and add few files in it


##After successfully creating the bucket create the following files:
##1.credentials file: add the following content to your credentials file, replacing <YOUR_ACCESS_KEY_ID> and <YOUR_SECRET_ACCESS_KEY> with your credentials
aws_access_key_id = <YOUR_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY> 
##2.config file: 
[profile chaitra]
region=us-east-1
##you can replace your name with chaitra and any region which you want and same should be canged in the code

##run the following code to set environment for aws sdk for golang
export AWS_REGION=us-east-1
export AWS_SDK_LOAD_CONFIG=true

##Finally Run the code
go run main.go awsgolang

#you can see objects of your s3 bucket
