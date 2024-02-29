onboard-flight-data-aggregator
==============================

Jest to projekt realizowany w ramach przedmiotu "Projektowanie systemów obiektowych i rozproszonych" na Politechnice Łódzkiej przez Piotra Krzemińskiego.

Założenia projektu
------------------

Celem projektu jest stworzenie systemu agregującego dane z systemu wbudowanego komputera pokładowego. System ma za zadanie zbierać dane z takiego komputera, przetwarzać je i zapisywać w bazie danych. Dane mają być wykorzystywane do analizy lotów w danej kampani testowej, a także do wyświetlania na interfejsie użytkownika w dowolnie przygotowanym przez niego panelu. Do projektu dołączony jest szkielet aplikacji webowej, który może być wykorzystany do stworzenia pełnego serwisu webowego.

Co zostało zrealizowane
-----------------------

W ramach projektu został stworzony system agregujący dane z systemu wbudowanego komputera pokładowego. System składa się z 6 kontenerów Dockerowych:

- `Backend` - serwer aplikacyjny napisany w języku Go, który przetwarza interakcje użytkownia z serisem webowym. Aktualnie realizuje rejestrowanie użytkowników, logowanie oraz walidację tokenów JWT. Realizuje zapytania do lokalnej bazy danych sqlite3. Aplikacja serwera budowana i kompilowana jest na kontenerze `golang:latest`, a uruchamiana na `alpine:latest`. Wykorzystanie kontenerów pozwala na uniezależnienie od systemu operacyjnego, na którym jest uruchamiany serwer aplikacyjny oraz pozwala na ograniczenie wielkości końcowego kontenera aplikacji.
- `InfluxDB` - baza danych czasowych, która przechowuje dane z systemu wbudowanego komputera pokładowego. Baza danych jest dostępna pod adresem `http://localhost:8086` z domyślnymi danymi logowania `admin:adminpass`. Baza danych obsługuje ładownanie do niej danych z plików tekstowych, które są generowane przez system wbudowany komputera pokładowego.
- `Grafana` - kontener zajmujący się wizualizacją danych z bazy danych InfluxDB. Grafana jest dostępna pod adresem `http://localhost:3000` z domyślnymi danymi logowania `admin:admin`.
- `Telegraf` - kontener zajmujący się zbieraniem danych z hosta, na którym uruchomiona jest cała aplikacja. Telegraf może być skonfigurowany na zbieranie metryk bezpośrednio . Dane są zbierane co 5 sekund.
- `Kapacitor` - kontener zajmujący się przetwarzaniem danych z bazy danych InfluxDB. Kapacitor jest obecnie skonfigurowany do przetwarzania użycia procesora i wysyłania alertów, jeśli przekroczony został średni próg 80% w ciągu 1 minuty.

Wykorzystane technologie
------------------------

- Docker
- Docker Compose
- Go
- InfluxDB
- Grafana
- Telegraf
- Kapacitor
- JWT
- Alpine Linux
  
<img src="https://www.docker.com/wp-content/uploads/2023/05/symbol_blue-docker-logo.png" height="50">  
<img src="https://github.com/docker/compose/blob/main/logo.png?raw=true" height="50">  
<img src="https://hub.docker.com/api/media/repos_logo/v1/library%2Fgolang" height="50">  
<img src="https://influxdata.github.io/branding/img/downloads/influxdata-logo--full--castle.svg" height="50">  
<img src="https://miro.medium.com/v2/resize:fit:720/format:webp/1*4M4OghuybPhjRsLxhrNsGA.png" height="50">  
<img src="https://raw.githubusercontent.com/influxdata/telegraf/master/assets/TelegrafTiger.png" height="50">  
<img src="https://influxdata.github.io/branding/img/downloads/influxdata-logo--full--castle.svg" height="50">  
<img src="https://jwt.io/img/pic_logo.svg" height="50">  
<img src="https://www.alpinelinux.org/alpinelinux-logo.svg" height="50">

Diagram przepływu danych
------------------------

![Diagram przepływu danych](./docs/dfd.png)
