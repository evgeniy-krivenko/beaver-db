package compute

type Query struct {
	cmdID CommandID
	args  []string
}

func NewQuery(cmdID CommandID, args []string) Query {
	return Query{cmdID: cmdID, args: args}
}

func (q Query) CommandID() CommandID {
	return q.cmdID
}

func (q Query) Args() []string {
	return q.args
}
