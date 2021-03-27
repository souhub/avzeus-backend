-- Create actresses table
CREATE TABLE IF NOT EXISTS actresses
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    image_path VARCHAR(255),
    vector VARCHAR(1000)
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
    'https://an-image-path-to-actress1',
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
    'https://an-image-path-to-actress2',
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
    'https://an-image-path-to-actress3',
    '[6,33,22,67,887,22,4,7,89,8]'
);
