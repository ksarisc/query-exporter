# query-exporter
Simple tool in go to generate text file from database query

### Configuration Sections
* Database
  * Connection
    * string
    * SQL Server (only database for now) connection string
  * BuildSql
    * string
    * 1st query run that will build/load the table with records
    * optional: if empty 
  * GetSql
  * SetSql
* OutputPath
