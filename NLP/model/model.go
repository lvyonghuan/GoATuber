package model

import "sync"

type Msg struct {
	Msg   string
	Name  string
	IsUse bool
	Mu    sync.Mutex
}
