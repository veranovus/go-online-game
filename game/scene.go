package game

import "log"

// SceneStates
const (
	SceneStateOk = iota
	SceneStateFail
	SceneStateChange
)

type SceneState int

type SceneInterface interface {
	GetName() string

	GetID() int
	SetID(int)

	Load() bool
	UnLoad()
	Update(float64) SceneState
}

type Scene struct {
	Name string
	ID   int
}

func NewScene(name string) Scene {
	return Scene{
		Name: name,
	}
}

type SceneManager struct {
	Scenes []SceneInterface
	Active SceneInterface
	NextID int
}

func NewSceneManager() SceneManager {
	return SceneManager{
		Scenes: nil,
		Active: nil,
		NextID: SceneStateChange + 1,
	}
}

// AddScene adds a new scene to the list and gives that scene a new id
func (sm *SceneManager) AddScene(scene SceneInterface) {
	scene.SetID(sm.NextID)
	sm.NextID += 1

	sm.Scenes = append(sm.Scenes, scene)
}

// RemoveScene is used the free the memory allocated for the scene,
// doesn't really remove the scene from the list
func (sm *SceneManager) RemoveScene(id int) {
	if id >= len(sm.Scenes) || id < 0 {
		log.Fatal("[ERROR] Failed to remove scene because id was out of range.")
	}

	sm.Scenes[id-SceneStateChange] = nil
}

func (sm *SceneManager) SetScene(name string) bool {
	for _, s := range sm.Scenes {
		if s.GetName() == name {

			// Unload the previous scene
			if sm.Active != nil {
				sm.Active.UnLoad()
			}

			// Set and load the new scene
			sm.Active = s
			sm.Active.Load()

			return true
		}
	}

	log.Println("[WARNING] Failed to set the scene to '{}.'", name)
	return false
}

func (sm *SceneManager) Update(dt float64) {
	sm.Active.Update(dt)
}
