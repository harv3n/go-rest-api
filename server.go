package main

import (
	"fmt"
	"net/http"

	"go-rest-api/controller"
	router "go-rest-api/http"
	"go-rest-api/repository"
	"go-rest-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	httpRouter     router.Router             = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	const port string = ":8000"

	httpRouter.Get("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})
	httpRouter.Get("/posts", postController.GetPosts)
	httpRouter.Post("/posts", postController.AddPosts)

	httpRouter.Serve(port)
}
