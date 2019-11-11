package matching
import (
	"github.com/gorilla/mux"
)

var logger = log.New(os.Stderr, "Matching\t", log.Lshortfile)

func StartEngine() {
	for _, product := range products {
		matchEngine.Start()
	}

	logger.Println("match engine ok")
}