// Code generated by "requestgen -type PlaceOrderRequest ./example/api"; DO NOT EDIT.

package api

func (p *PlaceOrderRequest) ClientOrderID(clientOrderID string) *PlaceOrderRequest {
	p.clientOrderID = &clientOrderID
	return p
}

func (p *PlaceOrderRequest) Symbol(symbol string) *PlaceOrderRequest {
	p.symbol = symbol
	return p
}

func (p *PlaceOrderRequest) Tag(tag string) *PlaceOrderRequest {
	p.tag = &tag
	return p
}

func (p *PlaceOrderRequest) Side(side SideType) *PlaceOrderRequest {
	p.side = side
	return p
}

func (p *PlaceOrderRequest) OrdType(ordType OrderType) *PlaceOrderRequest {
	p.ordType = ordType
	return p
}

func (p *PlaceOrderRequest) Size(size string) *PlaceOrderRequest {
	p.size = size
	return p
}

func (p *PlaceOrderRequest) Price(price string) *PlaceOrderRequest {
	p.price = &price
	return p
}

func (p *PlaceOrderRequest) TimeInForce(timeInForce TimeInForceType) *PlaceOrderRequest {
	p.timeInForce = &timeInForce
	return p
}

func (p *PlaceOrderRequest) ComplexArg(complexArg ComplexArg) *PlaceOrderRequest {
	p.complexArg = complexArg
	return p
}

func (p *PlaceOrderRequest) getParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	if p.clientOrderID != nil {
		a := *p.clientOrderID
		if len(a) == 0 {
			return params, fmt.Errorf("clientOid is required, empty string given")
		}
		params["clientOid"] = a
	}
	symbol := p.symbol
	if len(symbol) == 0 {
		return params, fmt.Errorf("symbol is required, empty string given")
	}
	params["symbol"] = symbol
	if p.tag != nil {
		a := *p.tag
		params["tag"] = a
	}
	side := p.side
	if len(side) == 0 {
		return params, fmt.Errorf("side is required, empty string given")
	}
	switch side {
	case "buy", "sell":
		params["side"] = side

	default:
		return params, fmt.Errorf("side value %v is not valid", a)

	}
	params["side"] = side
	ordType := p.ordType
	switch ordType {
	case "limit", "market":
		params["ordType"] = ordType

	default:
		return params, fmt.Errorf("ordType value %v is not valid", a)

	}
	params["ordType"] = ordType
	size := p.size
	params["size"] = size
	if p.price != nil {
		a := *p.price
		params["price"] = a
	}
	if p.timeInForce != nil {
		a := *p.timeInForce
		switch a {
		case "GTC", "GTT", "FOK":
			params["timeInForce"] = a

		default:
			return params, fmt.Errorf("timeInForce value %v is not valid", a)

		}
		params["timeInForce"] = a
	}
	complexArg := p.complexArg
	params["complexArg"] = complexArg
	return params, nil
}