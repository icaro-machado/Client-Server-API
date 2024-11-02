package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Price struct {
	Bid string `json:"bid"`
}

type PriceDB struct {
	ID   uint `gorm:"primaryKey"`
	Bid  string
	Data time.Time
}

func main() {
	http.HandleFunc("/cotacao", handlePrice)
	fmt.Println("Servidor executando na porta 8080!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePrice(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	price, err := findPrice(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erro ao buscar cotação:", err)
		return
	}

	dbCtx, dbCancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer dbCancel()

	if err := savePriceDb(dbCtx, price.Bid); err != nil {
		log.Println("Erro ao salvar cotação no banco:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": price.Bid})
}

func findPrice(ctx context.Context) (*Price, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resultado map[string]Price
	if err := json.NewDecoder(resp.Body).Decode(&resultado); err != nil {
		return nil, err
	}

	price, ok := resultado["USDBRL"]
	if !ok {
		return nil, fmt.Errorf("Cambio não encontrado")
	}
	return &price, nil
}

func savePriceDb(ctx context.Context, bid string) error {
	dsn := "root:root@tcp(localhost:3306)/pricedb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.WithContext(ctx).AutoMigrate(&PriceDB{}); err != nil {
		return err
	}

	price := PriceDB{Bid: bid, Data: time.Now()}
	return db.WithContext(ctx).Create(&price).Error
}
