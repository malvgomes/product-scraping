DROP DATABASE IF EXISTS scraper;

CREATE DATABASE scraper character set utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE scraper.products (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    description TEXT NOT NULL,
    url VARCHAR(255) NOT NULL,
    date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY url (url) USING BTREE,
    UNIQUE KEY unique_url (url) USING BTREE,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO scraper.products (title, image_url, price, description, url) VALUES
    (
        'Echo Show 8 (1ª Geração): Smart Speaker com tela de 8\" e Alexa - Cor Preta',
        'https://m.media-amazon.com/images/I/61xX62L2hGL._AC_SY300_SX300_.jpg',
        66405,
        '',
        'https://www.amazon.com.br/dp/B07SG8F1QF'
    ),
    (
        'A Rainha Vermelha - Vol. 1',
        'https://lojasaraiva.vteximg.com.br/arquivos/ids/12054910-287-426/1008976886.jpg?v=637141926959370000',
        3390,
        'O mundo de Mare Barrow é dividido pelo sangue: vermelho ou prateado. Mare e sua família são vermelhos: plebeus, humildes, destinados a servir uma elite prateada cujos poderes sobrenaturais os tornam quase deuses.Mare rouba o que pode para ajudar sua família a sobreviver e não tem esperanças de escapar do vilarejo miserável onde mora. Entretanto, numa reviravolta do destino, ela consegue um emprego no palácio real, onde, em frente ao rei e a toda a nobreza, descobre que tem um poder misterioso… Mas como isso seria possível, se seu sangue é vermelho?Em meio às intrigas dos nobres prateados, as ações da garota vão desencadear uma dança violenta e fatal, que colocará príncipe contra príncipe — e Mare contra seu próprio coração.',
        'https://www.saraiva.com.br/a-rainha-vermelha-vol-1-8884222/p'
    )
