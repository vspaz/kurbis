package task

type Runtime struct {
	ContainerId string
}

type Config struct {
	Name          string
	AttachStdin   bool
	AttachStdout  bool
	AttachStderr  bool
	Cmd           []string
	Image         string
	Memory        int64
	Disk          []string
	Env           []string
	RestartPolicy string
	Runtime       Runtime
}
