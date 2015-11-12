-- phpMyAdmin SQL Dump
-- version 4.4.2
-- http://www.phpmyadmin.net
--
-- Host: mariadb
-- Generation Time: Nov 12, 2015 at 10:17 PM
-- Server version: 5.6.26
-- PHP Version: 5.6.7-1~dotdeb.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `stalk`
--

-- --------------------------------------------------------

--
-- Table structure for table `dailyLog`
--

DROP TABLE IF EXISTS `dailyLog`;
CREATE TABLE IF NOT EXISTS `dailyLog` (
  `logID` int(11) NOT NULL,
  `userID` varchar(80) NOT NULL,
  `date` date NOT NULL,
  `value` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `hallOfFame`
--

DROP TABLE IF EXISTS `hallOfFame`;
CREATE TABLE IF NOT EXISTS `hallOfFame` (
  `id` int(11) NOT NULL,
  `userID` varchar(80) NOT NULL,
  `username` varchar(80) NOT NULL,
  `firstName` text NOT NULL,
  `lastName` text NOT NULL,
  `gain` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `portfolio`
--

DROP TABLE IF EXISTS `portfolio`;
CREATE TABLE IF NOT EXISTS `portfolio` (
  `id` int(11) NOT NULL,
  `userID` varchar(80) NOT NULL,
  `ticker` varchar(10) NOT NULL,
  `quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `userID` varchar(80) NOT NULL,
  `username` varchar(80) NOT NULL,
  `firstName` text NOT NULL,
  `lastName` text NOT NULL,
  `turnips` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `dailyLog`
--
ALTER TABLE `dailyLog`
  ADD PRIMARY KEY (`logID`),
  ADD KEY `log_userID` (`userID`);

--
-- Indexes for table `hallOfFame`
--
ALTER TABLE `hallOfFame`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `portfolio`
--
ALTER TABLE `portfolio`
  ADD PRIMARY KEY (`id`),
  ADD KEY `portfolio_userID` (`userID`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`userID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `dailyLog`
--
ALTER TABLE `dailyLog`
  MODIFY `logID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `hallOfFame`
--
ALTER TABLE `hallOfFame`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `portfolio`
--
ALTER TABLE `portfolio`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `dailyLog`
--
ALTER TABLE `dailyLog`
  ADD CONSTRAINT `dailyLog_userID_key` FOREIGN KEY (`userID`) REFERENCES `user` (`userID`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `portfolio`
--
ALTER TABLE `portfolio`
  ADD CONSTRAINT `portfolio_userID` FOREIGN KEY (`userID`) REFERENCES `user` (`userID`) ON DELETE CASCADE ON UPDATE CASCADE;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
