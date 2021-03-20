package mock

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
		}
		needle := model.Header{
			"a": []string{"aaa"},
			"b": []string{"bbb"},

		}
		assert.Equal(t, isHeaderValid(hey, needle), true)
	})
	t.Run("should return true when all headers of fake request are present in request", func(t *testing.T) {
		request := model.Header{
			"a": []string{"aaa"},
			"b": []string{"bbb"},
			"c": []string{"ccc"},
		}
		fakeRequest := model.Header{
			"a": []string{"aaa"},
			"b": []string{"bbb"},
		}
		assert.Equal(t, isHeaderValid(request, fakeRequest), true)
	})
	t.Run("should return false when all headers of fake request are not present in request", func(t *testing.T) {
		request := model.Header{
			"a": []string{"aaa"},
			"b": []string{"bbb"},
		}
		fakeRequest := model.Header{
			"a": []string{"aaa"},
			"b": []string{"bbb"},
			"c": []string{"ccc"},
		}
		assert.Equal(t, isHeaderValid(request, fakeRequest), false)
	})
	t.Run("should return false when all headers are present but value is different", func(t *testing.T) {
		request := model.Header{
			"a": []string{"aaa"},
			"b": []string{"bbb", "abc"},
		}
		fakeRequest := model.Header{
			"a": []string{"aaa"},
			"b": []string{"bbb"},
		}
		assert.Equal(t, isHeaderValid(request, fakeRequest), false)
	})

}

func TestIsRequestValid(t *testing.T) {
	t.Run("should return true when both param are same", func(t *testing.T) {
		var param1, param2 interface{}
		param1 = map[string]string{
			"a": "a",
			"b": "b",
		}
		param2 = map[string]string{
			"a": "a",
			"b": "b",
		}
		assert.Equal(t, isRequestValid(param1, param2), true)
	})

	t.Run("should return false when both param are not same", func(t *testing.T) {
		var param1, param2 interface{}
		param1 = map[string]string{
			"a": "a",
			"b": "b",
		}
		param2 = map[string]string{
			"a": "a",
			"b": "c",
		}
		assert.Equal(t, isRequestValid(param1, param2), false)
	})
}