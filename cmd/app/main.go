package main

import (
    "fmt"
    "os"
    "GoFundamentals/src/core_concepts/constants"
    "GoFundamentals/src/core_concepts/packages_and_modules"
    "GoFundamentals/src/core_concepts/structs"
    "GoFundamentals/src/core_concepts/envs"
)

func main() {
    // Access environment variables
    envVar1 := os.Getenv("ENV_VAR1")
    envVar2 := os.Getenv("ENV_VAR2")

    fmt.Println("ENV_VAR1:", envVar1)
    fmt.Println("ENV_VAR2:", envVar2)

    // Print all environment variables for debugging
    envs.PrintAllEnv()

    constants.Execute()
    packages_and_modules.Execute()
    structs.Execute()
    envs.Execute()

}
