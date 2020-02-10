# SANNTIDSPROGRAMMERING

## Prioritet

>Hver enkelt node har et nummer 1,2,...,n. Master er nr 1, mens Backup er nr 2. Resterende heiser har nr 3 ,...,n.

## Kommunikasjon
>Alle heisene snakker sammen over TCP, med master som koordinerende aktør. All informasjon samles inn til master, sendes til backup og sendes deretter ut til alle noder. Ved en handling informeres master og backup FØR noen handling utføres.
