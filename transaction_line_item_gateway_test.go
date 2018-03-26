// +build integration

package braintree

import (
	"context"
	"testing"
)

func TestTransactionWithLineItemsSingleOnlyRequiredFields(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:        "Name #1",
				Kind:        TransactionLineItemKindDebit,
				Quantity:    NewDecimal(10232, 4),
				UnitAmount:  NewDecimal(451232, 4),
				TotalAmount: NewDecimal(4515, 2),
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	lineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(lineItems), 1; g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}

	l := lineItems[0]
	if g, w := l.Name, "Name #1"; g != w {
		t.Errorf("got name %q, want %q", g, w)
	}
	if g, w := l.Kind, TransactionLineItemKindDebit; g != w {
		t.Errorf("got kind %q, want %q", g, w)
	}
	if g, w := l.Quantity, NewDecimal(10232, 4); g.Cmp(w) != 0 {
		t.Errorf("got quantity %q, want %q", g, w)
	}
	if g, w := l.UnitAmount, NewDecimal(451232, 4); g.Cmp(w) != 0 {
		t.Errorf("got unit amount %q, want %q", g, w)
	}
	if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
		t.Errorf("got total amount %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsSingleZeroAmountFields(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(0, 0),
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(0, 0),
				DiscountAmount: NewDecimal(0, 0),
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	lineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(lineItems), 1; g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}

	l := lineItems[0]
	if g, w := l.Name, "Name #1"; g != w {
		t.Errorf("got name %q, want %q", g, w)
	}
	if g, w := l.Kind, TransactionLineItemKindDebit; g != w {
		t.Errorf("got kind %q, want %q", g, w)
	}
	if g, w := l.Quantity, NewDecimal(10232, 4); g.Cmp(w) != 0 {
		t.Errorf("got quantity %q, want %q", g, w)
	}
	if g, w := l.UnitAmount, NewDecimal(451232, 4); g.Cmp(w) != 0 {
		t.Errorf("got unit amount %q, want %q", g, w)
	}
	if g, w := l.UnitTaxAmount, NewDecimal(0, 0); g.Cmp(w) != 0 {
		t.Errorf("got unit tax amount %q, want %q", g, w)
	}
	if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
		t.Errorf("got total amount %q, want %q", g, w)
	}
	if g, w := l.TaxAmount, NewDecimal(0, 0); g.Cmp(w) != 0 {
		t.Errorf("got tax amount %q, want %q", g, w)
	}
	if g, w := l.DiscountAmount, NewDecimal(0, 0); g.Cmp(w) != 0 {
		t.Errorf("got discount amount %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsSingle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	lineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(lineItems), 1; g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}

	l := lineItems[0]
	if g, w := l.Name, "Name #1"; g != w {
		t.Errorf("got name %q, want %q", g, w)
	}
	if g, w := l.Description, "Description #1"; g != w {
		t.Errorf("got description %q, want %q", g, w)
	}
	if g, w := l.Kind, TransactionLineItemKindDebit; g != w {
		t.Errorf("got kind %q, want %q", g, w)
	}
	if g, w := l.Quantity, NewDecimal(10232, 4); g.Cmp(w) != 0 {
		t.Errorf("got quantity %q, want %q", g, w)
	}
	if g, w := l.UnitAmount, NewDecimal(451232, 4); g.Cmp(w) != 0 {
		t.Errorf("got unit amount %q, want %q", g, w)
	}
	if g, w := l.UnitTaxAmount, NewDecimal(123, 2); g.Cmp(w) != 0 {
		t.Errorf("got unit tax amount %q, want %q", g, w)
	}
	if g, w := l.UnitOfMeasure, "gallon"; g != w {
		t.Errorf("got unit of measure %q, want %q", g, w)
	}
	if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
		t.Errorf("got total amount %q, want %q", g, w)
	}
	if g, w := l.TaxAmount, NewDecimal(455, 2); g.Cmp(w) != 0 {
		t.Errorf("got tax amount %q, want %q", g, w)
	}
	if g, w := l.DiscountAmount, NewDecimal(102, 2); g.Cmp(w) != 0 {
		t.Errorf("got discount amount %q, want %q", g, w)
	}
	if g, w := l.ProductCode, "23434"; g != w {
		t.Errorf("got product code %q, want %q", g, w)
	}
	if g, w := l.CommodityCode, "9SAASSD8724"; g != w {
		t.Errorf("got commodity code %q, want %q", g, w)
	}
	if g, w := l.URL, "https://example.com/products/23434"; g != w {
		t.Errorf("got url %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsMultiple(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:          "Name #2",
				Description:   "Description #2",
				Kind:          TransactionLineItemKindCredit,
				Quantity:      NewDecimal(202, 2),
				UnitAmount:    NewDecimal(5, 0),
				UnitOfMeasure: "gallon",
				TotalAmount:   NewDecimal(4515, 2),
				TaxAmount:     NewDecimal(455, 2),
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	lineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(lineItems), 2; g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}

	{
		l := lineItems[0]
		if g, w := l.Name, "Name #1"; g != w {
			t.Errorf("got name %q, want %q", g, w)
		}
		if g, w := l.Description, "Description #1"; g != w {
			t.Errorf("got description %q, want %q", g, w)
		}
		if g, w := l.Kind, TransactionLineItemKindDebit; g != w {
			t.Errorf("got kind %q, want %q", g, w)
		}
		if g, w := l.Quantity, NewDecimal(10232, 4); g.Cmp(w) != 0 {
			t.Errorf("got quantity %q, want %q", g, w)
		}
		if g, w := l.UnitAmount, NewDecimal(451232, 4); g.Cmp(w) != 0 {
			t.Errorf("got unit amount %q, want %q", g, w)
		}
		if g, w := l.UnitTaxAmount, NewDecimal(123, 2); g.Cmp(w) != 0 {
			t.Errorf("got unit tax amount %q, want %q", g, w)
		}
		if g, w := l.UnitOfMeasure, "gallon"; g != w {
			t.Errorf("got unit of measure %q, want %q", g, w)
		}
		if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
			t.Errorf("got total amount %q, want %q", g, w)
		}
		if g, w := l.TaxAmount, NewDecimal(455, 2); g.Cmp(w) != 0 {
			t.Errorf("got tax amount %q, want %q", g, w)
		}
		if g, w := l.DiscountAmount, NewDecimal(102, 2); g.Cmp(w) != 0 {
			t.Errorf("got discount amount %q, want %q", g, w)
		}
		if g, w := l.ProductCode, "23434"; g != w {
			t.Errorf("got product code %q, want %q", g, w)
		}
		if g, w := l.CommodityCode, "9SAASSD8724"; g != w {
			t.Errorf("got commodity code %q, want %q", g, w)
		}
		if g, w := l.URL, "https://example.com/products/23434"; g != w {
			t.Errorf("got url %q, want %q", g, w)
		}
	}

	{
		l := lineItems[1]
		if g, w := l.Name, "Name #2"; g != w {
			t.Errorf("got name %q, want %q", g, w)
		}
		if g, w := l.Description, "Description #2"; g != w {
			t.Errorf("got description %q, want %q", g, w)
		}
		if g, w := l.Kind, TransactionLineItemKindCredit; g != w {
			t.Errorf("got kind %q, want %q", g, w)
		}
		if g, w := l.Quantity, NewDecimal(202, 2); g.Cmp(w) != 0 {
			t.Errorf("got quantity %q, want %q", g, w)
		}
		if g, w := l.UnitAmount, NewDecimal(5, 0); g.Cmp(w) != 0 {
			t.Errorf("got unit amount %q, want %q", g, w)
		}
		if g, w := l.UnitOfMeasure, "gallon"; g != w {
			t.Errorf("got unit of measure %q, want %q", g, w)
		}
		if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
			t.Errorf("got total amount %q, want %q", g, w)
		}
		if g, w := l.TaxAmount, NewDecimal(455, 2); g.Cmp(w) != 0 {
			t.Errorf("got tax amount %q, want %q", g, w)
		}
		if g, w := l.DiscountAmount, (*Decimal)(nil); g != nil {
			t.Errorf("got discount amount %q, want %q", g, w)
		}
		if g, w := l.ProductCode, ""; g != w {
			t.Errorf("got product code %q, want %q", g, w)
		}
		if g, w := l.CommodityCode, ""; g != w {
			t.Errorf("got commodity code %q, want %q", g, w)
		}
		if g, w := l.URL, ""; g != w {
			t.Errorf("got url %q, want %q", g, w)
		}
	}
}

func TestTransactionWithLineItemsCommodityCodeTooLong(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "0123456789123",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("Transaction Sale got no error, want error")
	}
	t.Logf("%#v", err.(*BraintreeError).Errors.TransactionErrors.For("line-items"))
	errors := err.(*BraintreeError).Errors.TransactionErrors.For("line-items").For("index-1").On("commodity_code")
	if len(errors) != 1 {
		t.Fatalf("Transaction Sale got %d errors, want 1 error", len(errors))
	}
	if g, w := errors[0].Code, "95801"; g != w {
		t.Errorf("Transaction Sale got error code %s, want %s", g, w)
	}
	if g, w := errors[0].Message, "Commodity code is too long."; g != w {
		t.Errorf("Transaction Sale got error message %s, want %s", g, w)
	}
}
