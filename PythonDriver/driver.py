import psycopg2
import json
from decimal import Decimal
def default(o):
        if isinstance(o, Decimal):
            return int(o)
        raise TypeError
class JSONDriver:
    def __init__(self):
        pass
    def connect(self,data):
        return JSONConnection(self,data)
    
class JSONConnection:
    def __init__(self,driver,data):
        self.driver = driver
        self.data = data
        self.conn = psycopg2.connect(host=self.data.get("host"),database=self.data.get("database"),user=self.data.get("user"),password=self.data.get("password"))
    def get_property_info(self):
        props = []
        prop = {"name": "user", "value": self.data.get("user"), "description": "Database user"}
        props.append(prop)
        prop = {"name": "password", "value": self.data.get("password"), "description": "Database password"}
        props.append(prop)
        return props
    def query_db(self,query,args=(),one=False):
        cur = self.conn.cursor()
        cur.execute(query,args)
        rows = [dict((cur.description[i][0], value) 
            for i, value in enumerate(row)) for row in cur.fetchall()] 
        rows = json.dumps(rows,default=default)
        return rows
    def close(self):
        print("Closing connection")
        return self.conn.cursor().connection.close() 
           
            

