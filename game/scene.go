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
	Name        string
	ID          int
	ReturnState SceneState
}

func NewScene(name string) Scene {
	return Scene{
		Name:        name,
		ReturnState: SceneStateOk,
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
		NextID: SceneStateChange,
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
// NOTE: This function is temporarily inactive.
func (sm *SceneManager) RemoveScene(id int) {
	if id >= len(sm.Scenes) || id < 0 {
		log.Fatal("[ERROR] Failed to remove scene because id was out of range.")
	}

	panic("[ERROR] SceneManager@RemoveScene is temporarily disabled.")
	//sm.Scenes[id-SceneStateChange] = nil
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

func (sm *SceneManager) SetSceneWithID(id int) bool {
	id = id - SceneStateChange

	if id < 0 || id >= len(sm.Scenes) {
		log.Fatal("[ERROR] Failed to set scene because id was out of range.")
		return false
	}

	// Unload the previous scene
	if sm.Active != nil {
		sm.Active.UnLoad()
	}

	// Set and load the current scene
	sm.Active = sm.Scenes[id]
	sm.Active.Load()

	return true
}

func (sm *SceneManager) Update(dt float64) {
	result := sm.Active.Update(dt)
	if result >= SceneStateChange {
		sm.SetSceneWithID(int(result))
	}
}
