package type_conversion

import (
    "fmt"
    "strconv"
)

func Execute() {
    // Integer to float
    var intVal int = 42
    var floatVal float64 = float64(intVal)
    fmt.Println("Integer to float:", floatVal)

    // Float to integer
    var floatVal2 float64 = 42.58
    var intVal2 int = int(floatVal2)
    fmt.Println("Float to integer:", intVal2)

    // String to integer
    var strVal string = "42"
    var intVal3, err = strconv.Atoi(strVal)
    if err == nil {
        fmt.Println("String to integer:", intVal3)
    } else {
        fmt.Println("Error converting string to integer:", err)
    }
}
