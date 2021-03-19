package service

import (
	"github.com/magiconair/properties/assert"
	"self_initializing_fake/internal/model"
	"testing"
)

func TestIsHeaderValid(t *testing.T) {
	t.Run("should return true when all headers exist", func(t *testing.T) {
		hey := model.Header{
			"a": []string{"aaa"},
			"b": []string{"bbb"},
			"c": []string{"ccc"},
		}
		needle := model.Header{
			"a": []string{"aaa"},
			"b": []string{"bbb"},

		}
		assert.Equal(t, isHeaderValid(hey, needle), true)
	})
}
