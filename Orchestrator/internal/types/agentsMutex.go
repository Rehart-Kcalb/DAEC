package types

import (
	"sync"
	"time"
)

type AgentsMutex struct {
	Agents map[string]time.Time
	*sync.Mutex
}
