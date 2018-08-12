# taxes


Paleidimas:
1. npm install
2. gulp run - paleidžią appsą.

Komentarai:
kadangi pagal užduotį nerekomenduojama naudoti duomenų bazių tai visus duomenis laikau data.json faile. Turiu httpWorker'į, kuris priklausomai nuo URL užklausos kreipiasi į:
- addController
- findController
- importController
- updateController

Duomenų saugojimo pavyzdys:
{
  "taxes": [
    {
      "municipalityID": "5b6be696fe7fa5bf3825511f",
      "municipality": "Vilnius",
      "taxRates": [
        {
          "taxRateID": 0,
          "type": "yearly",
          "startDate": "2016-01-01",
          "endDate": "2016-12-31",
          "taxRate": 0.2
        },
        {
          "taxRateID": 1,
          "type": "monthly",
          "startDate": "2016-05-01",
          "endDate": "2016-05-31",
          "taxRate": 0.3
        }
    ]
}

1. Papildomų duomenų įvedimas:
Užklausos pvz:
http://localhost:8081/taxes/add?municipality=Vilnius&type=monthly&startDate=2016-11-01&endDate=2016-11-30&taxRateValue=0.6
Reikia paduoti:
- municipality
- type (yearly/monthly/weekly/daily)
- startDate
- endDate
- taxRateValue

Consolėje išspausdinu api URL kuriuo galima patikrinti ar duomenys įsirašė į 3rd party servisą

2. Duomenų importas iš failo. Kode nurodyt, akd duomenis imtų iš data.csv failo.
Užklausa:
http://localhost:8081/taxes/import

3. Paieška. Privaloma paduoti du parametrus: municipality ir date. Jeigu pagal paduotą datą atitinka keli rezultatai tai rodomas tas, kurio datų rėžis yra mažesnis. Tai yra daily tax bus rodomas prioriteto tvarka, jei neatitinka daily tax tai bus rodomas weekly tax ir t.t.

4. Update
Užklausos pvz:
http://localhost:8081/taxes/update?municipalityID=5b6be696fe7fa5bf3825511f&taxRateID=6&type=daily&startDate=2016-09-29&endDate=2016-09-29&taxRateValue=0.78
Pirmi du parametrai nurodo kurį įrašą keisim (privalo egzistuoti toks įrašas, t.y. su tokiomis parametrų reikšmėmis):
- municipalityID
-taxRateID

Visi kiti parametrai yra tai kuo updatinsim data.json įrašą:
- type (yearly/monthly/weekly/daily)
- startDate
- endDate
- taxRateValue
