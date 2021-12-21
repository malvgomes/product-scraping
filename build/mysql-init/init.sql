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
        'Batata Monalisa Aprox. 1 Kg',
        'https://carrefourbr.vtexassets.com/arquivos/ids/14648737-720-720?v=637511892564730000',
        499,
        'Toda a linha de hortifruti sempre fresquinhos, saudáveis e com qualidade garantida de origem',
        'https://mercado.carrefour.com.br/batata-monalisa-aprox--1-kg-46922/p'
    ),
    (
        'Macarrão Instantâneo Nissin Talharim Sabor Bolonhesa 99g',
        'https://carrefourbr.vtexassets.com/arquivos/ids/195542-720-720?v=637272435389600000',
        183,
        'Há dias que nem sempre é possível encontrar tempo para preparar aquela refeição saborosa para toda a família e a solução é prepara uma comida rápida. A melhor forma de fazer com sabor é com o Macarrão Instantâneo Talharim Sabor Bolonhesa 99g, da Nissin Food. Além de ficar pronto em 3 minutos de preparo, o produto contém tempero aromatizando artificialmente para dar um sabor extra à sua refeição.',
        'https://mercado.carrefour.com.br/macarrao-instantaneo-nissin-talharim-sabor-bolonhesa-99g/p'
    ),
    (
        'Arroz Branco Longo-fino Tipo 1 Tio João 2Kg',
        'https://carrefourbr.vtexassets.com/arquivos/ids/203188-720-720?v=637272452216400000',
        1519,
        '',
        'https://mercado.carrefour.com.br/arroz-branco-longo-fino-tipo-1-tio-joao-2kg-115657/p'
    ),
    (
        'Feijão Carioca Tipo 1 Camil Todo Dia 1 Kg',
        'https://carrefourbr.vtexassets.com/arquivos/ids/195175-720-720?v=637272434435130000',
        749,
        'Camil apresenta o feijão carioca, o mais popular do Brasil! Parte da história, da cultura e do dia a dia de milhares de brasileiros, o feijão carioca é riquíssimo em ferro e proteína vegetal. As proteínas auxiliam na manutenção e no ganho de músculos. E, para completar, o feijão carioca é fonte de fibras que contribuem para a saciedade e auxiliam no controle da liberação do açúcar no sangue e no funcionamento intestinal. Compre aqui no Carrefour o Feijão Carioca Camil e engrosse esse caldo com sabor e um toque de carinho e cuidado em todos os momentos ao redor da mesa.',
        'https://mercado.carrefour.com.br/feijao-carioca-tipo-1-camil-todo-dia-1-kg-871281/p'
    );
