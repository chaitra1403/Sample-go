###The following commend creates a go.mod file in which dependencies you add will be listed for tracking
go mod init example/Restexample

###To run main.go use the following commend
go get .
#After executing the above commend the server will be started navigate to http://localhost:8080/Student  
#or just run the following commend to get the details of student using GET request

curl http://localhost:8080/Student

##Now open another terminal DONOT CLOSE THE EXISTING SERVER
##pen new terminal and use the following commend to use POST request to add details of student
curl http://localhost:8080/Student --include --header "Content-Type: application/json" --request "POST" --data '{"USN": "4","name": "Nethra","course": "Bcom","CGPA": 4.99}'


#After executing the above commend navigate to http://localhost:8080/Student  
#or just run the following commend to get "updated" details of student 
curl http://localhost:8080/Student

##TO view student by USN navigate to http://localhost:8080/Student/<StudentUSN>  give the USN of student
##or just run the following commend
curl http://localhost:8080/Student/<StudentUSN>

