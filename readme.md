# Projekt zaliczeniowy - Sieci Komputerowe 2
## Jan Fedoruk, 143780

Aplikacja została napisana w języku go. W celu uruchomienia należy zbudować appkę poleceniem "go build ." i uruchomić powstałą binarkę. Aplikacja nasłuchuje na porcie 8080. Wymagana wersja go: >= 1.20. Kompilator golang należy zainstalować zgodnie z poleceniami: https://go.dev/doc/install.

Aplikacja realizuje serwer HTTP (zgodnie z HTTP/1.1 per RFC2616), który obsługuje żądania dotyczące zarządzania plikami; POST - utwórz, PUT - aktualizuj zawartość, GET - pobierz zawartość, HEAD - nagłówki, DELETE - usuń plik.
