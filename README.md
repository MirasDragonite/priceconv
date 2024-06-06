
# Go priceconv

`Priceconv` is a package for converting currency values using exchange rates. The package fetches daily exchange rates from the Central Bank of Russia and provides methods for currency conversion


## Features

- Fetch daily exchange rates from the Central Bank of Russia.
- Convert between different currencies.
- Convert formatted string prices to a target currency.

## Installation

To install the package, use `go get`:

```
go get github.com/MirasDragonite/priceconv@latest
```

## Usage
### Initialization
First, initialize the ValuteModule:


```
package main

import (
    "fmt"
    "priceconv"
)

func main() {
    valuteConverter := priceconv.CreateValuteConverter()
    valuteConverter.GetDailyRate()
}
```

### Fetching Daily Rates

Fetch the latest exchange rates:

```
valuteConverter.GetDailyRate()
```


### Converting Currency
Convert an amount from one currency to another:

```
amount, err := valuteConverter.ConvertCurrency(100, "USD", "EUR")
if err != nil {
    fmt.Println("Error converting currency:", err)
} else {
    fmt.Printf("Converted amount: %.2f\n", amount)
}
```
- `amount`: The amount to be converted.
- `fromCode`: The source currency code.
- `toCode`: The target currency code.
Returns the converted amount and an error if the conversion fails.

### Converting String Prices
Convert a formatted string price (e.g., "100 USD") to a target currency (EUR):

```
convertedPrice, err := valuteConverter.ConverStringPrices("100 USD")
if err != nil {
    fmt.Println("Error converting string price:", err)
} else {
    fmt.Printf("Converted price in EUR: %.2f\n", convertedPrice)
}
```

- `price`: The formatted price string (e.g., "100 USD").
Returns the converted price in EUR and an error if the conversion fails.

### Access Specific Rate
Access the specific currency rate from the Valutes map using the currency code (e.g., "USD" for US Dollar).

```
usdRate := result.Valutes["USD"]
```

The usdRate is of type Valute, which contains the following fields:

```
type Valute struct {
    NumCode   string   `xml:"NumCode"`
    CharCode  string   `xml:"CharCode"`
    Nominal   string   `xml:"Nominal"`
    Name      string   `xml:"Name"`
    Value     string   `xml:"Value"`
    VunitRate string   `xml:"VunitRate"`
}
```
- `NumCode`: The numeric code representing the currency (e.g., "840" for USD).

- `CharCode`: The three-character code representing the currency (e.g., "USD" for US Dollar).

- `Nominal`: The nominal value of the currency unit (e.g., "1" for 1 US Dollar, "10" for 10 Russian Rubles).

- `Name`: The full name of the currency (e.g., "US Dollar").

- `Value`: The exchange rate value of the currency as a string (e.g., "73.25" for the value of 1 USD in RUB).

- `VunitRate`: The rate per unit of the currency, which is typically the same as Value and is represented as a string.


