package menu

import "fmt"

func buyStone(Money *int) {
    fmt.Println("1. stone helmet for 10 roupies")
    fmt.Println("2. stone armor for 20 roupies")
    fmt.Println("3. stone boots for 20 roupies")
    
    var choice int
    fmt.Print("Choose an item: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        if *Money >= 10 {
            *Money -= 10
            fmt.Println("You bought the  stone helmet")
        } else {
            fmt.Println("Not enough money.")
        }
    case 2:
        if *Money >= 20 {
            *Money -= 20
            fmt.Println("You bought the stone armor")
        } else {
            fmt.Println("Not enough money.")
        }
    case 3:
        if *Money >= 20 {
            *Money -= 20
            fmt.Println("You bought the stone boots")
        } else {
            fmt.Println("Not enough money.")
        }
    default:
        fmt.Println("Invalid choice.")
    }

    fmt.Println("Remaining money:", *Money)
}
