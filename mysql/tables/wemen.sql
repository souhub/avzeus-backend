-- Create wemen table
CREATE TABLE IF NOT EXISTS wemen
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50),
    image_path VARCHAR(255),
    PRIMARY KEY(id)
);

-- Insert 9 data
INSERT INTO wemen
(
    name,
    image_path
)
VALUES
(
    '織田真子',
    'https://storage.googleapis.com/avzeus/wemen/%E7%B9%94%E7%94%B0%E7%9C%9F%E5%AD%90.jpg'
);

INSERT INTO wemen
(
    name,
    image_path
)
VALUES
(
    '桃乃木かな',
    'https://storage.googleapis.com/avzeus/wemen/%E6%A1%83%E4%B9%83%E6%9C%A8%E3%81%8B%E3%81%AA.jpg'
);

INSERT INTO wemen
(
    name,
    image_path
)
VALUES
(
    '丘えりな',
    'https://storage.googleapis.com/avzeus/wemen/%E4%B8%98%E3%81%88%E3%82%8A%E3%81%AA.jpg'
);

INSERT INTO wemen
(
    name,
    image_path
)
VALUES
(
    '希島あいり',
    'https://storage.googleapis.com/avzeus/wemen/%E5%B8%8C%E5%B3%B6%E3%81%82%E3%81%84%E3%82%8A.jpg'
);

INSERT INTO wemen
(
    name,
    image_path
)
VALUES
(
    '加美杏奈',
    'https://storage.googleapis.com/avzeus/wemen/%E5%8A%A0%E7%BE%8E%E6%9D%8F%E5%A5%88.jpg'
);

INSERT INTO wemen
(
    name,
    image_path
)
VALUES
(
    '美谷朱里',
    'https://storage.googleapis.com/avzeus/wemen/%E7%BE%8E%E8%B0%B7%E6%9C%B1%E9%87%8C.jpg'
);

INSERT INTO wemen
(
    name,
    image_path
)
VALUES
(
    '小野六花',
    'https://storage.googleapis.com/avzeus/wemen/%E5%B0%8F%E9%87%8E%E5%85%AD%E8%8A%B1.jpg'
);

INSERT INTO wemen
(
    name,
    image_path
)
VALUES
(
    '神木サラ',
    'https://storage.googleapis.com/avzeus/wemen/%E7%A5%9E%E6%9C%A8%E3%82%B5%E3%83%A9.jpg'
);

INSERT INTO wemen
(
    name,
    image_path
)
VALUES
(
    '篠田ゆう',
    'https://storage.googleapis.com/avzeus/wemen/%E7%AF%A0%E7%94%B0%E3%82%86%E3%81%86.jpg'
);
