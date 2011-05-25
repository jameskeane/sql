package sql

import "http"


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


func Connect(dsn string) (Connection, Error) {
	url, err := http.ParseURLReference(dsn)
	if err != nil {
		return nil, NewError("Invalid DSN URL: " + dsn)	
	}

	driver, found := drivers[url.Scheme]
	if !found {
		return nil, NewError("No driver found: " + url.Scheme)	
	}

	return driver.Connect(url)
}

func initDriversMap() {
	if drivers == nil {
		drivers = make(map[string]Driver)	
	}
}

func init() {
	initDriversMap()
}
