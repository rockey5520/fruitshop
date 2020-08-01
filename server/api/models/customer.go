package models

import "github.com/jinzhu/gorm"

/*Customer Customer table represents the customer of the fruit store
swagger:model Customer
*/
type Customer struct {
	// Login ID of the customer
	LoginId string `json:"loginid" gorm:"unique_index"`
	// First name of the customer
	FirstName string `json:"firstname"`
	// Last name of the customer
	LastName string `json:"lastname"`
	Cart     Cart   `gorm:"foreignkey:CustomerId;association_foreignkey:ID"`

	gorm.Model
}

// Error Bad Request
// swagger:response badReq
type H map[string]interface{}

// HTTP status code 200 and Customer model in data
// swagger:response userResp
type swaggCustomerResp struct {
	// in:body
	Body struct {
		// HTTP status code 200
		Code int `json:"code"`
		// User model
		Data Customer `json:"data"`
	}
}

func (c *Customer) Validate(action string) error {
	switch strings.ToLower(action) {
	
		if c.FirstName == "" {
			return errors.New("Required FirstName")
		}
		if u.LastName == "" {
			return errors.New("Required LastName")
		}
		return nil
	}
}

func (c *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {

	var err error
	// update cart to cart array in the customer table
	newcart := models.Cart{
		Total:  0.0,
		Status: "OPEN",
	}

	customer := models.Customer{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		LoginId:   input.LoginId,
		Cart:      newcart,
	}
	
	if err := db.Create(&customer).Error; err != nil {
		return &Customer{}, err
		
	}
	return c, nil
}


func (c *Customer) FindCustomerByID(db *gorm.DB, uid uint32) (*Customer, error) {
	var err error
	var customer models.Customer
	db := c.MustGet("db").(*gorm.DB)
	err := db.Where("login_id = ?", c.Param("login_id")).First(&customer).Error

	if gorm.IsRecordNotFoundError(err) {
		return &Customer{}, errors.New("Customer record Not Found")
	}

	var cart models.Cart
	db.Where("customer_id = ? AND status = ?", customer.ID, "OPEN").Find(&cart)
	var cartItem []models.CartItem
	db.Where("cart_id = ?", cart.ID).Find(&cartItem)
	var payment models.Payment
	db.Where("cart_id = ?", cart.ID).Find(&payment)
	var appliedDualItemDiscount []models.AppliedDualItemDiscount
	db.Where("cart_id = ?", cart.ID).Find(&appliedDualItemDiscount)
	var appliedSingleItemDiscount []models.AppliedSingleItemDiscount
	db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemDiscount)
	var appliedSingleItemCoupon []models.AppliedSingleItemCoupon
	db.Where("cart_id = ?", cart.ID).Find(&appliedSingleItemCoupon)
	customer.Cart = cart
	customer.Cart.CartItem = cartItem
	customer.Cart.Payment = payment
	customer.Cart.AppliedDualItemDiscount = appliedDualItemDiscount
	customer.Cart.AppliedSingleItemCoupon = appliedSingleItemCoupon
	customer.Cart.AppliedSingleItemDiscount = appliedSingleItemDiscount
	
	err = db.Debug().Model(Customer{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Customer{}, err
	}
	
	return c, err
}

