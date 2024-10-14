CREATE TABLE `companies`
(
    `id`   INT PRIMARY KEY,
    `name` VARCHAR(255)
);

CREATE TABLE `offices`
(
    `id`         INT PRIMARY KEY,
    `name`       VARCHAR(255),
    `company_id` INT,
    CONSTRAINT fk_offices_company_id
        FOREIGN KEY (`company_id`)
            REFERENCES companies (`id`)
            ON DELETE RESTRICT
            ON UPDATE CASCADE
);