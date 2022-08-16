from flask import Flask, jsonify, request
app= Flask(__name__)
userDb={1:{'name':'Chaitra','address':'Tumkur','phnno':'986543333'},
2:{'name':'Nethra','address':'Bangalore','phnno':'7539518240'},
3:{'name':'Suhas','address':'Mysore','phnno':'8523697410'}}
     
@app.route("/",methods=['GET'])
def connect():
    return "Welcome to Roost.ai"

@app.route("/users",methods=['GET'])
def get():
    return jsonify({'users':userDb})

@app.route("/users/<int:userId>",methods=['GET'])
def get_userDb(userId):
    if userId not in userDb:
        return "ERROR NO USER FOUND"
    else:
        return jsonify({userId:userDb[userId]})

@app.route("/users",methods=['POST'])
def create():
    note = request.get_data(as_text=True)
    idx = max(userDb.keys()) + 1
    userDb[idx] = note
    return jsonify({'users':userDb})

if __name__=="__main__":
    app.run(debug=True)
