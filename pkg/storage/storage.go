package storage

import (
	"context"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Storage
type Storage struct {
	database string
	client   driver.Client
}

// Opts
type Opts struct {
	Endpoint string
	User     string
	Password string
	Database string
}

// Connect connect to database
func Connect(opts *Opts) (*Storage, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{opts.Endpoint},
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(opts.User, opts.Password),
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	storage := &Storage{database: opts.Database, client: client}

	if err = storage.Init(); err != nil {
		return nil, err
	}

	return storage, nil
}

// Close close connection with database
func (s *Storage) Close() {}

// Init initialize storage
func (s *Storage) Init() error {
	var (
		db driver.Database
	)

	if err := s.InitDatabase(&db); err != nil {
		return err
	}

	if err := s.InitCollections(&db); err != nil {
		return err
	}

	return nil
}

// InitDatabase init storage database
func (s *Storage) InitDatabase(db *driver.Database) error {
	ctx := context.Background()

	ok, err := s.client.DatabaseExists(ctx, s.database)
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		*db, err = s.client.CreateDatabase(ctx, s.database, nil)
		if err != nil {
			log.Error(err)
			return err
		}
	} else {
		*db, err = s.Database()
		if err != nil {
			return err
		}
	}

	return nil
}

// InitCollections init collections
func (s *Storage) InitCollections(db *driver.Database) error {
	ctx := context.Background()

	ok, err := (*db).CollectionExists(ctx, CAccount)
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		_, err = (*db).CreateCollection(ctx, CAccount, nil)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	ok, err = (*db).CollectionExists(ctx, CCertificates)
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		_, err := (*db).CreateCollection(ctx, CCertificates, nil)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

// Clean clean database
func (s *Storage) Clean() error {
	ctx := context.Background()

	db, err := s.Database()
	if err != nil {
		return err
	}

	err = db.Remove(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Database returns database
func (s *Storage) Database() (driver.Database, error) {
	ctx := context.Background()

	db, err := s.client.Database(ctx, s.database)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}

// Collection returns collection
func (s *Storage) Collection(name string) (driver.Collection, error) {
	ctx := context.Background()

	db, err := s.Database()
	if err != nil {
		return nil, err
	}

	col, err := db.Collection(ctx, name)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return col, nil
}

// Exec exec query and returns document meta
func (s *Storage) Exec(query string, vars map[string]interface{}, document interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()

	db, err := s.Database()
	if err != nil {
		return nil, err
	}

	cursor, err := db.Query(driver.WithQueryCount(ctx), query, vars)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cursor.Close()

	if cursor.Count() == 0 {
		return nil, errors.NotFound
	}

	for {
		meta, err := cursor.ReadDocument(ctx, document)
		if err != nil {
			if driver.IsNoMoreDocuments(err) {
				return &meta, nil
			}
			log.Error(err)
			return nil, err
		}
	}
}

func (s *Storage) Read(key, collection string, document interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()

	db, err := s.Database()
	if err != nil {
		return nil, err
	}

	col, err := db.Collection(ctx, collection)
	if err != nil {
		return nil, nil
	}

	meta, err := col.ReadDocument(ctx, key, document)
	if err != nil {
		if driver.IsNotFound(err) {
			return nil, errors.NotFound
		}

		return nil, err
	}

	return &meta, nil
}

// Write write document to collection
func (s *Storage) Write(collection string, document interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()

	col, err := s.Collection(collection)
	if err != nil {
		return nil, err
	}

	meta, err := col.CreateDocument(ctx, document)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &meta, nil
}

// Update update document by collection and key
func (s *Storage) Update(collection, key string, document interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()

	col, err := s.Collection(collection)
	if err != nil {
		return nil, err
	}

	meta, err := col.UpdateDocument(ctx, key, document)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &meta, nil
}
