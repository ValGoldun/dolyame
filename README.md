# dolyame
Go клиент для интеграции с Долями

Создание клиента:
```
cert, err := tls.LoadX509KeyPair("cert.pem", "private.key")
if err != nil {
	return err
}

tlsConfig := &tls.Config{
	Certificates: []tls.Certificate{cert},
}

transport := &http.Transport{TLSClientConfig: tlsConfig}

httpClient := &http.Client{Transport: transport} 

client := dolyame.NewClient().SetHost("").SetLogin("").SetPassword("").SetHttpClient(httpClient)
```

Создание заказа:
```
order := dolyame.NewOrder().
	SetAmount(amount).
	SetID(orderID).
	SetPrepaidAmount(prepaidAmount)
order.AddItem(
    dolyame.NewItem().
	    SetName(name).
		SetPrice(price).
		SetQuantity(quantity),
)
clientInfo := dolyame.NewClientInfo().
			SetEmail(email).
			SetPhone(phone).
			SetBirthDate(birthdate).
			SetFirstName(fname).
			SetLastName(lname).
			SetMiddleName(mname)

createRequest := dolyame.NewCreateRequest().SetOrder(order).SetClientInfo(clientInfo)
correlationID := uuid.New().String()
createResponse, err := client.Create(createRequest, correlationID)
```

Документация в процессе