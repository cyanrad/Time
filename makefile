path = "postgresql://root:secret@localhost:5432/main?sslmode=disable"
migratepath = ./backend/SQL

# migrates database into newest iteration 
migrateup:
	migrate -path ${migratepath} -database ${path} -verbose up

# migrates database down (complete reset)
migratedown:
	migrate -path ${migratepath} -database ${path} -verbose down
	
migratenew:
	migrate create -ext sql -dir ./backend/SQL -seq schema

.PHONY: migrateup migratedown migratenew