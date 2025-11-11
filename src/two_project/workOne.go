package main

import (
    "fmt"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    // 加载.env文件中的环境变量
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
    // 使用环境变量
    myVar := os.Getenv("MY_ENV_VAR")
    fmt.Println("My Environment Variable:", myVar)
}
