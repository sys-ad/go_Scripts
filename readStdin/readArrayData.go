package readStdin

import "fmt"

func readData() ([]int, error) {                               // reads data (array) from stdin
        var length int
        
        _, err := fmt.Scanf("%d", &length)
        if err != nil {
                log.Fatal(err)
        }
        
        data := make([]int, length)
        
        for i := range data {
                _, err := fmt.Scanf("%d", &data[i])
                if err != {
                        return nil, err
                }
        }
        return data, nil
}
