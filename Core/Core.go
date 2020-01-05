package Core

// define all of the basic interfaces here

type Owner int

const (
	SYSTEM Owner = iota
	SERVER
	LOCAL
	REMOTE
)

type Entity struct {
	Active      bool
	Id          string
	EntityOwner Owner
}

type ComponentStore struct {
	// add in a map for each type of component
	Transforms    map[string]Transform
	AudioSources  map[string]AudioClip
	RigidBodies   map[string]RigidBody
	MeshRenderers map[string]MeshRenderer
}

type EntityStore struct {
	Entities map[string]Entity
}

type System interface {
	Simulate() error
}
