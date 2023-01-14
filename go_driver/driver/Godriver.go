package Godriver

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/lib/pq"
)

func scan(rows *sql.Rows) ([]map[string]interface{}, error) {
	// Get the column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Create the data slice
	data := make([]map[string]interface{}, 0)

	// Create a slice for the values
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for rows.Next() {
		// Assign pointers to the values slice
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		// Scan the row
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		// Create a map for the row
		row := make(map[string]interface{})

		// Add the values to the map
		for i, col := range columns {
			var v interface{}
			b, ok := values[i].([]byte)
			if ok {
				v = string(b)
			} else {
				v = values[i]
			}
			row[col] = v
		}

		// Add the row to the data slice
		data = append(data, row)
	}

	return data, nil
}
func Connect(user string, password string, dbname string) *sql.DB {
	// Connect to the database
	db, err := sql.Open("postgres", "user="+user+" "+"password="+password+" "+"dbname="+dbname+" "+"sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ReturnJSON(db *sql.DB, query string) []byte {
	// Prepare the query
	// query := "SELECT user_id, name FROM user_table"
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	data, err := scan(rows)
	if err != nil {
		return nil
	}

	// JSON encode the data
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	return jsonData
}

// JSON encode the data
