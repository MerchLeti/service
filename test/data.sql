-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               PostgreSQL 16.0 on x86_64-pc-linux-musl, compiled by gcc (Alpine 12.2.1_git20220924-r10) 12.2.1 20220924, 64-bit
-- Server OS:                    
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table public.categories
CREATE TABLE IF NOT EXISTS "categories" (
	"id" BIGINT NOT NULL DEFAULT 'nextval(''categories_id_seq''::regclass)',
	"name" TEXT NOT NULL,
	"parent" BIGINT NULL DEFAULT NULL,
	PRIMARY KEY ("id")
);

-- Dumping data for table public.categories: -1 rows
/*!40000 ALTER TABLE "categories" DISABLE KEYS */;
INSERT INTO "categories" ("id", "name", "parent") VALUES
	(1, 'Одежда', NULL),
	(2, 'Футболки', 1),
	(3, 'Худи', 1);
/*!40000 ALTER TABLE "categories" ENABLE KEYS */;

-- Dumping structure for table public.images
CREATE TABLE IF NOT EXISTS "images" (
	"id" BIGINT NOT NULL DEFAULT 'nextval(''images_id_seq''::regclass)',
	"item" BIGINT NOT NULL,
	"position" INTEGER NOT NULL,
	"url" TEXT NOT NULL,
	PRIMARY KEY ("id"),
	INDEX "images_item" ("item")
);

-- Dumping data for table public.images: -1 rows
/*!40000 ALTER TABLE "images" DISABLE KEYS */;
INSERT INTO "images" ("id", "item", "position", "url") VALUES
	(1, 1, 1, 'https://raw.githubusercontent.com/MerchLeti/images/main/hoodies/1.png'),
	(2, 2, 1, 'https://raw.githubusercontent.com/MerchLeti/images/main/hoodies/2.png'),
	(3, 3, 1, 'https://raw.githubusercontent.com/MerchLeti/images/main/hoodies/3.png'),
	(4, 4, 1, 'https://raw.githubusercontent.com/MerchLeti/images/main/hoodies/4.png'),
	(5, 5, 1, 'https://raw.githubusercontent.com/MerchLeti/images/main/hoodies/5.png'),
	(6, 6, 1, 'https://raw.githubusercontent.com/MerchLeti/images/main/hoodies/6.png'),
	(7, 7, 1, 'https://raw.githubusercontent.com/MerchLeti/images/main/hoodies/7.png');
/*!40000 ALTER TABLE "images" ENABLE KEYS */;

-- Dumping structure for table public.items
CREATE TABLE IF NOT EXISTS "items" (
	"id" BIGINT NOT NULL DEFAULT 'nextval(''items_id_seq''::regclass)',
	"name" TEXT NOT NULL,
	"description" TEXT NOT NULL,
	"category" BIGINT NULL DEFAULT NULL,
	PRIMARY KEY ("id"),
	INDEX "items_category" ("category")
);

-- Dumping data for table public.items: -1 rows
/*!40000 ALTER TABLE "items" DISABLE KEYS */;
INSERT INTO "items" ("id", "name", "description", "category") VALUES
	(1, 'Черный худак', 'Свитшоты давно выросли из футбольной формы и появились
в гардеробах множества людей. Благодаря свободному крою уютные теплые кофты не сковывают движений, а легкая ткань позволяет пропускать воздух и не препятствует дыханию кожи. Свитшот позволяет чувствовать себя комфортно и на прогулке,
и на работе, и на тренировке в зале', 3),
	(2, 'Красный худак', 'Свитшоты давно выросли из футбольной формы и появились в гардеробах множества людей. Благодаря свободному крою уютные теплые кофты не сковывают движений, а легкая ткань позволяет пропускать воздух и не препятствует дыханию кожи. Свитшот позволяет чувствовать себя комфортно и на прогулке, и на работе, и на тренировке в зале', 3),
	(3, 'Белый худак', 'Свитшоты давно выросли из футбольной формы и появились в гардеробах множества людей. Благодаря свободному крою уютные теплые кофты не сковывают движений, а легкая ткань позволяет пропускать воздух и не препятствует дыханию кожи. Свитшот позволяет чувствовать себя комфортно и на прогулке, и на работе, и на тренировке в зале', 3),
	(4, 'Черный жэнский худак', 'Свитшоты давно выросли из футбольной формы и появились в гардеробах множества людей. Благодаря свободному крою уютные теплые кофты не сковывают движений', 3),
	(5, 'Голубой худак', 'Свитшоты давно выросли из футбольной формы и появились в гардеробах множества людей. Благодаря свободному крою уютные теплые кофты не сковывают движений', 3),
	(6, 'Синий жэнский худак', 'Свитшоты давно выросли из футбольной формы и появились в гардеробах множества людей. Благодаря свободному крою уютные теплые кофты не сковывают движений', 3),
	(7, 'Синий худак', 'Свитшоты давно выросли из футбольной формы и появились в гардеробах множества людей. Благодаря свободному крою уютные теплые кофты не сковывают движений', 3);
/*!40000 ALTER TABLE "items" ENABLE KEYS */;

-- Dumping structure for table public.properties
CREATE TABLE IF NOT EXISTS "properties" (
	"id" BIGINT NOT NULL DEFAULT 'nextval(''properties_id_seq''::regclass)',
	"item" BIGINT NOT NULL,
	"name" TEXT NOT NULL,
	"value" TEXT NOT NULL,
	INDEX "properties_item" ("item"),
	PRIMARY KEY ("id"),
	INDEX "types_item" ("item")
);

-- Dumping data for table public.properties: 56 rows
/*!40000 ALTER TABLE "properties" DISABLE KEYS */;
INSERT INTO "properties" ("id", "item", "name", "value") VALUES
	(1, 1, 'Тип ткани', 'Футер трехнитка премиум пенье с начесом Penie 330гр'),
	(3, 1, 'Качество ткани', 'Premium | Penie'),
	(4, 1, 'Внутренняя сторона', 'Ворсованная / С начесом'),
	(5, 1, 'Плетение ткани', '3х-ниточное'),
	(6, 1, 'Плотность (гр/м.кв)', '320'),
	(7, 1, 'Состав', '70% хлопок, 30% полиэстер'),
	(8, 1, 'Рекомендации по уходу', 'Стирка 40 градусов. Отжим до 800 оборотов. Сушка на вешало'),
	(9, 2, 'Тип ткани', 'Футер трехнитка премиум пенье с начесом Penie 330гр'),
	(11, 2, 'Качество ткани', 'Premium | Penie'),
	(12, 2, 'Внутренняя сторона', 'Ворсованная / С начесом'),
	(13, 2, 'Плетение ткани', '3х-ниточное'),
	(14, 2, 'Плотность (гр/м.кв)', '320'),
	(15, 2, 'Состав', '70% хлопок, 30% полиэстер'),
	(16, 2, 'Рекомендации по уходу', 'Стирка 40 градусов. Отжим до 800 оборотов. Сушка на вешало'),
	(17, 3, 'Тип ткани', 'Футер трехнитка премиум пенье с начесом Penie 330гр'),
	(19, 3, 'Качество ткани', 'Premium | Penie'),
	(20, 3, 'Внутренняя сторона', 'Ворсованная / С начесом'),
	(21, 3, 'Плетение ткани', '3х-ниточное'),
	(22, 3, 'Плотность (гр/м.кв)', '320'),
	(23, 3, 'Состав', '70% хлопок, 30% полиэстер'),
	(24, 3, 'Рекомендации по уходу', 'Стирка 40 градусов. Отжим до 800 оборотов. Сушка на вешало'),
	(25, 4, 'Тип ткани', 'Футер трехнитка премиум пенье с начесом Penie 330гр'),
	(27, 4, 'Качество ткани', 'Premium | Penie'),
	(28, 4, 'Внутренняя сторона', 'Ворсованная / С начесом'),
	(29, 4, 'Плетение ткани', '3х-ниточное'),
	(30, 4, 'Плотность (гр/м.кв)', '320'),
	(31, 4, 'Состав', '70% хлопок, 30% полиэстер'),
	(32, 4, 'Рекомендации по уходу', 'Стирка 40 градусов. Отжим до 800 оборотов. Сушка на вешало'),
	(33, 5, 'Тип ткани', 'Футер трехнитка премиум пенье с начесом Penie 330гр'),
	(35, 5, 'Качество ткани', 'Premium | Penie'),
	(36, 5, 'Внутренняя сторона', 'Ворсованная / С начесом'),
	(37, 5, 'Плетение ткани', '3х-ниточное'),
	(38, 5, 'Плотность (гр/м.кв)', '320'),
	(39, 5, 'Состав', '70% хлопок, 30% полиэстер'),
	(40, 5, 'Рекомендации по уходу', 'Стирка 40 градусов. Отжим до 800 оборотов. Сушка на вешало'),
	(41, 6, 'Тип ткани', 'Футер трехнитка премиум пенье с начесом Penie 330гр'),
	(43, 6, 'Качество ткани', 'Premium | Penie'),
	(44, 6, 'Внутренняя сторона', 'Ворсованная / С начесом'),
	(45, 6, 'Плетение ткани', '3х-ниточное'),
	(46, 6, 'Плотность (гр/м.кв)', '320'),
	(47, 6, 'Состав', '70% хлопок, 30% полиэстер'),
	(48, 6, 'Рекомендации по уходу', 'Стирка 40 градусов. Отжим до 800 оборотов. Сушка на вешало'),
	(49, 7, 'Тип ткани', 'Футер трехнитка премиум пенье с начесом Penie 330гр'),
	(51, 7, 'Качество ткани', 'Premium | Penie'),
	(52, 7, 'Внутренняя сторона', 'Ворсованная / С начесом'),
	(53, 7, 'Плетение ткани', '3х-ниточное'),
	(54, 7, 'Плотность (гр/м.кв)', '320'),
	(55, 7, 'Состав', '70% хлопок, 30% полиэстер'),
	(56, 7, 'Рекомендации по уходу', 'Стирка 40 градусов. Отжим до 800 оборотов. Сушка на вешало'),
	(50, 7, 'Цвет', 'Синий'),
	(42, 6, 'Цвет', 'Темно-синий'),
	(26, 4, 'Цвет', 'Черный'),
	(34, 5, 'Цвет', 'Голубой'),
	(18, 3, 'Цвет', 'Белый'),
	(10, 2, 'Цвет', 'Красный'),
	(2, 1, 'Цвет', 'Черный');
/*!40000 ALTER TABLE "properties" ENABLE KEYS */;

-- Dumping structure for table public.types
CREATE TABLE IF NOT EXISTS "types" (
	"id" TEXT NOT NULL,
	"item" BIGINT NOT NULL,
	"name" TEXT NULL DEFAULT NULL,
	"price" INTEGER NOT NULL,
	"available" INTEGER NOT NULL DEFAULT '0',
	PRIMARY KEY ("id", "item")
);

-- Dumping data for table public.types: -1 rows
/*!40000 ALTER TABLE "types" DISABLE KEYS */;
INSERT INTO "types" ("id", "item", "name", "price", "available") VALUES
	('xs', 1, 'XS', 1000, 10),
	('s', 1, 'S', 1500, 11),
	('m', 1, 'M', 2000, 12),
	('l', 1, 'L', 2500, 13),
	('xl', 1, 'XL', 3000, 14),
	('xs', 2, 'XS', 1100, 9),
	('s', 2, 'S', 1600, 8),
	('m', 2, 'M', 2100, 7),
	('l', 2, 'L', 2600, 6),
	('xl', 2, 'XL', 3100, 5),
	('xs', 3, 'XS', 1200, 8),
	('s', 3, 'S', 1700, 7),
	('m', 3, 'M', 2200, 6),
	('l', 3, 'L', 2700, 5),
	('xl', 3, 'XL', 3200, 4),
	('xs', 4, 'XS', 1300, 7),
	('s', 4, 'S', 1800, 6),
	('m', 4, 'M', 2300, 5),
	('l', 4, 'L', 2800, 4),
	('xl', 4, 'XL', 3300, 3),
	('xs', 5, 'XS', 1400, 6),
	('s', 5, 'S', 1900, 5),
	('m', 5, 'M', 2400, 4),
	('l', 5, 'L', 2900, 3),
	('xl', 5, 'XL', 3400, 2),
	('s', 6, 'S', 2000, 4),
	('l', 6, 'L', 3000, 2),
	('xl', 7, 'XL', 3600, 0),
	('xs', 7, 'XS', 1600, 0),
	('s', 7, 'S', 2100, 0),
	('m', 6, 'M', 2500, 0),
	('m', 7, 'M', 2600, 0),
	('xl', 6, 'XL', 3500, 0),
	('xs', 6, 'XS', 1500, 0),
	('l', 7, 'L', 3100, 0);
/*!40000 ALTER TABLE "types" ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
