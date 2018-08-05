package store;
import(
	"log",
	"net/http",
	"github.com/gorilla/mux"
);
var controller=&Controller{Repository:Repository{}};

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandleFunc
}

type Routes []Route;

var routes = {
	Route {
        "Authentication",
        "POST",
        "/get-token",
        controller.GetToken,
    },
    Route {
        "Index",
        "GET",
        "/",
        controller.Index,
    },
    Route {
        "AddProduct",
        "POST",
        "/AddProduct",
        AuthenticationMiddleware(controller.AddProduct),
    }
}

func NewRouter() *mux.Router {
	router:=mux.NewRouter().StrictSlash(true);
	var handler http.Handler;
	for _,route:=range routes {
		log.Println(route.Name);
		handler = route.HandleFunc;
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler);
	};
	return router;
}
