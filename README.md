# query-exporter
Simple tool in go to generate text file from database query

### Configuration Sections
* Database
  * Connection
    * string
    * SQL Server (only database for now) connection string
  * Provider
    * string
    * mssql (currently)
      * whatever refers to library MUST import appropriate driver
      * using https://github.com/denisenkom/go-mssqldb
      * install via "go get github.com/denisenkom/go-mssqldb"
  * Build SQL
    * string
    * 1st query run that will build/load the table with records
    * OPTIONAL
  * Get SQL
    * string
    * 2nd query run that will get the records
    * REQUIRED
  * Get Fields
    * NOT CURRENTLY DESIGNED
  * Set SQL
    * string
    * 3rd query run that marks the records written
    * OPTIONAL
  * Set Fields
    * NOT CURRENTLY DESIGNED
* Output Path
  * string
  * directory/path where file will be generated
