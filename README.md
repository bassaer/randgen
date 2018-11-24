# randgen

```
import (
	"fmt"
	"github.com/bassaer/randgen"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(randgen.RandNum(1000000000, 9999999999))
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Println(randgen.RandStr(10))
	}
}
```

```
❯ go run cmd/randgen/main.go
6820525524
2909274254
7632124103
7653645410
4192549892
9015342648
9694299177
2491489227
2332337826
2017478884

EINCgwWBPa
EELtPPWZWS
CxcqzZHBcy
lJytFMRuhn
XclnzLMWcw
MCrsvWwdHH
okZkEJqEBA
CZMdXsbWtr
YAgOhmJIWg
sdgyaJvHDW
```
