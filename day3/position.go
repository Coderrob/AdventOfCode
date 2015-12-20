package main

import "fmt"

type Position struct {
    xPosition int
    yPosition int   
}

func (position *Position) move(direction byte) string {
    switch(direction){
	case '<':
        position.xPosition--
    case '>':
		position.xPosition++
	case '^':
		position.yPosition++
	case 'v':
		position.yPosition--
	}
    return fmt.Sprintf("%d,%d", position.xPosition, position.yPosition)
}
