package ioc

var Api = NewMapContainer("api")
var Controller = NewMapContainer("controller")
var Config = NewMapContainer("config")
var Default = NewMapContainer("default")

func NewMapContainer(name string) Container {
	return &MapContainer{
		name:    name,
		storage: map[string]Object{},
	}
}

// ioc容器
type MapContainer struct {
	name    string
	storage map[string]Object
}

func (c *MapContainer) Get(name string) Object {
	return c.storage[name]
}
func (c *MapContainer) Registry(name string, obj Object) {
	c.storage[name] = obj
}

func (m *MapContainer) Init() error {
	for _, v := range m.storage {
		if err := v.Init(); err != nil {
			return err
		}
	}

	return nil
}
