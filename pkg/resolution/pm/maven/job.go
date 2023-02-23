package maven

type Job struct {
	file       string
	cmdFactory ICmdFactory
	err        error
	status     chan string
}

func NewJob(file string, cmdFactory ICmdFactory) *Job {
	return &Job{file: file, cmdFactory: cmdFactory, status: make(chan string)}
}

func (j *Job) File() string {
	return j.file
}

func (j *Job) Error() error {
	return j.err
}

func (j *Job) Status() chan string {
	return j.status
}

func (j *Job) Run() {
	cmd, err := j.cmdFactory.MakeDependencyTreeCmd()
	if err != nil {
		j.err = err

		return
	}
	j.status <- "creating dependency graph"
	_, err = cmd.Output()
	if err != nil {
		j.err = err

		return
	}
}
