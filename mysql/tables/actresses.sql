-- Create actresses table
CREATE TABLE IF NOT EXISTS actresses
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    image_path VARCHAR(255),
    vector VARCHAR(1000),
    PRIMARY KEY(id)
);

-- Insert some data
INSERT INTO actresses
(
    name,
    image_path,
    vector
)
VALUES
(
    '明日花キララ',
    'https://storage.googleapis.com/avzeus-dev/actresses/%E6%98%8E%E6%97%A5%E8%8A%B1%E3%82%AD%E3%83%A9%E3%83%A9.jpeg',
    '[3,6,77,83,2,44,78,32,1,678,21,5]'
);

INSERT INTO actresses
(
    name,
    image_path,
    vector
)
VALUES
(
    '深田えいみ',
    'https://storage.googleapis.com/avzeus-dev/actresses/%E6%B7%B1%E7%94%B0%E3%81%88%E3%81%84%E3%81%BF.jpeg',
    '[3,6,77,83,2,44,78,32,1,678,21,5]'
);

INSERT INTO actresses
(
    name,
    image_path,
    vector
)
VALUES
(
    '橋本ありな',
    'https://storage.googleapis.com/avzeus-dev/actresses/%E6%A9%8B%E6%9C%AC%E3%81%82%E3%82%8A%E3%81%AA.jpeg',
    '[6,33,22,67,887,22,4,7,89,8]'
);
