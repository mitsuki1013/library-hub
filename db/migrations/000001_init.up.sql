CREATE TABLE `companies`
(
    `id`   CHAR(26) PRIMARY KEY,
    `name` VARCHAR(255)
);

CREATE TABLE `offices`
(
    `id`         CHAR(26) PRIMARY KEY,
    `name`       VARCHAR(255),
    `company_id` CHAR(26),
    CONSTRAINT fk_offices_company_id
        FOREIGN KEY (`company_id`)
            REFERENCES `companies` (`id`)
            ON DELETE RESTRICT
            ON UPDATE CASCADE
);

CREATE TABLE `library`
(
    `id`        CHAR(26) PRIMARY KEY,
    `name`      VARCHAR(255),
    `office_id` CHAR(26),
    `isbn`      CHAR(21),
    CONSTRAINT fk_library_office_id
        FOREIGN KEY (`office_id`)
            REFERENCES `offices` (`id`)
            ON DELETE CASCADE
            ON UPDATE CASCADE,
    CONSTRAINT `uk_library_office_id_isbn`
        UNIQUE (`office_id`, `isbn`)
);