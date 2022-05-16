package game

import "github.com/inkyblackness/imgui-go"

func SetImGUIStyle() {

	imgui.PushStyleColor(imgui.StyleColorText, imgui.Vec4{1.00, 1.00, 1.00, 1.00})
	imgui.PushStyleColor(imgui.StyleColorTextDisabled, imgui.Vec4{0.50 + 0.19, 0.50 + 0.19, 0.50 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorWindowBg, imgui.Vec4{0.13 + 0.19, 0.14 + 0.19, 0.15 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorChildBg, imgui.Vec4{0.13 + 0.19, 0.14 + 0.19, 0.15 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorPopupBg, imgui.Vec4{0.13 + 0.19, 0.14 + 0.19, 0.15 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorBorder, imgui.Vec4{0.43 + 0.19, 0.43 + 0.19, 0.50 + 0.19, 0.50})
	imgui.PushStyleColor(imgui.StyleColorBorderShadow, imgui.Vec4{0.00 + 0.19, 0.00 + 0.19, 0.00 + 0.19, 0.00})
	imgui.PushStyleColor(imgui.StyleColorFrameBg, imgui.Vec4{0.25 + 0.19, 0.25 + 0.19, 0.25 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorFrameBgHovered, imgui.Vec4{0.38 + 0.19, 0.38 + 0.19, 0.38 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorFrameBgActive, imgui.Vec4{0.67 + 0.19, 0.67 + 0.19, 0.67 + 0.19, 0.39})
	imgui.PushStyleColor(imgui.StyleColorTitleBg, imgui.Vec4{0.08 + 0.19, 0.08 + 0.19, 0.09 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorTitleBgActive, imgui.Vec4{0.08 + 0.19, 0.08 + 0.19, 0.09 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorTitleBgCollapsed, imgui.Vec4{0.00 + 0.19, 0.00 + 0.19, 0.00 + 0.19, 0.51})
	imgui.PushStyleColor(imgui.StyleColorMenuBarBg, imgui.Vec4{0.14 + 0.19, 0.14 + 0.19, 0.14 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorScrollbarBg, imgui.Vec4{0.02 + 0.19, 0.02 + 0.19, 0.02 + 0.19, 0.53})
	imgui.PushStyleColor(imgui.StyleColorScrollbarGrab, imgui.Vec4{0.31 + 0.19, 0.31 + 0.19, 0.31 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorScrollbarGrabHovered, imgui.Vec4{0.41 + 0.19, 0.41 + 0.19, 0.41 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorScrollbarGrabActive, imgui.Vec4{0.51 + 0.19, 0.51 + 0.19, 0.51 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorCheckMark, imgui.Vec4{0.64, 0.44, 0.87, 0.95})
	imgui.PushStyleColor(imgui.StyleColorSliderGrab, imgui.Vec4{0.64, 0.44, 0.87, 0.95})
	imgui.PushStyleColor(imgui.StyleColorSliderGrabActive, imgui.Vec4{0.57, 0.44, 0.87, 0.95})
	imgui.PushStyleColor(imgui.StyleColorButton, imgui.Vec4{0.25 + 0.19, 0.25 + 0.19, 0.25 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorButtonHovered, imgui.Vec4{0.38 + 0.19, 0.38 + 0.19, 0.38 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorButtonActive, imgui.Vec4{0.67 + 0.19, 0.67 + 0.19, 0.67 + 0.19, 0.39})
	imgui.PushStyleColor(imgui.StyleColorHeader, imgui.Vec4{0.22 + 0.19, 0.22 + 0.19, 0.22 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorHeaderHovered, imgui.Vec4{0.25 + 0.19, 0.25 + 0.19, 0.25 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorHeaderActive, imgui.Vec4{0.67 + 0.19, 0.67 + 0.19, 0.67 + 0.19, 0.39})
	imgui.PushStyleColor(imgui.StyleColorSeparator, imgui.Vec4{0.43 + 0.19, 0.43 + 0.19, 0.50 + 0.19, 0.50})
	imgui.PushStyleColor(imgui.StyleColorSeparatorHovered, imgui.Vec4{0.41 + 0.19, 0.42 + 0.19, 0.44 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorSeparatorActive, imgui.Vec4{0.64, 0.43, 0.87, 0.95})
	imgui.PushStyleColor(imgui.StyleColorResizeGrip, imgui.Vec4{0.00 + 0.19, 0.00 + 0.19, 0.00 + 0.19, 0.00})
	imgui.PushStyleColor(imgui.StyleColorResizeGripHovered, imgui.Vec4{0.29 + 0.19, 0.30 + 0.19, 0.31 + 0.19, 0.67})
	imgui.PushStyleColor(imgui.StyleColorResizeGripActive, imgui.Vec4{0.64, 0.44, 0.87, 0.95})
	imgui.PushStyleColor(imgui.StyleColorTab, imgui.Vec4{0.08 + 0.19, 0.08 + 0.19, 0.09 + 0.19, 0.83})
	imgui.PushStyleColor(imgui.StyleColorTabHovered, imgui.Vec4{0.33 + 0.19, 0.34 + 0.19, 0.36 + 0.19, 0.83})
	imgui.PushStyleColor(imgui.StyleColorTabActive, imgui.Vec4{0.23 + 0.19, 0.23 + 0.19, 0.24 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorTabUnfocused, imgui.Vec4{0.08 + 0.19, 0.08 + 0.19, 0.09 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorTabUnfocusedActive, imgui.Vec4{0.13 + 0.19, 0.14 + 0.19, 0.15 + 0.19, 1.00})
	//imgui.PushStyleColor(imgui.StyleColorDockingPreview,imgui.Vec4{0.26+0.1,0.59+0.1,0.98+0.1,0.70})
	//imgui.PushStyleColor(imgui.StyleColorDockingEmptyBg,imgui.Vec4{0.20+0.1,0.20+0.1,0.20+0.1,1.00})
	imgui.PushStyleColor(imgui.StyleColorPlotLines, imgui.Vec4{0.61 + 0.19, 0.61 + 0.19, 0.61 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorPlotLinesHovered, imgui.Vec4{1.00 + 0.19, 0.43 + 0.19, 0.35 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorPlotHistogram, imgui.Vec4{0.90 + 0.19, 0.70 + 0.19, 0.00 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorPlotHistogramHovered, imgui.Vec4{1.00 + 0.19, 0.60 + 0.19, 0.00 + 0.19, 1.00})
	imgui.PushStyleColor(imgui.StyleColorTextSelectedBg, imgui.Vec4{0.64, 0.44, 0.87, 0.36})
	imgui.PushStyleColor(imgui.StyleColorDragDropTarget, imgui.Vec4{0.73, 0.46, 0.88, 0.95})
	imgui.PushStyleColor(imgui.StyleColorNavHighlight, imgui.Vec4{0.64, 0.44, 0.87, 0.95})
	imgui.PushStyleColor(imgui.StyleColorNavWindowingHighlight, imgui.Vec4{1.00 + 0.19, 1.00 + 0.19, 1.00 + 0.19, 0.70})
	//imgui.PushStyleColor(imgui.StyleColorNavWindowingDimBg,imgui.Vec4{0.80+0.1,0.80+0.1,0.80+0.1,0.20})
	//imgui.PushStyleColor(imgui.StyleColorModalWindowDimBg,imgui.Vec4{0.80+0.1,0.80+0.1,0.80+0.1,0.35})

	//style.GrabRounding                           = style.FrameRounding = 2.3f;
	imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 0)
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0)
	//imgui.PushStyleVarFloat(imgui.StyleVarFrameBorderSize, 1)
}
