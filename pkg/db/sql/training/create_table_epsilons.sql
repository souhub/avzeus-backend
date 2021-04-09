CREATE TABLE IF NOT EXISTS epsilons
(
    id INT NOT NULL AUTO_INCREMENT,
    val DOUBLE,
    training_id INT,
    PRIMARY KEY(id)
);
