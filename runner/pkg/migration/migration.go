package migration

type Migration struct {
	Id             int
	Status         int
	Host           string
	Port           int
	Database       string
	Table          string
	AlterStatement string
}
