package Module

type Moduler interface {
    Init()
    Breath()
    Run()
    Stop()

    IsSelfRun() bool

    Load() error
    Unload() error
}


