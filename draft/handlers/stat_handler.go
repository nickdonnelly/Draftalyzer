package handlers

import (
  "draftalyzer/draft"
)

type StatHandler struct{
  Handler
  status StatHandlerStatus
}

func NewStatHandler() *StatHandler {
  var handler = StatHandler{}
  // handler.AddField("status", STATUS_IDLE)
  return &handler
}

type StatHandlerStatus uint

const (
  STATUS_IDLE = iota
  STATUS_CALCULATING
  STATUS_COMPLETE
  STATUS_UNKNOWN = 999
)

func (handler *StatHandler) SetHeroById(hero_id uint) bool{
  if handler.statusFree(){
    handler.hero_id = hero_id
    return true
  }else{
    return false
  }
}

func (handler *StatHandler) GetHeroId() uint{
  return handler.hero_id
}

func (handler *StatHandler) GetHeroName() string{
  return draft.GetHeroNameFromId(handler.hero_id)
}

func (handler *StatHandler) GetStatus() StatHandlerStatus{
  return handler.status
}

func (handler *StatHandler) Reset() bool{
  if handler.status == STATUS_CALCULATING || handler.status == STATUS_UNKNOWN{
    return false
  }else{
    // TODO: Make sure this stays up to date with any additional fields, could cause silent bugs
    handler.hero_id = 0
    handler.status = STATUS_IDLE
    return true
  }
}


func (handler *StatHandler) statusFree() bool{
  switch(handler.status){
  case STATUS_CALCULATING:
    return false
  case STATUS_UNKNOWN:
    return false
  case STATUS_IDLE:
    fallthrough
  case STATUS_COMPLETE:
    return true
  default:
    return false
  }
}
