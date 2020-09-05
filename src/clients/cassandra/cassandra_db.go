package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// FIXME: move out to env file or config file
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

// GetSession returns the cassandra session
func GetSession() *gocql.Session {
	return session
}
