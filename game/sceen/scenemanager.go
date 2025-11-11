package sceen

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game/input"
)

type Scene interface {
	Update(input *input.Input)
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current         Scene // currently active scene
	next            Scene // next scene to transition to
	transitionCount int   // remaining frames for scene transition (fade effect)
}

func NewSceneManager() *SceneManager {
	return &SceneManager{}
}

func (s *SceneManager) Update(i *input.Input) {
	s.current.Update(i)
}

func (s *SceneManager) Draw(r *ebiten.Image) {
	s.current.Draw(r)
}

// GoTo switches to the given scene.
func (s *SceneManager) GoTo() {
	s.current = NewGameScene()
}
