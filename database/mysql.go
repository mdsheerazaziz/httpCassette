package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dbConn  *sql.DB = nil

/**
 * Check the connection and make it if it is lost or not opened yet.
 * @return	error
 */

func checkConnection() error {
	if dbConn !=nil {
		return nil
	}
	conn, err := openDatabase()
	if err != nil {
		return err
	}
	dbConn = conn
	return nil
}
/**
 * openDatabase function opens a MySql database connection
 * @param 	*sql.DB
 * @param	Error
 */
func openDatabase() (*sql.DB, error){
	db, err := sql.Open("mysql", config.Getenv("CONNECTION_STRING"))
	if err != nil {
		utilities.Log("MySQL: DB connection failed - " + err.Error())
	}
	utilities.Log("MySQL: DB connection successful")
	return db, nil
}


/**
 * SqlQuery function runs a select query to MySql
 * @param 	string
 * @param	[]string
 */
func SqlQuery(sqlQuery string, sqlArgument []string) (*sql.Rows, error) {

	err := checkConnection()
	if err != nil {
		utilities.Log("MySQL: DB Query failed. " + err.Error())
		return nil, err
	}
	//	defer db.Close()

	sqlInterface 	:= utilities.StringArrayToInterface(sqlArgument)

	// query
	rows, err 	:= dbConn.Query(sqlQuery, sqlInterface...)
	if err != nil {
		utilities.Log("MySQL: SQL query failed. "+ err.Error())
		return nil, err
	}

	return rows, nil
}

/**
 * SqlUpdate function runs a update query to MySql
 * @param 	string
 * @param	[]string
 */
func SqlUpdate(sqlQuery string, sqlArgument []string) (int, error) {

	err := checkConnection()
	if err != nil {
		utilities.Log("MySQL: DB Update failed. " + err.Error())
		return 0, err
	}
	//	defer db.Close()

	sqlInterface := utilities.StringArrayToInterface(sqlArgument)

	// Prepare
	stmt, err := dbConn.Prepare(sqlQuery)
	if err != nil {
		utilities.Log("MySQL: SQL update prepare failed. "+ err.Error())
		return 0, err
	}

	// Execute
	res, err := stmt.Exec(sqlInterface...)
	if err != nil {
		utilities.Log("MySQL: SQL update execute failed. "+ err.Error())
		return 0, err
	}

	// Get rows affected
	affect, err := res.RowsAffected()
	if err != nil {
		utilities.Log("MySQL: SQL update rows affected failed. "+ err.Error())
		return 0, err
	}

	return int(affect), nil
}


/**
 * SqlDelete function runs a delete query to MySql
 * @param 	string
 * @param	[]string
 */
func SqlDelete(sqlQuery string, sqlArgument []string) (int, error) {

	err := checkConnection()
	if err != nil {
		utilities.Log("MySQL: DB Delete failed. " + err.Error())
		return 0, err
	}
	//	defer db.Close()

	sqlInterface := utilities.StringArrayToInterface(sqlArgument)

	// Prepare
	stmt, err := dbConn.Prepare(sqlQuery)
	if err != nil {
		utilities.Log("MySQL: SQL delete prepare failed. "+ err.Error())
		return 0, err
	}

	// Execute
	res, err := stmt.Exec(sqlInterface...)
	if err != nil {
		utilities.Log("MySQL: SQL delete execute failed. "+ err.Error())
		return 0, err
	}

	// Get rows affected
	affect, err := res.RowsAffected()
	if err != nil {
		utilities.Log("MySQL: SQL delete rows affected failed. "+ err.Error())
		return 0, err
	}

	return int(affect), nil
}

func SqlGetRecord(sqlQuery string, sqlArgument []string, scanned ...interface{}) error {

	// Get the rows
	rows, err := SqlQuery(sqlQuery, sqlArgument)
	if err != nil {
		return err
	}

	// Close Query connection in end
	defer rows.Close()

	// Try to scan the row
	if rows.Next() {
		err = rows.Scan(scanned...)
		if err != nil {
			return err
		}
	}

	// If record not found
	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
