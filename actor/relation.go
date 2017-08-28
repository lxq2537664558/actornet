package actor

type Relation interface {
	ParentPID() *PID

	addChild(*PID)

	setParentPID(ppid *PID)

	broadcast(interface{})
}

type RelationImplement struct {
	ppid *PID // 父级

	childs []*PID // 子级

	proc Process

	dm *Domain
}

func (self *RelationImplement) Domain() *Domain {
	return self.dm
}

func (self *RelationImplement) ParentPID() *PID {
	return self.ppid
}

func (self *RelationImplement) setParentPID(ppid *PID) {
	self.ppid = ppid
}

func (self *RelationImplement) addChild(pid *PID) {

	childProc := pid.ref()

	if childProc == nil {
		panic("child can not be nil when add child")
	}

	childProc.setParentPID(self.proc.PID())

	self.childs = append(self.childs, pid)
}

func (self *RelationImplement) broadcast(data interface{}) {

	for _, c := range self.childs {
		c.Tell(data)
	}
}

func NewRelation(proc Process, dm *Domain) *RelationImplement {
	return &RelationImplement{
		proc: proc,
		dm:   dm,
	}
}
