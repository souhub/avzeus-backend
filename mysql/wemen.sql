-- Create wemen table
CREATE TABLE IF NOT EXISTS wemen
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    image_path VARCHAR(255),
    vector VARCHAR(1000)
);

-- Insert some data
INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus-dev/wemen/woman1.jpeg',
    '[1,45,754,467,426,4456,467,445]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus-dev/wemen/woman2.jpeg',
    '[34,74,5,68,32,25,67,9,0,63,222,68,2]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus-dev/wemen/woman3.jpeg',
    '[5,7,3,7,8,99,3,2,5,8,99,442,6,7]'
);
