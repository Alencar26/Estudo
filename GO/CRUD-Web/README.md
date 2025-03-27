# Anotações


```go
type Coffee struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Ingredients Ingredients `json:"ingredients"`
	Image       string      `json:"image"`
}
```
Para a struct Coffee, foi criado uma tabela no Postgres
```sql
create table coffee(
id integer primary key,
  title text not null,
  description text,
  price numeric(10,2) not null,
  ingredients jsonb not null,
  image text
);
```

Insert dos dados iniciais:
```SQL
create table coffee(
id integer primary key,
  title text not null,
  description text,
  price numeric(10,2) not null,
  ingredients jsonb not null,
  image text
);


INSERT INTO coffee (id, title, description, price, ingredients, image) VALUES
(1, 'Black Coffee', 'A simple, bold brew made without milk or sugar. It’s rich, aromatic, and perfect for those who enjoy a pure coffee flavor.', 1.5, '["Coffee", "Water"]'::JSONB, 'https://images.unsplash.com/photo-1494314671902-399b18174975?auto=format&fit=crop&q=80&w=1887&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(2, 'Latte', 'A smooth espresso-based drink with steamed milk and a light layer of foam. It’s creamy, mildly sweet, and perfect for a balanced coffee experience.', 2, '["Coffee", "Water"]'::JSONB, 'https://i.imgur.com/L3r6o58.jpeg'),
(3, 'Caramel Latte', 'A rich espresso drink blended with steamed milk and sweet caramel syrup. It’s smooth, creamy, and topped with a light foam for a perfect balance of sweetness.', 2, '["Coffee", "Water"]'::JSONB, 'https://images.unsplash.com/photo-1599398054066-846f28917f38?auto=format&fit=crop&q=80&w=1887&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(4, 'Cappuccino', 'A bold espresso drink topped with equal parts steamed milk and velvety foam. It’s rich, creamy, and perfect for those who love a strong coffee flavor.', 2, '["Espresso", "Ångad mjölk", "Foam"]'::JSONB, 'https://images.unsplash.com/photo-1557006021-b85faa2bc5e2?auto=format&fit=crop&q=80&w=1887&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(5, 'Espresso', 'A bold, concentrated coffee shot with a rich crema on top. It’s smooth, intense, and the foundation of many coffee drinks.', 1.2, '["Espresso"]'::JSONB, 'https://images.unsplash.com/photo-1579992357154-faf4bde95b3d?auto=format&fit=crop&q=80&w=1887&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(6, 'Machiatto', 'A bold espresso topped with a small amount of frothy milk. It’s rich, smooth, and perfect for those who enjoy a strong coffee flavor with a hint of creaminess.', 2.1, '["Coffee", "Water"]'::JSONB, 'https://images.pexels.com/photos/5097056/pexels-photo-5097056.jpeg'),
(7, 'Mocha', 'A delicious blend of espresso, steamed milk, and chocolate syrup. It’s rich, creamy, and perfect for those who love a sweet coffee with a chocolatey twist.', 2.5, '["Coffee", "Water"]'::JSONB, 'https://images.pexels.com/photos/15021332/pexels-photo-15021332/free-photo-of-black-coffee-in-cup-in-cozy-decor.jpeg'),
(9, 'Chai Latte', 'A warm, spiced blend of black tea, steamed milk, and aromatic spices like cinnamon and cardamom. It’s creamy, comforting, and perfect for a cozy treat.', 3, '["Coffee", "Water"]'::JSONB, 'https://www.teaheritage.fr/cdn/shop/articles/fbfd23eb1812c26fe623eebd7b4a5249.jpg?v=1674638932'),
(10, 'Matcha Latte', 'A smooth, earthy drink made with finely ground green tea powder and steamed milk. It’s rich in antioxidants and offers a naturally sweet, slightly grassy flavor.', 5, '["Coffee", "Water"]'::JSONB, 'https://images.pexels.com/photos/28704749/pexels-photo-28704749/free-photo-of-cozy-matcha-latte-with-autumn-decor-on-rustic-table.jpeg');

```