-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 25 Mar 2024 pada 20.39
-- Versi server: 5.7.42-0ubuntu0.18.04.1-log
-- Versi PHP: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `items`
--

CREATE TABLE `items` (
  `id_item` int(11) NOT NULL,
  `id_order` int(11) NOT NULL,
  `id_product` int(11) NOT NULL,
  `harga` double NOT NULL,
  `jumlah` int(11) NOT NULL,
  `sub_total` double NOT NULL,
  `created_by` int(11) NOT NULL,
  `date_created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` int(11) NOT NULL,
  `date_updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `items`
--

INSERT INTO `items` (`id_item`, `id_order`, `id_product`, `harga`, `jumlah`, `sub_total`, `created_by`, `date_created`, `updated_by`, `date_updated`) VALUES
(1, 1, 1, 1000, 5, 5000, 2, '2024-03-25 12:45:13', 2, '2024-03-25 12:45:13'),
(2, 1, 2, 1000, 6, 6000, 2, '2024-03-25 12:45:14', 2, '2024-03-25 12:45:14'),
(5, 1, 1, 1000, 10, 10000, 2, '2024-03-25 13:32:13', 2, '2024-03-25 13:32:13'),
(6, 1, 2, 1000, 20, 20000, 2, '2024-03-25 13:32:13', 2, '2024-03-25 13:32:13');

-- --------------------------------------------------------

--
-- Struktur dari tabel `jwts`
--

CREATE TABLE `jwts` (
  `ids_jwt` int(11) NOT NULL,
  `headers` mediumtext,
  `ip_address` varchar(255) NOT NULL,
  `token` text NOT NULL,
  `expire_at` datetime NOT NULL,
  `expired` enum('YA','TIDAK') NOT NULL,
  `keterangan` enum('LOGIN','LOGOUT','EXPIRED BY SYSTEM') NOT NULL,
  `date_created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `date_updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `jwts`
--

INSERT INTO `jwts` (`ids_jwt`, `headers`, `ip_address`, `token`, `expire_at`, `expired`, `keterangan`, `date_created`, `date_updated`) VALUES
(1, 'PostmanRuntime/7.37.0', '::1', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlrIjoxNzExMjk3NzU5LCJpZCI6MSwibmFtYSI6IkFsZmkgR3VzbWFuIiwidXNlcm5hbWUiOiJhbGZpLmd1c21hbiIsInJvbGVzIjoiU1VQRVIgQURNSU4iLCJleHAiOjE3MTEzODQxNTksImlhdCI6MTcxMTI5Nzc1OX0.OF5-yqqttezsx9IvVmm-EoKMPfVDMhDsribZ7bBy3n4', '2024-03-25 23:29:19', 'TIDAK', 'LOGIN', '2024-03-24 16:29:19', '2024-03-24 16:29:19'),
(2, 'PostmanRuntime/7.37.0', '::1', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlrIjoxNzExMjk3ODYxLCJpZCI6MiwibmFtYSI6IlRlc3QgQWRtaW4gMDEiLCJ1c2VybmFtZSI6ImFrdW4wMDEiLCJyb2xlcyI6IkFETUlOIiwiZXhwIjoxNzExMzg0MjYxLCJpYXQiOjE3MTEyOTc4NjF9.dDkZKnvfJygAW_m654YtEE3e7wMU-SfGAFtgr3cbm1Q', '2024-03-25 23:31:01', 'TIDAK', 'LOGIN', '2024-03-24 16:31:01', '2024-03-24 16:31:01'),
(3, 'PostmanRuntime/7.37.0', '::1', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlrIjoxNzExMjk3OTAwLCJpZCI6MywibmFtYSI6IlRlc3QgQWRtaW4gMDIiLCJ1c2VybmFtZSI6ImFrdW4wMDIiLCJyb2xlcyI6IkFETUlOIiwiZXhwIjoxNzExMzg0MzAwLCJpYXQiOjE3MTEyOTc5MDB9.1N7H7zdxe67WjdvY5thlKpHoYYyxDyul2sM-ujl0e5w', '2024-03-25 23:31:40', 'TIDAK', 'LOGIN', '2024-03-24 16:31:40', '2024-03-24 16:31:40');

-- --------------------------------------------------------

--
-- Struktur dari tabel `orders`
--

CREATE TABLE `orders` (
  `id_order` int(11) NOT NULL,
  `customer_name` varchar(255) NOT NULL,
  `tmp_total` double NOT NULL,
  `potongan` double NOT NULL,
  `total` double NOT NULL,
  `created_by` int(11) NOT NULL,
  `date_created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` int(11) NOT NULL,
  `date_updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `orders`
--

INSERT INTO `orders` (`id_order`, `customer_name`, `tmp_total`, `potongan`, `total`, `created_by`, `date_created`, `updated_by`, `date_updated`) VALUES
(1, 'Test 2 (Update)', 30000, 3000, 27000, 2, '2024-03-25 11:59:41', 2, '2024-03-25 13:32:14'),
(2, 'Test 2', 15000, 1000, 14000, 2, '2024-03-25 13:29:54', 2, '2024-03-25 13:29:54');

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `id_product` int(11) NOT NULL,
  `product` varchar(255) NOT NULL,
  `harga_beli` double NOT NULL,
  `harga_jual` double NOT NULL,
  `stok` int(11) NOT NULL,
  `created_by` int(11) NOT NULL,
  `date_created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` int(11) NOT NULL,
  `date_updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `products`
--

INSERT INTO `products` (`id_product`, `product`, `harga_beli`, `harga_jual`, `stok`, `created_by`, `date_created`, `updated_by`, `date_updated`) VALUES
(1, 'Teh Gelas', 800, 1000, 100, 2, '2024-03-25 01:49:51', 2, '2024-03-25 01:49:51'),
(2, 'Panter', 800, 1000, 50, 2, '2024-03-25 01:53:52', 2, '2024-03-25 01:53:52');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id_user` int(11) NOT NULL,
  `nama` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `roles` enum('SUPER ADMIN','ADMIN') NOT NULL,
  `date_created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `date_updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id_user`, `nama`, `username`, `password`, `roles`, `date_created`, `date_updated`) VALUES
(1, 'Alfi Gusman', 'alfi.gusman', '$2a$06$4QSB.34BYdj1kjcGO05ulelSyB5PQEzGBqyn294sSGSV/ODqllalO', 'SUPER ADMIN', '2024-03-20 13:43:52', '2024-03-20 13:54:04'),
(2, 'Test Admin 01', 'akun001', '$2a$06$4QSB.34BYdj1kjcGO05ulelSyB5PQEzGBqyn294sSGSV/ODqllalO', 'ADMIN', '2024-03-20 15:59:35', '2024-03-25 13:22:59'),
(3, 'Test Admin 02', 'akun002', '$2a$06$/rVwYDsmmChEjuJ.UfWMSuZxT47AeQmYFDDsF2TRu9/Gs5CzzCKeu', 'ADMIN', '2024-03-24 15:10:58', '2024-03-24 15:10:58'),
(4, 'Test Admin 03', 'akun003', '$2a$06$Vz80Xlynt51RaQYdeuzZ0uODOe4eKGtYVVrwBSDuZib1/QXxmx30y', 'ADMIN', '2024-03-24 16:23:21', '2024-03-24 16:23:21'),
(5, 'Test Admin 13', 'akun013', '$2a$06$hz6Cx1lFK.fzvI31oZjlrOwjnWLznY4MpdH/uulYOaOTYrpxqT.Ea', 'ADMIN', '2024-03-25 13:22:05', '2024-03-25 13:22:05');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`id_item`),
  ADD KEY `craeted_by` (`created_by`),
  ADD KEY `updated_by` (`updated_by`),
  ADD KEY `id_item` (`id_item`),
  ADD KEY `id_order` (`id_order`),
  ADD KEY `id_product` (`id_product`);

--
-- Indeks untuk tabel `jwts`
--
ALTER TABLE `jwts`
  ADD PRIMARY KEY (`ids_jwt`),
  ADD KEY `ids_jwt` (`ids_jwt`);

--
-- Indeks untuk tabel `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id_order`),
  ADD KEY `created_by` (`created_by`),
  ADD KEY `updated_by` (`updated_by`),
  ADD KEY `id_order` (`id_order`);

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id_product`),
  ADD KEY `created_by` (`created_by`),
  ADD KEY `updated_by` (`updated_by`),
  ADD KEY `id_product` (`id_product`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`),
  ADD KEY `username` (`username`),
  ADD KEY `id_user` (`id_user`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `items`
--
ALTER TABLE `items`
  MODIFY `id_item` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `jwts`
--
ALTER TABLE `jwts`
  MODIFY `ids_jwt` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `orders`
--
ALTER TABLE `orders`
  MODIFY `id_order` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `products`
--
ALTER TABLE `products`
  MODIFY `id_product` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id_user` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `items`
--
ALTER TABLE `items`
  ADD CONSTRAINT `items_ibfk_1` FOREIGN KEY (`id_order`) REFERENCES `orders` (`id_order`) ON UPDATE CASCADE,
  ADD CONSTRAINT `items_ibfk_2` FOREIGN KEY (`id_product`) REFERENCES `products` (`id_product`) ON UPDATE CASCADE,
  ADD CONSTRAINT `items_ibfk_3` FOREIGN KEY (`created_by`) REFERENCES `users` (`id_user`) ON UPDATE CASCADE,
  ADD CONSTRAINT `items_ibfk_4` FOREIGN KEY (`updated_by`) REFERENCES `users` (`id_user`) ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `users` (`id_user`) ON UPDATE CASCADE,
  ADD CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`updated_by`) REFERENCES `users` (`id_user`) ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `users` (`id_user`) ON UPDATE CASCADE,
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`updated_by`) REFERENCES `users` (`id_user`) ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
