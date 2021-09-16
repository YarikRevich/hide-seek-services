package tools

import (
	"strings"

	"github.com/YarikRevich/lru/pkg/interfaces"
	"github.com/YarikRevich/lru/pkg/lru"
)

var (
	LRU = lru.New(20)
)

func IsService(path, serviceName string) bool {
	if cache := LRU.Get(path); cache != nil {
		return cache.(bool)
	} 

	split := strings.Split(strings.TrimSpace(path), "/")
	if len(split) != 2 {
		return false
	}
	
	r := split[len(split)-1] == serviceName
	LRU.Set(interfaces.Cell{path, r})
	return r
}
