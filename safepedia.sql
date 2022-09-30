-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Waktu pembuatan: 30 Sep 2022 pada 08.35
-- Versi server: 5.7.34
-- Versi PHP: 8.0.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `safepedia`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `karyawan`
--

CREATE TABLE `karyawan` (
  `id` int(11) NOT NULL,
  `Nama` varchar(200) NOT NULL,
  `TanggalLahir` date NOT NULL,
  `JumlahAnak` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `karyawan`
--

INSERT INTO `karyawan` (`id`, `Nama`, `TanggalLahir`, `JumlahAnak`) VALUES
(1, 'Anang', '1992-09-16', 2),
(2, 'Dani', '1998-04-07', 1),
(3, 'Jono', '1959-09-06', 3),
(4, 'Ani', '1995-09-06', 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `KendaraanPerusahaan`
--

CREATE TABLE `KendaraanPerusahaan` (
  `id` int(11) NOT NULL,
  `Tipe` varchar(50) NOT NULL,
  `Warna` varchar(50) NOT NULL,
  `TanggalBeli` datetime NOT NULL,
  `JumlahRoda` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `KendaraanPerusahaan`
--

INSERT INTO `KendaraanPerusahaan` (`id`, `Tipe`, `Warna`, `TanggalBeli`, `JumlahRoda`) VALUES
(1, 'Fortuner', 'hitam', '2022-09-29 14:01:02', 4),
(2, 'Agya', 'Merah', '2022-09-08 14:01:02', 4),
(3, 'Avanza', 'Hitam', '2022-09-29 14:01:27', 4),
(4, 'Rush', 'Hitam', '2022-09-29 14:01:27', 4);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `karyawan`
--
ALTER TABLE `karyawan`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `KendaraanPerusahaan`
--
ALTER TABLE `KendaraanPerusahaan`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `karyawan`
--
ALTER TABLE `karyawan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `KendaraanPerusahaan`
--
ALTER TABLE `KendaraanPerusahaan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
