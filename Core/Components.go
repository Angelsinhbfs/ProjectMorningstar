package Core

// all of the component types should be listed here
// all components are simple data objects with no functionality

type Vector2 struct {
	X float32
	Y float32
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

type Transform struct {
	Parent        Transform
	Position      Vector3
	Rotation      Vector4
	Scale         Vector3
	LocalPosition Vector3
	LocalRotation Vector3
	Children      []Transform
}

type AudioClip struct {
	Playing          bool
	AttenuationRange float32
	Volume           float32
	StartTime        float32
	EndTime          float32
	CurrentTime      float32
	// actual audio clip
}

type Animation struct {
	// some fbx or gltf data?
}

type AnimationClip struct {
	Anim        Animation
	StartTime   float32
	EndTime     float32
	CurrentTime float32
	Loop        bool
}

type RigidBody struct {
	Acceleration        Vector3
	Velocity            Vector3
	AngularAcceleration Vector4
	AngularVelocity     Vector4
}

type MeshRenderer struct {
}
