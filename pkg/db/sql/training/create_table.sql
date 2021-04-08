CREATE TABLE IF NOT EXISTS training
(
    id INT NOT NULL AUTO_INCREMENT,
    states VARCHAR(255),
    epsilons VARCHAR(1000),
    selected_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
)
