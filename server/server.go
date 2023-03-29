package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	cart_repository "github.com/barizalhaq/fita_shopping_api/cart/repository"
	cart_usecase "github.com/barizalhaq/fita_shopping_api/cart/usecase"
	"github.com/barizalhaq/fita_shopping_api/config"
	"github.com/barizalhaq/fita_shopping_api/domain"
	"github.com/barizalhaq/fita_shopping_api/graph"
	"github.com/barizalhaq/fita_shopping_api/graph/resolver"
	product_repository "github.com/barizalhaq/fita_shopping_api/product/repository"
	product_usecase "github.com/barizalhaq/fita_shopping_api/product/usecase"
	promo_repository "github.com/barizalhaq/fita_shopping_api/promo/repository"
	user_middleware "github.com/barizalhaq/fita_shopping_api/user/delivery/middleware"
	user_repository "github.com/barizalhaq/fita_shopping_api/user/repository"
	user_usecase "github.com/barizalhaq/fita_shopping_api/user/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DefaultPort = ":8080"

func init() {
	config.Init()
}

type resolverDeps struct {
	userUC    domain.UserUsecaseInterface
	productUC domain.ProductUsecaseInterface
	cartUC    domain.CartUsecaseInterface
}

// Defining the Graphql handler
func graphqlHandler(deps resolverDeps) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &resolver.Resolver{
					UserUsecase:    deps.userUC,
					ProductUsecase: deps.productUC,
					CartUsecase:    deps.cartUC,
				},
			},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Initiate DB with GORM
	dsn := config.C.Database.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Initiate repositories
	userRepo := user_repository.NewPGUserRepository(db)
	productRepo := product_repository.NewPGProductRepository(db)
	cartRepo := cart_repository.NewPGCartRepository(db)
	promoRepo := promo_repository.NewPromoRepository()

	// Inititate Usecases
	userUC := user_usecase.NewUserUsecase(userRepo)
	productUC := product_usecase.NewProductUsecase(productRepo)
	cartUC := cart_usecase.NewCartUsecase(cartRepo, productRepo, promoRepo)

	// Initiate Gin for http route
	r := gin.Default()

	r.Use(user_middleware.InitMiddleware())
	r.Use(user_middleware.Authenticated(userRepo))

	r.POST("/query", graphqlHandler(
		resolverDeps{
			userUC:    userUC,
			productUC: productUC,
			cartUC:    cartUC,
		},
	))
	r.GET("/", playgroundHandler())
	r.Run(DefaultPort)
}
