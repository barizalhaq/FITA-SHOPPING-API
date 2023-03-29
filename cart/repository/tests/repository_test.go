package repository_test

import (
	cart_repository "github.com/barizalhaq/fita_shopping_api/cart/repository"
	"github.com/barizalhaq/fita_shopping_api/domain"
	product_repository "github.com/barizalhaq/fita_shopping_api/product/repository"
	user_repository "github.com/barizalhaq/fita_shopping_api/user/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	dummyUser domain.User = domain.User{
		Username: "DUMMY_USER",
		Cart: domain.Cart{
			ID: 1,
		},
	}
	dummyCart     domain.Cart      = dummyUser.Cart
	dummyProducts []domain.Product = []domain.Product{
		domain.Product{
			SKU:   "DUMMY_SKU_1",
			Name:  "DUMMY PRODUCT 1",
			Price: 599.98,
			Qty:   15,
		},
		domain.Product{
			SKU:   "DUMMY_SKU_2",
			Name:  "DUMMY PRODUCT 2",
			Price: 123.98,
			Qty:   56,
		},
		domain.Product{
			SKU:   "DUMMY_SKU_3",
			Name:  "DUMMY PRODUCT 3",
			Price: 9876,
			Qty:   37,
		},
	}
	userRepo    domain.UserRepositoryInterface
	cartRepo    domain.CartRepositoryInterface
	productRepo domain.ProductRepositoryInterface
)

var _ = Describe("[Cart] Init", func() {
	BeforeEach(func() {
		err := user_repository.Migrate(Db)
		Ω(err).To(Succeed())

		err = product_repository.Migrate(Db)
		Ω(err).To(Succeed())

		err = cart_repository.Migrate(Db)
		Ω(err).To(Succeed())

		err = Db.AutoMigrate(&domain.CartItem{})
		Ω(err).To(Succeed())

		userRepo = user_repository.NewPGUserRepository(Db)
		cartRepo = cart_repository.NewPGCartRepository(Db)
		productRepo = product_repository.NewPGProductRepository(Db)

		err = Db.Create(&dummyUser).Error
		Ω(err).To(Succeed())

		err = Db.Create(&dummyProducts).Error
		Ω(err).To(Succeed())
	})
	Context("[Cart] View", func() {
		It("Found", func() {
			cart, err := cartRepo.View(&dummyUser)

			Ω(err).To(Succeed())
			Ω(cart.UserID).To(Equal(dummyUser.ID))
			Ω(cart.ID).To(Equal(dummyCart.ID))
		})
		It("Not Found", func() {
			notCreatedUser := &domain.User{
				ID: 999,
			}
			cart, err := cartRepo.View(notCreatedUser)

			Ω(err).To(Succeed())
			Ω(cart).To(BeNil())
		})
	})
	Context("[Cart] Create", func() {
		It("Succeed", func() {
			createdCart, err := cartRepo.Create(&dummyUser)

			Ω(err).To(Succeed())
			Ω(createdCart).ToNot(BeNil())
			Ω(createdCart.UserID).To(Equal(dummyUser.ID))
		})

		It("Failed", func() {
			createdCart, err := cartRepo.Create(&domain.User{ID: 100})

			Ω(err).ToNot(Succeed())
			Ω(createdCart).To(BeNil())
		})
	})
	Context("[Cart] Add/Subtracted/Removed Products", func() {
		It("Saved/Subtracted/Removed", func() {
			products, err := productRepo.GetProducts()
			Ω(err).To(Succeed())
			Ω(len(products)).To(Equal(3))
			Ω(products).ToNot(BeNil())

			// Add products
			productsToAdd := map[domain.Product]int{
				products[0]: 5,
				products[1]: 6,
				products[2]: 10,
			}

			mapQty := map[uint64]int{}
			for _, product := range products {
				mapQty[product.ID] = productsToAdd[product]
			}

			cart, err := cartRepo.AddProducts(&dummyUser, productsToAdd)
			Ω(err).To(Succeed())
			Ω(cart).ToNot(BeNil())
			Ω(len(cart.Items)).To(Equal(len(productsToAdd)))
			Ω(len(cart.Products)).To(Equal(len(productsToAdd)))
			for _, item := range cart.Items {
				Ω(item.Qty).To(Equal(mapQty[item.ProductID]))
			}

			// Subtract desired products
			productsToRemove := map[domain.Product]int{
				products[0]: 3,
				products[1]: 2,
				products[2]: 1,
			}

			for _, product := range products {
				mapQty[product.ID] = productsToAdd[product] - productsToRemove[product]
			}

			cart, err = cartRepo.SubtractProducts(&dummyUser, productsToRemove)
			Ω(err).To(Succeed())
			Ω(cart).ToNot(BeNil())
			Ω(len(cart.Items)).To(Equal(len(productsToRemove)))
			Ω(len(cart.Products)).To(Equal(len(productsToRemove)))
			for _, item := range cart.Items {
				Ω(item.Qty).To(Equal(mapQty[item.ProductID]))
			}

			// Remove desired products
			cart, err = cartRepo.RemoveProducts(&dummyUser, []int{int(products[0].ID)})
			Ω(err).To(Succeed())
			Ω(cart).ToNot(BeNil())
			Ω(len(cart.Items)).To(Equal(len(productsToAdd) - 1))

			cart, err = cartRepo.RemoveProducts(&dummyUser, []int{int(products[1].ID), int(products[2].ID)})
			Ω(err).To(Succeed())
			Ω(cart).ToNot(BeNil())
			Ω(len(cart.Items)).To(Equal(0))
		})
	})
})
