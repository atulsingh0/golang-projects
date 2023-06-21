## Generating UUID

```
package main

import (
 "fmt"

 "github.com/google/uuid"
)

func main() {

 // generating uuid
 fmt.Println("UUID:", uuid.NewString())
}
```
