package models

import (
	// "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	uuid "github.com/satori/go.uuid"
)

/*Student struct
 */
type Student struct {
	Name  string    `bson:"name"`
	Grade int       `bson:"grade"`
	UUID  uuid.UUID `bson:"uuid"`
}
