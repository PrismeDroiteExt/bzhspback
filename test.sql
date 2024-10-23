-- Users
INSERT INTO users (mail, password, register_date, name, last_name) VALUES
('john@example.com', 'hashed_password_1', '2023-01-01', 'John', 'Doe'),
('jane@example.com', 'hashed_password_2', '2023-01-02', 'Jane', 'Smith'),
('alice@example.com', 'hashed_password_3', '2023-01-03', 'Alice', 'Johnson'),
('bob@example.com', 'hashed_password_4', '2023-01-04', 'Bob', 'Williams'),
('emma@example.com', 'hashed_password_5', '2023-01-05', 'Emma', 'Brown');

-- Categories
INSERT INTO categories (id, name, picture_url, description, sub_category_id) VALUES
(1, 'Electronics', 'https://example.com/electronics.jpg', 'Electronic devices and accessories', NULL),
(2, 'Smartphones', 'https://example.com/smartphones.jpg', 'Mobile phones and accessories', 1),
(3, 'Laptops', 'https://example.com/laptops.jpg', 'Portable computers', 1),
(4, 'Clothing', 'https://example.com/clothing.jpg', 'Apparel and fashion items', NULL),
(5, 'T-shirts', 'https://example.com/tshirts.jpg', 'Casual and formal t-shirts', 4),
(6, 'Jeans', 'https://example.com/jeans.jpg', 'Denim pants for all occasions', 4);

-- Brands
INSERT INTO brands (name) VALUES
('Apple'),
('Samsung'),
('Dell'),
('Levi''s'),
('Nike'),
('Adidas');

-- Products
INSERT INTO products (title, description, category_id, brand_id, price, discount, colors, sizes, picture_url) VALUES
('iPhone 12', 'Latest iPhone model', 2, 1, 999.99, 0.10, 'Black,White,Red', '64GB,128GB,256GB', 'https://example.com/iphone12.jpg'),
('MacBook Pro', '13-inch MacBook Pro with M1 chip', 3, 1, 1299.99, 0.05, 'Silver,Space Gray', '8GB,16GB', 'https://example.com/macbookpro.jpg'),
('Samsung Galaxy S21', 'Flagship Android smartphone', 2, 2, 799.99, 0.15, 'Phantom Black,Phantom Silver,Phantom Violet', '128GB,256GB', 'https://example.com/galaxys21.jpg'),
('Dell XPS 13', 'Ultra-thin and light laptop', 3, 3, 999.99, 0.08, 'Platinum Silver,Frost White', '8GB,16GB,32GB', 'https://example.com/dellxps13.jpg'),
('Classic Cotton T-Shirt', 'Comfortable everyday t-shirt', 5, 5, 19.99, 0.00, 'White,Black,Navy,Gray', 'S,M,L,XL', 'https://example.com/cottontshirt.jpg'),
('Slim Fit Jeans', 'Modern slim fit denim jeans', 6, 4, 59.99, 0.20, 'Blue,Black,Gray', '28,30,32,34,36', 'https://example.com/slimfitjeans.jpg');


-- Carts
INSERT INTO carts (user_id, is_processed) VALUES
(1, false),
(2, true),
(3, false),
(4, true),
(5, false);

-- Cart Items
INSERT INTO cart_items (cart_id, product_id, quantity) VALUES
(1, 1, 1),
(1, 2, 1),
(2, 1, 2),
(3, 3, 1),
(3, 5, 3),
(4, 4, 1),
(4, 6, 2),
(5, 2, 1),
(5, 5, 2);

-- Orders
INSERT INTO orders (cart_id, status) VALUES
(2, 'Completed'),
(4, 'Processing');

-- Payments
INSERT INTO payments (order_id, user_id, status, payment_method, transaction_id) VALUES
(1, 2, 'Completed', 'Credit Card', 'TXN123456789'),
(2, 4, 'Processing', 'PayPal', 'TXN987654321');