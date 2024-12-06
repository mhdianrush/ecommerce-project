# E-Commerce API Documentation 🛒

This API used to maintain products and brands in the e-commerce app. This documentation is to describe how to use all of the endpoint.

## 📋 **Table of Contents**
- [Prerequisites](#prerequisites)
- [Base URL](#base-url)
- [Endpoints](#endpoints)
  - [Products](#products)
    - [Create Product](#create-product)
    - [Get Product by ID](#get-product-by-id)
    - [Update Product](#update-product)
    - [Delete Product](#delete-product)
    - [Get All Products with Pagination](#get-all-products-with-pagination)
  - [Brands](#brands)
    - [Create Brand](#create-brand)
    - [Delete Brand](#delete-brand)
- [How to Run](#how-to-run)
- [Notes](#notes)

---

## 🛠️ **Prerequisites**
- A running instance of the API (default: `http://localhost:8000`).
- Tools to test the API:
  - [Postman](https://www.postman.com/)
  - `curl` (command line tool)
  - HTTP client library (e.g., Axios, Fetch).

---

## 🌐 **Base URL**

```plaintext
http://localhost:8000

## 🔥 Endpoints Overview

| Method | Endpoint           | Description                     |
|--------|--------------------|---------------------------------|
| POST   | `/products`        | Create a new product            |
| GET    | `/products/{id}`   | Get product details by ID       |
| PUT    | `/products/{id}`   | Update product details by ID    |
| DELETE | `/products/{id}`   | Delete a product by ID          |
| GET    | `/products`        | Get all products with pagination|
| POST   | `/brands`          | Create a new brand              |
| DELETE | `/brands/{id}`     | Delete a brand (if not in use)  |

---

## 🚀 How to Run
- Clone repository:
  - git clone <repo-url>
  - cd <project-folder>

- run local server:
  - go run main.go

