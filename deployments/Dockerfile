FROM httpd:bookworm

RUN apt update; apt upgrade -y;
RUN apt install golang -y

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /usr/local/apache2/cgi-bin/app.cgi apache2-cgi-template/cmd/app

RUN chmod 755 /usr/local/apache2/cgi-bin/app.cgi
RUN /usr/local/apache2/cgi-bin/app.cgi
COPY configs/httpd.conf /usr/local/apache2/conf/httpd.conf
EXPOSE 80
