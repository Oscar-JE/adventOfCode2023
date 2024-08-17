package lens

type Lens struct {
	focalLength int
	label       string
}

func Init(label string, focalLength int) Lens {
	return Lens{focalLength: focalLength, label: label}
}

func (l Lens) Label() string {
	return l.label
}

func (l Lens) FocalLength() int {
	return l.focalLength
}
