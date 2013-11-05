package Module

type Moduler interface {
    Init()
    Breath()
    Run()
    Stop()

    IsSelfRun() bool

    Load()
    Unload()
}


