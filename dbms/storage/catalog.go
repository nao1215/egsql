package storage

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"path/filepath"
	"sync"

	"github.com/nao1215/egsql/dbms/meta"
	"github.com/nao1215/egsql/misc/errfmt"
)

// jsonMarshal is a variable used to change json.Marshal to a stub at test time.
var jsonMarshal = json.Marshal

// catalogName os catlog file name.
const catalogName = "catalog.db"

// Catalog is the top-level structure for data storage.
// Data storage is organized in units of catalogs, schemas, and tables, in order from top to bottom
type Catalog struct {
	Schemes []*meta.Scheme
	mutex   *sync.RWMutex
}

// NewEmtpyCatalog return Catalog pointer. Only setup mutex, not setup any schema.
func NewEmtpyCatalog() *Catalog {
	return &Catalog{
		mutex: &sync.RWMutex{},
	}
}

// LoadCatalog reads a catalog file and returns its contents as a Catalog pointer.
// If the catalog file does not exist, a new empty catalog pointer is returned.
func LoadCatalog(egsqlHomePath string) (*Catalog, error) {
	b, err := ioutil.ReadFile(path.Join(egsqlHomePath, catalogName))
	if err != nil {
		return NewEmtpyCatalog(), nil
	}

	var catalog Catalog
	err = json.Unmarshal(b, &catalog)
	if err != nil {
		return nil, errfmt.Wrap(ErrParseCatalogFile, err.Error())
	}

	catalog.mutex = &sync.RWMutex{}
	return &catalog, nil
}

// SaveCatalog persists the system catalog as `catalog.db`.
// `catalog.db` has a simple json format like key/value.
func SaveCatalog(egsqlHomePath string, c *Catalog) (err error) {
	jsonStr, err := jsonMarshal(c)
	if err != nil {
		return errfmt.Wrap(ErrSaveCatalogFile, err.Error())
	}

	err = ioutil.WriteFile(filepath.Join(egsqlHomePath, catalogName), jsonStr, 0644)
	if err != nil {
		return errfmt.Wrap(ErrSaveCatalogFile, err.Error())
	}
	return nil
}

// Add is to add the new scheme into a memory.
// Be careful not to persist the disk.
func (c *Catalog) Add(scheme *meta.Scheme) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Schemes = append(c.Schemes, scheme)
}

// HasScheme returns whether a schema with the specified table name exists.
func (c *Catalog) HasScheme(tableName string) bool {
	return c.FetchScheme(tableName) != nil
}

// FetchScheme returns the schema with the specified table name,
// if one exists. If no schema exists, nil is returned.
func (c *Catalog) FetchScheme(tableName string) *meta.Scheme {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// TODO : Fix O(n) to O(1)
	for _, s := range c.Schemes {
		if s.TableName == tableName {
			return s
		}
	}
	return nil
}
