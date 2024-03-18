from cassandra.cluster import Cluster
from faker import Faker
from uuid import uuid4
from datetime import datetime
import os
import json  # Para serializar custom_attributes


# Acessar variáveis de ambiente
cassandra_host = os.getenv("CASSANDRA_HOST")
cassandra_port = os.getenv("CASSANDRA_PORT")

# Formar string de conexão
connection_string = f"cassandra://{cassandra_host}:{cassandra_port}"

# Conectar ao CassandraDB
cluster = Cluster(['cassandra'])
session = cluster.connect('product_keyspace')

if __name__ == "__main__":
    faker = Faker()

    for _ in range(500):
        product_id = str(uuid4())
        created_at = datetime.now()
        
        # Dados do produto
        product_data = (
            product_id,
            faker.name(),  # Nome do produto
            [faker.word() for _ in range(faker.random_int(min=1, max=5))],  # Categorias
            json.dumps({'key_' + str(i): 'value_' + str(i) for i in range(faker.random_int(min=1, max=5))}),  # Custom_attributes serializado
            created_at,
            faker.boolean(),  # Is_blocked
            [faker.image_url() for _ in range(faker.random_int(min=1, max=5))]  # Images
        )

        # Inserção na tabela products
        insert_product_query = """
        INSERT INTO products (id, name, categories, custom_attributes, created_at, is_blocked, images)
        VALUES (%s, %s, %s, %s, %s, %s, %s)
        """
        session.execute(insert_product_query, product_data)
        
        # Dados de buy_options (considerando múltiplas opções por produto)
        for _ in range(faker.random_int(min=1, max=3)):  # Gerando de 1 a 3 opções de compra por produto
            buy_option_data = (
                product_id,
                str(faker.random_int(min=1, max=25)),
                faker.random_int(min=1000, max=10000),
                faker.random_int(min=0, max=50),
                faker.random_int(min=0, max=100)
            )

            # Inserção na tabela buy_options
            insert_buy_option_query = """
            INSERT INTO buy_options (product_id, seller_id, price, promotional_discount, stock)
            VALUES (%s, %s, %s, %s, %s)
            """
            session.execute(insert_buy_option_query, buy_option_data)

    # Fechando conexão
    session.shutdown()
    cluster.shutdown()
