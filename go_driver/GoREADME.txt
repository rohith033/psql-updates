hear i am updating the pq(github.com/lib/pq) package which is a pure go postgres driver 

setup :
    make sure you have pq(github.com/lib/pq) installed
    make a folder named driver and put Godriver.go file inside it
    driver folder should be in go/src folder
functions:
   in Godriver.go file there are two function
   1) Connect which takes  username,password and database as input returns a Databaseconnection
   2) Returnjson which takes Databaseconnection and Query as input and puts the data in json format

