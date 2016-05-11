package handlers


type StatHandler struct {
  status StatHandlerStatus
}

type StatHandlerStatus uint

const (
  STATUS_IDLE = iota
  STATUS_CALCULATING
  STATUS_COMPLETE
  STATUS_UNKNOWN = 999
)

func (handler *StatHandler) GetStatus() StatHandlerStatus{
  // TODO

  return STATUS_UNKNOWN
}

func (handler *StatHandler) Reset() bool{
  if handler.status == STATUS_CALCULATING || handler.status == STATUS_UNKNOWN{
    return false
  }else{
    // TODO 
    handler.status = STATUS_IDLE
    return true
  }
}
