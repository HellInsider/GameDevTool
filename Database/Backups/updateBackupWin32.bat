"C:\Program Files\PostgreSQL\15\bin\psql.exe" -U postgres -a -f recreateDB.sql
"C:\Program Files\PostgreSQL\15\bin\pg_restore.exe" -U postgres -d GameDevDB < GameDevBackup