-- Create wemen table
CREATE TABLE IF NOT EXISTS wemen
(
    id INT NOT NULL AUTO_INCREMENT,
    image_path VARCHAR(255),
    vector VARCHAR(1000),
    PRIMARY KEY(id)
);

-- Insert 9 data
INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus/wemen/%E7%B9%94%E7%94%B0%E7%9C%9F%E5%AD%90.jpg',
    '[1,45,754,467,426,4456,467,445]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus/wemen/%E6%A1%83%E4%B9%83%E6%9C%A8%E3%81%8B%E3%81%AA.jpg',
    '[34,74,5,68,32,25,67,9,0,63,222,68,2]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus/wemen/%E4%B8%98%E3%81%88%E3%82%8A%E3%81%AA.jpg',
    '[5,7,3,7,8,99,3,2,5,8,99,442,6,7]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus/wemen/%E5%B8%8C%E5%B3%B6%E3%81%82%E3%81%84%E3%82%8A.jpg',
    '[5,7,3,7,8,99,3,2,5,8,99,442,6,7]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus/wemen/%E5%8A%A0%E7%BE%8E%E6%9D%8F%E5%A5%88.jpg',
    '[5,7,3,7,8,99,3,2,5,8,99,442,6,7]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus/wemen/%E7%BE%8E%E8%B0%B7%E6%9C%B1%E9%87%8C.jpg',
    '[5,7,3,7,8,99,3,2,5,8,99,442,6,7]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus/wemen/%E5%B0%8F%E9%87%8E%E5%85%AD%E8%8A%B1.jpg',
    '[5,7,3,7,8,99,3,2,5,8,99,442,6,7]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus/wemen/%E7%A5%9E%E6%9C%A8%E3%82%B5%E3%83%A9.jpg',
    '[5,7,3,7,8,99,3,2,5,8,99,442,6,7]'
);

INSERT INTO wemen
(
    image_path,
    vector
)
VALUES
(
    'https://storage.googleapis.com/avzeus/wemen/%E7%AF%A0%E7%94%B0%E3%82%86%E3%81%86.jpg',
    '[5,7,3,7,8,99,3,2,5,8,99,442,6,7]'
);
