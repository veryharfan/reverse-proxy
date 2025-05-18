user-service:
- register user [v]
- login user [v]
- create token [v]

product-service:
- get product by id 
- get list product
    - via db
    - via search engine (todo)
- caching stock
- creating product
- updating product and status
- emit event to stock.init?
- emit event for product.created?

shop-service:
- creating shop
- call warehouse to create warehouse
- create product 
    - call product-service to create product
    - call warehouse to init product stock after creating product

warehouse-service:
- create warehouse
- get warehouse by shop_id
- updating product stock
- get product stock
- receive api from shop-service to init product-stock
    - emit event stock.init, consumed by product-service to cache the stock
- emit event stock.deducted, stock.reserved, stock.returned, stock.init?
- api/consume event to deduct stock by product-id after when order received?