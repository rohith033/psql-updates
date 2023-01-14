from driver import JSONDriver

def user_input(data):
    for key in data.keys():
        print(f'enter {key}')
        data[key] = str(input())

def make_query(conn):
    print("type the query")
    return conn.query_db(str(input()))


driver = JSONDriver()
data = {"host":"","database":"","user":"","password":""}
user_input(data)
conn = driver.connect(data)
result = make_query(conn)
print(result)
conn.close()
