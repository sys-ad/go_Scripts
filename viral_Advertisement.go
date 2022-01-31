package main

import (
    "bufio"
    "fmt"n
    "io"
    "os"
    "strconv"s
    "strings"
)

/*
 * Complete the 'viralAdvertising' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func viralAdvertising(n int32) int32 {
    var cumulative, last_count, new_count, days, i int32
    days := 1
    if days == n {
        return 2 
    }
    for i := 0; i <= days; i++ {
        last_count := (n % 2) * 3
        
        new_count := last_count * 3
        cumulative := last_count + new_count       
    i++
    }
    return cumulative
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    result := viralAdvertising(n)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
