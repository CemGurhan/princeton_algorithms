package quickfind

import "errors"

type QuickFind []int

func (qf *QuickFind) Find(componentOne int, componentTwo int) bool {
	if len(*qf) == 0 {
		return false
	}

	if (*qf)[componentOne] == (*qf)[componentTwo] {
		return true
	}

	return false
}

func (qf *QuickFind) Union(componentOne int, componentTwo int) error {
	if len(*qf) == 0 {
		return errors.New("QuickFind array empty")
	}

	componentOneEntry := (*qf)[componentOne]
	for i := 0; i < len(*qf); i++ {
		if (*qf)[i] == componentOneEntry {
			(*qf)[i] = (*qf)[componentTwo]
		}
	}

	return nil
}
