package draft

type DefinitionUpdaterMode int
type DefinitionUpdaterStatus int

const (
  ALL_HEROES DefinitionUpdaterMode = iota
  SINGLE_HERO
  MULTI_HERO // this is for an array of heroes
)

const (
  STATUS_IDLE DefinitionUpdaterStatus = iota
  STATUS_UPDATING_IN_PROGRESS
  STATUS_COMPLETE
)

type DefinitionUpdater struct{
  mode DefinitionUpdaterMode
  status DefinitionUpdaterStatus // will default to idle thanks to iota
}

func (updater *DefinitionUpdater) getMode() DefinitionUpdaterMode{
  return updater.mode
}

func (updater *DefinitionUpdater) setMode(newMode DefinitionUpdaterMode) bool{
  updater.mode = newMode
  if updater.mode == newMode {
    return true
  }else{
    return false
  }
}

func (updater *DefinitionUpdater) getStatus() DefinitionUpdaterStatus{
  return updater.status
}


func (updater *DefinitionUpdater) Update() bool{
  return false // false = not succesful
}
