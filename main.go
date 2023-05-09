package main

import (
	"fmt"
)

func pow(n, m int) int {
	if m == 0 {
		return 1
	}

	res := n
	for i := 2; i <= m; i++ {
		res *= n
	}

	return res
}

const (
	lines = 8
	rows = 4
	size = 34
)

type LifeGenerator struct {
	config [lines][rows]int
}

func (l *LifeGenerator) Step() {
	prevConf := l.config
	l.config = [lines][rows]int{}
	for i := 0; i < lines; i++ {
		for j := 0; j < rows; j++ {
			lifeCount := 0
			for k1 := i-1; k1 < i+2; k1++ {
				for k2 := j-1; k2 < j+2; k2++ {
					if k1 == i && k2 == j {
						continue
					}
					i1 := (k1 + lines) % lines
					j1 := (k2 + rows) % rows
					lifeCount += prevConf[i1][j1]
				}
			}
			if prevConf[i][j] == 0 && lifeCount == 3 {
				l.config[i][j] = 1
			} else if prevConf[i][j] == 1 && (lifeCount == 3 || lifeCount == 2) {
				l.config[i][j] = 1
			}
		}
 	}
}

func (l *LifeGenerator) ToNum() int {
	binNum := [lines*rows]int{}
	k := 0
	for i := 0; i < lines; i++ {
		for j := 0; j < rows; j++ {
			binNum[k] = l.config[i][j]
			k++
		}
	}

	num := 0

	for i := 0; i < lines*rows; i++ {
		if binNum[i] == 1 {
			num += pow(2*binNum[i], lines*rows-1 - i)
		}
	}
	return num;
}

func (l *LifeGenerator) String() string {
	res := ""
	for i := 0; i < lines; i++ {
		res += fmt.Sprint(l.config[i]) + "\n"
	}
	return res
}

func (l *LifeGenerator) Generate(count int) []int {
	res := make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = l.ToNum()
		l.Step()
	}
	return res
}

func NewLife() LifeGenerator {
	conf := [lines][rows]int{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 1},
	}

	return LifeGenerator{config: conf}
}

type LinearLifeGen struct {
	config [size]int
}

func (l *LinearLifeGen) Step() {
	prevConf := l.config
	l.config = [size]int{}
	for i := 0; i < size; i++ {
		i1 := (i - 1 + size) % size
		i2 := (i + 1 + size) % size
		switch {
		case prevConf[i1] == 1 && prevConf[i] == 1 && prevConf[i2] == 1:
			l.config[i] = 0
		case prevConf[i1] == 1 && prevConf[i] == 1 && prevConf[i2] == 0:
			l.config[i] = 1
		case prevConf[i1] == 1 && prevConf[i] == 0 && prevConf[i2] == 1:
			l.config[i] = 1
		case prevConf[i1] == 1 && prevConf[i] == 0 && prevConf[i2] == 0:
			l.config[i] = 0
		case prevConf[i1] == 0 && prevConf[i] == 1 && prevConf[i2] == 1:
			l.config[i] = 1
		case prevConf[i1] == 0 && prevConf[i] == 1 && prevConf[i2] == 0:
			l.config[i] = 1
		case prevConf[i1] == 0 && prevConf[i] == 0 && prevConf[i2] == 1:
			l.config[i] = 1
		case prevConf[i1] == 0 && prevConf[i] == 0 && prevConf[i2] == 0:
			l.config[i] = 0
		}
	}
}

func (l *LinearLifeGen) ToNum() int {
	num := 0

	for i := 0; i < size; i++ {
		if l.config[i] == 1 {
			num += pow(2, size - 1 - i)
		}
	}
	return num
}

func (l *LinearLifeGen) Generate(count int) []int {
	res := make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = l.ToNum()
		l.Step()
	}
	return res
}

func NewLinearGen(conf [size]int) LinearLifeGen {
	return LinearLifeGen{config: conf}
}

// 08.01.2001
// 0 - 0000
// 8 - 1000
// 0 - 0000
// 1 - 0001
// 2 - 0010
// 0 - 0000
// 0 - 0000
// 1 - 0000
func main() {
	// life := NewLife()
	// fmt.Println(life.String())
	// life.Step()
	// fmt.Println(life.String())
	conf := [size]int{0, 0,0,0,0, 1,0,0,0, 0,0,0,0, 0,0,0,1, 0,0,1,0, 0,0,0,0, 0,0,0,0, 0,0,0,1, 0}
	life := NewLinearGen(conf)
	for _, num := range life.Generate(100) {
		fmt.Println(num)
	}
	// for _, num := range life.Generate(100) {
	// 	fmt.Println(num)
	// }
}