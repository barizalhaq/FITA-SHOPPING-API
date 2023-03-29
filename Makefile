DB_USER = postgres
DB_PASS = postgres
DB_NAME = fita_shopping
DB_SSL = disable

migration-up:
	goose -dir=./migrations postgres "host=localhost port=5432 user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} sslmode=${DB_SSL}" up

db-status:
	goose postgres "host=localhost port=5432 user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} sslmode=${DB_SSL}" status

migration-reset:
	goose -dir=./migrations postgres "host=localhost port=5432 user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} sslmode=${DB_SSL}" reset

mock-generate:
	mockgen -source=./domain/user.go -destination=./domain/mocks/mock_user.go -package=mocks UserRepositoryInterface,UserUsecaseInterface
	mockgen -source=./domain/product.go -destination=./domain/mocks/mock_product.go -package=mocks ProductUsecaseInterface,ProductRepositoryInterface
	mockgen -source=./domain/cart.go -destination=./domain/mocks/mock_cart.go -package=mocks CartUsecaseInterface,CartRepositoryInterface
	mockgen -source=./domain/promo.go -destination=./domain/mocks/mock_promo.go -package=mocks PromoUsecaseInterface,PromoRepositoryInterface
