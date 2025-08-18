package ioc

var containers = []Container{
	Api,
	Controller,
	Config,
	Default,
}

func Init() {
	for _, c := range containers {
		if err := c.Init(); err != nil {
			panic(err)
		}
	}
}

type Container interface {
	Registry(name string, obj Object)
	Get(name string) Object
	Init() error
}

type Object interface {
	Init() error
}

type ObjectImpl struct{}

func (o *ObjectImpl) Init() error {
	return nil
}
