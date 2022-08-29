###The following command creates a go.mod file in which dependencies you add will be listed for tracking
go mod init example/Restexample

###To run main.go use the following command
go get .
#After executing the above commend the server will be started navigate to http://localhost:8080/Student  
#or just run the following command to get the details of the student using GET request

curl http://localhost:8080/Student

##Now open another terminal DO NOT CLOSE THE EXISTING SERVER
##pen new terminal and use the following command to use POST request to add details of student
curl http://localhost:8080/Student --include --header "Content-Type: application/json" --request "POST" --data '{"USN": "4","name": "Nethra","course": "Bcom","CGPA": 4.99}'


#After executing the above command navigate to http://localhost:8080/Student  
#or just run the following command to get "updated" details of student 
curl http://localhost:8080/Student

##TO view a student by USN navigate to http://localhost:8080/Student/<StudentUSN>  give the USN of the student
##or just run the following command
curl http://localhost:8080/Student/<StudentUSN>
