package value

const No = 0

const (
	Sent       = 1 << iota
	Approved   = 1 << iota
	Decided    = 1 << iota
	Controlled = 1 << iota
)

const (
	Nsfw = 1 << iota
)
