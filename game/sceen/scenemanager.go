package sceen

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game"
	"github.com/xiao-dong-li/tennis/game/input"
)

var transitionTo = ebiten.NewImage(game.ScreenWidth, game.ScreenHeight)

type Scene interface {
	Update(input *GameState)
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current         Scene   // currently active scene
	next            Scene   // next scene to transition to
	transitionCount float32 // remaining frames for scene transition (fade effect)
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
	if s.transitionCount == 0 {
		s.current.Update(NewGameState(s, i))
		return
	}

	s.transitionCount--
	if s.transitionCount == 0 {
		s.current, s.next = s.next, nil
	}
}

func (s *SceneManager) Draw(r *ebiten.Image) {
	s.current.Draw(r)

	if s.transitionCount == 0 || s.next == nil {
		return
	}

	transitionTo.Clear()
	s.next.Draw(transitionTo)

	alpha := 1 - s.transitionCount/game.TransitionFrames
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(alpha)
	r.DrawImage(transitionTo, op)
}

// GoTo switches to the given scene.
func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = game.TransitionFrames
	}
}
