package app

type Invoice struct {
	DocNo               string        `json:"doc_no"`
	DocDate             string        `json:"doc_date"`
	ArCode              string        `json:"ar_code"`
	BillType            int           `json:"bill_type"`
	ArName              string        `json:"ar_name"`
	SumOfItemAmount     float64       `json:"sum_of_item_amount"`
	SumOfDiscountAmount float64       `json:"sum_of_discount_amount"`
	BeForeTaxAmount     float64       `json:"be_fore_tax_amount"`
	TaxAmount           float64       `json:"tax_amount"`
	TotalAmount         float64       `json:"total_amount"`
	NetDebtAmount       float64       `json:"net_debt_amount"`
	BillBalance         float64       `json:"bill_balance"`
	InvoiceSub          Invoicesub    `json:"invoice_sub""`
	Payment             PaymentDetail `json:"payment"`
	TaxDetail           TaxDetail     `json:"tax_detail"`
}

type Invoicesub struct {
	ItemCode       string  `json:"item_code"`
	Qty            float32 `json:"qty"`
	UnitCode       string  `json:"unit_code"`
	WhCode         string  `json:"wh_code"`
	ShelfCode      string  `json:"shelf_code"`
	Price          float64 `json:"price"`
	DisCountAmount float64 `json:"dis_count_amount"`
	Amount         float64 `json:"amount"`
	NetAmount      float64 `json:"net_amount"`
	ItemName       string  `json:"item_name"`
}

type PaymentDetail struct {
	Seq         int     `json:"seq"`
	PaymentType int     `json:"payment_type"`
	PayAmount   float64 `json:"pay_amount"`
	Refno       string  `json:"refno"`
}

type TaxDetail struct {
	TaxNo   string  `json:"tax_no"`
	TaxDate string  `json:"tax_date"`
	TaxType int     `json:"tax_type"`
	TaxRate float32 `json:"tax_rate"`
}
