package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/sergio/go-hexagonal/adapters/web/dto"
	"github.com/sergio/go-hexagonal/application/product"
)

func ProductHandlers(r *mux.Router, n *negroni.Negroni, service product.IProductService) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/product/", n.With(
		negroni.Wrap(getProducts(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/product/", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PATCH", "OPTIONS")
	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PATCH", "OPTIONS")
}

func createProduct(service product.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var newProduct product.Product
		w.Header().Set("Content-Type", "application/json")

		productDTO := dto.NewProductDTO()
		err := json.NewDecoder(r.Body).Decode(productDTO)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError("json payload is required"))
			return
		}

		productDTO.Bind(&newProduct)
		product, err := service.Create(newProduct.Name, newProduct.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func enableProduct(service product.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		service.Enable(product)
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("[ERROR] getProducts:", err.Error())
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})
}

func disableProduct(service product.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		service.Disable(product)
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("[ERROR] getProducts:", err.Error())
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})
}

func getProducts(service product.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		dbProducts, err := service.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("[ERROR] getProducts:", err.Error())
			return
		}
		err = json.NewEncoder(w).Encode(dbProducts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("[ERROR] getProducts:", err.Error())
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})
}

func getProduct(service product.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		id := vars["id"]

		dbProduct, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(JsonError("product not found!"))
			return
		}

		err = json.NewEncoder(w).Encode(dbProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("[ERROR] getProduct:", err.Error())
			return
		}

		w.WriteHeader(http.StatusAccepted)
	})
}
