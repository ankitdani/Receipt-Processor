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

```git clone https://github.com/ankitdani/Receipt-Processor.git
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
      "total": number,
      "items": [
          {
              "shortDescription": "string",
              "price": number
          }
      ],
      "purchaseDate": "YYYY-MM-DD",
      "purchaseTime": "HH:MM"
  }
  ```

- **Endpoint**: `http:localhost:8080/receipts/{id}/points`
- **Method**: `GET`
```
{
    "id": "unique_receipt_id"
}
```

- **Response**
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
