package cmd

import (
	"github.com/limingxinleo/logtail/cmd/handler"
	"github.com/spf13/viper"
	"log"
	"reflect"
)

var handlerMap map[string]interface{}

func init() {
	handlerMap = make(map[string]interface{})
	handlerMap["stdout"] = handler.Stdout{}
}

func push(text string) {
	handlers := viper.GetStringMap("push")

	for name := range handlers {
		if handlerMap[name] == nil {
			log.Println("Handler" + name + "is not found.")
			return
		}

		t := reflect.ValueOf(handlers[name]).Type()
		h := reflect.New(t).Elem().Interface().(handler.HandlerInterface)

		h.Handle(text)
	}
}
