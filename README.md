# Receipt-Processor

# Receipt Points Calculator

This is a golang application that calculates points based on submitted receipts. The points are determined by retailer name, total amount, number of items, and purchase time.

## Features

- Process receipts and calculate points based on specific rules.
- Retrieve points for a processed receipt using its ID.


## Requirements

- Go
- Docker

## Installation

1. Clone the repository

```
git clone https://github.com/ankitdani/Receipt-Processor.git
```

2. Navigate to directory
```
cd Receipt-Processor-main
```

3. Run using docker compose

```
docker-compose up --build
```

## Endpoint structure

- **Endpoint**: `http:localhost:8080/receipts/process`
- **Method**: `POST`
- **Request Body**: 
  ```json
  {
      "retailer": "string",
      "purchaseDate": "YYYY-MM-DD",
      "purchaseTime": "HH:MM",
      "items": [
          {
              "shortDescription": "string",
              "price": number
          }
      ],
      "total": number
  }
  ```

-**Test using curl**:
```
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '<your-POST-body>'  
```

- **Sample POST body**:
```
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

- **Response**:
```
{
    "id": "unique_receipt_id"
}
```

- **Endpoint**: `http:localhost:8080/receipts/{id}/points`
- **Method**: `GET`
- **Test using curl**:
```
curl http://localhost:8080/receipts/<your-uuid>/points
```
- **Response**:
```
{
    "points": number
}
```
If the receipt is not found:
```
{
    "error": "Receipt not found"
}
```
