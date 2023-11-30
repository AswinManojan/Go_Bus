package interfaces

import "gobus/entities"

type AdminRepository interface {
	FindUserById(id int) (*entities.User, error)
	FindUserByEmail(mail string) (*entities.User, error)
	FindAllUsers() ([]*entities.User, error)
	EditUser(id int, user *entities.User) (*entities.User, error)
	DeleteUser(id int) (*entities.User, error)
	BlockUser(id int) (*entities.User, error)
	UnBlockUser(id int) (*entities.User, error)
	FindProviderById(id int) (*entities.ServiceProvider, error)
	FindAllProviders() ([]*entities.ServiceProvider, error)
	EditProvider(id int, provider *entities.ServiceProvider) (*entities.ServiceProvider, error)
	DeleteProvider(id int) (*entities.ServiceProvider, error)
	BlockProvider(id int) (*entities.ServiceProvider, error)
	UnBlockProvider(id int) (*entities.ServiceProvider, error)
	FindStationById(id int) (*entities.Stations, error)
	FindStationByName(name string) (*entities.Stations, error)
	FindAllStations() ([]*entities.Stations, error)
	EditStation(id int, station *entities.Stations) (*entities.Stations, error)
	DeleteStation(id int) (*entities.Stations, error)
	AddStation(station *entities.Stations) (*entities.Stations, error)
}
