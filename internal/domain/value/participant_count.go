package value

type ParticipantCount struct {
	ValueObject[uint32]
}

func NewParticipantCount(v uint32) *ParticipantCount {
	return &ParticipantCount{NewValueObject(v)}
}
