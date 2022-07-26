package provider_test

type Database struct {
	Name string
}

type DatabasePostgreeSql Database
type DatabaseMongoDB Database

func NewDatabasePostgreeSQL() *DatabasePostgreeSql {
	return (*DatabasePostgreeSql)(&Database{Name: "Postgree SQL"})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "Mongo DB"})
}

type DatabaseRepository struct {
	DatabasePostgreeSQL *DatabasePostgreeSql
	DatabaseMongoDB     *DatabaseMongoDB
}

func NewDatabaseRepository(databasePostgreeSQL *DatabasePostgreeSql, databaseMongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreeSQL: databasePostgreeSQL, DatabaseMongoDB: databaseMongoDB}
}
