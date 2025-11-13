package sceen

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game/input"
)

type Scene interface {
	Update(input *GameState)
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current         Scene // currently active scene
	next            Scene // next scene to transition to
	transitionCount int   // remaining frames for scene transition (fade effect)
}

type GameState struct {
	SceneManager *SceneManager
	Input        *input.Input
}

func NewSceneManager() *SceneManager {
	return &SceneManager{}
}

func NewGameState(sceneManager *SceneManager, i *input.Input) *GameState {
	return &GameState{
		SceneManager: sceneManager,
		Input:        i,
	}
}

func (s *SceneManager) Update(i *input.Input) {
	s.current.Update(NewGameState(s, i))
}

func (s *SceneManager) Draw(r *ebiten.Image) {
	s.current.Draw(r)
}

// GoTo switches to the given scene.
func (s *SceneManager) GoTo(scene Scene) {
	s.current = scene
}
