package routes

import (
	"api/infra/controllers"
	"api/infra/database"
	"api/infra/middleware"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func addRoutes(muxR *mux.Router, pool *pgxpool.Pool) {
	NewUserRouter(controllers.NewUserController(database.NewUserRepositoryImp(pool))).Load(muxR)
	NewCategoryRouter(controllers.NewCategoryController(database.NewCategoryRepository(pool))).Load(muxR)
	NewAuthorRouter(controllers.NewAuthorController(database.NewAuthorRepository(pool))).Load(muxR)
	NewBookRouter(controllers.NewBookController(database.NewBookRepository(pool))).Load(muxR)
	NewAuthRouter(controllers.NewAuthController(database.NewUserRepositoryImp(pool))).Load(muxR)
	muxR.Use(mux.CORSMethodMiddleware(muxR))
}

func NewServer(env map[string]string) *mux.Router {
	mux := mux.NewRouter()

	connect := database.NewConnect(env["DATABASE_URL"])

	addRoutes(mux, connect)

	mux.Use(middleware.ApplicationTypeSet)

	return mux
}
