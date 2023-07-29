package cstack

type FrameInfo struct {
	Func string `json:"func"`
	File string `json:"file"`
	Line int    `json:"line"`
}
