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
  heroes []string
}

func (updater *DefinitionUpdater) getMode() DefinitionUpdaterMode{
  return updater.mode
}

func (updater *DefinitionUpdater) setMode(newMode DefinitionUpdaterMode) bool{
  updater.Reset()
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

func (updater *DefinitionUpdater) Reset() bool{
  if(updater.status == STATUS_IDLE || updater.status == STATUS_COMPLETE){
    updater.mode = ALL_HEROES
    updater.heroes = nil
    updater.status = STATUS_IDLE
    return true
  }else{
    return false
  }
}

func (updater *DefinitionUpdater) Update() bool{
  return false // false = not succesful
}

func (updater *DefinitionUpdater) AddHero(heroName string) bool {
  if updater.mode == ALL_HEROES{
    return false
  }else if updater.mode == SINGLE_HERO && len(updater.heroes) >= 1{
    // error out if they are trying to append to the list and
    // they selected single hero mode
    return false
  }else{
    updater.heroes = append(updater.heroes, heroName)
    return true
  }
}
