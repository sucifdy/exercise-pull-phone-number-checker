package main

import (
    "fmt"
    "strings"
)

func PhoneNumberChecker(number string, result *string) {
    // Cek apakah nomor diawali dengan "62" atau "08"
    if !(strings.HasPrefix(number, "62") || strings.HasPrefix(number, "08")) {
        fmt.Println("Bukan 62 atau 08")
        *result = "invalid"
        return
    }

    // Cek panjang nomor, minimal 10 angka jika diawali dengan "08"
    // dan minimal 11 angka jika diawali dengan "62"
    if strings.HasPrefix(number, "62") && len(number) < 11 {
        fmt.Println("Panjangnya kurang dari 11")
        *result = "invalid"
        return
    } else if strings.HasPrefix(number, "08") && len(number) < 10 {
        fmt.Println("Panjangnya kurang dari 10")
        *result = "invalid"
        return
    }

    // Cek prefix nomor untuk menentukan provider
    if strings.HasPrefix(number, "62811") || strings.HasPrefix(number, "0811") ||
        strings.HasPrefix(number, "62812") || strings.HasPrefix(number, "0812") ||
        strings.HasPrefix(number, "62813") || strings.HasPrefix(number, "0813") ||
        strings.HasPrefix(number, "62814") || strings.HasPrefix(number, "0814") ||
        strings.HasPrefix(number, "62815") || strings.HasPrefix(number, "0815") {
        *result = "Telkomsel"
    } else if strings.HasPrefix(number, "62816") || strings.HasPrefix(number, "0816") ||
        strings.HasPrefix(number, "62817") || strings.HasPrefix(number, "0817") ||
        strings.HasPrefix(number, "62818") || strings.HasPrefix(number, "0818") ||
        strings.HasPrefix(number, "62819") || strings.HasPrefix(number, "0819") {
        *result = "Indosat"
    } else if strings.HasPrefix(number, "62821") || strings.HasPrefix(number, "0821") ||
        strings.HasPrefix(number, "62822") || strings.HasPrefix(number, "0822") ||
        strings.HasPrefix(number, "62823") || strings.HasPrefix(number, "0823") {
        *result = "XL"
    } else if strings.HasPrefix(number, "62827") || strings.HasPrefix(number, "0827") ||
        strings.HasPrefix(number, "62828") || strings.HasPrefix(number, "0828") ||
        strings.HasPrefix(number, "62829") || strings.HasPrefix(number, "0829") {
        *result = "Tri"
    } else if strings.HasPrefix(number, "62852") || strings.HasPrefix(number, "0852") ||
        strings.HasPrefix(number, "62853") || strings.HasPrefix(number, "0853") {
        *result = "AS"
    } else if strings.HasPrefix(number, "62881") || strings.HasPrefix(number, "0881") ||
        strings.HasPrefix(number, "62882") || strings.HasPrefix(number, "0882") ||
        strings.HasPrefix(number, "62883") || strings.HasPrefix(number, "0883") ||
        strings.HasPrefix(number, "62884") || strings.HasPrefix(number, "0884") ||
        strings.HasPrefix(number, "62885") || strings.HasPrefix(number, "0885") ||
        strings.HasPrefix(number, "62886") || strings.HasPrefix(number, "0886") ||
        strings.HasPrefix(number, "62887") || strings.HasPrefix(number, "0887") ||
        strings.HasPrefix(number, "62888") || strings.HasPrefix(number, "0888") {
        *result = "Smartfren"
    } else {
        *result = "invalid"
    }
}

func main() {
    var number = "081211111111" // Ganti dengan nomor yang ingin diuji
    var result string

    PhoneNumberChecker(number, &result)
    fmt.Println(result)
}
