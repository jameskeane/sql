package sql

var drivers map[string]Driver


func RegisterDriver(name string, driver Driver) Error {
	// A drivers package may be loaded before init gets called
	// so we must ensure that the drivers map is initialized
	initDriversMap()

	// TODO: Should we allow more recently loaded drivers to 
	//       overwrite others?
	drivers[name] = driver
	return nil
}

func UnregisterDriver(name string) {
	drivers[name] = nil, false
}

func Connect(dsn_string string) (Connection, Error) {
	dsn, err := parseDSN(dsn_string)
	if err != nil {
		return nil, err
	}

	driver, found := drivers[dsn.Driver]
	if !found {
		return nil, NewError("sql: No driver found: " + dsn.Driver)	
	}

	return driver.Connect(dsn)
}

func initDriversMap() {
	if drivers == nil {
		drivers = make(map[string]Driver)	
	}
}

func init() {
	initDriversMap()
}
