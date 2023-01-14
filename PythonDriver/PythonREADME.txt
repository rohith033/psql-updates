hear i am updating the psycopg2 package which is a python postgres driver 

setup :
    makesure psycopg2 is installed
    make sure both driver and main are in same folder if not chage the import in main accordingly
    we can import JSONDriver class like this "from driver import JSONDriver"
to run:
  just run the main.py file and connect to database 
functions:
   1)we can instantiate a new instance of JSONDriver by calling newdriver = JSONDriver()
   2) we can connect to any postgres database by calling newdriver.connect() which takes input in following format:
  "data = {"host":" ","databasename":" ","user":" ","password":" "}" it will return a connection to the postgres database
   3) connection.dbquery will take query as input and returns the json output
   4) connection.close will close the connection


   