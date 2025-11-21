package main

type line struct {
    num int
    content string
}

type VisibleLines struct {
    firstLine int
    total int
    content []line
    linesOnScreen []line
}

func NewVisibleLines(firstLine int, total int, content []line) *VisibleLines {
    vl :=  &VisibleLines{
		content: content,
	}
    vl.UpdateVisibleLines(firstLine, total)
    return vl
}

func (vl *VisibleLines) UpdateContent(content []line) {
    vl.content = content
}

func (vl *VisibleLines) UpdateVisibleLines(firstLine int, total int) {
    vl.firstLine = firstLine
    vl.total = total 

    // clear slice
    vl.linesOnScreen = vl.linesOnScreen[:0]

    // copy the lines that should be on the screen
    for _, line := range(vl.content) {
        if line.num >= firstLine && len(vl.linesOnScreen) < total {
            vl.linesOnScreen = append(vl.linesOnScreen, line)
        }
    }
}
