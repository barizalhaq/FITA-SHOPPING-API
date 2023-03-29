package repository_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docker Suite")
}

var Db *gorm.DB
var cleanupDocker func()

const (
	dbName = "test"
	passwd = "test"
)

var _ = BeforeSuite(func() {
	// setup *gorm.Db with docker
	Db, cleanupDocker = setupGormWithDocker()
})

var _ = AfterSuite(func() {
	// cleanup resource
	cleanupDocker()
})

var _ = BeforeEach(func() {
	// clear db tables before each test
	err := Db.Exec(`DROP SCHEMA public CASCADE;CREATE SCHEMA public;`).Error
	Ω(err).To(Succeed())
})

func setupGormWithDocker() (*gorm.DB, func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}

	runDockerOpt := &dockertest.RunOptions{
		Repository: "postgres", // image
		Tag:        "14",       // version
		Env:        []string{"POSTGRES_PASSWORD=" + passwd, "POSTGRES_DB=" + dbName, "POSTGRES_USER=postgres", "listen_addresses = '*'"},
	}

	fnConfig := func(config *docker.HostConfig) {
		config.AutoRemove = true                                // set AutoRemove to true so that stopped container goes away by itself
		config.RestartPolicy = docker.RestartPolicy{Name: "no"} // don't restart container
	}

	resource, err := pool.RunWithOptions(runDockerOpt, fnConfig)
	if err != nil {
		panic(err)
	}
	// call clean up function to release resource
	fnCleanup := func() {
		err := resource.Close()
		if err != nil {
			panic(err)
		}
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://postgres:%s@%s/%s?sslmode=disable", passwd, hostAndPort, dbName)

	var gdb *gorm.DB
	// retry until db server is ready
	err = pool.Retry(func() error {
		gdb, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	// container is ready, return *gorm.Db for testing
	return gdb, fnCleanup
}
