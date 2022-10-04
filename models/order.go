package models

import "gorm.io/gorm"

type Order struct {
	Id         uint        `json:"id"`
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`
	Total      float32     `json:"total" gorm:"-"`
	Email      string      `json:"email"`
	CreatedAt  string      `json:"created_at"`
	UpdatedAt  string      `json:"updated_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(Order{}).Count(&total)
	return total
}

func (order *Order) Take(db *gorm.DB, offset int, limit int) interface{} {
	var orders []Order
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	for i, _ := range orders {
		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
		var total float32 = 0

		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float32(orderItem.Quantity)
		}
		// frontend will require this later
		orders[i].Total = total
	}
	return orders
}
