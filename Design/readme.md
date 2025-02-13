# SANNTIDSPROGRAMMERING

## Prioritet

>Hver enkelt node har et nummer 1,2,...,n. Master er nr 1, mens Backup er nr 2. Resterende heiser har nr 3 ,...,n.

## Kommunikasjon
>Alle heisene snakker sammen over UDP, med master som koordinerende aktør. All informasjon samles inn til master, sendes til backup og sendes deretter ut til alle noder. Ved en handling informeres master og backup FØR noen handling utføres. Formatet for kommunikasjon mellom noder finnes under 'update.txt' og 'state.txt'. En heismodul mottar kun et tillegg i sin individuelle kø som må håndteres, og trenger derfor ingen voldsom utvidelse i forhold til heisen som ble laget i Tilpassede Datasystemer. Videre må heisen kunne sende melding på formatet gitt i 'update.txt', men dette kan håndteres utenfra selve heismodulen.

### Bekreftelsesmeldinger
>Ved en sendt melding fra master bekreftes meldingen før en handling gjennomføres. Master trenger derimot ingen bekreftelse for at en node skal utføre en handling (for eksempel å skru på et lys). Dette er fordi at med en gang master har fått melding om en handling fra en node vil enhver bestilling følges opp. Det kritiske objektet, enhver bestilling skal betjenes, er med dette oppfylt ved å sende en ny heis til målet dersom en viss tid har gått. Dersom en melding ikke bekreftes godtatt av master kan heller ingen bestilling gjennomføres.

## Vaktbikkje
> Fungerende master purrer gjevnlig på alle nodene. Dersom bekreftelse ikke mottas fra en node registreres denne som inaktiv. Handlingskøen fra denne noden vil da flyttes fra aktiv i den aktuelle noden til inaktiv, slik at andre noder kan ta over arbeidet. Dersom noden fortsetter arbeidet vil de aktive bestillinger kunne betjenes 2 ganger, men sett fra nettverket kun en gang.

## Toleranse
> Master sitter på enhver tid på all informasjon, og oppdaterer backup. Dersom master svikter vil backup ved hjelp av en timer registrere at den ikke får livstegn fra master. Da vil backup ta over som master, og gjevnlig forsøke å vekke master tilbake til live. Når tidligere master er på nett igjen fungerer den som backup. Alle heiser har samme programvare, slik at master kan gå fra 1 til 2 til 3 og nedover.
Sett fra nodenes perspektiv vil den eneste endringen være adressen ved et master-bytte.
> På denne måten vil en svikt hvor som helst i systemet ikke kunne ta ned hele driften fordi to noder til enhver tid sitter på en oversikt som kan gjenopptas av en backup ved eventuell svikt.

## NETWORK
> Network er en generell modul som alle noder har. Disse har alle tilgang til et dokument, som er hoveddokumentet distribuert av master. Kun master kan endre på dette dokumentet, og alle andre noder sender inn handlinger til master via sin network-node. I alle diagrammer er alle kommunikative handlinger implisitt via network.


## Threads
> Hver node antas i forkant av utviklingen å ha behov for 3 tråder. En for nettverkshåndtering, en for vaktbikkje og en for selve heishåndteringen. I flere av nodene kan vaktbikkje sløyfes, men bør være der i tilfelle overtagelse av masterrollen.


## Beregning av kostnad

> Når en heis er på vei oppover, og registrerer en bestilling på veien som går nedover er denne kostnaden uendelig frem til heisen har stoppet, eller er på vei nedover igjen. Kun bestillinger i samme retning kan vurderes. En etasje unna koster 1, 2 etasjer koster 2 osv... På denne måten vet alle heiser hvilken som er best skikket til å ta en bestilling fra inaktiv, til sin respektive aktiv-liste. Også dette rapporteres til master før handlingen utføres.
