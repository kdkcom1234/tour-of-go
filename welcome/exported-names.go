package welcome

import (
	"fmt"
	"math"
)

func ExportedNamesGo() {
	// 대문자로 시작하는 이름만 export됨
	fmt.Println(math.Pi)
}
