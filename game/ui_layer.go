package game

import (
	"github.com/dusk125/pixelui"
)

type UILayer func(*pixelui.UI)

type UISetting struct {
	Render []int
}

type UILayerStack struct {
	Layers        []UILayer
	Settings      map[string]UISetting
	Active        *UISetting
	UI            *pixelui.UI
	SkipIteration bool
}

func NewUILayerStack(uiInstance *pixelui.UI) UILayerStack {
	return UILayerStack{
		Layers:        nil,
		Settings:      make(map[string]UISetting),
		Active:        nil,
		UI:            uiInstance,
		SkipIteration: false,
	}
}

func (s *UILayerStack) PushLayer(layer UILayer) {
	s.Layers = append(s.Layers, layer)
}

func (s *UILayerStack) PopLayer() {
	if len(s.Layers) == 0 {
		panic("[ERROR] There is no layer to pop from the stack.")
	}

	s.Layers = s.Layers[:len(s.Layers)-1]
}

func (s *UILayerStack) AddSetting(name string, setting UISetting) {
	s.Settings[name] = setting
}

func (s *UILayerStack) SetSetting(name string) {
	setting, ok := s.Settings[name]
	if ok {
		s.Active = &setting
	}
}

func (s *UILayerStack) Update() {

	renderPtr := 0
	renderStart := s.Active.Render[renderPtr]
	renderFinal := s.Active.Render[renderPtr+1]

	for i := range s.Layers {

		update := true
		if i < renderStart || i > renderFinal {
			update = false
		}

		if s.SkipIteration {
			s.SkipIteration = false
			break
		}

		if update {
			s.Layers[i](s.UI)
		}

		if i == renderFinal && renderPtr+1 < len(s.Active.Render)/2 {
			renderPtr += 1
			renderStart = s.Active.Render[renderPtr]
			renderFinal = s.Active.Render[renderPtr+1]
		}
	}
}
