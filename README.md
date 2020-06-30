# Utilizando o docker para desenvolvimento
docker-compose up backend-dev

# Utilizando o docker para produção
docker-compose up backend-prod

# Comentários
O ambiente de desenvolvimento esta utilizando o live-reload CompileDaemon que funciona tanto no Windows quanto Linux

O ambiente de produção depente do banco de dados postgres para funcionar e sempre que for gerado uma nova versão o comando acima deverá ser executado.

# Próximos desafios
1 - Utilizar um ORM (gorm)

2 - Autenticar com token JWT e criar um sistema de autorização RBAC

3 - Implementar o backend utilizando o protocolo gRPC

4 - Construir um frontend que consuma serviços via gRPC
