rm go.mod
go mod init gocashflow
go mod tidy
swag init 
go run main.go
go get "go.mongodb.org/mongo-driver/bson"
go install "go.mongodb.org/mongo-driver/bson"

***************
migrate on mongo

use db

db.customers.updateMany({ status: { $exists: false } },{ $set: { status: true } })
***************
