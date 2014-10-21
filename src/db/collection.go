// Collection definition and methods to work with collection objects
package db

import . "types"

const (
	FIELD_ID      = "id"
	FIELD_VERSION = "__version__"
)

type Hash [20]byte // SHA1 hash

type ObjectIndex map[Hash][]int

// Object instance
type ObjectVersion struct {
	Offset int // Offset from the beginning of object data file
	Len    int // Number of bytes to read
}

// List of object instances. index is version number.
type ObjectPointer map[int]ObjectVersion

type Collection struct {
	Name             string                 // Collection/class name
	Objects          map[int]ObjectPointer  // Objects. map index is object ID
	Indices          map[string]ObjectIndex // collection of indices
	DataFile         *DbFile                // Objects storage
	IndexFile        map[string]*DbFile     // List of indices
	freeSlotOffset   int
	IndexPointerFile string
	ObjectIndexFlush chan (bool)
}

// Returns pointer to the start of unallocated file space
func (c *Collection) getFreeSpaceOffset() int {
	return c.freeSlotOffset
}

// Returns the number of object versions
func (c *Collection) GetObjectVersions(p GetObjectVersions) (ObjectVersions, error) {
	return ObjectVersions(len(c.Objects[p.Id])), nil
}
