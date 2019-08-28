# Query Exporter
This is a simple tool written in go to generate pipe separated text file from database query(ies).

### Configuration Sections
* Database
  * Connection
    * string
    * SQL Server (only database for now) connection string
  * Provider
    * string
    * mssql (currently)
      * whatever refers to library MUST import appropriate driver
        * currently using https://github.com/denisenkom/go-mssqldb
        * install via "go get github.com/denisenkom/go-mssqldb"
        * change this in main.go for other library
  * Build SQL
    * string
    * 1st query run that will build/load the table with records
    * OPTIONAL
  * Get SQL
    * string
    * 2nd query run that will get the records
    * REQUIRED
  * Get Fields
    * string array
      * field names
      * currently only names... thinking about switching array of objects
    * OPTIONAL
      * if NOT specified, all fields will be returned
  * Set SQL
    * string
    * 3rd query run that marks the records written
    * OPTIONAL
* Output Path
  * string
  * directory/path where file will be generated

### Notes
* Placing a file name and path with the .sql extension in Build, Get, or Set SQL strings will cause the configuration to attempt to retrieve the text from the file as the SQL statement.

### ODDITIES
* **TCP/IP MUST be enabled in the SQL Server Configuration Manager for the go-mssqldb driver**
* ~~The connection string must include the IP Address (as the hostname will NOT be properly resolved)~~

# License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
