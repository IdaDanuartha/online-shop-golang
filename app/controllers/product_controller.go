package controllers

import (
	"net/http"
	"strconv"

	"github.com/IdaDanuartha/online-shop-golang/app/models"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
		Extensions: []string{".html", ".tmpl"},
	})

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage := 9

	productModel := models.Product{}
	products, totalRows, err := productModel.GetProducts(server.DB, perPage, page)
	if err != nil {
		return
	}

	pagination, _ := GetPaginationLinks(server.AppConfig, PaginationParams{
		Path:        "products",
		TotalRows:   int32(totalRows),
		PerPage:     int32(perPage),
		CurrentPage: int32(page),
	})

	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products":   products,
		"pagination": pagination,
	})
}

func (server *Server) GetProductBySlug(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
        Layout: "layout",
		Extensions: []string{".html", ".tmpl"},
    })

    vars := mux.Vars(r)
    slug := vars["slug"]

    productModel := models.Product{}
    product, err := productModel.FindBySlug(server.DB, slug)
    if err != nil {
        return
    }

    _ = render.HTML(w, http.StatusOK, "product", map[string]interface{}{
        "product": product,
    })
}
