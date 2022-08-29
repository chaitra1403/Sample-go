##run the below command to download required packages 
##which are defined inside import
go run .
##if you get error downloading the package 
##try the below command
go get github.com/lestrrat/go-file-rotatelogs

##to run the program use the following command
go run . 
##or using
go run fileRotate.go

##after running the above command Files are rotated every second.
#Max age of the file is 10 seconds. So every file will be deleted after 10 seconds.
