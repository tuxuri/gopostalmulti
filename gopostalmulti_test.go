package gopostalmulti

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"testing"
)

func TestParse(t *testing.T) {
	log.Printf("%v", l.Parse("1265 East Fort Union Blvd., suite 250, Cotton Wood Hights, Utah 84047"))
}

var streets = []string{
	"Route 10",
	"Elizabeth Street",
	"Hill Street",
	"Orchard Lane",
	"3rd Avenue",
	"Oak Lane",
	"Mechanic Street",
	"Cedar Lane",
	"Route 4",
	"Cross Street",
	"Fairview Avenue",
	"Route 20",
	"Catherine Street",
	"1st Avenue",
	"Linden Street",
	"Brown Street",
	"Broad Street West",
	"Lafayette Avenue",
	"Jackson Street",
	"New Street",
	"Belmont Avenue",
	"Tanglewood Drive",
	"Magnolia Avenue",
	"Fieldstone Drive",
	"Woodland Road",
	"2nd Street North",
	"Surrey Lane",
	"Country Club Drive",
	"Washington Avenue",
	"Route 30",
	"Circle Drive",
	"Mill Street",
	"Hartford Road",
	"Route 1",
	"4th Avenue",
	"Rosewood Drive",
	"13th Street",
	"Elm Avenue",
	"Cambridge Court",
	"Woodland Avenue",
	"River Road",
	"Lexington Court",
	"Bayberry Drive",
	"Schoolhouse Lane",
	"Meadow Street",
	"Walnut Street",
	"Madison Street",
	"Arch Street",
	"Vine Street",
	"6th Street West",
}

var states = []string{
	"alabama",
	"al",
	"alaska",
	"ak",
	"arizona",
	"az",
	"arkansas",
	"ar",
	"california",
	"ca",
	"colorado",
	"co",
	"connecticut",
	"ct",
	"delaware",
	"de",
	"florida",
	"fl",
	"georgia",
	"ga",
	"hawaii",
	"hi",
	"idaho",
	"id",
	"illinois",
	"il",
	"indiana",
	"in",
	"iowa",
	"ia",
	"kansas",
	"ks",
	"kentucky",
	"ky",
	"louisiana",
	"la",
	"maine",
	"me",
	"maryland",
	"md",
	"massachusetts",
	"ma",
	"michigan",
	"mi",
	"minnesota",
	"mn",
	"mississippi",
	"ms",
	"missouri",
	"mo",
	"montana",
	"mt",
	"nebraska",
	"ne",
	"nevada",
	"nv",
	"new hampshire",
	"nh",
	"new jersey",
	"nj",
	"new mexico",
	"nm",
	"new york",
	"ny",
	"north carolina",
	"nc",
	"north dakota",
	"nd",
	"ohio",
	"oh",
	"oklahoma",
	"ok",
	"oregon",
	"or",
	"pennsylvania",
	"pa",
	"rhode island",
	"ri",
	"south carolina",
	"sc",
	"south dakota",
	"sd",
	"tennessee",
	"tn",
	"texas",
	"tx",
	"utah",
	"ut",
	"vermont",
	"vt",
	"virginia",
	"va",
	"washington",
	"wa",
	"west virginia",
	"wv",
	"wisconsin",
	"wi",
	"wyoming",
	"wy",
}

var cities = []string{
	"Babol",
	"Babruysk",
	"Bacabal",
	"Bacău",
	"Bacolod",
	"Bacoor",
	"Badajoz",
	"Badalona",
	"Badlapur",
	"Bafoussam",
	"Bagaha",
	"Bagalkot",
	"Bagé",
	"Baghdad",
	"Baghlan",
	"Bago",
	"Baguio",
	"Bahadurgarh",
	"Baharampur",
	"Bahawalnagar",
	"Bahawalpur",
	"Bahía",
	"Bahir",
	"Bahraich",
	"Baia",
	"Baicheng",
	"Baidoa",
	"Baidyabati",
	"Baise",
	"Baishan",
	"Baiyin",
	"Bajos",
	"Bakersfield",
	"Baku",
	"Balakovo",
	"Balashikha",
	"Baleshwar",
	"Balikpapan",
	"Baliuag",
	"Balıkesir",
	"Ballarat",
	"Ballia",
	"Bally,",
	"Balneário",
	"Bălți",
	"Baltimore",
	"Balurghat",
	"Bamako",
	"Bamenda",
	"Banda",
	"Bandar",
	"Bandar-e",
	"Bandırma",
	"Bandung",
	"Bangalore",
	"Bangaon",
	"Bangkok",
	"Bangui",
	"Banha",
	"Bani",
	"Baní",
	"Banja",
	"Banjarmasin",
	"Banjul",
	"Bankura",
	"Bansberia",
	"Banswara",
	"Baoding",
	"Baoji",
	"Baoshan",
	"Baotou",
	"Baqubah",
	"Barakaldo",
	"Baraki",
	"Baran",
	"Baranagar",
	"Baranovichi",
	"Barasat",
	"Baraut",
	"Barbacena",
	"Barcarena",
	"Barcelona",
	"Barcelona,",
	"Bardhaman",
	"Bareilly",
	"Bari",
	"Barika",
	"Bariloche",
	"Barinas",
	"Baripada",
	"Barisal",
	"Barka",
	"Barnala",
	"Barnaul",
	"Barquisimeto",
	"Barra",
	"Barrackpur",
	"Barrancabermeja",
	"Barranquilla",
	"Barreiras",
	"Barretos",
	"Barrie",
	"Barshi",
	"Barueri",
	"Baruta",
	"Barysaw",
	"Basel",
	"Bashiqa",
	"Basirhat",
	"Basra",
	"Basti",
	"Bat",
	"Batala",
	"Batam",
	"Batangas",
	"Bataysk",
	"Bathinda",
	"Batman",
	"Batna",
	"Baton",
	"Battambang",
	"Batu",
	"Batumi",
	"Bauru",
	"Bawshar",
	"Bayambang",
	"Bayamo",
	"Bayamón",
	"Bayannur",
	"Bayawan",
	"Baybay",
	"Bazhou",
	"Beau-Bassin",
	"Beaumont",
	"Beawar",
	"Béchar",
	"Beed",
	"Beersheba",
	"Begusarai",
	"Behbahan",
	"Bei'an",
	"Beihai",
	"Beijing",
	"Beipiao",
	"Beira",
	"Beirut",
	"Beit",
	"Béjaïa",
	"Bekasi",
	"Belém",
	"Belfast",
	"Belford",
	"Belgaum",
	"Belgorod",
	"Belgrade",
	"Bellary",
	"Bellevue",
	"Bello",
	"Belo",
	"Bengbu",
	"Benghazi",
	"Bengkulu",
	"Benguela",
	"Beni",
	"Benito",
	"Benoni",
	"Bento",
	"Benxi",
	"Beppu",
	"Berbera",
	"Berdyansk",
	"Berezniki",
	"Bergamo",
	"Bergen",
	"Bergisch",
	"Berhampur",
	"Berkane",
	"Berkeley",
	"Berlin",
	"Bern",
	"Berrechid",
	"Besançon",
	"Betim",
	"Bettiah",
	"Betul",
	"Bhadrak",
	"Bhadravati",
	"Bhadreswar",
	"Bhagalpur",
	"Bhalswa",
	"Bhalwal",
	"Bharatpur,",
	"Bharuch",
	"Bhatpara",
	"Bhavnagar",
	"Bhilai",
	"Bhilwara",
	"Bhimavaram",
	"Bhimdatta",
	"Bhind",
	"Bhiwadi",
	"Bhiwandi",
	"Bhiwani",
	"Bhopal",
	"Bhubaneswar",
	"Bhuj",
	"Bhusawal",
	"Białystok",
	"Bidar",
	"Bidhannagar",
	"Bielefeld",
	"Bielsko-Biała",
	"Bihar",
	"Bijapur",
	"Bikaner",
	"Bila",
	"Bilaspur",
	"Bilbao",
	"Billings",
	"Biñan",
	"Binangonan",
	"Binjai",
	"Bintulu",
	"Binzhou",
	"Biratnagar",
	"Birgunj",
	"Birigui",
	"Birjand",
	"Birmingham",
	"Birmingham,",
	"Bishkek",
	"Bishoftu",
	"Biskra",
	"Bissau",
	"Bitung",
	"Biysk",
	"Bizerte",
	"Blackburn",
	"Blackpool",
	"Blagoveshchensk",
	"Blantyre",
	"Blida",
	"Blitar",
	"Bloemfontein",
	"Blumenau",
	"Bnei",
	"Bo",
	"Boa",
	"Bobo-Dioulasso",
	"Boca",
	"Bocaue",
	"Bochum",
	"Bogor",
	"Bogota",
	"Bogra",
	"Boise",
	"Bojnord",
	"Bokaro",
	"Boksburg",
	"Bole",
	"Bologna",
	"Bolton",
	"Bolu",
	"Bolzano",
	"Boma",
	"Bonao",
	"Bongao",
	"Bonn",
	"Bordeaux",
	"Bordj",
	"Borujerd",
	"Bosaso",
	"Boston",
	"Botad",
	"Botoșani",
	"Botou",
	"Botshabelo",
	"Bottrop",
	"Botucatu",
	"Bou",
	"Bouaké",
	"Bouïra",
	"Boulder",
	"Boulogne-Billancourt",
	"Bournemouth",
	"Bozhou",
	"Bradford",
	"Braga",
	"Bragança",
	"Brăila",
	"Brampton",
	"Brasília",
	"Brașov",
	"Bratislava",
	"Bratsk",
	"Braunschweig",
	"Brazzaville",
	"Breda",
	"Bremen",
	"Bremerhaven",
	"Brescia",
	"Brest,",
	"Bridgeport",
	"Brighton",
	"Brisbane",
	"Bristol",
	"Brno",
	"Brownsville",
	"Bruges",
	"Brusque",
	"Brussels",
	"Bryansk",
	"Bucaramanga",
	"Bucharest",
	"Bucheon",
	"Budapest",
	"Budaun",
	"Buenaventura",
	"Buenos",
	"Buffalo",
	"Bujumbura",
	"Bukan",
	"Bukavu",
	"Bukhara",
	"Bulandshahr",
	"Bulawayo",
	"Bundi",
	"Buôn",
	"Buraidah",
	"Burao",
	"Burari",
	"Burbank",
	"Burewala",
	"Burgas",
	"Burgos",
	"Burhanpur",
	"Burlington,",
	"Burnaby",
	"Bursa",
	"Busan",
	"Bushehr",
	"Butuan",
	"Butwal",
	"Buxar",
	"Buzău",
	"Bydgoszcz",
	"Bytom",
}

var  l =Libpostal{
	MaxBackends: 2,
}

func init() {
	/* load test data */
	l.Init()
}

func benchmarkParse(numAddresses int, b *testing.B) {
	addressIn := make(chan string, 1000)
	for x := 0; x < runtime.NumCPU(); x++ {
		go func() {
			for add := range addressIn {
				_ = l.Parse(add)
			}
		}()
	}
	var addresses = make([]string, numAddresses)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for x := 0; x < numAddresses; x++ {
			rnd := rand.Int()
			addresses[x] = fmt.Sprintf("%d %s, %s, %s %d", rnd%10000, streets[rnd%len(streets)], cities[rnd%len(cities)], states[rnd%len(states)], rnd%10000)
		}
		b.StartTimer()
		for x := 0; x < numAddresses; x++ {
			addressIn <- addresses[x%len(addresses)]
		}
	}
	close(addressIn)
}

func BenchmarkParse50(b *testing.B) {
	benchmarkParse(50, b)
}

func BenchmarkParse500(b *testing.B) {
	benchmarkParse(500, b)
}

func BenchmarkParse5000(b *testing.B) {
	benchmarkParse(5000, b)
}

func BenchmarkParse50000(b *testing.B) {
	benchmarkParse(50000, b)
}
