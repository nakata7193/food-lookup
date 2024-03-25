CREATE TABLE IF NOT EXISTS foods (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description VARCHAR(255) NOT NULL,
    carbs DECIMAL(10,2) NOT NULL,
    protein DECIMAL(10,2) NOT NULL,
    fat DECIMAL(10,2) NOT NULL
);

DROP TABLE foods;

INSERT INTO foods (description, carbs, protein, fat) VALUES 
('fish', 0.0, 20.0, 13.0);

SELECT * FROM foods WHERE id = 2;

SELECT * FROM foods;