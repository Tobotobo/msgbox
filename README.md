# msgbox
golang msgbox

一番簡単な使い方
``` go
package main

import (
	"github.com/Tobotobo/msgbox"
)

func main() {
	msgbox.Show("テキスト")
}
```
情報アイコン、OKCancel、タイトルあり、判定あり
``` go
if msgbox.Info().OKCancel().Show("テキスト", "タイトル").IsOK() {
	fmt.Println("OK")
} else {
	fmt.Println("Cancel")
}
```
![image](https://user-images.githubusercontent.com/46508541/127505449-e0c28877-087f-4df7-b8ce-ffe975798ad4.png)
