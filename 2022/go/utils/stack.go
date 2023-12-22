package utils

type Stack struct {
	elements []string
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Push(str string) {
	if str != " " {
		s.elements = append(s.elements, str)
	}
}

func (s *Stack) PushMultiple(crates []string) {
	s.elements = append(s.elements, crates...)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	tmp := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return tmp
}

func (s *Stack) PopMultiple(n int) []string {
	if s.IsEmpty() {
		return []string{}
	}
	tmp := s.elements[len(s.elements)-n:]
	s.elements = s.elements[:len(s.elements)-n]
	return tmp
}
