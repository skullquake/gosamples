package main
import(
	"fmt"
	"os"
)
func main(){
	defer fmt.Println("never prints!")
	os.Exit(3)
}
