/* https://coderun.yandex.ru/selections/backend/problems/merge-jsons-2/description */

package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Offer struct {
    OfferID   string `json:"offer_id"`
    MarketSKU int `json:"market_sku"`
    Price     int `json:"price"`
}

type Feed struct {
    Offers []Offer `json:"offers"`
}

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Ошибка при открытии файла:", err)
        return
    }
    defer inputFile.Close()
/*     outputFile, err := os.Create("output.json")
    if err != nil {
        fmt.Println("Ошибка при создании выходного файла:", err)
        return
    }
    defer outputFile.Close() */
    writer := bufio.NewWriter(os.Stdout)
    scanner := bufio.NewScanner(inputFile)
    scanner.Scan()
    firstLine := scanner.Text()
    parts := strings.Fields(firstLine)
    numFeeds, _ := strconv.Atoi(parts[0])
    maxOffers, _ := strconv.Atoi(parts[1])
    var allOffers []Offer
    for i := 0; i < numFeeds && scanner.Scan(); i++ {
        line := scanner.Text()
        var feed Feed
        err := json.Unmarshal([]byte(line), &feed)
        if err != nil {
            fmt.Println("Ошибка при десериализации JSON:", err)
            return
        }
        allOffers = append(allOffers, feed.Offers...)
    }
    if len(allOffers) > maxOffers {
        allOffers = allOffers[:maxOffers]
    }
    finalFeed := Feed{Offers: allOffers}
    jsonData, err := json.Marshal(finalFeed)
    if err != nil {
        fmt.Println("Ошибка при преобразовании в JSON:", err)
        return
    }

    _, err = writer.WriteString(string(jsonData) + "\n")
    if err != nil {
        fmt.Println("Ошибка при записи в файл:", err)
        return
    }
    writer.Flush()
}
