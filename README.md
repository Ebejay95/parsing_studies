Recursive Descent Parsing - Die Denkweise
1. Grundprinzip: Grammar → Functions
Schritt 1: Definiere deine Grammatik
JSON     := Value
Value    := Object | Array | String | Number | Boolean | Null
Object   := '{' (KeyValue (',' KeyValue)*)? '}'
KeyValue := String ':' Value
Array    := '[' (Value (',' Value)*)? ']'
String   := '"' ... '"'
Number   := [0-9]+
Boolean  := 'true' | 'false'
Null     := 'null'
Schritt 2: Jede Grammatik-Regel wird zu einer Funktion

JSON → parseJSON()
Value → parseValue()
Object → parseObject()
Array → parseArray()
etc.

2. Die mentale Herangehensweise
Top-Down Denken:

"Was erwarte ich hier?" - Schaue auf das aktuelle Zeichen
"Welche Regel passt?" - Entscheide basierend auf dem ersten Zeichen
"Delegiere die Arbeit" - Rufe die passende Funktion auf
"Vertraue der Rekursion" - Die aufgerufene Funktion macht ihren Job

Beispiel-Denkprozess für {"name": "John"}:
parseValue() sieht '{'
  → "Das ist ein Object!"
  → Rufe parseObject() auf

parseObject() sieht '"name"'
  → "Das ist ein String-Key!"
  → Rufe parseString() auf
  → Erwarte ':'
  → Rufe parseValue() für den Wert auf

parseValue() sieht '"John"'
  → "Das ist ein String!"
  → Rufe parseString() auf
3. Das Recursive Descent Pattern
Template für jede Parser-Funktion:
gofunc (p *Parser) parseXXX() error {
    // 1. Eingabe validieren
    if p.current() != erwartetes_zeichen {
        return error
    }

    // 2. Zeichen konsumieren
    p.advance()

    // 3. Inhalt parsen (oft rekursiv)
    for ... {
        if err := p.parseYYY(); err != nil {
            return err
        }
    }

    // 4. Abschluss validieren
    if p.current() != abschluss_zeichen {
        return error
    }
    p.advance()

    return nil
}
4. Schritt-für-Schritt Aufbau-Strategie
Phase 1: Grundgerüst
gotype Parser struct {
    input string
    pos   int
}

func (p *Parser) current() byte { /* ... */ }
func (p *Parser) advance() { /* ... */ }
func (p *Parser) skipWhitespace() { /* ... */ }
Phase 2: Einfache Werte
gofunc (p *Parser) parseString() error {
    // Nur Strings parsen
}

func (p *Parser) parseNumber() error {
    // Nur Numbers parsen
}
Phase 3: Zentrale Dispatch-Funktion
gofunc (p *Parser) parseValue() error {
    switch p.current() {
    case '"': return p.parseString()
    case '0'..'9', '-': return p.parseNumber()
    case '{': return p.parseObject()  // ← Hier kommt Rekursion!
    case '[': return p.parseArray()   // ← Hier kommt Rekursion!
    }
}
Phase 4: Rekursive Strukturen
gofunc (p *Parser) parseObject() error {
    // Konsumiere '{'
    // Parse Key (String)
    // Konsumiere ':'
    // Parse Value (REKURSION!) ← parseValue() ruft sich selbst auf
    // Konsumiere '}'
}
5. Debugging-Mindset
Denke in Ebenen:
parseJSON()
  ├─ parseValue()
       ├─ parseObject()
            ├─ parseString() (key)
            └─ parseValue() ← REKURSION
                 └─ parseString() (value)
Jede Funktion hat einen klaren Job:

Input: Wo bin ich im String?
Output: Parsing erfolgreich + neue Position
Verantwortung: Nur EINE Grammatik-Regel

6. Häufige Denkfehler vermeiden
❌ Falsch: "Ich parse alles in einer großen Funktion"
✅ Richtig: "Jede Grammatik-Regel = eine Funktion"
❌ Falsch: "Ich tracke den kompletten Zustand global"
✅ Richtig: "Jede Funktion kennt nur ihre lokale Verantwortung"
❌ Falsch: "Ich implementiere erstmal alles komplett"
✅ Richtig: "Schritt für Schritt: Erst einfache Werte, dann Rekursion"
7. Mental Model für Rekursion
Stelle dir vor: Jede Funktion ist ein "Spezialist"

parseObject() ist der "Object-Experte"
parseArray() ist der "Array-Experte"
parseValue() ist der "Dispatcher" der entscheidet: "Wer ist hier zuständig?"

Vertraue dem System:

Wenn parseObject() auf einen Wert stößt, sagt es: "Hey parseValue(), du bist dran!"
parseValue() schaut und sagt: "Das ist ein Array, hey parseArray(), mach du das!"
parseArray() sagt: "Ok, für jedes Element rufe ich parseValue() auf!"

Das ist die Magie: Jede Funktion löst nur ihr kleines Problem und vertraut darauf, dass andere Funktionen ihre Probleme lösen.
8. Praktisches Vorgehen

Starte klein: Implementiere erstmal nur parseString()
Teste früh: Schreibe Tests für jede einzelne Funktion
Erweitere schrittweise: Füge eine Grammatik-Regel nach der anderen hinzu
Debug systematisch: Bei Fehlern: Welche Funktion ist für diese Stelle zuständig?

Die Kernidee: Grammar Rules → Functions → Recursive Calls