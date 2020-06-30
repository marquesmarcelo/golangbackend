# Utilizando o docker para desenvolvimento
docker-compose up backend-dev

# Utilizando o docker para produção
docker-compose up backend-prod

# Comentários
O ambiente de desenvolvimento esta utilizando o live-reload CompileDaemon que funciona tanto no Windows quanto Linux

Deverá ser execute o comando para gerar a produção sempre que uma nova versão estiver concluída. A opção de live-reload funciona apenas no desenvolvimento.

# Próximos desafios
1 - Utilizar um ORM

2 - Gerar token JWT

3 - Ativar o protocolo gRPC

4 - Construir um frontend que consuma serviços via gRPC