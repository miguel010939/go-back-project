package factories

import (
	"fmt"
	"main.go/models"
	"math/rand"
)

func randomUser() *models.UserSignUpForm {
	username := names[rand.Intn(len(names))]
	email := fmt.Sprintf("%s%d@gmail.com", username, rand.Intn(69420)) // 2 users cant have the same email
	password := passwords[rand.Intn(len(passwords))]
	return &models.UserSignUpForm{
		Username: username,
		Email:    email,
		Password: password,
	}
}

func randomProduct() *models.ProductForm {
	name := productNames[rand.Intn(len(productNames))]
	var description string
	for i := 0; i < 9; i++ {
		description += words[rand.Intn(len(words))] + " "
	}
	image := images[rand.Intn(len(images))]
	return &models.ProductForm{
		Name:        name,
		Description: description,
		ImageUrl:    image,
	}
}

var (
	names = []string{
		"manuel", "jose", "luis", "adrian", "rafael", "ivan", "juan",
		"ana", "eva", "elena", "alicia", "mireia", "uxia", "celia",
		"john", "michael", "william", "james", "robert", "mary", "patricia",
		"jennifer", "linda", "elizabeth", "barbara", "susan", "jessica", "sarah",
		"david", "richard", "thomas", "charles", "christopher", "daniel", "matthew",
		"anthony", "donald", "mark", "paul", "steven", "andrew", "kenneth",
		"george", "joshua", "kevin", "brian", "edward", "ronald", "timothy",
		"jason", "jeffrey", "ryan", "jacob", "gary", "nicholas", "eric",
		"stephen", "jonathan", "larry", "scott",
	}

	productNames = []string{
		"car", "airplane", "bicycle", "bus", "train", "motorcycle", "truck", "boat", "scooter", "helicopter",
		"phone", "laptop", "tablet", "camera", "television", "printer", "monitor", "keyboard", "mouse", "speakers",
		"chair", "table", "sofa", "bed", "lamp", "bookshelf", "desk", "cabinet", "stool", "drawer",
		"refrigerator", "oven", "microwave", "toaster", "blender", "coffee maker", "washing machine", "dryer", "vacuum cleaner", "dishwasher",
		"shoes", "shirt", "pants", "jacket", "hat", "gloves", "scarf", "belt", "socks", "umbrella",
	}
	words = []string{
		"awesome", "goose", "is", "has", "a", "opening", "safe", "children", "snow", "sweet", "unhealthy",
		"cat", "runs", "in", "wall", "tree", "river", "jumps", "quickly", "soft", "blue", "apple", "cloud", "plays",
		"sun", "bright", "happy", "green", "sings", "under", "moon", "light", "strong", "tiny", "frog", "flies",
		"pond", "quiet", "yellow", "garden", "bright", "loud", "big", "small", "heavy", "light", "fast", "slow",
		"red", "colorful", "night", "day", "sunny", "rainy", "stormy", "breeze", "warm", "cold", "hot", "cool",
		"windy", "bouncy", "circle", "square", "triangle", "oval", "smooth", "rough", "softly", "quick", "lazy",
		"sharp", "dull", "wild", "calm", "noisy", "quietly", "tall", "short", "thin", "thick", "wide", "narrow",
		"curved", "straight", "round", "flat", "deep", "shallow", "rich", "poor", "clean", "dirty", "fresh", "stale",
		"brightly", "dimly", "warmly", "coldly", "joyful", "sad", "funny", "serious", "brave", "cowardly", "fierce", "gentle",
		"polite", "rude", "kind", "mean", "clever", "dull", "strongly", "weakly", "happily", "angrily", "patient", "impatient",
		"generous", "stingy", "greedy", "selfish", "helpful", "unhelpful", "friendly", "unfriendly", "hopeful", "hopeless", "beautiful", "ugly",
		"graceful", "clumsy", "careful", "careless", "thoughtful", "thoughtless", "honest", "dishonest", "loyal", "disloyal", "reliable", "unreliable",
		"creative", "unimaginative", "productive", "unproductive", "organized", "disorganized", "efficient", "inefficient", "responsible", "irresponsible",
		"confident", "insecure", "sincere", "insincere", "genuine", "fake", "real", "imaginary", "magical", "ordinary", "mysterious", "obvious",
		"ancient", "modern", "antique", "new", "classic", "trendy", "popular", "unpopular", "common", "rare", "unique", "standard",
		"familiar", "unfamiliar", "interesting", "boring", "exciting", "dull", "dynamic", "static", "active", "inactive", "lively", "lifeless",
		"car", "dog", "house", "flower", "bird", "fish", "rock", "star", "tree", "book", "bottle", "cup",
		"table", "chair", "computer", "phone", "shoe", "hat", "clock", "pen", "pencil", "paper", "notebook", "bag",
		"shirt", "pants", "dress", "coat", "boat", "plane", "train", "bus", "bike", "truck", "road", "path",
		"bridge", "river", "lake", "sea", "ocean", "forest", "mountain", "valley", "hill", "field", "farm", "garden",
		"park", "city", "town", "village", "building", "room", "door", "window", "wall", "roof", "floor", "ceiling",
		"plum", "pencil", "bandage", "doll", "notepad", "fence", "potato", "corn", "bench", "statue", "waterfall", "fountain",
		"factory", "library", "school", "college", "university", "hospital", "clinic", "market", "store", "mall", "shop", "restaurant",
		"cafe", "bar", "hotel", "motel", "inn", "house", "home", "apartment", "flat", "condo", "office", "warehouse",
		"plant", "flower", "bush", "tree", "wood", "forest", "jungle", "swamp", "desert", "plain", "meadow", "prairie",
		"barn", "stable", "silo", "shed", "garage", "dock", "pier", "wharf", "airport", "station", "depot", "terminal",
	}
	passwords = []string{"12", "123", "1234", "12345", "123456", "contrasenha", "password", "nomoreideas"}

	images = []string{"https://upload.wikimedia.org/wikipedia/commons/6/66/Openlogo-debianV2.svg",
		"https://upload.wikimedia.org/wikipedia/commons/a/ab/Logo-ubuntu_cof-orange-hex.svg",
		"https://upload.wikimedia.org/wikipedia/commons/3/3f/Logo_Linux_Mint.png",
		"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fedora_logo.svg",
		"https://upload.wikimedia.org/wikipedia/commons/4/46/Pop%21_OS_Icon.svg"}
)
