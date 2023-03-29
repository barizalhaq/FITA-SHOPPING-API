<h1 align="center" height="800" font-weight="bold">FITA SHOPPING API</h1>

# Requirements
- GO >= 1.19
- GO MOD
- DOCKER
- DOCKER COMPOSE
- GOOSE (for database migration)
- GOMOCK
- POSTMAN

# Installation
- Clone from the git repository
- Open cloned directory
- Run `go mod download`
- Run `docker-compose -f docker-compose.dev.yml up` for running applications inside a docker container
- Run `make migration-up` to populate database schema
- Run 
```SQL
    INSERT INTO products (sku,name,price,qty)
    VALUES
    ('120P90', 'Google Home', 49.99, 10),
    ('43N23P', 'MacBook Pro', 5399.99, 5),
    ('A304SD', 'Alexa Speaker', 109.50, 10),
    ('234234', 'Raspberry Pi B', 30, 2);
```
to populate products

- Open up POSTMAN and you're good to go

# API's
Shopping API built with GraphQL API. This kind of backend service is used for shopping purposes with an interesting promo in it. APIs serve several features such as registering users to secure their cart, authentication, creating/viewing carts, adding/removing or even decreasing product quantity within their cart, and last but not least checkout. Checkout API will give a summarized invoice that comes with an invoice consisting of several pieces of information, for example, original price (total price without promo discount), actual price (the total price which already discounts the user should pay), price discount (discount total), and currency. To simplify the testing purposes here's the POSTMAN Collection.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/4551391-054eb2d9-1982-4040-8aa6-dde1e0a1e16d?action=collection%2Ffork&collection-url=entityId%3D4551391-054eb2d9-1982-4040-8aa6-dde1e0a1e16d%26entityType%3Dcollection%26workspaceId%3D6ceb5ec7-69f0-41df-b9d3-806e5455d740)

<br />

<h5 font-weight="bold">OR WITH COLLECTION</h5>

<br />

```json
{
	"info": {
		"_postman_id": "054eb2d9-1982-4040-8aa6-dde1e0a1e16d",
		"name": "fita shopping api",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "4551391"
	},
	"item": [
		{
			"name": "mutations",
			"item": [
				{
					"name": "register",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "graphql",
							"graphql": {
								"query": "mutation register($input: RegisterInput!) {\n    register(input: $input)\n    {\n        user {\n            ID\n            username\n        }\n    }\n}",
								"variables": "{\r\n    \"input\": {\r\n        \"username\": \"user\",\r\n        \"password\": \"password\"\r\n    }\r\n}"
							}
						},
						"url": {
							"raw": "{{url}}/query",
							"host": [
								"{{url}}"
							],
							"path": [
								"query"
							]
						}
					},
					"response": []
				},
				{
					"name": "enCart",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "graphql",
							"graphql": {
								"query": "mutation enCart ($input: EncartInput!) {\n    enCart (input: $input) {\n        ID\n        owner {\n            ID\n            username\n        }\n        products {\n            ID\n            name\n            qty\n            sku\n            price {\n                originalPrice\n                currency\n            }\n        }\n    }\n}",
								"variables": "{\n  \"input\": {\n    \"productsToAdd\": {\n      \"productID\": 4,\n      \"qty\": 1\n    }\n  }\n}"
							}
						},
						"url": {
							"raw": "{{url}}/query",
							"host": [
								"{{url}}"
							],
							"path": [
								"query"
							]
						}
					},
					"response": []
				},
				{
					"name": "deCart",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "graphql",
							"graphql": {
								"query": "mutation deCart ($input: DecartInput!) {\n    deCart (input: $input) {\n        ID\n        owner {\n            ID\n            username\n        }\n        products {\n            ID\n            name\n        }\n    }\n}",
								"variables": "{\n  \"input\": {\n    \"productIDs\": [\n      1\n    ]\n  }\n}"
							}
						},
						"url": {
							"raw": "{{url}}/query",
							"host": [
								"{{url}}"
							],
							"path": [
								"query"
							]
						}
					},
					"response": []
				},
				{
					"name": "decreaseCartProductQty",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "graphql",
							"graphql": {
								"query": "mutation decreaseCartProductQty ($input: DecreaseCartProductQtyInput!) {\n    decreaseCartProductQty (input: $input) {\n        ID\n    }\n}",
								"variables": "{\n  \"input\": {\n    \"productsToAdd\": {\n      \"productID\": 0,\n      \"qty\": 0\n    }\n  }\n}"
							}
						},
						"url": {
							"raw": "{{url}}/query",
							"host": [
								"{{url}}"
							],
							"path": [
								"query"
							]
						}
					},
					"response": []
				},
				{
					"name": "checkout",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "graphql",
							"graphql": {
								"query": "mutation checkout {\n    checkout {\n        cart {\n            ID\n            owner {\n                ID\n                username\n            }\n            products {\n                ID\n                name\n                price {\n                    currency\n                    originalPrice\n                }\n                qty\n                sku\n            }\n        }\n        total_price {\n            actualCurrentPrice\n            originalTotalPrice\n            priceDiscount\n            currency\n        }\n    }\n}",
								"variables": "{}"
							}
						},
						"url": {
							"raw": "{{url}}/query",
							"host": [
								"{{url}}"
							],
							"path": [
								"query"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "queries",
			"item": [
				{
					"name": "authenticate",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "graphql",
							"graphql": {
								"query": "query authenticate ($input: AuthenticateInput!) {\n    authenticate (input: $input) {\n        authenticated\n    }\n}",
								"variables": "{\n  \"input\": {\n    \"username\": \"user\",\n    \"password\": \"user\"\n  }\n}"
							}
						},
						"url": {
							"raw": "{{url}}/query",
							"host": [
								"{{url}}"
							],
							"path": [
								"query"
							]
						}
					},
					"response": []
				},
				{
					"name": "products",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "graphql",
							"graphql": {
								"query": "query products {\n    products {\n        ID\n        sku\n        name\n        qty\n    }\n}",
								"variables": "{}"
							}
						},
						"url": {
							"raw": "{{url}}/query",
							"host": [
								"{{url}}"
							],
							"path": [
								"query"
							]
						}
					},
					"response": []
				},
				{
					"name": "cart",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "graphql",
							"graphql": {
								"query": "query cart {\n    cart {\n        ID\n        owner {\n            ID\n            username\n        }\n        products {\n            ID\n            name\n            qty\n        }\n    }\n}",
								"variables": "{}"
							}
						},
						"url": {
							"raw": "{{url}}/query",
							"host": [
								"{{url}}"
							],
							"path": [
								"query"
							]
						}
					},
					"response": []
				}
			]
		}
	],
	"variable": [
		{
			"key": "url",
			"value": "",
			"type": "any",
			"description": "URL for the request."
		}
	]
}
```

# GraphQL Schema
```
type User {
    ID: Int!
    username: String!
}

type Product {
    ID: Int!
    sku: String!
    name: String!
    price: ProductPrice!
    qty: Int!
}

type Cart {
    ID: Int!
    owner: User!
    products: [Product]!
}

type ProductPrice {
    originalPrice: Float!
    currency: String!
}

type Invoice {
    cart: Cart!
    total_price: InvoicePrice!
}

type InvoicePrice {
    originalTotalPrice: Float!
    actualCurrentPrice: Float!
    priceDiscount: Float!
    currency: String!
}

input RegisterInput {
    username: String!
    password: String!
}

input AuthenticateInput {
    username: String!
    password: String!
}

input EncartInput {
    productsToAdd: [ProductWithQty!]!
}

input ProductWithQty {
    productID: Int!
    qty: Int!
}

input DecreaseCartProductQtyInput {
    productsToAdd: [ProductWithQty!]!
}

input DecartInput {
    productIDs: [Int!]!
}

type RegisterResponse {
    user: User!
}

type AuthenticateResponse {
    authenticated: Boolean!
}

type Query {
    authenticate(input: AuthenticateInput!): AuthenticateResponse!
    products: [Product]!
    cart: Cart!
}

type Mutation {
    register(input: RegisterInput!): RegisterResponse!
    enCart(input: EncartInput!): Cart!
    deCart(input: DecartInput!): Cart!
    decreaseCartProductQty(input: DecreaseCartProductQtyInput!): Cart!
    checkout: Invoice!
}
```