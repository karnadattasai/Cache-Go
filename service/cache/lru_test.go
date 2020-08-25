package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lru_Read(t *testing.T) {

	type initCacheArg struct { // Holds the arguments to call Write() for initializing cache
		key   int
		value int
	}
	tests := []struct {
		name          string
		initCacheArgs []initCacheArg
		key           int
		returnValue   int
	}{
		{
			name:          "Empty cache",
			initCacheArgs: []initCacheArg{},
			key:           1,
			returnValue:   -1,
		},
		{
			name: "Cache Full and key not present",
			initCacheArgs: []initCacheArg{
				{
					key:   7,
					value: 7,
				},
				{
					key:   0,
					value: 0,
				},
				{
					key:   1,
					value: 1,
				},
			},
			key:         2,
			returnValue: -1,
		},
		{
			name: "Cache Full and key present",
			initCacheArgs: []initCacheArg{
				{
					key:   7,
					value: 7,
				},
				{
					key:   0,
					value: 0,
				},
				{
					key:   1,
					value: 1,
				},
			},
			key:         0,
			returnValue: 0,
		},
		{
			name: "Cache Full, removing least recently used",
			initCacheArgs: []initCacheArg{
				{
					key:   7,
					value: 7,
				},
				{
					key:   0,
					value: 0,
				},
				{
					key:   1,
					value: 1,
				},
				{
					key:   7,
					value: 7,
				},
				{
					key:   1,
					value: 1,
				},
				{
					key:   2,
					value: 2,
				},
			},
			key:         0,
			returnValue: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Cache
			c = NewLRUCache()
			// Initialize cache
			for _, arg := range tt.initCacheArgs {
				c.Write(arg.key, arg.value)
			}
			assert.Equal(t, tt.returnValue, c.Read(tt.key), "should be same")
		})
	}
}

func Test_lru_Write(t *testing.T) {

	type initCacheArg struct { // Holds the arguments to call Write() for initializing cache
		key   int
		value int
	}
	tests := []struct {
		name          string
		initCacheArgs []initCacheArg
		key           int
		value         int
	}{
		{
			name: "Write Success",
			initCacheArgs: []initCacheArg{
				{
					key:   7,
					value: 7,
				},
				{
					key:   0,
					value: 0,
				},
				{
					key:   1,
					value: 1,
				},
			},
			key: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Cache
			c = NewLRUCache()
			// Initialize cache
			for _, arg := range tt.initCacheArgs {
				c.Write(arg.key, arg.value)
			}
			c.Write(tt.key, tt.value)
		})
	}
}
