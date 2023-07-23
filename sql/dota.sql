-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Июл 23 2023 г., 15:29
-- Версия сервера: 5.7.39-log
-- Версия PHP: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `dota`
--

-- --------------------------------------------------------

--
-- Структура таблицы `Posted`
--

CREATE TABLE `Posted` (
  `Postedid` int(11) NOT NULL,
  `Username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `Title` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `Content` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Дамп данных таблицы `Posted`
--

INSERT INTO `Posted` (`Postedid`, `Username`, `Title`, `Content`) VALUES
(15, 'deadinside', 'ghoul 1000-7', 'tru nickname for true ghoul 1000-7'),
(16, 'ghoul', 'SF-Coil-Coil-Coil-Z-X-C', 'NickName for SF pickers');

-- --------------------------------------------------------

--
-- Структура таблицы `Users`
--

CREATE TABLE `Users` (
  `Username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `Password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Дамп данных таблицы `Users`
--

INSERT INTO `Users` (`Username`, `Password`) VALUES
('deadinside', 'deadinside'),
('ghoul', '10007');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `Posted`
--
ALTER TABLE `Posted`
  ADD PRIMARY KEY (`Postedid`),
  ADD KEY `Username` (`Username`);

--
-- Индексы таблицы `Users`
--
ALTER TABLE `Users`
  ADD PRIMARY KEY (`Username`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `Posted`
--
ALTER TABLE `Posted`
  MODIFY `Postedid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `Posted`
--
ALTER TABLE `Posted`
  ADD CONSTRAINT `posted_ibfk_1` FOREIGN KEY (`Username`) REFERENCES `Users` (`Username`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
