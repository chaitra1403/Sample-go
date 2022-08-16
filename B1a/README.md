#I am considering "Getting user details "as an example

###Installation guide:

#Step 1: open VSCode / pycharm or any editor which supports Python and open the folder which you unzipped and then run the following command
    pip install requirements.txt

#Step 2: in the terminal run the app.py file using the following command
    python3 app.py

#Step 3:The server will be started no browse in the browser paste the below URL and run
    http://127.0.0.1:5000
#you will see the welcome screen

#Step 4:To get a list of users navigate to  http://127.0.0.1:5000/users
#Step 5:To get a single user to navigate to  http://127.0.0.1:5000/users/<userId>

#######UPDATED###########

#Step 6 To Create a new user do the following
    First Run: Python3 app.py
    Then open another terminal don't close the server and run: 
        curl -X POST http://localhost:5000/users -d "<the details of new users can be given here>"
    Now navigate to http://127.0.0.1:5000/users you should be able to see the new user details


