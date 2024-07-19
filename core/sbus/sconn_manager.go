package sbus

type SConnManager interface {
	Add(SConnection)                                                        // Add connection
	Remove(SConnection)                                                     // Remove connection
	Get(uint64) (SConnection, error)                                        // Get a connection by ConnID
	Get2(string) (SConnection, error)                                       // Get a connection by string ConnID
	Len() int                                                               // Get current number of connections
	ClearConn()                                                             // Remove and stop all connections
	GetAllConnID() []uint64                                                 // Get all connection IDs
	GetAllConnIdStr() []string                                              // Get all string connection IDs
	Range(func(uint64, SConnection, interface{}) error, interface{}) error  // Traverse all connections
	Range2(func(string, SConnection, interface{}) error, interface{}) error // Traverse all connections 2
}
