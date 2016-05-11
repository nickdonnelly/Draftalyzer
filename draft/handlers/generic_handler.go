package handlers

import (
  "log"
)

type Handler struct {
  hero_id uint
  fields map[string]interface{}
}

func (handler *Handler) fieldExists(key string) bool{
  if _, ok := handler.fields[key]; ok {
    return true
  }else{
    return false
  }
}

func (handler *Handler) removeField(key string) bool{
  if handler.fieldExists(key){
    delete(handler.fields, key)
    return true
  }else{
    return false
  }
}

func (handler *Handler) AddField(key string, value interface{}) bool{
  if val, ok := handler.fields[key]; ok{
    log.Printf("[Error] Attempt to add key %v to fields map with value %v when %k already exists with value %val", key, value, key, val)
    return false
  }else{
    handler.fields[key] = value
    return true
  }
}
